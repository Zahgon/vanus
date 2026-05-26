// Copyright 2023 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package command

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"

	"github.com/vanus-labs/vanus/api/meta"
	metapb "github.com/vanus-labs/vanus/api/meta"
)

func NewUserCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func createUserCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func deleteUserCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func getUserInfoCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func listUserInfoCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func getUserRolesCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func printUser(cmd *cobra.Command, showNo bool, data ...*metapb.User) {
	_ = "STUB: not implemented"
	return
}

func getUserHeader(showNo bool) table.Row { _ = "STUB: not implemented"; return *new(table.Row) }

func getUserRow(user *meta.User) []interface{} { _ = "STUB: not implemented"; return nil }

func printUserRoles(cmd *cobra.Command, showNo bool, data ...*metapb.UserRole) {
	_ = "STUB: not implemented"
	return
}

func getUserRoleHeader(showNo bool) table.Row { _ = "STUB: not implemented"; return *new(table.Row) }

func getUserRoleRow(user *meta.UserRole) []interface{} { _ = "STUB: not implemented"; return nil }
