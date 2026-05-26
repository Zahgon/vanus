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
	"context"

	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	cloudEventDataRowLength = 4
	httpPrefix              = "http://"
	xceVanusDeliveryTime    = "xvanusdeliverytime"
)

func NewEventCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func putEventCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func sendOne(ctx context.Context, cmd *cobra.Command, ceClient v2.Client) {
	_ = "STUB: not implemented"
	return
}

// validate event delivery time

// validate event delay time

func sendFile(ctx context.Context, cmd *cobra.Command, ceClient v2.Client) {
	_ = "STUB: not implemented"
	return
}

func getEventCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func format(value *wrapperspb.BytesValue) string { _ = "STUB: not implemented"; return "" }

func queryEventCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

type QueryOutput struct {
	Eventlog string `json:"eventlog"`
	Offset   int64  `json:"offset"`
	Event    string `json:"event"`
}
