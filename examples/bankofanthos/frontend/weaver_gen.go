// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package frontend

import (
	"context"
	"github.com/thunur/weaver"
	"github.com/thunur/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:      "github.com/thunur/weaver/Main",
		Iface:     reflect.TypeOf((*weaver.Main)(nil)).Elem(),
		Impl:      reflect.TypeOf(server{}),
		Listeners: []string{"bank"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return main_local_stub{impl: impl.(weaver.Main), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any { return main_client_stub{stub: stub} },
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return main_server_stub{impl: impl.(weaver.Main), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return main_reflect_stub{caller: caller}
		},
		RefData: "⟦583f439b:wEaVeReDgE:github.com/thunur/weaver/Main→github.com/thunur/weaver/examples/bankofanthos/balancereader/T⟧\n⟦01efa328:wEaVeReDgE:github.com/thunur/weaver/Main→github.com/thunur/weaver/examples/bankofanthos/contacts/T⟧\n⟦285db949:wEaVeReDgE:github.com/thunur/weaver/Main→github.com/thunur/weaver/examples/bankofanthos/ledgerwriter/T⟧\n⟦c236fa3b:wEaVeReDgE:github.com/thunur/weaver/Main→github.com/thunur/weaver/examples/bankofanthos/transactionhistory/T⟧\n⟦0906345d:wEaVeReDgE:github.com/thunur/weaver/Main→github.com/thunur/weaver/examples/bankofanthos/userservice/T⟧\n⟦969790bc:wEaVeRlIsTeNeRs:github.com/thunur/weaver/Main→bank⟧\n",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[weaver.Main] = (*server)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*server)(nil)

// Local stub implementations.

type main_local_stub struct {
	impl   weaver.Main
	tracer trace.Tracer
}

// Check that main_local_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_local_stub)(nil)

// Client stub implementations.

type main_client_stub struct {
	stub codegen.Stub
}

// Check that main_client_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_client_stub)(nil)

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

type main_server_stub struct {
	impl    weaver.Main
	addLoad func(key uint64, load float64)
}

// Check that main_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*main_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s main_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	default:
		return nil
	}
}

// Reflect stub implementations.

type main_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that main_reflect_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_reflect_stub)(nil)

