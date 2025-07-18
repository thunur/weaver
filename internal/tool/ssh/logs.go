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

package ssh

import (
	"context"

	"github.com/thunur/weaver/internal/tool/ssh/impl"
	"github.com/thunur/weaver/runtime/logging"
	"github.com/thunur/weaver/runtime/tool"
)

var logsSpec = tool.LogsSpec{
	Tool: "weaver ssh",
	Source: func(context.Context) (logging.Source, error) {
		return logging.FileSource(impl.LogDir), nil
	},
}
