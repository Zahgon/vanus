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

package block

import (
	// standard libraries.
	"context"

	// first-party libraries.
	"github.com/vanus-labs/vanus/pkg/raft"
	"github.com/vanus-labs/vanus/pkg/raft/raftpb"
)

func (a *appender) step(msg *raftpb.Message) { _ = "STUB: not implemented"; return }

func (a *appender) propose(pds ...raft.ProposeData) { _ = "STUB: not implemented"; return }

func (a *appender) reportStateStatus(_ context.Context, term, vote uint64) {
	_ = "STUB: not implemented"
	return
}

func (a *appender) reportLogStatus(_ context.Context, index, term uint64) {
	_ = "STUB: not implemented"
	return
}

func (a *appender) reportApplyStatus(_ context.Context, index uint64) {
	_ = "STUB: not implemented"
	return
}

func (a *appender) reportUnreachable(id uint64) { _ = "STUB: not implemented"; return }

func (a *appender) tick() bool { _ = "STUB: not implemented"; return false }

func (a *appender) bootstrap(peers []raft.Peer) error { _ = "STUB: not implemented"; return nil }

// FIXME(james.yin): appender is stopped when bootstrap.

func (a *appender) applyConfChange(cc raftpb.ConfChangeI) *raftpb.ConfState {
	_ = "STUB: not implemented"
	return nil
}
