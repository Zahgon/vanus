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
	"math/rand"
	"sync"
	"time"

	"github.com/vanus-labs/vanus/pkg/raft/raftpb"
)

// a network interface
type iface interface {
	send(m raftpb.Message)
	recv() chan raftpb.Message
	disconnect()
	connect()
}

type raftNetwork struct {
	rand         *rand.Rand
	mu           sync.Mutex
	disconnected map[uint64]bool
	dropmap      map[conn]float64
	delaymap     map[conn]delay
	recvQueues   map[uint64]chan raftpb.Message
}

type conn struct {
	from, to uint64
}

type delay struct {
	d    time.Duration
	rate float64
}

func newRaftNetwork(nodes ...uint64) *raftNetwork { _ = "STUB: not implemented"; return nil }

func (rn *raftNetwork) nodeNetwork(id uint64) iface { _ = "STUB: not implemented"; return *new(iface) }

func (rn *raftNetwork) send(m raftpb.Message) { _ = "STUB: not implemented"; return }

// TODO: shall we dl without blocking the send call?

// use marshal/unmarshal to copy message to avoid data race.

// drop messages when the receiver queue is full.

func (rn *raftNetwork) recvFrom(from uint64) chan raftpb.Message {
	_ = "STUB: not implemented"
	return nil
}

func (rn *raftNetwork) drop(from, to uint64, rate float64) { _ = "STUB: not implemented"; return }

func (rn *raftNetwork) delay(from, to uint64, d time.Duration, rate float64) {
	_ = "STUB: not implemented"
	return
}

func (rn *raftNetwork) disconnect(id uint64) { _ = "STUB: not implemented"; return }

func (rn *raftNetwork) connect(id uint64) { _ = "STUB: not implemented"; return }

type nodeNetwork struct {
	id uint64
	*raftNetwork
}

func (nt *nodeNetwork) connect() { _ = "STUB: not implemented"; return }

func (nt *nodeNetwork) disconnect() { _ = "STUB: not implemented"; return }

func (nt *nodeNetwork) send(m raftpb.Message) { _ = "STUB: not implemented"; return }

func (nt *nodeNetwork) recv() chan raftpb.Message { _ = "STUB: not implemented"; return nil }
