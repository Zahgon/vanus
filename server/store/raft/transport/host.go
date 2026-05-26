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
	"sync"

	// first-party libraries.

	"github.com/vanus-labs/vanus/pkg/raft/raftpb"
)

type SendCallback func(error)

type Host interface {
	Sender
	Demultiplexer

	Stop()
	Register(node uint64, r Receiver)
	// TODO(james.yin): Unregister
}

type host struct {
	peers     sync.Map
	receivers sync.Map
	resolver  Resolver
	callback  string
	lo        Multiplexer
}

// Make sure host implements Host.
var _ Host = (*host)(nil)

func NewHost(resolver Resolver, callback string) Host { _ = "STUB: not implemented"; return *new(Host) }

func (h *host) Stop() { _ = "STUB: not implemented"; return }

func (h *host) Send(ctx context.Context, msg *raftpb.Message, to uint64, endpoint string, cb SendCallback) {
	_ = "STUB: not implemented"
	return
}

func (h *host) resolveMultiplexer(_ context.Context, to uint64, endpoint string) Multiplexer {
	_ = "STUB: not implemented"
	return *new(Multiplexer)
}

// Receive implements Demultiplexer.
func (h *host) Receive(ctx context.Context, msg *raftpb.Message, endpoint string) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *host) Register(node uint64, r Receiver) {
	_ = "STUB: not implemented"
	// TODO(james.yin): Handles the case where the receiver already exists.
	return
}
