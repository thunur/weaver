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

package reflection_test

import (
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/thunur/weaver/internal/reflection"
)

func TestType(t *testing.T) {
	for _, test := range []struct {
		want string
		t    reflect.Type
	}{
		{"int", reflection.Type[int]()},
		{"io.Reader", reflection.Type[io.Reader]()},
	} {
		t.Run(test.want, func(t *testing.T) {
			if got := fmt.Sprint(test.t); got != test.want {
				t.Errorf("reflection.Type() = %s, expecting %s", got, test.want)
			}
		})
	}
}
