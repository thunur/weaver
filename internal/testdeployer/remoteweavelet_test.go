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

package testdeployer

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/thunur/weaver/internal/reflection"
	"github.com/thunur/weaver/internal/weaver"
	"github.com/thunur/weaver/runtime"
	"github.com/thunur/weaver/runtime/codegen"
	"github.com/thunur/weaver/runtime/deployers"
	"github.com/thunur/weaver/runtime/envelope"
	"github.com/thunur/weaver/runtime/logging"
	"github.com/thunur/weaver/runtime/protomsg"
	"github.com/thunur/weaver/runtime/protos"
	"github.com/google/pprof/profile"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

// TODO(mwhittaker): In addition to the tests that are currently failing, here
// are some situations where it's unclear what a weavelet should do.
//
// - A component method call panics.
// - A component's Init method fails.

var (
	componenta = "github.com/thunur/weaver/internal/testdeployer/a"
	componentb = "github.com/thunur/weaver/internal/testdeployer/b"
	componentc = "github.com/thunur/weaver/internal/testdeployer/c"
	componentd = "github.com/thunur/weaver/internal/testdeployer/d"
	colocated  = map[string][]string{"1": {componenta, componentb, componentc}}
)

// A weavelet is a connection to a RemoteWeavelet running in this process.
type weavelet struct {
	cancel  context.CancelFunc     // shuts down the weavelet
	env     *envelope.Envelope     // envelope
	wlet    *weaver.RemoteWeavelet // weavelet
	threads *errgroup.Group        // background threads
}

// spawn spawns a weavelet with the provided info and handler.
func spawn(ctx context.Context, info *protos.WeaveletArgs, handler envelope.EnvelopeHandler, log *slog.Logger, tmpDir string) (*weavelet, error) {
	// envelope.NewEnvelope blocks performing a handshake with the weavelet, so
	// we have to run it in a separate goroutine.
	ctx, cancel := context.WithCancel(ctx)
	threads, ctx := errgroup.WithContext(ctx)
	errs := make(chan error)
	child := envelope.NewInProcessChild()
	var env *envelope.Envelope
	go func() {
		var err error
		env, err = envelope.NewEnvelope(ctx, info, &protos.AppConfig{},
			envelope.Options{
				TmpDir: tmpDir,
				Logger: log,
				Child:  child,
			})
		errs <- err
	}()

	// Create the weavelet.
	wlet, err := weaver.NewRemoteWeavelet(
		ctx,
		codegen.Registered(),
		runtime.Bootstrap{
			Args: child.Args(),
		},
		weaver.RemoteWeaveletOptions{},
	)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("spawn: NewRemoteWeavelet: %w", err)
	}

	// Wait for the EnvelopeConn to finish the handshake.
	if err := <-errs; err != nil {
		cancel()
		return nil, fmt.Errorf("spawn: NewEnvelopeConn: %v", err)
	}

	// Monitor the envelope and weavelet in background threads. Discard errors
	// after the context has been cancelled, as those are expected.
	threads.Go(func() error {
		if err := wlet.Wait(); err != nil && ctx.Err() == nil {
			return err
		}
		return nil
	})
	threads.Go(func() error {
		if err := env.Serve(handler); err != nil && ctx.Err() == nil {
			return err
		}
		return nil
	})

	return &weavelet{
		cancel:  cancel,
		env:     env,
		wlet:    wlet,
		threads: threads,
	}, nil
}

// deployer is a simple testing deployer that spawns all weavelets in the
// current process.
type deployer struct {
	t         *testing.T         // underlying unit test
	ctx       context.Context    // context used to spawn weavelets
	cancel    context.CancelFunc // shuts down the deployer and all weavelets
	info      *protos.WeaveletArgs
	logger    *logging.TestLogger  // logger
	threads   *errgroup.Group      // background threads
	placement map[string][]string  // weavelet -> components
	placedAt  map[string][]string  // component -> weavelets
	weavelets map[string]*weavelet // weavelets

	// A unit test can override the following envelope methods to do things
	// like inject errors or return invalid values.
	mu                 sync.Mutex
	activateComponent  func(context.Context, *protos.ActivateComponentRequest) (*protos.ActivateComponentReply, error)
	getListenerAddress func(context.Context, *protos.GetListenerAddressRequest) (*protos.GetListenerAddressReply, error)
	exportListener     func(context.Context, *protos.ExportListenerRequest) (*protos.ExportListenerReply, error)
}

// deploy creates a new test deployer.
//
// # Placement
//
// placement is a map from weavelet names to component names. It describes how
// many weavelets to run, what to name them, and which components the weavelets
// should host. For example, the following placement creates a single weavelet
// that hosts components a, b, and c:
//
//	map[string][]string{"1": {componenta, componentb, componentc}}
//
// The following placement creates three weavelets, each with its own
// component:
//
//	map[string][]string{
//	    "1": {componenta},
//	    "2": {componentb},
//	    "3": {componentc},
//	}
//
// The following placement replicates a, b, and c across two weavelets:
//
//	map[string][]string{
//	    "1": {componenta, componentb, componentc},
//	    "2": {componenta, componentb, componentc},
//	}
func deploy(t *testing.T, ctx context.Context, placement map[string][]string) *deployer {
	return deployWithInfo(t, ctx, placement, &protos.WeaveletArgs{
		App:             "remoteweavelet_test.go",
		DeploymentId:    fmt.Sprint(os.Getpid()),
		InternalAddress: "localhost:0",
	})
}

// deployWithInfo is identical to deploy but with an additional WeaveletArgs
// argument.
func deployWithInfo(t *testing.T, ctx context.Context, placement map[string][]string, info *protos.WeaveletArgs) *deployer {
	t.Helper()

	// Invert placement.
	placedAt := map[string][]string{}
	for name, components := range placement {
		for _, component := range components {
			placedAt[component] = append(placedAt[component], name)
		}
	}

	// Create the deployer.
	ctx, cancel := context.WithCancel(ctx)
	threads, ctx := errgroup.WithContext(ctx)
	d := &deployer{
		t:         t,
		info:      protomsg.Clone(info),
		ctx:       ctx,
		cancel:    cancel,
		logger:    logging.NewTestLogger(t, testing.Verbose()),
		threads:   threads,
		placement: placement,
		placedAt:  placedAt,
		weavelets: map[string]*weavelet{},
	}

	// Handle calls from weavelets.
	logger := slog.New(&logging.LogHandler{Write: d.logger.Log})

	// Spawn the weavelets.
	tmpDir := t.TempDir()
	for name := range placement {
		info := d.info
		info.Id = uuid.New().String()
		weavelet, err := spawn(ctx, info, d, logger, tmpDir)
		if err != nil {
			t.Fatal(err)
		}
		d.weavelets[name] = weavelet
	}
	return d
}

// shutdown shuts down a deployer and its weavelets.
func (d *deployer) shutdown() {
	d.cancel()
	for _, weavelet := range d.weavelets {
		if err := weavelet.threads.Wait(); err != nil {
			d.t.Fatal(err)
		}
	}
}

// LogBatch implements the control.DeployerControl interface.
func (d *deployer) LogBatch(ctx context.Context, batch *protos.LogEntryBatch) error {
	for _, entry := range batch.Entries {
		d.logger.Log(entry)
	}
	return nil
}

// ActivateComponent implements the EnvelopeHandler interface.
func (d *deployer) ActivateComponent(ctx context.Context, req *protos.ActivateComponentRequest) (*protos.ActivateComponentReply, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.activateComponent != nil {
		return d.activateComponent(ctx, req)
	}

	// Start the requested component.
	components := &protos.UpdateComponentsRequest{Components: []string{req.Component}}
	replicas := []string{}
	for _, name := range d.placedAt[req.Component] {
		weavelet := d.weavelets[name]
		if _, err := weavelet.wlet.UpdateComponents(ctx, components); err != nil {
			return nil, err
		}
		replicas = append(replicas, weavelet.env.WeaveletAddress())
	}

	// For simplicity, route locally if there is a single weavelet, and route
	// remotely otherwise. We also report routing info to all weavelets, not
	// just those that called ActivateComponent.
	routing := &protos.UpdateRoutingInfoRequest{}
	if len(d.placement) == 1 {
		routing.RoutingInfo = &protos.RoutingInfo{
			Component: req.Component,
			Local:     true,
		}
	} else {
		routing.RoutingInfo = &protos.RoutingInfo{
			Component: req.Component,
			Replicas:  replicas,
		}
	}
	for _, weavelet := range d.weavelets {
		if _, err := weavelet.wlet.UpdateRoutingInfo(ctx, routing); err != nil {
			return nil, err
		}
	}
	return &protos.ActivateComponentReply{}, nil
}

// GetListenerAddress implements the EnvelopeHandler interface.
func (d *deployer) GetListenerAddress(ctx context.Context, req *protos.GetListenerAddressRequest) (*protos.GetListenerAddressReply, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.getListenerAddress != nil {
		return d.getListenerAddress(ctx, req)
	}
	return &protos.GetListenerAddressReply{Address: ":0"}, nil
}

// ExportListenerAddress implements the EnvelopeHandler interface.
func (d *deployer) ExportListener(ctx context.Context, req *protos.ExportListenerRequest) (*protos.ExportListenerReply, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.exportListener != nil {
		return d.exportListener(ctx, req)
	}
	return &protos.ExportListenerReply{}, nil
}

// HandleTraceSpans implements the EnvelopeHandler interface.
func (d *deployer) HandleTraceSpans(context.Context, *protos.TraceSpans) error {
	return nil
}

// GetSelfCertificate implements the EnvelopeHandler interface.
func (d *deployer) GetSelfCertificate(context.Context, *protos.GetSelfCertificateRequest) (*protos.GetSelfCertificateReply, error) {
	d.t.Fatal("unimplemented")
	return nil, nil
}

// VerifyClientCertificate implements the EnvelopeHandler interface.
func (d *deployer) VerifyClientCertificate(context.Context, *protos.VerifyClientCertificateRequest) (*protos.VerifyClientCertificateReply, error) {
	d.t.Fatal("unimplemented")
	return nil, nil
}

// VerifyServerCertificate implements the EnvelopeHandler interface.
func (d *deployer) VerifyServerCertificate(context.Context, *protos.VerifyServerCertificateRequest) (*protos.VerifyServerCertificateReply, error) {
	d.t.Fatal("unimplemented")
	return nil, nil
}

// testComponents tests that the components spawned by d are working properly.
func testComponents(dep *deployer) {
	dep.t.Helper()
	for _, name := range dep.placedAt[componenta] {
		x, err := dep.weavelets[name].wlet.GetIntf(reflection.Type[a]())
		if err != nil {
			dep.t.Fatal(err)
		}
		const want = 42
		got, err := x.(a).A(dep.ctx, want)
		if err != nil {
			dep.t.Fatal(err)
		}
		if got != want {
			dep.t.Fatalf("A(%d): got %d, want %d", want, got, want)
		}
	}
	for _, name := range dep.placedAt[componentd] {
		x, err := dep.weavelets[name].wlet.GetIntf(reflection.Type[d]())
		if err != nil {
			dep.t.Fatal(err)
		}
		got, err := x.(d).D(dep.ctx)
		if err != nil {
			dep.t.Fatal(err)
		}
		want := dep.info.DeploymentId
		if got != want {
			dep.t.Fatalf("D(): got %s, want %s", got, want)
		}
	}
}

func TestLocalhostWeaveletAddress(t *testing.T) {
	// Start the weavelet with internal address "localhost:12345".
	d := deployWithInfo(t, context.Background(), colocated, &protos.WeaveletArgs{
		App:             "remoteweavelet_test.go",
		DeploymentId:    fmt.Sprint(os.Getpid()),
		ControlSocket:   deployers.NewUnixSocketPath(t.TempDir()),
		InternalAddress: "localhost:12345",
	})
	defer d.shutdown()
	got := d.weavelets["1"].env.WeaveletAddress()
	const want = "tcp://127.0.0.1:12345"
	if got != want {
		t.Fatalf("DialAddr: got %q, want %q", got, want)
	}
}

func TestHostnameWeaveletAddress(t *testing.T) {
	// Start the weavelet with internal address "$HOSTNAME:12345".
	hostname, err := os.Hostname()
	if err != nil {
		t.Fatal(err)
	}
	ips, err := net.LookupIP(hostname)
	if err != nil {
		t.Fatalf("net.LookupIP(%q): %v", hostname, err)
	}
	if len(ips) == 0 {
		t.Fatalf("net.LookupIP(%q): no IPs", hostname)
	}

	d := deployWithInfo(t, context.Background(), colocated, &protos.WeaveletArgs{
		App:             "remoteweavelet_test.go",
		DeploymentId:    fmt.Sprint(os.Getpid()),
		ControlSocket:   deployers.NewUnixSocketPath(t.TempDir()),
		InternalAddress: net.JoinHostPort(ips[0].String(), "12345"),
	})
	defer d.shutdown()
	got := d.weavelets["1"].env.WeaveletAddress()
	want := fmt.Sprintf("tcp://%s", net.JoinHostPort(ips[0].String(), "12345"))
	if got != want {
		t.Fatalf("DialAddr: got %q, want %q", got, want)
	}
}

func TestErrorFreeColocatedExecution(t *testing.T) {
	d := deploy(t, context.Background(), colocated)
	defer d.shutdown()
	testComponents(d)
}

func TestErrorFreeDistributedExecution(t *testing.T) {
	placement := map[string][]string{
		"1": {componenta, componentb},
		"2": {componentb, componentc},
		"3": {componenta, componentc},
	}
	d := deploy(t, context.Background(), placement)
	defer d.shutdown()
	testComponents(d)
}

func TestFailActivateComponent(t *testing.T) {
	d := deploy(t, context.Background(), colocated)
	defer d.shutdown()

	// Fail ActivateComponent a number of times.
	const n = 3
	failures := map[string]int{}
	d.activateComponent = func(ctx context.Context, req *protos.ActivateComponentRequest) (*protos.ActivateComponentReply, error) {
		if failures[req.Component] < n {
			failures[req.Component]++
			return nil, fmt.Errorf("simulated ActivateComponent(%q) failure", req.Component)
		}

		routing := &protos.UpdateRoutingInfoRequest{RoutingInfo: &protos.RoutingInfo{Component: req.Component, Local: true}}
		if _, err := d.weavelets["1"].wlet.UpdateRoutingInfo(ctx, routing); err != nil {
			return nil, err
		}
		components := &protos.UpdateComponentsRequest{Components: []string{req.Component}}
		if _, err := d.weavelets["1"].wlet.UpdateComponents(ctx, components); err != nil {
			return nil, err
		}
		return &protos.ActivateComponentReply{}, nil
	}

	testComponents(d)
}

func TestFailGetListenerAddress(t *testing.T) {
	t.Skip("TODO(mwhittaker): Make this test pass.")

	d := deploy(t, context.Background(), colocated)
	defer d.shutdown()

	// Fail GetListenerAddress a number of times.
	const n = 3
	failures := map[string]int{}
	d.getListenerAddress = func(ctx context.Context, req *protos.GetListenerAddressRequest) (*protos.GetListenerAddressReply, error) {
		if failures[req.Name] < n {
			failures[req.Name]++
			return nil, fmt.Errorf("simulated GetListenerAddress(%q) failure", req.Name)
		}
		return &protos.GetListenerAddressReply{Address: ":0"}, nil
	}

	testComponents(d)
}

func TestGetListenerAddressReturnsInvalidAddress(t *testing.T) {
	t.Skip("TODO(mwhittaker): Make this test pass.")

	d := deploy(t, context.Background(), colocated)
	defer d.shutdown()

	// Return an invalid listener a number of times.
	const n = 3
	failures := map[string]int{}
	d.getListenerAddress = func(ctx context.Context, req *protos.GetListenerAddressRequest) (*protos.GetListenerAddressReply, error) {
		if failures[req.Name] < n {
			failures[req.Name]++
			return &protos.GetListenerAddressReply{Address: "this is not a valid address"}, nil
		}
		return &protos.GetListenerAddressReply{Address: ":0"}, nil
	}

	testComponents(d)
}

func TestGetListenerAddressReturnsAddressAlreadyInUse(t *testing.T) {
	t.Skip("TODO(mwhittaker): Make this test pass.")

	// Listen on port 45678.
	lis, err := net.Listen("tcp", "localhost:45678")
	if err != nil {
		t.Fatal(err)
	}
	defer lis.Close()

	d := deploy(t, context.Background(), colocated)
	defer d.shutdown()

	// Tell the weavelet to listen on port 45678 a number of times.
	const n = 3
	failures := map[string]int{}
	d.getListenerAddress = func(ctx context.Context, req *protos.GetListenerAddressRequest) (*protos.GetListenerAddressReply, error) {
		if failures[req.Name] < n {
			failures[req.Name]++
			return &protos.GetListenerAddressReply{Address: "localhost:45678"}, nil
		}
		return &protos.GetListenerAddressReply{Address: ":0"}, nil
	}
	testComponents(d)
}

func TestFailExportListener(t *testing.T) {
	d := deploy(t, context.Background(), colocated)
	defer d.shutdown()

	// Fail ExportListener a number of times.
	const n = 3
	failures := map[string]int{}
	d.exportListener = func(ctx context.Context, req *protos.ExportListenerRequest) (*protos.ExportListenerReply, error) {
		if failures[req.Listener] < n {
			failures[req.Listener]++
			return nil, fmt.Errorf("simulated ExportListener(%q) error", req.Listener)
		}
		return &protos.ExportListenerReply{}, nil
	}

	testComponents(d)
}

func TestExportListenerReturnsError(t *testing.T) {
	t.Skip("TODO(mwhittaker): Make this test pass.")

	d := deploy(t, context.Background(), colocated)
	defer d.shutdown()

	// Return an error from ExportListener a number of times.
	const n = 3
	failures := map[string]int{}
	d.exportListener = func(ctx context.Context, req *protos.ExportListenerRequest) (*protos.ExportListenerReply, error) {
		if failures[req.Listener] < n {
			failures[req.Listener]++
			return &protos.ExportListenerReply{Error: fmt.Sprintf("simulated ExportListener(%q) error", req.Listener)}, nil
		}
		return &protos.ExportListenerReply{}, nil
	}

	testComponents(d)
}

func TestUpdateMissingComponents(t *testing.T) {
	d := deploy(t, context.Background(), colocated)
	defer d.shutdown()

	// Update the weavelet with components that don't exist.
	components := &protos.UpdateComponentsRequest{Components: []string{"foo", "bar"}}
	if _, err := d.weavelets["1"].wlet.UpdateComponents(context.Background(), components); err == nil {
		t.Fatal("unexpected success")
	}

	testComponents(d)
}

func TestUpdateExistingComponents(t *testing.T) {
	d := deploy(t, context.Background(), colocated)
	defer d.shutdown()
	testComponents(d)

	// Update the weavelet with components that have already been started.
	components := &protos.UpdateComponentsRequest{
		Components: []string{componenta, componentb, componentc},
	}
	if _, err := d.weavelets["1"].wlet.UpdateComponents(context.Background(), components); err != nil {
		t.Fatal(err)
	}

	testComponents(d)
}

func TestUpdateNilRoutingInfo(t *testing.T) {
	ctx := context.Background()
	d := deploy(t, ctx, colocated)
	defer d.shutdown()

	// Update the weavelet with a nil routing info.
	routing := &protos.UpdateRoutingInfoRequest{}
	if _, err := d.weavelets["1"].wlet.UpdateRoutingInfo(ctx, routing); err == nil {
		t.Fatal("UpdateRoutingInfo: unexpected success")
	}

	testComponents(d)
}

func TestUpdateRoutingInfoMissingComponent(t *testing.T) {
	ctx := context.Background()
	d := deploy(t, ctx, colocated)
	defer d.shutdown()

	// Update the weavelet with routing info for a component that doesn't
	// exist.
	routing := &protos.UpdateRoutingInfoRequest{
		RoutingInfo: &protos.RoutingInfo{
			Component: "foo",
			Local:     true,
		},
	}
	if _, err := d.weavelets["1"].wlet.UpdateRoutingInfo(ctx, routing); err == nil {
		t.Fatal("UpdateRoutingInfo: unexpected success")
	}

	testComponents(d)
}

func TestUpdateRoutingInfoNotStartedComponent(t *testing.T) {
	ctx := context.Background()
	d := deploy(t, ctx, colocated)
	defer d.shutdown()

	// Update the weavelet with routing info for a component that has hasn't
	// started yet.
	routing := &protos.UpdateRoutingInfoRequest{
		RoutingInfo: &protos.RoutingInfo{
			Component: componenta,
			Local:     true,
		},
	}
	if _, err := d.weavelets["1"].wlet.UpdateRoutingInfo(ctx, routing); err != nil {
		t.Fatal(err)
	}
	testComponents(d)
}

func TestUpdateLocalRoutingInfoWithNonLocal(t *testing.T) {
	ctx := context.Background()
	d := deploy(t, ctx, colocated)
	defer d.shutdown()
	testComponents(d)

	// Update the weavelet with non-local routing info for a component, even
	// though the component has already started with local routing info. Today,
	// that is not allowed and should fail.
	routing := &protos.UpdateRoutingInfoRequest{
		RoutingInfo: &protos.RoutingInfo{
			Component: componenta,
		},
	}
	if _, err := d.weavelets["1"].wlet.UpdateRoutingInfo(ctx, routing); err == nil {
		t.Fatal("UpdateRoutingInfo: unexpected success")
	}
	testComponents(d)
}

func TestFailReplica(t *testing.T) {
	ctx := context.Background()
	placement := map[string][]string{
		"1": {componenta},
		"2": {componentb, componentc},
		"3": {componentb, componentc},
	}
	d := deploy(t, ctx, placement)
	defer d.shutdown()
	testComponents(d)

	// Kill replica 3.
	d.weavelets["3"].cancel()
	if err := d.weavelets["3"].threads.Wait(); err != nil {
		t.Fatal(err)
	}

	// Update routing info to exclude replica 3.
	for _, wlet := range []string{"1", "2"} {
		for _, component := range []string{componentb, componentc} {
			routing := &protos.UpdateRoutingInfoRequest{
				RoutingInfo: &protos.RoutingInfo{
					Component: component,
					Replicas:  []string{d.weavelets["2"].env.WeaveletAddress()},
				},
			}
			if _, err := d.weavelets[wlet].wlet.UpdateRoutingInfo(ctx, routing); err != nil {
				t.Fatal(err)
			}
		}
	}

	testComponents(d)
}

func TestUpdateBadRoutingInfo(t *testing.T) {
	ctx := context.Background()
	placement := map[string][]string{
		"1": {componenta},
		"2": {componentb},
		"3": {componentc},
	}
	d := deploy(t, ctx, placement)
	defer d.shutdown()

	// When activating components, provide incorrect routing info. Later, send
	// the correct routing info. Incorrect routing info should stall, but not
	// crash the weavelets.
	var mu sync.Mutex
	var activateErr error
	d.activateComponent = func(ctx context.Context, req *protos.ActivateComponentRequest) (*protos.ActivateComponentReply, error) {
		// Update the component.
		components := &protos.UpdateComponentsRequest{Components: []string{req.Component}}
		weavelets := map[string]*weavelet{
			componenta: d.weavelets["1"],
			componentb: d.weavelets["2"],
			componentc: d.weavelets["3"],
		}
		weavelet := weavelets[req.Component]
		if _, err := weavelet.wlet.UpdateComponents(ctx, components); err != nil {
			return nil, err
		}

		// Provide incorrect routing info.
		routing := &protos.UpdateRoutingInfoRequest{
			RoutingInfo: &protos.RoutingInfo{
				Component: req.Component,
				Replicas:  []string{"tcp://1.1.1.1:9999"},
			},
		}
		for _, weavelet := range d.weavelets {
			if _, err := weavelet.wlet.UpdateRoutingInfo(ctx, routing); err != nil {
				return nil, err
			}
		}

		// Update to the correct routing info after a short delay.
		routing.RoutingInfo.Replicas = []string{weavelet.env.WeaveletAddress()}
		go func() {
			time.Sleep(100 * time.Millisecond)
			for _, weavelet := range d.weavelets {
				if _, err := weavelet.wlet.UpdateRoutingInfo(ctx, routing); err != nil {
					mu.Lock()
					activateErr = err
					mu.Unlock()
				}
			}
		}()

		return &protos.ActivateComponentReply{}, nil
	}

	testComponents(d)
	d.shutdown()

	if activateErr != nil {
		t.Fatal(activateErr)
	}
}

func TestProfile(t *testing.T) {
	// Ensure things are started.
	ctx := context.Background()
	placement := map[string][]string{
		"1": {componenta},
		"2": {componentb, componentc},
		"3": {componentb, componentc},
	}
	d := deploy(t, ctx, placement)
	defer d.shutdown()
	testComponents(d)

	target := d.weavelets["2"]
	for _, typ := range []protos.ProfileType{protos.ProfileType_Heap, protos.ProfileType_CPU} {
		typ := typ
		t.Run(typ.String(), func(t *testing.T) {
			// Send a profiling request and wait for a reply.
			start := time.Now()
			req := &protos.GetProfileRequest{
				ProfileType:   typ,
				CpuDurationNs: int64(100 * time.Millisecond / time.Nanosecond),
			}
			reply, err := target.wlet.GetProfile(ctx, req)
			if err != nil {
				t.Fatal(err)
			}
			end := time.Now()

			// Small sanity check of the profile.
			p, err := profile.ParseData(reply.Data)
			if err != nil {
				t.Fatal(err)
			}
			when := time.Unix(0, p.TimeNanos)
			if when.Before(start) || end.Before(when) {
				t.Errorf("profile timestamp %v is not in profiling time range [%v,%v]", when, start, end)
			}
		})
	}
}

func TestMetrics(t *testing.T) {
	// Ensure a component is started.
	ctx := context.Background()
	const group = "1"
	placement := map[string][]string{group: {componentc}}
	d := deploy(t, ctx, placement)
	defer d.shutdown()
	target := d.weavelets[group]

	// Helper that calls a component method.
	call := func() {
		t.Helper()
		comp, err := target.wlet.GetIntf(reflection.Type[c]())
		if err != nil {
			t.Fatal(err)
		}
		const want = 42
		got, err := comp.(c).C(d.ctx, want)
		if err != nil {
			t.Fatal(err)
		}
		if got != want {
			t.Fatalf("C(%d): got %d, want %d", want, got, want)
		}
	}

	// Helper for maintaining fetched metrics.
	names := map[uint64]string{}    // map from metric id to name
	metrics := map[string]float64{} // map from metric name to value
	readMetrics := func() {
		t.Helper()
		m, err := target.wlet.GetMetrics(d.ctx, &protos.GetMetricsRequest{})
		if err != nil {
			t.Fatal("GetMetrics failed", err)
		}
		for _, def := range m.Update.Defs {
			names[def.Id] = def.Name
		}
		for _, val := range m.Update.Values {
			metrics[names[val.Id]] = val.Value
		}
	}
	readMetrics()
	start := int(metrics["c_calls"])

	for i := 0; i < 5; i++ {
		readMetrics()
		if m := int(metrics["c_calls"]); m != start+i {
			t.Fatalf("c_calls = %d, expecting %d", m, start+i)
		}
		call()
	}
}
