// Copyright 2022 Linkall Inc.
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

package gateway

import (
	// standard libraries.
	"context"
	"net"

	// third-party libraries.
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/protocol"
	cehttp "github.com/cloudevents/sdk-go/v2/protocol/http"

	// first-party libraries.

	"github.com/vanus-labs/vanus/api/cluster"
	vanus "github.com/vanus-labs/vanus/api/vsr"
	"github.com/vanus-labs/vanus/pkg/observability/tracing"

	// this project.

	"github.com/vanus-labs/vanus/server/gateway/proxy"
)

var requestDataFromContext = cehttp.RequestDataFromContext

type EventData struct {
	EventID string   `json:"event_id"`
	BusID   vanus.ID `json:"eventbus_id"`
}

type ceGateway struct {
	config     Config
	proxySrv   *proxy.ControllerProxy
	tracer     *tracing.Tracer
	ceListener net.Listener
	ctrl       cluster.Cluster
}

func NewGateway(config Config) *ceGateway { _ = "STUB: not implemented"; return nil }

func (ga *ceGateway) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (ga *ceGateway) Stop() { _ = "STUB: not implemented"; return }

func (ga *ceGateway) startCloudEventsReceiver(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (ga *ceGateway) receive(ctx context.Context, event v2.Event) (re *v2.Event, result protocol.Result) {
	_ = "STUB: not implemented"
	return nil, *new(protocol.Result)
}

const (
	httpRequestPrefix = "/gateway"
)

func (ga *ceGateway) getEventbusFromPath(ctx context.Context, reqData *cehttp.RequestData) (vanus.ID, error) {
	_ = "STUB: not implemented"
	// TODO validate
	return *new(vanus.ID), nil
}

// Deprecated, just for compatibility of older than v0.7.0

// namespaces/:namespace_name/eventbus/:eventbus_name/events
