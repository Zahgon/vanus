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

package raft

import pb "github.com/vanus-labs/vanus/pkg/raft/raftpb"

// unstable.entries[i] has raft log position i+unstable.offset.
// Note that unstable.offset may be less than the highest log
// position in storage; this means that the next write to storage
// might need to truncate the log before persisting unstable.entries.
type unstable struct {
	// the incoming unstable snapshot, if any.
	snapshot *pb.Snapshot
	// all entries that have not yet been written to storage.
	entries []pb.Entry
	offset  uint64

	logger Logger
}

// maybeFirstIndex returns the index of the first possible entry in entries
// if it has a snapshot.
func (u *unstable) maybeFirstIndex() (uint64, bool) { _ = "STUB: not implemented"; return 0, false }

// maybeLastIndex returns the last index if it has at least one
// unstable entry or snapshot.
func (u *unstable) maybeLastIndex() (uint64, bool) { _ = "STUB: not implemented"; return 0, false }

// maybeTerm returns the term of the entry at index i, if there
// is any.
func (u *unstable) maybeTerm(i uint64) (uint64, bool) { _ = "STUB: not implemented"; return 0, false }

func (u *unstable) stableTo(i, t uint64) bool { _ = "STUB: not implemented"; return false }

// Unstable entry missing. Ignore.

// Index matched unstable snapshot, not unstable entry. Ignore.

// Term mismatch between unstable entry and specified entry. Ignore.
// This is possible if part or all of the unstable log was replaced
// between that time that a set of entries started to be written to
// stable storage and when they finished.

// shrinkEntriesArray discards the underlying array used by the entries slice
// if most of it isn't being used. This avoids holding references to a bunch of
// potentially large entries that aren't needed anymore. Simply clearing the
// entries wouldn't be safe because clients might still be using them.
func (u *unstable) shrinkEntriesArray() {
	_ = "STUB: not implemented"
	// We replace the array if we're using less than half of the space in
	// it. This number is fairly arbitrary, chosen as an attempt to balance
	// memory usage vs number of allocations. It could probably be improved
	// with some focused tuning.
	return
}

func (u *unstable) stableSnapTo(i uint64) bool { _ = "STUB: not implemented"; return false }

func (u *unstable) restore(s pb.Snapshot) { _ = "STUB: not implemented"; return }

func (u *unstable) truncateAndAppend(ents []pb.Entry) (li uint64, truncated bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// start is the next index in the u.entries, so append directly.

// The log is being truncated to before our current offset
// portion, so set the offset and replace the entries.

// truncate to start and copy to u.entries then append.

func (u *unstable) slice(lo uint64, hi uint64) []pb.Entry { _ = "STUB: not implemented"; return nil }

// u.offset <= lo <= hi <= u.offset+len(u.entries)
func (u *unstable) mustCheckOutOfBounds(lo, hi uint64) { _ = "STUB: not implemented"; return }
