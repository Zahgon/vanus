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

func NewNamespaceCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func createNamespaceCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func deleteNamespaceCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func getNamespaceInfoCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func listNamespaceInfoCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func grantNamespaceCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func revokeNamespaceCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func printNamespace(cmd *cobra.Command, showNo bool, data ...*metapb.Namespace) {
	_ = "STUB: not implemented"
	return
}

func getNamespaceHeader(showNo bool) table.Row { _ = "STUB: not implemented"; return *new(table.Row) }

func getNamespaceRow(ns *meta.Namespace) []interface{} { _ = "STUB: not implemented"; return nil }
