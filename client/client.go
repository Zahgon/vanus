// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate mockgen -source=client.go -destination=mock_client.go -package=client
package client

import (
	// standard libraries.
	"context"
	"sync"

	// third-party libraries.

	// first-party libraries.

	"github.com/vanus-labs/vanus/pkg/observability/tracing"

	// this project.

	"github.com/vanus-labs/vanus/client/pkg/api"
)

type Client interface {
	Eventbus(ctx context.Context, opts ...api.EventbusOption) api.Eventbus
	Disconnect(ctx context.Context)
}

type client struct {
	// Endpoints is a list of URLs.
	Endpoints     []string
	eventbusCache sync.Map

	mu     sync.RWMutex
	tracer *tracing.Tracer
}

func (c *client) Eventbus(ctx context.Context, opts ...api.EventbusOption) api.Eventbus {
	_ = "STUB: not implemented"
	return *new(api.Eventbus)
}

// double check

func (c *client) Disconnect(ctx context.Context) { _ = "STUB: not implemented"; return }

func Connect(endpoints []string) Client { _ = "STUB: not implemented"; return *new(Client) }

func GetEventbusIDIfNotSet(ctx context.Context, endpoints []string, opts *api.EventbusOptions) error {
	_ = "STUB: not implemented"
	// the eventbus id does not exist, get the eventbus id first
	return nil
}

// get eventbus id from name
