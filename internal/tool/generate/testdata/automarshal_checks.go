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

// EXPECTED
// var _ codegen.AutoMarshal = (*Pair)(nil)
// weaver.AutoMarshal
// X int
// Y int
// }] struct{}
// var _ __is_Pair[Pair]

package foo

import "github.com/thunur/weaver"

type Pair struct {
	weaver.AutoMarshal
	X int
	Y int
}
