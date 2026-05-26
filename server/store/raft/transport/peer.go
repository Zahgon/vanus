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

package transport

import (
	// standard libraries.
	"context"
	"time"

	// third-party libraries.
	"google.golang.org/grpc"

	// first-party libraries.
	vsraftpb "github.com/vanus-labs/vanus/api/raft"
	"github.com/vanus-labs/vanus/pkg/raft/raftpb"
)

const (
	minConnectTimeout       = 100 * time.Millisecond
	defaultConnectTimeout   = 300 * time.Millisecond
	defaultMessageChainSize = 2048
)

type task struct {
	msg *raftpb.Message
	ctx context.Context
	cb  SendCallback
}

type peer struct {
	addr   string
	taskc  chan task
	stream vsraftpb.RaftServer_SendMessageClient
	closec chan struct{}
	donec  chan struct{}
}

// Make sure peer implements Multiplexer.
var _ Multiplexer = (*peer)(nil)

func newPeer(endpoint string, callback string) *peer { _ = "STUB: not implemented"; return nil }

func (p *peer) run(callback string) { _ = "STUB: not implemented"; return }

func (p *peer) processSendError(t task, err error) { _ = "STUB: not implemented"; return }

func (p *peer) Close() { _ = "STUB: not implemented"; return }

func (p *peer) Send(ctx context.Context, msg *raftpb.Message, cb SendCallback) {
	_ = "STUB: not implemented"
	return
}

func (p *peer) connect(ctx context.Context, opts ...grpc.DialOption) (vsraftpb.RaftServer_SendMessageClient, error) {
	_ = "STUB: not implemented"
	return *new(vsraftpb.RaftServer_SendMessageClient), nil
}
