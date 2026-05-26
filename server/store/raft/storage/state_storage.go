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

package storage

import (
	// standard libraries.
	"context"

	// third-party libraries.

	// first-party libraries.
	"github.com/vanus-labs/vanus/pkg/raft/raftpb"

	// this project.
	"github.com/vanus-labs/vanus/server/store/meta"
)

type stateStorage struct {
	stateStore *meta.SyncStore
	hintStore  *meta.AsyncStore

	prevHardSt raftpb.HardState
	prevConfSt raftpb.ConfState
	prevApply  uint64

	hsKey  []byte
	offKey []byte
	csKey  []byte
	appKey []byte
}

// InitialState returns the saved HardState and ConfState information.
func (s *Storage) InitialState() (raftpb.HardState, raftpb.ConfState, error) {
	_ = "STUB: not implemented"
	return *new(raftpb.HardState), *new(raftpb.ConfState), nil
}

// HardState returns the saved HardState.
// NOTE: HardState.Commit will always be 0, don't use it.
func (s *Storage) HardState() raftpb.HardState {
	_ = "STUB: not implemented"
	return *new(raftpb.HardState)
}

// SetHardState saves the current HardState.
func (s *Storage) SetHardState(ctx context.Context, hs raftpb.HardState, cb meta.StoreCallback) {
	_ = "STUB: not implemented"
	return
}

func (s *Storage) Commit() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *Storage) SetCommit(ctx context.Context, commit uint64) { _ = "STUB: not implemented"; return }

func (s *Storage) SetConfState(ctx context.Context, cs raftpb.ConfState, cb meta.StoreCallback) {
	_ = "STUB: not implemented"
	return
}

func (s *Storage) Applied() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *Storage) SetApplied(ctx context.Context, app uint64) { _ = "STUB: not implemented"; return }

func HardStateKey(id uint64) string { _ = "STUB: not implemented"; return "" }

func CommitKey(id uint64) string { _ = "STUB: not implemented"; return "" }

func ConfStateKey(id uint64) string { _ = "STUB: not implemented"; return "" }

func ApplyKey(id uint64) string { _ = "STUB: not implemented"; return "" }
