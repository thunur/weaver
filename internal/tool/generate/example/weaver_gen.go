// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package main

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
		Name:      "github.com/thunur/weaver/internal/tool/generate/example/A",
		Iface:     reflect.TypeOf((*A)(nil)).Elem(),
		Impl:      reflect.TypeOf(a{}),
		Routed:    true,
		Listeners: []string{"lis2", "renamed_listener"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return a_local_stub{impl: impl.(A), tracer: tracer, m1Metrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/internal/tool/generate/example/A", Method: "M1", Remote: false, Generated: true}), m2Metrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/internal/tool/generate/example/A", Method: "M2", Remote: false, Generated: true})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return a_client_stub{stub: stub, m1Metrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/internal/tool/generate/example/A", Method: "M1", Remote: true, Generated: true}), m2Metrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/internal/tool/generate/example/A", Method: "M2", Remote: true, Generated: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return a_server_stub{impl: impl.(A), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return a_reflect_stub{caller: caller}
		},
		RefData: "⟦627f661b:wEaVeReDgE:github.com/thunur/weaver/internal/tool/generate/example/A→github.com/thunur/weaver/internal/tool/generate/example/B⟧\n⟦26168bd7:wEaVeRlIsTeNeRs:github.com/thunur/weaver/internal/tool/generate/example/A→lis2,renamed_listener⟧\n",
	})
	codegen.Register(codegen.Registration{
		Name:      "github.com/thunur/weaver/internal/tool/generate/example/B",
		Iface:     reflect.TypeOf((*B)(nil)).Elem(),
		Impl:      reflect.TypeOf(b{}),
		Routed:    true,
		Listeners: []string{"lis2", "renamed_listener"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return b_local_stub{impl: impl.(B), tracer: tracer, m1Metrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/internal/tool/generate/example/B", Method: "M1", Remote: false, Generated: true}), m2Metrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/internal/tool/generate/example/B", Method: "M2", Remote: false, Generated: true})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return b_client_stub{stub: stub, m1Metrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/internal/tool/generate/example/B", Method: "M1", Remote: true, Generated: true}), m2Metrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/internal/tool/generate/example/B", Method: "M2", Remote: true, Generated: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return b_server_stub{impl: impl.(B), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return b_reflect_stub{caller: caller}
		},
		RefData: "⟦6971bce2:wEaVeReDgE:github.com/thunur/weaver/internal/tool/generate/example/B→github.com/thunur/weaver/internal/tool/generate/example/A⟧\n⟦c9c43570:wEaVeRlIsTeNeRs:github.com/thunur/weaver/internal/tool/generate/example/B→lis2,renamed_listener⟧\n",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[A] = (*a)(nil)
var _ weaver.InstanceOf[B] = (*b)(nil)

// weaver.Router checks.
var _ weaver.RoutedBy[router] = (*a)(nil)
var _ weaver.RoutedBy[router] = (*b)(nil)

// Component "a", router "router" checks.
var _ func(context.Context, int, string, bool, [10]int, []string, map[bool]int, message) routingKey = (&router{}).M1 // routed
var _ func(context.Context, int, string, bool, [10]int, []string, map[bool]int, message) routingKey = (&router{}).M2 // routed
// Component "b", router "router" checks.
var _ func(context.Context, int, string, bool, [10]int, []string, map[bool]int, message) routingKey = (&router{}).M1 // routed
var _ func(context.Context, int, string, bool, [10]int, []string, map[bool]int, message) routingKey = (&router{}).M2 // routed

// Local stub implementations.

type a_local_stub struct {
	impl      A
	tracer    trace.Tracer
	m1Metrics *codegen.MethodMetrics
	m2Metrics *codegen.MethodMetrics
}

// Check that a_local_stub implements the A interface.
var _ A = (*a_local_stub)(nil)

func (s a_local_stub) M1(ctx context.Context, a0 int, a1 string, a2 bool, a3 [10]int, a4 []string, a5 map[bool]int, a6 message) (r0 pair, err error) {
	// Update metrics.
	begin := s.m1Metrics.Begin()
	defer func() { s.m1Metrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.A.M1", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.M1(ctx, a0, a1, a2, a3, a4, a5, a6)
}

func (s a_local_stub) M2(ctx context.Context, a0 int, a1 string, a2 bool, a3 [10]int, a4 []string, a5 map[bool]int, a6 message) (r0 pair, err error) {
	// Update metrics.
	begin := s.m2Metrics.Begin()
	defer func() { s.m2Metrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.A.M2", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.M2(ctx, a0, a1, a2, a3, a4, a5, a6)
}

type b_local_stub struct {
	impl      B
	tracer    trace.Tracer
	m1Metrics *codegen.MethodMetrics
	m2Metrics *codegen.MethodMetrics
}

// Check that b_local_stub implements the B interface.
var _ B = (*b_local_stub)(nil)

func (s b_local_stub) M1(ctx context.Context, a0 int, a1 string, a2 bool, a3 [10]int, a4 []string, a5 map[bool]int, a6 message) (r0 pair, err error) {
	// Update metrics.
	begin := s.m1Metrics.Begin()
	defer func() { s.m1Metrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.B.M1", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.M1(ctx, a0, a1, a2, a3, a4, a5, a6)
}

func (s b_local_stub) M2(ctx context.Context, a0 int, a1 string, a2 bool, a3 [10]int, a4 []string, a5 map[bool]int, a6 message) (r0 pair, err error) {
	// Update metrics.
	begin := s.m2Metrics.Begin()
	defer func() { s.m2Metrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.B.M2", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.M2(ctx, a0, a1, a2, a3, a4, a5, a6)
}

// Client stub implementations.

type a_client_stub struct {
	stub      codegen.Stub
	m1Metrics *codegen.MethodMetrics
	m2Metrics *codegen.MethodMetrics
}

// Check that a_client_stub implements the A interface.
var _ A = (*a_client_stub)(nil)

func (s a_client_stub) M1(ctx context.Context, a0 int, a1 string, a2 bool, a3 [10]int, a4 []string, a5 map[bool]int, a6 message) (r0 pair, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.m1Metrics.Begin()
	defer func() { s.m1Metrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.A.M1", trace.WithSpanKind(trace.SpanKindClient))
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

	// Encode arguments.
	enc := codegen.NewEncoder()
	enc.Int(a0)
	enc.String(a1)
	enc.Bool(a2)
	serviceweaver_enc_array_10_int_03f98313(enc, &a3)
	serviceweaver_enc_slice_string_4af10117(enc, a4)
	serviceweaver_enc_map_bool_int_acb668fa(enc, a5)
	(a6).WeaverMarshal(enc)

	// Set the shardKey.
	var r router
	shardKey := _hashA(r.M1(ctx, a0, a1, a2, a3, a4, a5, a6))

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
	(&r0).WeaverUnmarshal(dec)
	err = dec.Error()
	return
}

func (s a_client_stub) M2(ctx context.Context, a0 int, a1 string, a2 bool, a3 [10]int, a4 []string, a5 map[bool]int, a6 message) (r0 pair, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.m2Metrics.Begin()
	defer func() { s.m2Metrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.A.M2", trace.WithSpanKind(trace.SpanKindClient))
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

	// Encode arguments.
	enc := codegen.NewEncoder()
	enc.Int(a0)
	enc.String(a1)
	enc.Bool(a2)
	serviceweaver_enc_array_10_int_03f98313(enc, &a3)
	serviceweaver_enc_slice_string_4af10117(enc, a4)
	serviceweaver_enc_map_bool_int_acb668fa(enc, a5)
	(a6).WeaverMarshal(enc)

	// Set the shardKey.
	var r router
	shardKey := _hashA(r.M2(ctx, a0, a1, a2, a3, a4, a5, a6))

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
	(&r0).WeaverUnmarshal(dec)
	err = dec.Error()
	return
}

type b_client_stub struct {
	stub      codegen.Stub
	m1Metrics *codegen.MethodMetrics
	m2Metrics *codegen.MethodMetrics
}

// Check that b_client_stub implements the B interface.
var _ B = (*b_client_stub)(nil)

func (s b_client_stub) M1(ctx context.Context, a0 int, a1 string, a2 bool, a3 [10]int, a4 []string, a5 map[bool]int, a6 message) (r0 pair, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.m1Metrics.Begin()
	defer func() { s.m1Metrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.B.M1", trace.WithSpanKind(trace.SpanKindClient))
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

	// Encode arguments.
	enc := codegen.NewEncoder()
	enc.Int(a0)
	enc.String(a1)
	enc.Bool(a2)
	serviceweaver_enc_array_10_int_03f98313(enc, &a3)
	serviceweaver_enc_slice_string_4af10117(enc, a4)
	serviceweaver_enc_map_bool_int_acb668fa(enc, a5)
	(a6).WeaverMarshal(enc)

	// Set the shardKey.
	var r router
	shardKey := _hashB(r.M1(ctx, a0, a1, a2, a3, a4, a5, a6))

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
	(&r0).WeaverUnmarshal(dec)
	err = dec.Error()
	return
}

func (s b_client_stub) M2(ctx context.Context, a0 int, a1 string, a2 bool, a3 [10]int, a4 []string, a5 map[bool]int, a6 message) (r0 pair, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.m2Metrics.Begin()
	defer func() { s.m2Metrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.B.M2", trace.WithSpanKind(trace.SpanKindClient))
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

	// Encode arguments.
	enc := codegen.NewEncoder()
	enc.Int(a0)
	enc.String(a1)
	enc.Bool(a2)
	serviceweaver_enc_array_10_int_03f98313(enc, &a3)
	serviceweaver_enc_slice_string_4af10117(enc, a4)
	serviceweaver_enc_map_bool_int_acb668fa(enc, a5)
	(a6).WeaverMarshal(enc)

	// Set the shardKey.
	var r router
	shardKey := _hashB(r.M2(ctx, a0, a1, a2, a3, a4, a5, a6))

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
	(&r0).WeaverUnmarshal(dec)
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
	impl    A
	addLoad func(key uint64, load float64)
}

// Check that a_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*a_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s a_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "M1":
		return s.m1
	case "M2":
		return s.m2
	default:
		return nil
	}
}

func (s a_server_stub) m1(ctx context.Context, args []byte) (res []byte, err error) {
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
	var a1 string
	a1 = dec.String()
	var a2 bool
	a2 = dec.Bool()
	var a3 [10]int
	serviceweaver_dec_array_10_int_03f98313(dec, &a3)
	var a4 []string
	a4 = serviceweaver_dec_slice_string_4af10117(dec)
	var a5 map[bool]int
	a5 = serviceweaver_dec_map_bool_int_acb668fa(dec)
	var a6 message
	(&a6).WeaverUnmarshal(dec)
	var r router
	s.addLoad(_hashA(r.M1(ctx, a0, a1, a2, a3, a4, a5, a6)), 1.0)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.M1(ctx, a0, a1, a2, a3, a4, a5, a6)

	// Encode the results.
	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s a_server_stub) m2(ctx context.Context, args []byte) (res []byte, err error) {
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
	var a1 string
	a1 = dec.String()
	var a2 bool
	a2 = dec.Bool()
	var a3 [10]int
	serviceweaver_dec_array_10_int_03f98313(dec, &a3)
	var a4 []string
	a4 = serviceweaver_dec_slice_string_4af10117(dec)
	var a5 map[bool]int
	a5 = serviceweaver_dec_map_bool_int_acb668fa(dec)
	var a6 message
	(&a6).WeaverUnmarshal(dec)
	var r router
	s.addLoad(_hashA(r.M2(ctx, a0, a1, a2, a3, a4, a5, a6)), 1.0)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.M2(ctx, a0, a1, a2, a3, a4, a5, a6)

	// Encode the results.
	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

type b_server_stub struct {
	impl    B
	addLoad func(key uint64, load float64)
}

// Check that b_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*b_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s b_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "M1":
		return s.m1
	case "M2":
		return s.m2
	default:
		return nil
	}
}

func (s b_server_stub) m1(ctx context.Context, args []byte) (res []byte, err error) {
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
	var a1 string
	a1 = dec.String()
	var a2 bool
	a2 = dec.Bool()
	var a3 [10]int
	serviceweaver_dec_array_10_int_03f98313(dec, &a3)
	var a4 []string
	a4 = serviceweaver_dec_slice_string_4af10117(dec)
	var a5 map[bool]int
	a5 = serviceweaver_dec_map_bool_int_acb668fa(dec)
	var a6 message
	(&a6).WeaverUnmarshal(dec)
	var r router
	s.addLoad(_hashB(r.M1(ctx, a0, a1, a2, a3, a4, a5, a6)), 1.0)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.M1(ctx, a0, a1, a2, a3, a4, a5, a6)

	// Encode the results.
	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s b_server_stub) m2(ctx context.Context, args []byte) (res []byte, err error) {
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
	var a1 string
	a1 = dec.String()
	var a2 bool
	a2 = dec.Bool()
	var a3 [10]int
	serviceweaver_dec_array_10_int_03f98313(dec, &a3)
	var a4 []string
	a4 = serviceweaver_dec_slice_string_4af10117(dec)
	var a5 map[bool]int
	a5 = serviceweaver_dec_map_bool_int_acb668fa(dec)
	var a6 message
	(&a6).WeaverUnmarshal(dec)
	var r router
	s.addLoad(_hashB(r.M2(ctx, a0, a1, a2, a3, a4, a5, a6)), 1.0)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.M2(ctx, a0, a1, a2, a3, a4, a5, a6)

	// Encode the results.
	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type a_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that a_reflect_stub implements the A interface.
var _ A = (*a_reflect_stub)(nil)

func (s a_reflect_stub) M1(ctx context.Context, a0 int, a1 string, a2 bool, a3 [10]int, a4 []string, a5 map[bool]int, a6 message) (r0 pair, err error) {
	err = s.caller("M1", ctx, []any{a0, a1, a2, a3, a4, a5, a6}, []any{&r0})
	return
}

func (s a_reflect_stub) M2(ctx context.Context, a0 int, a1 string, a2 bool, a3 [10]int, a4 []string, a5 map[bool]int, a6 message) (r0 pair, err error) {
	err = s.caller("M2", ctx, []any{a0, a1, a2, a3, a4, a5, a6}, []any{&r0})
	return
}

type b_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that b_reflect_stub implements the B interface.
var _ B = (*b_reflect_stub)(nil)

func (s b_reflect_stub) M1(ctx context.Context, a0 int, a1 string, a2 bool, a3 [10]int, a4 []string, a5 map[bool]int, a6 message) (r0 pair, err error) {
	err = s.caller("M1", ctx, []any{a0, a1, a2, a3, a4, a5, a6}, []any{&r0})
	return
}

func (s b_reflect_stub) M2(ctx context.Context, a0 int, a1 string, a2 bool, a3 [10]int, a4 []string, a5 map[bool]int, a6 message) (r0 pair, err error) {
	err = s.caller("M2", ctx, []any{a0, a1, a2, a3, a4, a5, a6}, []any{&r0})
	return
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*message)(nil)

type __is_message[T ~struct {
	weaver.AutoMarshal
	a int
	b string
	c bool
	d [10]int
	e []string
	f map[bool]int
}] struct{}

var _ __is_message[message]

func (x *message) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("message.WeaverMarshal: nil receiver"))
	}
	enc.Int(x.a)
	enc.String(x.b)
	enc.Bool(x.c)
	serviceweaver_enc_array_10_int_03f98313(enc, &x.d)
	serviceweaver_enc_slice_string_4af10117(enc, x.e)
	serviceweaver_enc_map_bool_int_acb668fa(enc, x.f)
}

func (x *message) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("message.WeaverUnmarshal: nil receiver"))
	}
	x.a = dec.Int()
	x.b = dec.String()
	x.c = dec.Bool()
	serviceweaver_dec_array_10_int_03f98313(dec, &x.d)
	x.e = serviceweaver_dec_slice_string_4af10117(dec)
	x.f = serviceweaver_dec_map_bool_int_acb668fa(dec)
}

func serviceweaver_enc_array_10_int_03f98313(enc *codegen.Encoder, arg *[10]int) {
	for i := 0; i < 10; i++ {
		enc.Int(arg[i])
	}
}

func serviceweaver_dec_array_10_int_03f98313(dec *codegen.Decoder, res *[10]int) {
	for i := 0; i < 10; i++ {
		res[i] = dec.Int()
	}
}

func serviceweaver_enc_slice_string_4af10117(enc *codegen.Encoder, arg []string) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.String(arg[i])
	}
}

func serviceweaver_dec_slice_string_4af10117(dec *codegen.Decoder) []string {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = dec.String()
	}
	return res
}

func serviceweaver_enc_map_bool_int_acb668fa(enc *codegen.Encoder, arg map[bool]int) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for k, v := range arg {
		enc.Bool(k)
		enc.Int(v)
	}
}

func serviceweaver_dec_map_bool_int_acb668fa(dec *codegen.Decoder) map[bool]int {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make(map[bool]int, n)
	var k bool
	var v int
	for i := 0; i < n; i++ {
		k = dec.Bool()
		v = dec.Int()
		res[k] = v
	}
	return res
}

var _ codegen.AutoMarshal = (*pair)(nil)

type __is_pair[T ~struct {
	weaver.AutoMarshal
	a message
	b message
}] struct{}

var _ __is_pair[pair]

func (x *pair) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("pair.WeaverMarshal: nil receiver"))
	}
	(x.a).WeaverMarshal(enc)
	(x.b).WeaverMarshal(enc)
}

func (x *pair) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("pair.WeaverUnmarshal: nil receiver"))
	}
	(&x.a).WeaverUnmarshal(dec)
	(&x.b).WeaverUnmarshal(dec)
}

var _ codegen.AutoMarshal = (*routingKey)(nil)

type __is_routingKey[T ~struct {
	weaver.AutoMarshal
	a int
	b string
	c float32
}] struct{}

var _ __is_routingKey[routingKey]

func (x *routingKey) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("routingKey.WeaverMarshal: nil receiver"))
	}
	enc.Int(x.a)
	enc.String(x.b)
	enc.Float32(x.c)
}

func (x *routingKey) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("routingKey.WeaverUnmarshal: nil receiver"))
	}
	x.a = dec.Int()
	x.b = dec.String()
	x.c = dec.Float32()
}

// Router methods.

// _hashA returns a 64 bit hash of the provided value.
func _hashA(r routingKey) uint64 {
	var h codegen.Hasher
	h.WriteInt(int(r.a))
	h.WriteString(string(r.b))
	h.WriteFloat32(float32(r.c))
	return h.Sum64()
}

// _orderedCodeA returns an order-preserving serialization of the provided value.
func _orderedCodeA(r routingKey) codegen.OrderedCode {
	var enc codegen.OrderedEncoder
	enc.WriteInt(r.a)
	enc.WriteString(r.b)
	enc.WriteFloat32(r.c)
	return enc.Encode()
}

// _hashB returns a 64 bit hash of the provided value.
func _hashB(r routingKey) uint64 {
	var h codegen.Hasher
	h.WriteInt(int(r.a))
	h.WriteString(string(r.b))
	h.WriteFloat32(float32(r.c))
	return h.Sum64()
}

// _orderedCodeB returns an order-preserving serialization of the provided value.
func _orderedCodeB(r routingKey) codegen.OrderedCode {
	var enc codegen.OrderedEncoder
	enc.WriteInt(r.a)
	enc.WriteString(r.b)
	enc.WriteFloat32(r.c)
	return enc.Encode()
}
