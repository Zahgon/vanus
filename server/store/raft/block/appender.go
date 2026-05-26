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

//go:generate mockgen -source=appender.go -destination=testing/mock_appender.go -package=testing
package block

import (
	// standard libraries.
	"context"
	"time"

	// first-party libraries.
	vanus "github.com/vanus-labs/vanus/api/vsr"
	"github.com/vanus-labs/vanus/pkg/raft"
	"github.com/vanus-labs/vanus/pkg/raft/raftpb"

	// this project.
	"github.com/vanus-labs/vanus/lib/executor"
	"github.com/vanus-labs/vanus/server/store/block"
	"github.com/vanus-labs/vanus/server/store/raft/storage"
	"github.com/vanus-labs/vanus/server/store/raft/transport"
)

const (
	defaultHintCapacity    = 2
	defaultTickInterval    = 100 * time.Millisecond
	defaultElectionTick    = 10
	defaultHeartbeatTick   = 3
	defaultMaxSizePerMsg   = 16 * 1024
	defaultMaxInflightMsgs = 256
	defaultSendTimeout     = 80 * time.Millisecond
)

type Peer struct {
	ID       vanus.ID
	Endpoint string
}

type ClusterStatus struct {
	Leader vanus.ID
	Term   uint64
}

type LeaderChangedListener = func(block, leader vanus.ID, term uint64)

type EntryAppendedListener = func(block vanus.ID)

type Appender interface {
	block.Appender

	Stop(ctx context.Context)
	Delete(ctx context.Context)
	Bootstrap(ctx context.Context, blocks []Peer) error
	Status() ClusterStatus
}

type appender struct {
	raw       block.Raw
	actx      block.AppendContext
	appendLis EntryAppendedListener

	leaderID  vanus.ID
	leaderLis LeaderChangedListener

	node    *raft.RawNode
	storage *storage.Storage
	host    transport.Host
	e       *Engine

	hint map[uint64]string

	raftExecutor      executor.ExecuteCloser
	appendExecutor    executor.ExecuteCloser
	commitExecutor    executor.ExecuteCloser
	persistExecutor   executor.ExecuteCloser
	applyExecutor     executor.ExecuteCloser
	transportExecutor executor.ExecuteCloser
}

// Make sure appender implements Appender.
var _ Appender = (*appender)(nil)

func (a *appender) ID() vanus.ID { _ = "STUB: not implemented"; return *new(vanus.ID) }

func (a *appender) Stop(ctx context.Context) {
	_ = "STUB: not implemented"
	// TODO(james.yin): waiting for acknowledgments from executors is unnecessary?
	return
}

func (a *appender) Delete(ctx context.Context) { _ = "STUB: not implemented"; return }

// FIXME(james.yin): wakeup inflight append calls?

func (a *appender) Bootstrap(_ context.Context, blocks []Peer) error {
	_ = "STUB: not implemented"
	return nil
}

// sort peers

func (a *appender) persistHardState(ctx context.Context, hs raftpb.HardState) {
	_ = "STUB: not implemented"
	return
}

func (a *appender) persistEntries(ctx context.Context, entries []raftpb.Entry) {
	_ = "STUB: not implemented"
	// log.Debug(ctx).Msg("Append entries to raft log.")
	// 	"node_id":        a.ID(),
	// 	"appended_index": entries[0].Index,
	// 	"entries_num":    len(entries),
	// })
	return
}

// FIXME(james.yin): report to raft?

// Report entries has been persisted.

func (a *appender) compactLog(ctx context.Context, index uint64) {
	_ = "STUB: not implemented"
	// log.Debug(ctx).Msg("Compact raft log.")
	//
	//		"node_id": a.ID(),
	//		"index":   index,
	//	})
	return
}

func (a *appender) applyEntries(ctx context.Context, committedEntries []raftpb.Entry) {
	_ = "STUB: not implemented"
	return
}

// Change membership.

// FIXME(james.yin): do not pass frag with nil value?

// log.Debug(ctx).Msg("Store applied offset.")
// 	"node_id":        a.ID(),
// 	"applied_offset": index,
// })

func (a *appender) onAppend(ctx context.Context, index uint64) { _ = "STUB: not implemented"; return }

func (a *appender) changeMembership(ctx context.Context, pbEntry *raftpb.Entry) {
	_ = "STUB: not implemented"
	return
}

// FIXME(james.yin): do not pass frag with nil value?

func (a *appender) changeConf(_ context.Context, pbEntry *raftpb.Entry) *raftpb.ConfState {
	_ = "STUB: not implemented"
	return nil
}

// TODO(james.yin): return error

// FIXME(james.yin): check it.

func (a *appender) becomeLeader(ctx context.Context) { _ = "STUB: not implemented"; return }

// Reset append context when become leader.

func (a *appender) onLeaderChanged() { _ = "STUB: not implemented"; return }

func (a *appender) resetAppendContext() { _ = "STUB: not implemented"; return }

func (a *appender) doReset() { _ = "STUB: not implemented"; return }

// unreachable

// Entry has been compacted.

// no normal entry

// Append implements block.Appender.
func (a *appender) Append(ctx context.Context, entries []block.Entry, cb block.AppendCallback) {
	_ = "STUB: not implemented"
	return
}

func (a *appender) doAppend(ctx context.Context, entries []block.Entry, cb block.AppendCallback) {
	_ = "STUB: not implemented"
	return
}

// FIXME(james.yin): revert archived if propose failed.

func (a *appender) Status() ClusterStatus { _ = "STUB: not implemented"; return *new(ClusterStatus) }

func (a *appender) leaderInfo() (vanus.ID, uint64) {
	_ = "STUB: not implemented"
	return *new(vanus.ID), 0
}

func (a *appender) isLeader() bool { _ = "STUB: not implemented"; return false }
