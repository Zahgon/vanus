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

func NewSubscriptionCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func createSubscriptionCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// subscription config

func getProtocol(cmd *cobra.Command) meta.Protocol {
	_ = "STUB: not implemented"
	return *new(meta.Protocol)
}

func getSinkCredential(cmd *cobra.Command) *meta.SinkCredential {
	_ = "STUB: not implemented"
	return nil
}

// expand value from env

func getFilters(cmd *cobra.Command) []*meta.Filter { _ = "STUB: not implemented"; return nil }

func getTransformer(cmd *cobra.Command) *meta.Transformer { _ = "STUB: not implemented"; return nil }

func getSubscriptionConfig(cmd *cobra.Command, config *meta.SubscriptionConfig) {
	_ = "STUB: not implemented"
	return
}

func updateSubscriptionCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// subscription config

func deleteSubscriptionCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func resumeSubscriptionCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func disableSubscriptionCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func resetOffsetCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func getSubscriptionCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func listSubscriptionCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func printSubscription(cmd *cobra.Command, showNo, showFilters, showTransformer bool, data ...*metapb.Subscription) {
	_ = "STUB: not implemented"
	return
}

var subscriptionHeaders = []interface{}{
	"id", "name", "disable", "eventbus", "namespace", "sink", "description", "protocol", "sinkCredential",
	"config", "offsets", "filter", "transformer", "created_at", "updated_at",
}

func getSubscriptionHeader(showNo bool) table.Row {
	_ = "STUB: not implemented"
	return *new(table.Row)
}

func getSubscriptionRow(sub *meta.Subscription) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func getSubscriptionColumnConfig(header table.Row) []table.ColumnConfig {
	_ = "STUB: not implemented"
	return nil
}
