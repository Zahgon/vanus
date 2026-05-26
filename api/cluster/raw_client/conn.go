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

package raw_client

import (
	// standard libraries.
	"context"
	"sync"

	// third-party libraries.
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	// first-party libraries.
	// this project.
)

const (
	vanusConnBypass = "VANUS_CONN_BYPASS"
)

type Conn struct {
	mutex        sync.Mutex
	leader       string
	leaderClient *grpc.ClientConn
	endpoints    []string
	credentials  credentials.TransportCredentials
	grpcConn     map[string]*grpc.ClientConn
	bypass       bool
}

func NewConnection(endpoints []string, credentials credentials.TransportCredentials) *Conn {
	_ = "STUB: not implemented"
	// TODO temporary implement
	return nil
}

func (c *Conn) invoke(ctx context.Context, method string, args, reply interface{}, opts ...grpc.CallOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) close() error { _ = "STUB: not implemented"; return nil }

func (c *Conn) makeSureClient(ctx context.Context, renew bool) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) getGRPCConn(ctx context.Context, addr string) *grpc.ClientConn {
	_ = "STUB: not implemented"
	return nil
}

// make sure it's closed

func isNeedRetry(err error) bool { _ = "STUB: not implemented"; return false }

func isConnectionOK(conn *grpc.ClientConn) bool { _ = "STUB: not implemented"; return false }
