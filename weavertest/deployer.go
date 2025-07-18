// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package weavertest

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"reflect"
	"sync"

	"github.com/thunur/weaver/internal/control"
	"github.com/thunur/weaver/internal/weaver"
	"github.com/thunur/weaver/runtime"
	"github.com/thunur/weaver/runtime/codegen"
	"github.com/thunur/weaver/runtime/envelope"
	"github.com/thunur/weaver/runtime/logging"
	"github.com/thunur/weaver/runtime/protos"
	"github.com/google/uuid"
	"golang.org/x/exp/maps"
	"golang.org/x/sync/errgroup"
)

// DefaultReplication is the default number of times a component is replicated.
//
// TODO(mwhittaker): Include this in the Options struct?
const DefaultReplication = 2

// deployer is the weavertest multiprocess deployer. Every multiprocess
// weavertest runs its own deployer. The main component is run in the same
// process as the deployer, which is the same process as the unit test. All
// other components are run in subprocesses.
//
// This deployer differs from 'weaver multi' in two key ways.
//
//  1. This deployer doesn't implement unneeded features (e.g., traces,
//     metrics, routing, health checking). This greatly simplifies the
//     implementation.
//  2. This deployer handles the fact that the main component is run in the
//     same process as the deployer. This is special to weavertests and
//     requires special care. See start() for more details.
type deployer struct {
	ctx        context.Context
	ctxCancel  context.CancelFunc
	tmpDir     string
	runner     Runner                 // holds runner-specific info like config
	wlet       *protos.WeaveletArgs   // info for subprocesses
	config     *protos.AppConfig      // application config
	colocation map[string]string      // maps component to group
	running    errgroup.Group         // collects errors from goroutines
	local      map[string]bool        // Components that should run locally
	log        func(*protos.LogEntry) // logs the passed in string
	sysLogger  *slog.Logger           // system message logger

	mu     sync.Mutex        // guards fields below
	groups map[string]*group // groups, by group name
	err    error             // error the test was terminated with, if any.
}

// A group contains information about a co-location group.
type group struct {
	name        string                               // group name
	controllers []control.WeaveletControl            // weavelet controllers
	components  map[string]bool                      // started components
	addresses   map[string]bool                      // weavelet addresses
	subscribers map[string][]control.WeaveletControl // routing info subscribers, by component
}

// handler handles a connection to a weavelet.
type handler struct {
	*deployer
	group      *group
	controller control.WeaveletControl
	subscribed map[string]bool // routing info subscriptions, by component
}

var _ envelope.EnvelopeHandler = &handler{}

// newDeployer returns a new weavertest multiprocess deployer. locals contains
// components that should be co-located with the main component and not
// replicated.
func newDeployer(ctx context.Context, wlet *protos.WeaveletArgs, config *protos.AppConfig, runner Runner, locals []reflect.Type, logWriter func(*protos.LogEntry), tmpDir string) *deployer {
	colocation := map[string]string{}
	for _, group := range config.Colocate {
		for _, c := range group.Components {
			colocation[c] = group.Components[0]
		}
	}
	ctx, cancel := context.WithCancel(ctx)
	d := &deployer{
		ctx:        ctx,
		ctxCancel:  cancel,
		tmpDir:     tmpDir,
		runner:     runner,
		wlet:       wlet,
		config:     config,
		colocation: colocation,
		groups:     map[string]*group{},
		local:      map[string]bool{},
		log:        logWriter,
	}
	d.sysLogger = slog.New(&logging.LogHandler{
		Opts: logging.Options{
			App:       d.wlet.App,
			Component: "deployer",
			Attrs:     []string{"serviceweaver/system", ""},
		},
		Write: d.log,
	})

	for _, local := range locals {
		name := fmt.Sprintf("%s/%s", local.PkgPath(), local.Name())
		d.local[name] = true
	}

	// Fakes need to be local as well.
	for _, fake := range runner.Fakes {
		name := fmt.Sprintf("%s/%s", fake.intf.PkgPath(), fake.intf.Name())
		d.local[name] = true
	}

	return d
}

func (d *deployer) start(opts weaver.RemoteWeaveletOptions) (*weaver.RemoteWeavelet, error) {
	// Run an envelope connection to the main co-location group.
	wlet := &protos.WeaveletArgs{
		App:             d.wlet.App,
		DeploymentId:    d.wlet.DeploymentId,
		Id:              uuid.New().String(),
		InternalAddress: "localhost:0",
	}

	child := envelope.NewInProcessChild()

	// We concurrently start the creation of the weavelet and the creation of the envelope
	// connection to the weavelet since they both talk to each other.
	type weaveletResult struct {
		weavelet *weaver.RemoteWeavelet
		err      error
	}
	wchan := make(chan weaveletResult)
	go func() {
		bootstrap := runtime.Bootstrap{
			Args: child.Args(), // Will block
		}
		weavelet, err := weaver.NewRemoteWeavelet(d.ctx, codegen.Registered(), bootstrap, opts)
		wchan <- weaveletResult{weavelet, err} // Give weavelet to caller
		wchan <- weaveletResult{weavelet, err} // Give weavelet to NewEnvelope goroutine
	}()

	// NOTE: NewEnvelope initiates a blocking handshake with the weavelet
	// and therefore we run the rest of the initialization in a goroutine which
	// will wait for weaver.Run to create a weavelet.
	go func() {
		e, err := envelope.NewEnvelope(d.ctx, wlet, d.config, envelope.Options{
			TmpDir: d.tmpDir,
			Logger: d.sysLogger,
			Child:  child,
		})
		if err != nil {
			d.stop(err)
			return
		}

		// Get weavelet info when it is ready.
		w := <-wchan
		if w.err != nil {
			d.stop(err)
			return
		}

		d.mu.Lock()
		defer d.mu.Unlock()
		g := d.group("main")
		g.controllers = append(g.controllers, w.weavelet)
		handler := &handler{
			deployer:   d,
			group:      g,
			subscribed: map[string]bool{},
			controller: w.weavelet,
		}
		d.running.Go(func() error {
			err := e.Serve(handler)
			d.stop(err)
			return err
		})
		if err := d.registerReplica(g, e.WeaveletAddress()); err != nil {
			d.stopLocked(fmt.Errorf(`cannot register the replica for "main": %w`, err))
		}
	}()

	w := <-wchan
	return w.weavelet, w.err
}

// stop stops the deployer.
// REQUIRES: d.mu is not held.
func (d *deployer) stop(err error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.stopLocked(err)
}

// stopLocked stops the deployer.
// REQUIRES: d.mu is held.
func (d *deployer) stopLocked(err error) {
	if d.err == nil && !errors.Is(err, context.Canceled) {
		d.err = err
	}
	d.ctxCancel()
}

// cleanup cleans up all of the running envelopes' state.
func (d *deployer) cleanup() error {
	d.ctxCancel()
	d.running.Wait()
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.err
}

// LogBatch implements the control.DeployerControl interface.
func (d *deployer) LogBatch(ctx context.Context, batch *protos.LogEntryBatch) error {
	for _, entry := range batch.Entries {
		d.log(entry)
	}
	return nil
}

// HandleTraceSpans implements the envelope.EnvelopeHandler interface.
func (d *deployer) HandleTraceSpans(context.Context, *protos.TraceSpans) error {
	// Ignore traces.
	return nil
}

// GetListenerAddress implements the envelope.EnvelopeHandler interface.
func (d *deployer) GetListenerAddress(_ context.Context, req *protos.GetListenerAddressRequest) (*protos.GetListenerAddressReply, error) {
	return &protos.GetListenerAddressReply{Address: "localhost:0"}, nil
}

// ExportListener implements the envelope.EnvelopeHandler interface.
func (d *deployer) ExportListener(_ context.Context, req *protos.ExportListenerRequest) (*protos.ExportListenerReply, error) {
	return &protos.ExportListenerReply{}, nil
}

func (*deployer) GetSelfCertificate(context.Context, *protos.GetSelfCertificateRequest) (*protos.GetSelfCertificateReply, error) {
	// This deployer doesn't enable mTLS.
	panic("unused")
}

func (*deployer) VerifyClientCertificate(context.Context, *protos.VerifyClientCertificateRequest) (*protos.VerifyClientCertificateReply, error) {
	// This deployer doesn't enable mTLS.
	panic("unused")
}

func (*deployer) VerifyServerCertificate(context.Context, *protos.VerifyServerCertificateRequest) (*protos.VerifyServerCertificateReply, error) {
	// This deployer doesn't enable mTLS.
	panic("unused")
}

// registerReplica registers the information about a colocation group replica
// (i.e., a weavelet).
func (d *deployer) registerReplica(g *group, replicaAddr string) error {
	// Update addresses.
	if g.addresses[replicaAddr] {
		// Replica already registered.
		return nil
	}
	g.addresses[replicaAddr] = true

	// Notify subscribers.
	for component := range g.components {
		update := &protos.UpdateRoutingInfoRequest{RoutingInfo: g.routing(component)}
		for _, sub := range g.subscribers[component] {
			if _, err := sub.UpdateRoutingInfo(d.ctx, update); err != nil {
				return err
			}
		}
	}
	return nil
}

// ActivateComponent implements the control.DeployerControl interface.
func (h *handler) ActivateComponent(ctx context.Context, req *protos.ActivateComponentRequest) (*protos.ActivateComponentReply, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	// Update the set of components in the target co-location group.
	target := h.deployer.group(req.Component)
	if !target.components[req.Component] {
		target.components[req.Component] = true

		// Notify the weavelets.
		update := &protos.UpdateComponentsRequest{Components: maps.Keys(target.components)}
		for _, controller := range target.controllers {
			if _, err := controller.UpdateComponents(ctx, update); err != nil {
				return nil, err
			}
		}

		// Notify the subscribers.
		routing := &protos.UpdateRoutingInfoRequest{RoutingInfo: target.routing(req.Component)}
		for _, sub := range target.subscribers[req.Component] {
			if _, err := sub.UpdateRoutingInfo(ctx, routing); err != nil {
				return nil, err
			}
		}
	}

	// Subscribe to the component's routing info.
	if !h.subscribed[req.Component] {
		h.subscribed[req.Component] = true

		if !h.runner.forceRPC && h.group.name == target.name {
			// Route locally.
			routing := &protos.UpdateRoutingInfoRequest{RoutingInfo: &protos.RoutingInfo{Component: req.Component, Local: true}}
			if _, err := h.controller.UpdateRoutingInfo(ctx, routing); err != nil {
				return nil, err
			}
		} else {
			// Route remotely.
			target.subscribers[req.Component] = append(target.subscribers[req.Component], h.controller)
			if _, err := h.controller.UpdateRoutingInfo(ctx, &protos.UpdateRoutingInfoRequest{RoutingInfo: target.routing(req.Component)}); err != nil {
				return nil, err
			}
		}
	}

	// Start the co-location group
	return &protos.ActivateComponentReply{}, h.deployer.startGroup(target)
}

// startGroup starts the provided co-location group in a subprocess, if it
// hasn't already been started.
//
// REQUIRES: d.mu is held.
func (d *deployer) startGroup(g *group) error {
	if len(g.controllers) > 0 {
		// Envelopes already started
		return nil
	}

	update := &protos.UpdateComponentsRequest{Components: maps.Keys(g.components)}
	for r := 0; r < DefaultReplication; r++ {
		// Start the weavelet.
		wlet := &protos.WeaveletArgs{
			App:             d.wlet.App,
			DeploymentId:    d.wlet.DeploymentId,
			Id:              uuid.New().String(),
			InternalAddress: "localhost:0",
		}
		handler := &handler{
			deployer:   d,
			group:      g,
			subscribed: map[string]bool{},
		}
		logger := slog.New(&logging.LogHandler{
			Opts:  logging.Options{Component: "envelope", Weavelet: wlet.Id},
			Write: d.log,
		})
		e, err := envelope.NewEnvelope(d.ctx, wlet, d.config, envelope.Options{
			Logger: logger,
		})
		if err != nil {
			return err
		}
		d.running.Go(func() error {
			err := e.Serve(handler)
			d.stop(err)
			return err
		})
		if err := d.registerReplica(g, e.WeaveletAddress()); err != nil {
			return err
		}
		wc := e.WeaveletControl()
		if _, err := wc.UpdateComponents(d.ctx, update); err != nil {
			return err
		}
		handler.controller = wc
		g.controllers = append(g.controllers, wc)
	}
	return nil
}

// group returns the group that corresponds to the given component.
//
// REQUIRES: d.mu is held.
func (d *deployer) group(component string) *group {
	var name string
	if !d.runner.multi {
		name = "main" // Everything is in one group.
	} else if d.local[component] {
		name = "main" // Run locally
	} else if x, ok := d.colocation[component]; ok {
		name = x // Use specified group
	} else {
		name = component // A group of its own
	}

	g, ok := d.groups[name]
	if !ok {
		g = &group{
			name:        name,
			components:  map[string]bool{},
			addresses:   map[string]bool{},
			subscribers: map[string][]control.WeaveletControl{},
		}
		d.groups[name] = g
	}
	return g
}

// routing returns the RoutingInfo for the provided component.
//
// REQUIRES: d.mu is held.
func (g *group) routing(component string) *protos.RoutingInfo {
	return &protos.RoutingInfo{
		Component: component,
		Replicas:  maps.Keys(g.addresses),
	}
}
