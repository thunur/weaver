// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package testdeployer

import (
	"context"
	"errors"
	"github.com/thunur/weaver"
	"github.com/thunur/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:      "github.com/thunur/weaver/internal/testdeployer/a",
		Iface:     reflect.TypeOf((*a)(nil)).Elem(),
		Impl:      reflect.TypeOf(aimpl{}),
		Listeners: []string{"lis"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return a_local_stub{impl: impl.(a), tracer: tracer, aMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/internal/testdeployer/a", Method: "A", Remote: false, Generated: true})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return a_client_stub{stub: stub, aMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/internal/testdeployer/a", Method: "A", Remote: true, Generated: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return a_server_stub{impl: impl.(a), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return a_reflect_stub{caller: caller}
		},
		RefData: "⟦d473cf51:wEaVeReDgE:github.com/thunur/weaver/internal/testdeployer/a→github.com/thunur/weaver/internal/testdeployer/b⟧\n⟦83f71f4e:wEaVeRlIsTeNeRs:github.com/thunur/weaver/internal/testdeployer/a→lis⟧\n",
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/thunur/weaver/internal/testdeployer/b",
		Iface: reflect.TypeOf((*b)(nil)).Elem(),
		Impl:  reflect.TypeOf(bimpl{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return b_local_stub{impl: impl.(b), tracer: tracer, bMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/internal/testdeployer/b", Method: "B", Remote: false, Generated: true})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return b_client_stub{stub: stub, bMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/internal/testdeployer/b", Method: "B", Remote: true, Generated: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return b_server_stub{impl: impl.(b), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return b_reflect_stub{caller: caller}
		},
		RefData: "⟦54fc5958:wEaVeReDgE:github.com/thunur/weaver/internal/testdeployer/b→github.com/thunur/weaver/internal/testdeployer/c⟧\n",
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/thunur/weaver/internal/testdeployer/c",
		Iface: reflect.TypeOf((*c)(nil)).Elem(),
		Impl:  reflect.TypeOf(cimpl{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return c_local_stub{impl: impl.(c), tracer: tracer, cMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/internal/testdeployer/c", Method: "C", Remote: false, Generated: true})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return c_client_stub{stub: stub, cMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/internal/testdeployer/c", Method: "C", Remote: true, Generated: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return c_server_stub{impl: impl.(c), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return c_reflect_stub{caller: caller}
		},
		RefData: "",
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/thunur/weaver/internal/testdeployer/d",
		Iface: reflect.TypeOf((*d)(nil)).Elem(),
		Impl:  reflect.TypeOf(dimpl{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return d_local_stub{impl: impl.(d), tracer: tracer, dMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/internal/testdeployer/d", Method: "D", Remote: false, Generated: true})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return d_client_stub{stub: stub, dMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/internal/testdeployer/d", Method: "D", Remote: true, Generated: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return d_server_stub{impl: impl.(d), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return d_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[a] = (*aimpl)(nil)
var _ weaver.InstanceOf[b] = (*bimpl)(nil)
var _ weaver.InstanceOf[c] = (*cimpl)(nil)
var _ weaver.InstanceOf[d] = (*dimpl)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*aimpl)(nil)
var _ weaver.Unrouted = (*bimpl)(nil)
var _ weaver.Unrouted = (*cimpl)(nil)
var _ weaver.Unrouted = (*dimpl)(nil)

// Local stub implementations.

type a_local_stub struct {
	impl     a
	tracer   trace.Tracer
	aMetrics *codegen.MethodMetrics
}

// Check that a_local_stub implements the a interface.
var _ a = (*a_local_stub)(nil)

func (s a_local_stub) A(ctx context.Context, a0 int) (r0 int, err error) {
	// Update metrics.
	begin := s.aMetrics.Begin()
	defer func() { s.aMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "testdeployer.a.A", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.A(ctx, a0)
}

type b_local_stub struct {
	impl     b
	tracer   trace.Tracer
	bMetrics *codegen.MethodMetrics
}

// Check that b_local_stub implements the b interface.
var _ b = (*b_local_stub)(nil)

func (s b_local_stub) B(ctx context.Context, a0 int) (r0 int, err error) {
	// Update metrics.
	begin := s.bMetrics.Begin()
	defer func() { s.bMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "testdeployer.b.B", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.B(ctx, a0)
}

type c_local_stub struct {
	impl     c
	tracer   trace.Tracer
	cMetrics *codegen.MethodMetrics
}

// Check that c_local_stub implements the c interface.
var _ c = (*c_local_stub)(nil)

func (s c_local_stub) C(ctx context.Context, a0 int) (r0 int, err error) {
	// Update metrics.
	begin := s.cMetrics.Begin()
	defer func() { s.cMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "testdeployer.c.C", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.C(ctx, a0)
}

type d_local_stub struct {
	impl     d
	tracer   trace.Tracer
	dMetrics *codegen.MethodMetrics
}

// Check that d_local_stub implements the d interface.
var _ d = (*d_local_stub)(nil)

func (s d_local_stub) D(ctx context.Context) (r0 string, err error) {
	// Update metrics.
	begin := s.dMetrics.Begin()
	defer func() { s.dMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "testdeployer.d.D", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.D(ctx)
}

// Client stub implementations.

type a_client_stub struct {
	stub     codegen.Stub
	aMetrics *codegen.MethodMetrics
}

// Check that a_client_stub implements the a interface.
var _ a = (*a_client_stub)(nil)

func (s a_client_stub) A(ctx context.Context, a0 int) (r0 int, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.aMetrics.Begin()
	defer func() { s.aMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "testdeployer.a.A", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Int(a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Int()
	err = dec.Error()
	return
}

type b_client_stub struct {
	stub     codegen.Stub
	bMetrics *codegen.MethodMetrics
}

// Check that b_client_stub implements the b interface.
var _ b = (*b_client_stub)(nil)

func (s b_client_stub) B(ctx context.Context, a0 int) (r0 int, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.bMetrics.Begin()
	defer func() { s.bMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "testdeployer.b.B", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Int(a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Int()
	err = dec.Error()
	return
}

type c_client_stub struct {
	stub     codegen.Stub
	cMetrics *codegen.MethodMetrics
}

// Check that c_client_stub implements the c interface.
var _ c = (*c_client_stub)(nil)

func (s c_client_stub) C(ctx context.Context, a0 int) (r0 int, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.cMetrics.Begin()
	defer func() { s.cMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "testdeployer.c.C", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Int(a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Int()
	err = dec.Error()
	return
}

type d_client_stub struct {
	stub     codegen.Stub
	dMetrics *codegen.MethodMetrics
}

// Check that d_client_stub implements the d interface.
var _ d = (*d_client_stub)(nil)

func (s d_client_stub) D(ctx context.Context) (r0 string, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.dMetrics.Begin()
	defer func() { s.dMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "testdeployer.d.D", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	var shardKey uint64

	// Call the remote method.
	var results []byte
	results, err = s.stub.Run(ctx, 0, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.String()
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][24]struct{}](`

ERROR: You generated this file with 'weaver generate' (devel) (codegen
version v0.24.0). The generated code is incompatible with the version of the
github.com/thunur/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/thunur/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/thunur/weaver@latest
    go install github.com/thunur/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/thunur/weaver/issues.

`)

// Server stub implementations.

type a_server_stub struct {
	impl    a
	addLoad func(key uint64, load float64)
}

// Check that a_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*a_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s a_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "A":
		return s.a
	default:
		return nil
	}
}

func (s a_server_stub) a(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 int
	a0 = dec.Int()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.A(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Int(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type b_server_stub struct {
	impl    b
	addLoad func(key uint64, load float64)
}

// Check that b_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*b_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s b_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "B":
		return s.b
	default:
		return nil
	}
}

func (s b_server_stub) b(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 int
	a0 = dec.Int()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.B(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Int(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type c_server_stub struct {
	impl    c
	addLoad func(key uint64, load float64)
}

// Check that c_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*c_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s c_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "C":
		return s.c
	default:
		return nil
	}
}

func (s c_server_stub) c(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 int
	a0 = dec.Int()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.C(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Int(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type d_server_stub struct {
	impl    d
	addLoad func(key uint64, load float64)
}

// Check that d_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*d_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s d_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "D":
		return s.d
	default:
		return nil
	}
}

func (s d_server_stub) d(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.D(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.String(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type a_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that a_reflect_stub implements the a interface.
var _ a = (*a_reflect_stub)(nil)

func (s a_reflect_stub) A(ctx context.Context, a0 int) (r0 int, err error) {
	err = s.caller("A", ctx, []any{a0}, []any{&r0})
	return
}

type b_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that b_reflect_stub implements the b interface.
var _ b = (*b_reflect_stub)(nil)

func (s b_reflect_stub) B(ctx context.Context, a0 int) (r0 int, err error) {
	err = s.caller("B", ctx, []any{a0}, []any{&r0})
	return
}

type c_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that c_reflect_stub implements the c interface.
var _ c = (*c_reflect_stub)(nil)

func (s c_reflect_stub) C(ctx context.Context, a0 int) (r0 int, err error) {
	err = s.caller("C", ctx, []any{a0}, []any{&r0})
	return
}

type d_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that d_reflect_stub implements the d interface.
var _ d = (*d_reflect_stub)(nil)

func (s d_reflect_stub) D(ctx context.Context) (r0 string, err error) {
	err = s.caller("D", ctx, []any{}, []any{&r0})
	return
}
