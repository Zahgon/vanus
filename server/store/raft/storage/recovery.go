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

	// first-party libraries.
	vanus "github.com/vanus-labs/vanus/api/vsr"
	"github.com/vanus-labs/vanus/pkg/raft/raftpb"

	// this project.
	"github.com/vanus-labs/vanus/server/store/meta"
	walog "github.com/vanus-labs/vanus/server/store/wal"
)

type storagesBuilder struct {
	ctx        context.Context
	stateStore *meta.SyncStore
	hintStore  *meta.AsyncStore
	storages   map[uint64]*Storage
}

var compactSuffix = []byte("/compact")

const (
	nodeIDStart      = 6 // block/
	compactSuffixLen = 8 // /compact
	minCompactKeyLen = nodeIDStart + compactSuffixLen + 1
)

func (sb *storagesBuilder) onMeta(key []byte, value interface{}) error {
	_ = "STUB: not implemented"
	// Filter compact key.
	return nil
}

//nolint:nilerr // skip

// unreachable

func (sb *storagesBuilder) onEntry(data []byte, r walog.Range) error {
	_ = "STUB: not implemented"
	return nil
}

var (
	rangeStartKey = []byte("block/\000")
	rangeEndKey   = []byte("block0")
)

func Recover(
	ctx context.Context, dir string, stateStore *meta.SyncStore, hintStore *meta.AsyncStore, opts ...walog.Option,
) (map[vanus.ID]*Storage, *WAL, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// convert, and set wal

// TODO(james.yin): move to compaction.go

// Start compaction after recover nodes.

func recoverStorage(
	nodeID vanus.ID, wal *WAL, stateStore *meta.SyncStore, hintStore *meta.AsyncStore, info []byte,
	snapOp SnapshotOperator,
) (*Storage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Storage) recoverState() error { _ = "STUB: not implemented"; return nil }

func (s *Storage) recoverHardState() (raftpb.HardState, error) {
	_ = "STUB: not implemented"
	return *new(raftpb.HardState), nil
}

func (s *Storage) recoverConfState() (raftpb.ConfState, error) {
	_ = "STUB: not implemented"
	return *new(raftpb.ConfState), nil
}

func (s *Storage) recoverApplied() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Storage) recoverCompactionInfo(info []byte) error { _ = "STUB: not implemented"; return nil }

// unreachable

func (s *Storage) appendInRecovery(_ context.Context, entry raftpb.Entry, so int64) error {
	_ = "STUB: not implemented"
	return nil
}

// Compacted entry, discard.

// All compacted entries are committed, and committed entries are immutable.

// Write to cache.

// Term will not roll back.

// Append entry.

// Truncate, then append entry.

// In the same term, index is monotonically increasing.

func unreachablePanic(reason string, nodeID vanus.ID, lastTerm, lastIndex, term, index uint64) {
	_ = "STUB: not implemented"
	return
}
