// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package generate

import (
	"context"
	"errors"
	"fmt"
	"github.com/thunur/weaver"
	"github.com/thunur/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/thunur/weaver/weavertest/internal/generate/testApp",
		Iface: reflect.TypeOf((*testApp)(nil)).Elem(),
		Impl:  reflect.TypeOf(impl{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return testApp_local_stub{impl: impl.(testApp), tracer: tracer, divModMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/weavertest/internal/generate/testApp", Method: "DivMod", Remote: false, Generated: true}), getMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/weavertest/internal/generate/testApp", Method: "Get", Remote: false, Generated: true}), incPointerMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/weavertest/internal/generate/testApp", Method: "IncPointer", Remote: false, Generated: true})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return testApp_client_stub{stub: stub, divModMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/weavertest/internal/generate/testApp", Method: "DivMod", Remote: true, Generated: true}), getMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/weavertest/internal/generate/testApp", Method: "Get", Remote: true, Generated: true}), incPointerMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/weavertest/internal/generate/testApp", Method: "IncPointer", Remote: true, Generated: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return testApp_server_stub{impl: impl.(testApp), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return testApp_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[testApp] = (*impl)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*impl)(nil)

// Local stub implementations.

type testApp_local_stub struct {
	impl              testApp
	tracer            trace.Tracer
	divModMetrics     *codegen.MethodMetrics
	getMetrics        *codegen.MethodMetrics
	incPointerMetrics *codegen.MethodMetrics
}

// Check that testApp_local_stub implements the testApp interface.
var _ testApp = (*testApp_local_stub)(nil)

func (s testApp_local_stub) DivMod(ctx context.Context, a0 int, a1 int) (r0 int, r1 int, err error) {
	// Update metrics.
	begin := s.divModMetrics.Begin()
	defer func() { s.divModMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "generate.testApp.DivMod", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.DivMod(ctx, a0, a1)
}

func (s testApp_local_stub) Get(ctx context.Context, a0 string, a1 behaviorType) (r0 int, err error) {
	// Update metrics.
	begin := s.getMetrics.Begin()
	defer func() { s.getMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "generate.testApp.Get", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Get(ctx, a0, a1)
}

func (s testApp_local_stub) IncPointer(ctx context.Context, a0 *int) (r0 *int, err error) {
	// Update metrics.
	begin := s.incPointerMetrics.Begin()
	defer func() { s.incPointerMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "generate.testApp.IncPointer", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.IncPointer(ctx, a0)
}

// Client stub implementations.

type testApp_client_stub struct {
	stub              codegen.Stub
	divModMetrics     *codegen.MethodMetrics
	getMetrics        *codegen.MethodMetrics
	incPointerMetrics *codegen.MethodMetrics
}

// Check that testApp_client_stub implements the testApp interface.
var _ testApp = (*testApp_client_stub)(nil)

func (s testApp_client_stub) DivMod(ctx context.Context, a0 int, a1 int) (r0 int, r1 int, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.divModMetrics.Begin()
	defer func() { s.divModMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "generate.testApp.DivMod", trace.WithSpanKind(trace.SpanKindClient))
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
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Int(a0)
	enc.Int(a1)
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
	r1 = dec.Int()
	err = dec.Error()
	return
}

func (s testApp_client_stub) Get(ctx context.Context, a0 string, a1 behaviorType) (r0 int, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.getMetrics.Begin()
	defer func() { s.getMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "generate.testApp.Get", trace.WithSpanKind(trace.SpanKindClient))
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
	size += (4 + len(a0))
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	enc.Int((int)(a1))
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 1, enc.Data(), shardKey)
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

func (s testApp_client_stub) IncPointer(ctx context.Context, a0 *int) (r0 *int, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.incPointerMetrics.Begin()
	defer func() { s.incPointerMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "generate.testApp.IncPointer", trace.WithSpanKind(trace.SpanKindClient))
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
	size += serviceweaver_size_ptr_int_98a2a745(a0)
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	serviceweaver_enc_ptr_int_98a2a745(enc, a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 2, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_ptr_int_98a2a745(dec)
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

type testApp_server_stub struct {
	impl    testApp
	addLoad func(key uint64, load float64)
}

// Check that testApp_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*testApp_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s testApp_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "DivMod":
		return s.divMod
	case "Get":
		return s.get
	case "IncPointer":
		return s.incPointer
	default:
		return nil
	}
}

func (s testApp_server_stub) divMod(ctx context.Context, args []byte) (res []byte, err error) {
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
	var a1 int
	a1 = dec.Int()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, r1, appErr := s.impl.DivMod(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Int(r0)
	enc.Int(r1)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s testApp_server_stub) get(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()
	var a1 behaviorType
	*(*int)(&a1) = dec.Int()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Get(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Int(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s testApp_server_stub) incPointer(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 *int
	a0 = serviceweaver_dec_ptr_int_98a2a745(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.IncPointer(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_ptr_int_98a2a745(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type testApp_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that testApp_reflect_stub implements the testApp interface.
var _ testApp = (*testApp_reflect_stub)(nil)

func (s testApp_reflect_stub) DivMod(ctx context.Context, a0 int, a1 int) (r0 int, r1 int, err error) {
	err = s.caller("DivMod", ctx, []any{a0, a1}, []any{&r0, &r1})
	return
}

func (s testApp_reflect_stub) Get(ctx context.Context, a0 string, a1 behaviorType) (r0 int, err error) {
	err = s.caller("Get", ctx, []any{a0, a1}, []any{&r0})
	return
}

func (s testApp_reflect_stub) IncPointer(ctx context.Context, a0 *int) (r0 *int, err error) {
	err = s.caller("IncPointer", ctx, []any{a0}, []any{&r0})
	return
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*customErrorValue)(nil)

type __is_customErrorValue[T ~struct {
	weaver.AutoMarshal
	key string
}] struct{}

var _ __is_customErrorValue[customErrorValue]

func (x *customErrorValue) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("customErrorValue.WeaverMarshal: nil receiver"))
	}
	enc.String(x.key)
}

func (x *customErrorValue) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("customErrorValue.WeaverUnmarshal: nil receiver"))
	}
	x.key = dec.String()
}
func init() { codegen.RegisterSerializable[*customErrorValue]() }

// Encoding/decoding implementations.

func serviceweaver_enc_ptr_int_98a2a745(enc *codegen.Encoder, arg *int) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		enc.Int(*arg)
	}
}

func serviceweaver_dec_ptr_int_98a2a745(dec *codegen.Decoder) *int {
	if !dec.Bool() {
		return nil
	}
	var res int
	res = dec.Int()
	return &res
}

// Size implementations.

// serviceweaver_size_ptr_int_98a2a745 returns the size (in bytes) of the serialization
// of the provided type.
func serviceweaver_size_ptr_int_98a2a745(x *int) int {
	if x == nil {
		return 1
	} else {
		return 1 + 8
	}
}
