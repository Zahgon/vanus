// Copyright 2015 The etcd Authors
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

package rafttest

import (
	"sync"

	"github.com/vanus-labs/vanus/pkg/raft"
	"github.com/vanus-labs/vanus/pkg/raft/raftpb"
)

type node struct {
	raft.Node
	id     uint64
	iface  iface
	stopc  chan struct{}
	pausec chan bool

	// stable
	storage *raft.MemoryStorage

	mu    sync.Mutex // guards state
	state raftpb.HardState
}

func startNode(id uint64, peers []raft.Peer, iface iface) *node {
	_ = "STUB: not implemented"
	return nil
}

func (n *node) start() { _ = "STUB: not implemented"; return }

// simulate async send, more like real world...

// step all pending messages

// stop stops the node. stop a stopped node might panic.
// All in memory state of node is discarded.
// All stable MUST be unchanged.
func (n *node) stop() { _ = "STUB: not implemented"; return }

// wait for the shutdown

// restart restarts the node. restart a started node
// blocks and might affect the future stop operation.
func (n *node) restart() {
	_ = "STUB: not implemented"
	// wait for the shutdown
	return
}

// pause pauses the node.
// The paused node buffers the received messages and replies
// all of them when it resumes.
func (n *node) pause() {
	_ = "STUB: not implemented"

	// resume resumes the paused node.
	return
}

func (n *node) resume() { _ = "STUB: not implemented"; return }
