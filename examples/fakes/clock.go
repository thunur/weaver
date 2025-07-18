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

package fakes

import (
	"context"
	"time"

	"github.com/thunur/weaver"
)

//go:generate ../../cmd/weaver/weaver generate

// Clock component interface.
type Clock interface {
	// UnixMicro returns the current time in microseconds since the unix epoch.
	UnixMicro(context.Context) (int64, error)
}

// Clock component implementation.
type clock struct {
	weaver.Implements[Clock]
}

// UnixMicro implements the Clock interface.
func (clock) UnixMicro(context.Context) (int64, error) {
	return time.Now().UnixMicro(), nil
}
