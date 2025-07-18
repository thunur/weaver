// Copyright 2022 Google LLC
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

// EXPECTED
// time
// A(ctx context.Context, a0 time.Duration)

// UNEXPECTED
// Preallocate

// Imported type with non-default package name.
package foo

import (
	"context"
	xt "time"

	"github.com/thunur/weaver"
)

type foo interface {
	A(context.Context, xt.Duration) error
}

type impl struct{ weaver.Implements[foo] }

func (l *impl) A(context.Context, xt.Duration) error {
	return nil
}
