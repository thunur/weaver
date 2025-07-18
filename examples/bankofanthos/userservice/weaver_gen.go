// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package userservice

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
		Name:  "github.com/thunur/weaver/examples/bankofanthos/userservice/T",
		Iface: reflect.TypeOf((*T)(nil)).Elem(),
		Impl:  reflect.TypeOf(impl{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return t_local_stub{impl: impl.(T), tracer: tracer, createUserMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/examples/bankofanthos/userservice/T", Method: "CreateUser", Remote: false, Generated: true}), loginMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/examples/bankofanthos/userservice/T", Method: "Login", Remote: false, Generated: true})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return t_client_stub{stub: stub, createUserMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/examples/bankofanthos/userservice/T", Method: "CreateUser", Remote: true, Generated: true}), loginMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/thunur/weaver/examples/bankofanthos/userservice/T", Method: "Login", Remote: true, Generated: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return t_server_stub{impl: impl.(T), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return t_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[T] = (*impl)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*impl)(nil)

// Local stub implementations.

type t_local_stub struct {
	impl              T
	tracer            trace.Tracer
	createUserMetrics *codegen.MethodMetrics
	loginMetrics      *codegen.MethodMetrics
}

// Check that t_local_stub implements the T interface.
var _ T = (*t_local_stub)(nil)

func (s t_local_stub) CreateUser(ctx context.Context, a0 CreateUserRequest) (err error) {
	// Update metrics.
	begin := s.createUserMetrics.Begin()
	defer func() { s.createUserMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "userservice.T.CreateUser", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.CreateUser(ctx, a0)
}

func (s t_local_stub) Login(ctx context.Context, a0 LoginRequest) (r0 string, err error) {
	// Update metrics.
	begin := s.loginMetrics.Begin()
	defer func() { s.loginMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "userservice.T.Login", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Login(ctx, a0)
}

// Client stub implementations.

type t_client_stub struct {
	stub              codegen.Stub
	createUserMetrics *codegen.MethodMetrics
	loginMetrics      *codegen.MethodMetrics
}

// Check that t_client_stub implements the T interface.
var _ T = (*t_client_stub)(nil)

func (s t_client_stub) CreateUser(ctx context.Context, a0 CreateUserRequest) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.createUserMetrics.Begin()
	defer func() { s.createUserMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "userservice.T.CreateUser", trace.WithSpanKind(trace.SpanKindClient))
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
	size += serviceweaver_size_CreateUserRequest_4ef79cd1(&a0)
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	(a0).WeaverMarshal(enc)
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
	err = dec.Error()
	return
}

func (s t_client_stub) Login(ctx context.Context, a0 LoginRequest) (r0 string, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.loginMetrics.Begin()
	defer func() { s.loginMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "userservice.T.Login", trace.WithSpanKind(trace.SpanKindClient))
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
	size += serviceweaver_size_LoginRequest_cbd66e76(&a0)
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	(a0).WeaverMarshal(enc)
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

type t_server_stub struct {
	impl    T
	addLoad func(key uint64, load float64)
}

// Check that t_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*t_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s t_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "CreateUser":
		return s.createUser
	case "Login":
		return s.login
	default:
		return nil
	}
}

func (s t_server_stub) createUser(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 CreateUserRequest
	(&a0).WeaverUnmarshal(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.CreateUser(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) login(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 LoginRequest
	(&a0).WeaverUnmarshal(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Login(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.String(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type t_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that t_reflect_stub implements the T interface.
var _ T = (*t_reflect_stub)(nil)

func (s t_reflect_stub) CreateUser(ctx context.Context, a0 CreateUserRequest) (err error) {
	err = s.caller("CreateUser", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) Login(ctx context.Context, a0 LoginRequest) (r0 string, err error) {
	err = s.caller("Login", ctx, []any{a0}, []any{&r0})
	return
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*CreateUserRequest)(nil)

type __is_CreateUserRequest[T ~struct {
	weaver.AutoMarshal
	Username       string
	Password       string
	PasswordRepeat string
	FirstName      string
	LastName       string
	Birthday       string
	Timezone       string
	Address        string
	State          string
	Zip            string
	Ssn            string
}] struct{}

var _ __is_CreateUserRequest[CreateUserRequest]

func (x *CreateUserRequest) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("CreateUserRequest.WeaverMarshal: nil receiver"))
	}
	enc.String(x.Username)
	enc.String(x.Password)
	enc.String(x.PasswordRepeat)
	enc.String(x.FirstName)
	enc.String(x.LastName)
	enc.String(x.Birthday)
	enc.String(x.Timezone)
	enc.String(x.Address)
	enc.String(x.State)
	enc.String(x.Zip)
	enc.String(x.Ssn)
}

func (x *CreateUserRequest) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("CreateUserRequest.WeaverUnmarshal: nil receiver"))
	}
	x.Username = dec.String()
	x.Password = dec.String()
	x.PasswordRepeat = dec.String()
	x.FirstName = dec.String()
	x.LastName = dec.String()
	x.Birthday = dec.String()
	x.Timezone = dec.String()
	x.Address = dec.String()
	x.State = dec.String()
	x.Zip = dec.String()
	x.Ssn = dec.String()
}

var _ codegen.AutoMarshal = (*LoginRequest)(nil)

type __is_LoginRequest[T ~struct {
	weaver.AutoMarshal
	Username string
	Password string
}] struct{}

var _ __is_LoginRequest[LoginRequest]

func (x *LoginRequest) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("LoginRequest.WeaverMarshal: nil receiver"))
	}
	enc.String(x.Username)
	enc.String(x.Password)
}

func (x *LoginRequest) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("LoginRequest.WeaverUnmarshal: nil receiver"))
	}
	x.Username = dec.String()
	x.Password = dec.String()
}

var _ codegen.AutoMarshal = (*User)(nil)

type __is_User[T ~struct {
	weaver.AutoMarshal
	AccountID string "gorm:\"column:accountid;primary_key\""
	Username  string "gorm:\"unique;not null\""
	Passhash  []byte "gorm:\"not null\""
	Firstname string "gorm:\"not null\""
	Lastname  string "gorm:\"not null\""
	Birthday  string "gorm:\"not null\""
	Timezone  string "gorm:\"not null\""
	Address   string "gorm:\"not null\""
	State     string "gorm:\"not null\""
	Zip       string "gorm:\"not null\""
	SSN       string "gorm:\"not null\""
}] struct{}

var _ __is_User[User]

func (x *User) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("User.WeaverMarshal: nil receiver"))
	}
	enc.String(x.AccountID)
	enc.String(x.Username)
	serviceweaver_enc_slice_byte_87461245(enc, x.Passhash)
	enc.String(x.Firstname)
	enc.String(x.Lastname)
	enc.String(x.Birthday)
	enc.String(x.Timezone)
	enc.String(x.Address)
	enc.String(x.State)
	enc.String(x.Zip)
	enc.String(x.SSN)
}

func (x *User) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("User.WeaverUnmarshal: nil receiver"))
	}
	x.AccountID = dec.String()
	x.Username = dec.String()
	x.Passhash = serviceweaver_dec_slice_byte_87461245(dec)
	x.Firstname = dec.String()
	x.Lastname = dec.String()
	x.Birthday = dec.String()
	x.Timezone = dec.String()
	x.Address = dec.String()
	x.State = dec.String()
	x.Zip = dec.String()
	x.SSN = dec.String()
}

func serviceweaver_enc_slice_byte_87461245(enc *codegen.Encoder, arg []byte) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.Byte(arg[i])
	}
}

func serviceweaver_dec_slice_byte_87461245(dec *codegen.Decoder) []byte {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = dec.Byte()
	}
	return res
}

// Size implementations.

// serviceweaver_size_CreateUserRequest_4ef79cd1 returns the size (in bytes) of the serialization
// of the provided type.
func serviceweaver_size_CreateUserRequest_4ef79cd1(x *CreateUserRequest) int {
	size := 0
	size += 0
	size += (4 + len(x.Username))
	size += (4 + len(x.Password))
	size += (4 + len(x.PasswordRepeat))
	size += (4 + len(x.FirstName))
	size += (4 + len(x.LastName))
	size += (4 + len(x.Birthday))
	size += (4 + len(x.Timezone))
	size += (4 + len(x.Address))
	size += (4 + len(x.State))
	size += (4 + len(x.Zip))
	size += (4 + len(x.Ssn))
	return size
}

// serviceweaver_size_LoginRequest_cbd66e76 returns the size (in bytes) of the serialization
// of the provided type.
func serviceweaver_size_LoginRequest_cbd66e76(x *LoginRequest) int {
	size := 0
	size += 0
	size += (4 + len(x.Username))
	size += (4 + len(x.Password))
	return size
}
