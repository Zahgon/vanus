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

	// first-party libraries.

	"github.com/vanus-labs/vanus/pkg/raft"
	"github.com/vanus-labs/vanus/pkg/raft/raftpb"
)

// Make sure appender implements raft.Keeper.
var _ raft.Keeper = (*appender)(nil)

func (a *appender) SetHardState(st raftpb.HardState) { _ = "STUB: not implemented"; return }

func (a *appender) CommitTo(index uint64) { _ = "STUB: not implemented"; return }

func (a *appender) SetSoftState(st raft.SoftState) {
	_ = "STUB: not implemented"
	// TODO(james.yin): dispatch to another goroutine.
	return
}

func (a *appender) TruncateAndAppend(ents []raftpb.Entry) { _ = "STUB: not implemented"; return }

func (a *appender) CompactTo(index uint64) { _ = "STUB: not implemented"; return }

func (a *appender) Apply(ents []raftpb.Entry) { _ = "STUB: not implemented"; return }

func (a *appender) Send(msg raftpb.Message) { _ = "STUB: not implemented"; return }
