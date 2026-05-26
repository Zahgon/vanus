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
	"errors"

	"github.com/vanus-labs/vanus/pkg/raft/raftpb"
	// this project.
)

var (
	ErrNoEntry   = errors.New("no entry")
	ErrBadEntry  = errors.New("bad entry")
	ErrCompacted = errors.New("appending entries has been compacted")
	ErrTruncated = errors.New("appending entries has been truncated")
)

type logStorage struct {
	// ents[0] is a dummy entry, which record compact information.
	// ents[i] has raft log position i+snapshot.Metadata.Index.
	ents []raftpb.Entry

	// offs[0] is a dummy entry, which records last offset where the barrier was set.
	// offs[i] is the start offset of ents[i] in WAL.
	offs []int64
	tail int64

	wal *WAL
}

func (s *Storage) Compacted() uint64 { _ = "STUB: not implemented"; return 0 }

// Entries returns a slice of log entries in the range [lo,hi).
// MaxSize limits the total size of the log entries returned, but
// Entries returns at least one entry if any.
func (s *Storage) Entries(lo, hi, maxSize uint64) ([]raftpb.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// no log entry

func limitSize(ents []raftpb.Entry, maxSize uint64) []raftpb.Entry {
	_ = "STUB: not implemented"
	return nil
}

// Term returns the term of entry i, which must be in the range
// [FirstIndex()-1, LastIndex()]. The term of the entry before
// FirstIndex is retained for matching purposes even though the
// rest of that entry may not be available.
func (s *Storage) Term(i uint64) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Storage) term(i uint64) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Storage) lastTerm() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *Storage) lastStableTerm() uint64 {
	_ = "STUB: not implemented" //nolint:unused // ok
	return 0
}

func (s *Storage) compactedTerm() uint64 { _ = "STUB: not implemented"; return 0 }

// LastIndex returns the index of the last entry in the log.
func (s *Storage) LastIndex() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// FIXME(james.yin): no entry

func (s *Storage) lastIndex() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *Storage) lastStableIndex() uint64 { _ = "STUB: not implemented"; return 0 }

// FirstIndex returns the index of the first log entry that is
// possibly available via Entries (older entries have been incorporated
// into the latest Snapshot; if storage only contains the dummy entry the
// first log entry is not available).
func (s *Storage) FirstIndex() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// FIXME(james.yin): no entry

func (s *Storage) firstIndex() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *Storage) compactedIndex() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *Storage) length() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *Storage) stableLength() uint64 { _ = "STUB: not implemented"; return 0 }

type AppendResult struct {
	Term  uint64
	Index uint64
}

type AppendCallback = func(AppendResult, error)

// Append appends the new entries to storage.
// After the call returns, all entries are readable. After the AppendCallback cb fires, all entries are persisted.
// NOTE: Synchronization is the responsibility of the caller.
func (s *Storage) Append(ctx context.Context, entries []raftpb.Entry, cb AppendCallback) {
	_ = "STUB: not implemented" //nolint:funlen // ok
	return
}

// entries[len(entries)-1].Index

// Shortcut if there is no new entry.

// Truncate compacted entries.

// append
//nolint:gocritic // assign below

// truncate then append: term > lastTerm

// truncate offsets

// Append to WAL.

// Mark barrier on the offset of first entry in WAL.

// FIXME(james.yin): appender is deleted.

func (s *Storage) postAppend(entries []raftpb.Entry, offsets []int64, tail int64, remark bool, cb AppendCallback) {
	_ = "STUB: not implemented"
	return
}

// Remove obsolete barrier from truncated entry.

// All entries has been truncated.

// append

// truncate then append: term > lastTerm

// FIXME(james.yin): real?

// Remove obsolete barrier.

// Record new barrier.

func (s *Storage) truncateObsoleteEntries(entries []raftpb.Entry) []raftpb.Entry {
	_ = "STUB: not implemented"
	return nil
}

func (s *Storage) prepareAppend(ctx context.Context, entries []raftpb.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Storage) marshalEntries(entries []raftpb.Entry) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// reset node ID.
