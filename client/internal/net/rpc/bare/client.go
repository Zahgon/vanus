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

package bare

import (
	// standard libraries.
	"context"
	"sync"
	"sync/atomic"

	// third-party libraries.

	"google.golang.org/grpc"

	// first-party libraries.

	"github.com/vanus-labs/vanus/pkg/observability/tracing"

	// this project.

	"github.com/vanus-labs/vanus/client/internal/net/rpc"
)

const (
	defaultConnectTimeoutMs = 300
)

func New(endpoint string, creator rpc.ClientCreator) rpc.Client {
	_ = "STUB: not implemented"
	return *new(rpc.Client)
}

// client is a generic client of gRPC.
type client struct {
	endpoint string

	closed  atomic.Bool
	mu      sync.RWMutex
	conn    *grpc.ClientConn
	client  interface{}
	creator rpc.ClientCreator
	tracer  *tracing.Tracer
}

// make sure client implements rpc.Client.
var _ rpc.Client = (*client)(nil)

func (c *client) Endpoint() string { _ = "STUB: not implemented"; return "" }

func (c *client) Get(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) cachedClient() interface{} { _ = "STUB: not implemented"; return nil }

func (c *client) refreshClient(ctx context.Context, force bool) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: close previous connection

// TODO: connect with opts

func (c *client) Ready() bool { _ = "STUB: not implemented"; return false }

func (c *client) Close() { _ = "STUB: not implemented"; return }

func (c *client) doClose() { _ = "STUB: not implemented"; return }
