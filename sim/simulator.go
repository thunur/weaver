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

// Package sim implements deterministic simulation.
//
// [Deterministic simulation][1] is a type of randomized testing in which
// millions of random operations are run against a system (with randomly
// injected failures) in an attempt to find bugs. See
// serviceweaver.dev/blog/testing.html for an overview of determistic
// simulation and its implementation in the sim package.
//
// # Generators
//
// A key component of deterministic simulation is the ability to
// deterministically generate "random" values. We accomplish this with the
// [Generator] interface:
//
//	type Generator[T any] interface {
//	    Generate(*rand.Rand) T
//	}
//
// A Generator[T] generates random values of type T. For example, the [Int]
// function returns a Generator[int] that generates random integers.
//
// While random, a Generator is also deterministic. Given a random number
// generator with a particular seed, a Generator will always produce the same
// value:
//
//	// x and y are always equal.
//	var gen Gen[int] = ...
//	x := gen.Generate(rand.New(rand.NewSource(42)))
//	y := gen.Generate(rand.New(rand.NewSource(42)))
//
// The sim package includes generators that generate booleans, ints,
// floats, runes, strings, slices, and maps (e.g., [Flip], [Int], [Float64],
// [Rune], [String], [Range], [Map]). It also contains generator combinators
// that combine existing generators into new generators (e.g., [OneOf],
// [Weight], [Filter]). You can also implement your own custom generators by
// implementing the Generator interface.
//
// # Workloads
//
// Deterministic simulation verifies a system by running random operations
// against the system, checking for invariant violations along the way. A
// workload defines the set of operations to run and the set of invariants to
// check.
//
// Concretely, a workload is a struct that implements the [Workload] interface.
// When a simulator executes a workload, it will randomly call the exported
// methods of the struct with randomly generated values. We call these methods
// *operations*. If an operation ever encounters an invariant violation, it
// returns a non-nil error and the simulation is aborted.
//
// Consider the following workload as an example.
//
//	func even(x int) bool {
//		return x%2 == 0
//	}
//
//	type EvenWorkload struct {
//		x int
//	}
//
//	func (e *EvenWorkload) Add(_ context.Context, y int) error {
//		e.x = e.x + y
//		if !even(e.x) {
//			return fmt.Errorf("%d is not even", e.x)
//		}
//		return nil
//	}
//
//	func (e *EvenWorkload) Multiply(_ context.Context, y int) error {
//		e.x = e.x * y
//		if !even(e.x) {
//			return fmt.Errorf("%d is not even", e.x)
//		}
//		return nil
//	}
//
// An EvenWorkload has an integer x as state and defines two operations: Add
// and Multiply. Add adds a value to x, and Multiply multiplies x. Both
// operations check the invariant that x is even. Of course, this invariant
// does not hold if we add arbitrary values to x.
//
// However, we control the arguments on which which operations are called.
// Specifically, we add an Init method that registers a set of generators. The
// simulator will call the workload's operations on values produced by these
// generators.
//
//	func (e *EvenWorkload) Init(r sim.Registrar) error {
//		r.RegisterGenerators("Add", sim.Filter(sim.Int(), even))
//		r.RegisterGenerators("Multiply", sim.Int())
//		return nil
//	}
//
// Note that we only call the Add operation on even integers. Finally, we can
// construct a simulator and simulate the EvenWorkload.
//
//	func TestEvenWorkload(t *testing.T) {
//		s := sim.New(t, &EvenWorkload{}, sim.Options{})
//		r := s.Run(5 * time.Second)
//		if r.Err != nil {
//			t.Fatal(r.Err)
//		}
//	}
//
// In this trivial example, our workload did not use any Service Weaver
// components, but most realistic workloads do. A workload can get a reference
// to a component using weaver.Ref. See serviceweaver.dev/blog/testing.html for
// a complete example.
//
// # Graveyard
//
// When the simulator runs a failed execution, it persists the failing inputs
// to disk. The collection of saved failing inputs is called a *graveyard*, and
// each individual entry is called a *graveyard entry*. When a simulator is
// created, the first thing it does is load and re-simulate all graveyard
// entries.
//
// We borrow the design of go's fuzzing library's corpus with only minor
// changes [2]. When a simulator runs as part of a test named TestFoo, it
// stores its graveyard entries in testdata/sim/TestFoo. Every graveyard entry
// is a JSON file. Filenames are derived from the hash of the contents of the
// graveyard entry. Here's an example testdata directory:
//
//	testdata/
//	└── sim
//	    ├── TestCancelledSimulation
//	    │   └── a52f5ec5f94e674d.json
//	    ├── TestSimulateGraveyardEntries
//	    │   ├── 2bfe847328319dae.json
//	    │   └── a52f5ec5f94e674d.json
//	    └── TestUnsuccessfulSimulation
//	        ├── 2bfe847328319dae.json
//	        └── a52f5ec5f94e674d.json
//
// As with go's fuzzing library, graveyard entries are never garbage collected.
// Users are responsible for manually deleting graveyard entries when
// appropriate.
//
// TODO(mwhittaker): Move things to the weavertest package.
//
// [1]: https://asatarin.github.io/testing-distributed-systems/#deterministic-simulation
// [2]: https://go.dev/security/fuzz
package sim

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/thunur/weaver/internal/reflection"
	"github.com/thunur/weaver/internal/weaver"
	swruntime "github.com/thunur/weaver/runtime"
	"github.com/thunur/weaver/runtime/codegen"
	"github.com/thunur/weaver/runtime/logging"
	"github.com/thunur/weaver/runtime/protos"
	"golang.org/x/exp/maps"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// FakeComponent is a copy of weavertest.FakeComponent. It's needed to access
// the unexported fields.
//
// TODO(mwhittaker): Remove this once we merge with weavertest.
type FakeComponent struct {
	intf reflect.Type
	impl any
}

// Fake is a copy of weavertest.Fake.
//
// TODO(mwhittaker): Remove this once we merge with the weavertest package.
func Fake[T any](impl any) FakeComponent {
	t := reflection.Type[T]()
	if _, ok := impl.(T); !ok {
		panic(fmt.Sprintf("%T does not implement %v", impl, t))
	}
	return FakeComponent{intf: t, impl: impl}
}

// A Generator[T] generates random values of type T.
type Generator[T any] interface {
	// Generate returns a randomly generated value of type T. While Generate is
	// "random", it must be deterministic. That is, given the same instance of
	// *rand.Rand, Generate must always return the same value.
	//
	// TODO(mwhittaker): Generate should maybe take something other than a
	// *rand.Rand?
	Generate(*rand.Rand) T
}

// A Registrar is used to register fakes and generators with a [Simulator].
type Registrar interface {
	// RegisterFake registers a fake implementation of a component.
	RegisterFake(FakeComponent)

	// RegisterGenerators registers generators for a workload method, one
	// generator per method argument. The number and type of the registered
	// generators must match the method. For example, given the method:
	//
	//     Foo(context.Context, int, bool) error
	//
	// we must register a Generator[int] and a Generator[bool]:
	//
	//     var r Registrar = ...
	//     var i Generator[int] = ...
	//     var b Generator[bool] = ...
	//     r.RegisterGenerators("Foo", i, b)
	//
	// TODO(mwhittaker): Allow people to register a func(*rand.Rand) T instead
	// of a Generator[T] for convenience.
	RegisterGenerators(method string, generators ...any)
}

// A Workload defines the set of operations to run as part of a simulation.
// Every workload is defined as a named struct. To execute a workload, a
// simulator constructs an instance of the struct, calls the struct's Init
// method, and then randomly calls the struct's exported methods. For example,
// the following is a simple workload:
//
//	type myWorkload struct {}
//	func (w *myWorkload) Init(r sim.Registrar) {...}
//	func (w *myWorkload) Foo(context.Context, int) error {...}
//	func (w *myWorkload) Bar(context.Context, bool, string) error {...}
//	func (w *myWorkload) baz(context.Context) error {...}
//
// When this workload is executed, its Foo and Bar methods will be called with
// random values generated by the generators registered in the Init method (see
// [Registrar] for details). Note that unexported methods, like baz, are
// ignored.
//
// Note that every exported workload method must receive a [context.Context] as
// its first argument and must return a single error value. A simulation is
// aborted when a method returns a non-nil error.
//
// TODO(mwhittaker): For now, the Init method is required. In the future, we
// could make it optional and use default generators for methods.
type Workload interface {
	// Init initializes a workload. The Init method must also register
	// generators for every exported method.
	Init(Registrar) error
}

// Options configure a Simulator.
type Options struct {
	// TOML config file contents.
	Config string

	// The number of executions to run in parallel. If Parallelism is 0, the
	// simulator picks the degree of parallelism.
	Parallelism int
}

// A Simulator deterministically simulates a Service Weaver application. See
// the package documentation for instructions on how to use a Simulator.
type Simulator struct {
	opts       Options                                // options
	t          testing.TB                             // underlying test
	w          reflect.Type                           // workload type
	regsByIntf map[reflect.Type]*codegen.Registration // components, by interface
	info       componentInfo                          // component metadata
	config     *protos.AppConfig                      // application config
}

// Results are the results of simulating a workload.
type Results struct {
	Err           error         // first non-nil error returned by an op
	History       []Event       // a history of the error inducing run, if Err is not nil
	NumExecutions int           // number of executions ran
	NumOps        int           // number of ops ran
	Duration      time.Duration // duration of simulation
}

// New returns a new Simulator that simulates the provided workload.
func New(t testing.TB, x Workload, opts Options) *Simulator {
	t.Helper()

	// Parse config.
	app := &protos.AppConfig{}
	if opts.Config != "" {
		var err error
		app, err = swruntime.ParseConfig("", opts.Config, codegen.ComponentConfigValidator)
		if err != nil {
			t.Fatalf("sim.New: parse config: %v", err)
		}
	}

	// Methods can have either value or pointer receivers. For example,
	// consider the following code:
	//
	//     type t struct{}
	//     func (t) ValueReceiver() {}
	//     func (*t) PointerReceiver() {}
	//
	// According to the Go spec, the method set of t includes only
	// ValueReceiver, while the method set of *t includes ValueReceiver and
	// PointerReceiver [1]. We want to call *every* exported method on a
	// workload struct, so we need to massage the type of x into a pointer if
	// it isn't already.
	//
	// [1]: https://go.dev/ref/spec#Method_sets
	w := reflect.TypeOf(x)
	if w.Kind() != reflect.Ptr {
		w = reflect.PointerTo(w)
	}

	// Validate the workload struct.
	if err := validateWorkload(w); err != nil {
		t.Fatalf("sim.New: invalid workload type %v: %v", w, err)
	}

	// Gather the set of registered components.
	//
	// TODO(mwhittaker): Only use the components actually referenced by the
	// workload.
	registered := map[reflect.Type]struct{}{}
	regsByIntf := map[reflect.Type]*codegen.Registration{}
	info := componentInfo{
		hasRefs:      map[reflect.Type]bool{},
		hasListeners: map[reflect.Type]bool{},
		hasConfig:    map[reflect.Type]bool{},
	}
	for _, reg := range codegen.Registered() {
		x := reflect.New(reg.Impl).Interface()
		registered[reg.Iface] = struct{}{}
		regsByIntf[reg.Iface] = reg
		info.hasRefs[reg.Iface] = weaver.HasRefs(x)
		info.hasListeners[reg.Iface] = weaver.HasListeners(x)
		info.hasConfig[reg.Iface] = weaver.HasConfig(x)
	}

	// Call Init and validate the registered fakes and generators.
	r := newRegistrar(t, w, registered)
	if err := x.Init(r); err != nil {
		t.Fatalf("sim.New: %v", err)
	}
	if err := r.finalize(); err != nil {
		t.Fatalf("sim.New: %v", err)
	}

	return &Simulator{opts, t, w, regsByIntf, info, app}
}

// validateWorkload validates a workload struct of the provided type.
func validateWorkload(w reflect.Type) error {
	var errs []error
	numOps := 0
	for i := 0; i < w.NumMethod(); i++ {
		m := w.Method(i)
		if m.Name == "Init" {
			continue
		}
		numOps++

		// Method should have type func(context.Context, ...) error.
		err := fmt.Errorf("method %s has type '%v' but should have type 'func(%v, context.Context, ...) error'", m.Name, m.Type, w)
		switch {
		case m.Type.NumIn() < 2:
			errs = append(errs, fmt.Errorf("%w: no arguments", err))
		case m.Type.In(1) != reflection.Type[context.Context]():
			errs = append(errs, fmt.Errorf("%w: first argument is not context.Context", err))
		case m.Type.NumOut() == 0:
			errs = append(errs, fmt.Errorf("%w: no return value", err))
		case m.Type.NumOut() > 1:
			errs = append(errs, fmt.Errorf("%w: too many return values", err))
		case m.Type.Out(0) != reflection.Type[error]():
			errs = append(errs, fmt.Errorf("%w: return value is not error", err))
		}
	}
	if numOps == 0 {
		errs = append(errs, fmt.Errorf("no exported methods"))
	}
	return errors.Join(errs...)
}

// newExecutor returns a new executor.
func (s *Simulator) newExecutor() *executor {
	return newExecutor(s.t, s.w, s.regsByIntf, s.info, s.config)
}

// graveyardDir returns the graveyard directory for this simulator.
func (s *Simulator) graveyardDir() string {
	// Test names often contain slashes ("/"). We replace "/" with "#" to
	// safely use the test name as a directory name.
	//
	// TODO(mwhittaker): This mapping is sensitive to collisions. A test named
	// "foo/bar" collides with a test named "foo#bar". I think in practice,
	// this will likely not be a big issue. But, if people are running into
	// problems, we can use a more complex collision resistant sanitization.
	sanitized := strings.ReplaceAll(s.t.Name(), "/", "#")
	return filepath.Join("testdata", "sim", sanitized)
}

// Run runs a simulation for the provided duration.
func (s *Simulator) Run(duration time.Duration) Results {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	s.t.Logf("Simulating workload %v for %v.", s.w, duration)
	stats := &stats{start: time.Now()}
	switch result, err := s.run(ctx, stats); {
	case err != nil && err == ctx.Err():
		// The simulation was cancelled.
		results := Results{
			NumExecutions: int(stats.numExecutions),
			NumOps:        int(stats.numOps),
			Duration:      time.Since(stats.start),
		}
		s.t.Log(results.summary())
		return results

	case err != nil:
		// The simulation failed to run properly.
		s.t.Fatalf("Simulator.Run: %v", err)
		return Results{}

	case result.err != nil:
		// The simulation found a failing execution.
		results := Results{
			Err:           result.err,
			History:       result.history,
			NumExecutions: int(stats.numExecutions),
			NumOps:        int(stats.numOps),
			Duration:      time.Since(stats.start),
		}
		s.t.Log(results.summary())

		entry := graveyardEntry{
			Version:     version,
			Seed:        result.params.Seed,
			NumReplicas: result.params.NumReplicas,
			NumOps:      result.params.NumOps,
			FailureRate: result.params.FailureRate,
			YieldRate:   result.params.YieldRate,
		}
		if filename, err := writeGraveyardEntry(s.graveyardDir(), entry); err == nil {
			s.t.Logf("Failing input written to %s.", filename)
		}
		return results

	default:
		// All executions passed.
		results := Results{
			NumExecutions: int(stats.numExecutions),
			NumOps:        int(stats.numOps),
			Duration:      time.Since(stats.start),
		}
		s.t.Log(results.summary())
		return results
	}
}

// stats contains simulation statistics.
type stats struct {
	start         time.Time // start of simulation
	numExecutions int64     // number of fully executed executions
	numOps        int64     // number of fully executed ops
}

// run runs a simulation until the provided context is cancelled. It returns
// the hyperparameters and result of a failing execution if any are found.
func (s *Simulator) run(ctx context.Context, stats *stats) (result, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Spawn a goroutine to periodically print progress.
	done := sync.WaitGroup{}
	done.Add(1)
	go func() {
		defer done.Done()
		s.printProgress(ctx, stats)
	}()

	// Execute the graveyard entries.
	if r, err := s.executeGraveyard(ctx, stats); err != nil || r.err != nil {
		return r, err
	}

	// Spawn n concurrent executors which read hyperparamters from the params
	// channel. Simulation ends when:
	//
	//     1. the context is cancelled;
	//     2. an execution fails to run properly (written to errs); or
	//     3. a failing execution is found (written to failing).
	n := s.opts.Parallelism
	if n == 0 {
		n = 10 * runtime.NumCPU()
	}
	params := make(chan hyperparameters, n)
	errs := make(chan error, n)
	failing := make(chan result, n)

	s.t.Logf("Executing with %d executors.", n)
	done.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer done.Done()
			switch r, err := s.execute(ctx, stats, params); {
			case err != nil && err == ctx.Err():
				return
			case err != nil:
				errs <- err
				return
			case r.err != nil:
				failing <- r
			}
		}()
	}

	// Spawn a goroutine that writes to the params channel.
	//
	// TODO(mwhittaker): Use a smarter algorithm to sweep over hyperparameters.
	done.Add(1)
	go func() {
		defer done.Done()
		seed := time.Now().UnixNano()
		for numOps := 1; ; numOps++ {
			for _, numReplicas := range []int{1, 2, 3} {
				for _, failureRate := range []float64{0.0, 0.01, 0.05, 0.1} {
					for _, yieldRate := range []float64{0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0} {
						for i := 0; i < 1000; i++ {
							seed++
							p := hyperparameters{
								Seed:        seed,
								NumOps:      numOps,
								NumReplicas: numReplicas,
								FailureRate: failureRate,
								YieldRate:   yieldRate,
							}
							select {
							case <-ctx.Done():
								return
							case params <- p:
							}
						}
					}
				}
			}
		}
	}()

	// Wait for the simulation to end.
	select {
	case <-ctx.Done():
		done.Wait()
		return result{}, ctx.Err()
	case err := <-errs:
		cancel()
		done.Wait()
		return result{}, err
	case r := <-failing:
		cancel()
		done.Wait()
		return r, nil
	}
}

// printProgress periodically prints the progress of the simulation.
func (s *Simulator) printProgress(ctx context.Context, stats *stats) {
	printer := message.NewPrinter(language.AmericanEnglish)
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			elapsed := time.Since(stats.start)
			truncated := elapsed.Truncate(time.Second)
			execs := atomic.LoadInt64(&stats.numExecutions)
			ops := atomic.LoadInt64(&stats.numOps)
			execRate := printer.Sprintf("%0.0f", float64(execs)/elapsed.Seconds())
			opRate := printer.Sprintf("%0.0f", float64(ops)/elapsed.Seconds())
			s.t.Logf("[%v] %s execs (%s execs/s), %s ops (%s ops/s)", truncated, printer.Sprint(execs), execRate, printer.Sprint(ops), opRate)
		}
	}
}

// executeGraveyardEntries executes graveyard entries serially. Executing
// graveyard entries serially is important to make errors repeatable. If we
// execute graveyard entries in multiple goroutines, the user might see a
// different error every time they run "go test", which is discombobulating.
func (s *Simulator) executeGraveyard(ctx context.Context, stats *stats) (result, error) {
	graveyard, err := readGraveyard(s.graveyardDir())
	if err != nil {
		return result{}, err
	}

	s.t.Logf("Executing %d graveyard entries.", len(graveyard))
	exec := s.newExecutor()
	for _, entry := range graveyard {
		p := hyperparameters{
			Seed:        entry.Seed,
			NumReplicas: entry.NumReplicas,
			NumOps:      entry.NumOps,
			FailureRate: entry.FailureRate,
			YieldRate:   entry.YieldRate,
		}
		r, err := exec.execute(ctx, p)
		if err != nil {
			return result{}, err
		}
		atomic.AddInt64(&stats.numExecutions, 1)
		atomic.AddInt64(&stats.numOps, int64(p.NumOps))
		if r.err != nil {
			return r, nil
		}
	}
	s.t.Log("Done executing graveyard entries.")
	return result{}, nil
}

// execute repeatedly performs executions until the provided context is
// cancelled or until a failing result is found. Hyperparameters for the
// executions are read from the provided params channel.
func (s *Simulator) execute(ctx context.Context, stats *stats, params <-chan hyperparameters) (result, error) {
	exec := s.newExecutor()
	for {
		select {
		case <-ctx.Done():
			return result{}, ctx.Err()
		case p := <-params:
			r, err := exec.execute(ctx, p)
			if err != nil {
				return result{}, err
			}
			atomic.AddInt64(&stats.numExecutions, 1)
			atomic.AddInt64(&stats.numOps, int64(p.NumOps))
			if r.err != nil {
				return r, nil
			}
		}
	}
}

// summary returns a human readable summary of the results.
func (r *Results) summary() string {
	duration := r.Duration.Truncate(time.Millisecond)
	printer := message.NewPrinter(language.AmericanEnglish)
	execRate := printer.Sprintf("%0.2f", float64(r.NumExecutions)/r.Duration.Seconds())
	opRate := printer.Sprintf("%0.2f", float64(r.NumOps)/r.Duration.Seconds())
	prefix := "No errors"
	if r.Err != nil {
		prefix = "Error"
	}
	return fmt.Sprintf("%s found after %s ops across %s executions in %v (%s execs/s, %s ops/s).",
		prefix, printer.Sprint(r.NumOps), printer.Sprint(r.NumExecutions), duration, execRate, opRate)
}

// Mermaid returns a [mermaid] diagram that illustrates an execution history.
//
// [mermaid]: https://mermaid.js.org/
func (r *Results) Mermaid() string {
	// TODO(mwhittaker): Arrange replicas in topological order.

	// Some abbreviations to save typing.
	shorten := logging.ShortenComponent
	commas := func(xs []string) string { return strings.Join(xs, ", ") }

	// Gather the set of all ops and replicas.
	type replica struct {
		component string
		replica   int
	}
	var ops []EventOpStart
	replicas := map[replica]struct{}{}
	calls := map[int]EventCall{}
	returns := map[int]EventReturn{}
	for _, event := range r.History {
		switch x := event.(type) {
		case EventOpStart:
			ops = append(ops, x)
		case EventCall:
			calls[x.SpanID] = x
		case EventDeliverCall:
			call := calls[x.SpanID]
			replicas[replica{call.Component, x.Replica}] = struct{}{}
		case EventReturn:
			returns[x.SpanID] = x
		}
	}

	// Create the diagram.
	var b strings.Builder
	fmt.Fprintln(&b, "sequenceDiagram")

	// Create ops.
	for _, op := range ops {
		fmt.Fprintf(&b, "    participant op%d as Op %d\n", op.TraceID, op.TraceID)
	}

	// Create component replicas.
	sorted := maps.Keys(replicas)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].component != sorted[j].component {
			return sorted[i].component < sorted[j].component
		}
		return sorted[i].replica < sorted[j].replica
	})
	for _, replica := range sorted {
		fmt.Fprintf(&b, "    participant %s%d as %s %d\n", replica.component, replica.replica, shorten(replica.component), replica.replica)
	}

	// Create events.
	for _, event := range r.History {
		switch x := event.(type) {
		case EventOpStart:
			fmt.Fprintf(&b, "    note right of op%d: [%d:%d] %s(%s)\n", x.TraceID, x.TraceID, x.SpanID, x.Name, commas(x.Args))
		case EventOpFinish:
			fmt.Fprintf(&b, "    note right of op%d: [%d:%d] return %s\n", x.TraceID, x.TraceID, x.SpanID, x.Error)
		case EventDeliverCall:
			call := calls[x.SpanID]
			fmt.Fprintf(&b, "    %s%d->>%s%d: [%d:%d] %s.%s(%s)\n", call.Caller, call.Replica, call.Component, x.Replica, x.TraceID, x.SpanID, shorten(call.Component), call.Method, commas(call.Args))
		case EventDeliverReturn:
			call := calls[x.SpanID]
			ret := returns[x.SpanID]
			fmt.Fprintf(&b, "    %s%d->>%s%d: [%d:%d] return %s\n", ret.Component, ret.Replica, call.Caller, call.Replica, x.TraceID, x.SpanID, commas(ret.Returns))
		case EventDeliverError:
			call := calls[x.SpanID]
			fmt.Fprintf(&b, "    note right of %s%d: [%d:%d] RemoteCallError\n", call.Caller, call.Replica, x.TraceID, x.SpanID)
		case EventPanic:
			stack := strings.ReplaceAll(x.Stack, "\n", "<br>")
			fmt.Fprintf(&b, "    note right of %s%d: [%d:%d] %s<br>%s\n", x.Panicker, x.Replica, x.TraceID, x.SpanID, x.Error, stack)
		}
	}
	return b.String()
}
