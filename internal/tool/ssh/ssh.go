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
	"github.com/thunur/weaver/internal/status"
	itool "github.com/thunur/weaver/internal/tool"
	"github.com/thunur/weaver/runtime/tool"
)

var (
	Commands = map[string]*tool.Command{
		"deploy":    &deployCmd,
		"logs":      tool.LogsCmd(&logsSpec),
		"dashboard": status.DashboardCommand(dashboardSpec),
		"version":   itool.VersionCmd("weaver ssh"),

		// Hidden commands.
		"babysitter": &babysitterCmd,
	}
)
