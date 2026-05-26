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

import (
	pb "github.com/vanus-labs/vanus/pkg/raft/raftpb"
)

type raftLog struct {
	// storage contains all stable entries since the last snapshot.
	storage Storage
	// keeper writes entries to the stable storage.
	keeper Keeper

	inflight inflight

	// unstable contains all unstable entries and snapshot.
	// they will be saved into storage.
	unstable unstable

	// persisting is the next log position that will be persisted to storage.
	// Invariant: unstable.offset <= persisting
	persisting uint64
	// committed is the highest log position that is known to be in
	// stable storage on a quorum of nodes.
	// Invariant: committed < unstable.offset + len(unstable.entries)
	committed uint64
	// Invariant: localCommitted = min(committed, unstable.offset)
	localCommitted uint64
	// applying is the highest log position that the application has
	// been instructed to be applying to its state machine.
	// Invariant: applying <= localCommitted
	applying uint64
	// applied is the highest log position that the application has
	// been instructed to apply to its state machine.
	// Invariant: applied <= applying
	applied uint64
	// compacted is the highest log position that the application can
	// delete safety.
	// Invariant: compacted <= applied
	compacted uint64

	logger Logger

	// maxNextEntsSize is the maximum number aggregate byte size of the messages
	// returned from calls to nextEnts.
	maxNextEntsSize uint64
}

// newLog returns log using the given storage and default options. It
// recovers the log to the state that it just commits and applies the
// latest snapshot.
func newLog(storage Storage, keeper Keeper, logger Logger) *raftLog {
	_ = "STUB: not implemented"
	return nil
}

// newLogWithSize returns a log using the given storage and max
// message size.
func newLogWithSize(storage Storage, keeper Keeper, logger Logger, maxNextEntsSize uint64) *raftLog {
	_ = "STUB: not implemented"
	return nil
}

// TODO(bdarnell)

// TODO(bdarnell)

// Initialize our committed and applied pointers to the time of the last compaction.

func (l *raftLog) String() string { _ = "STUB: not implemented"; return "" }

// maybeAppend returns false if the entries cannot be appended. Otherwise, it returns true.
func (l *raftLog) maybeAppend(index, logTerm, committed uint64, ents ...pb.Entry) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func (l *raftLog) append(ents ...pb.Entry) uint64 { _ = "STUB: not implemented"; return 0 }

// Reset pending when any entry being persisted is truncated.

// TODO(james.yin): limiting

// findConflict finds the index of the conflict.
// It returns the first pair of conflicting entries between the existing
// entries and the given entries, if there are any.
// If there is no conflicting entries, and the existing entries contains
// all the given entries, zero will be returned.
// If there is no conflicting entries, but the given entries contains new
// entries, the index of the first new entry will be returned.
// An entry is considered to be conflicting if it has the same index but
// a different term.
// The index of the given entries MUST be continuously increasing.
func (l *raftLog) findConflict(ents []pb.Entry) uint64 { _ = "STUB: not implemented"; return 0 }

// findConflictByTerm takes an (index, term) pair (indicating a conflicting log
// entry on a leader/follower during an append) and finds the largest index in
// log l with a term <= `term` and an index <= `index`. If no such index exists
// in the log, the log's first index is returned.
//
// The index provided MUST be equal to or less than l.lastIndex(). Invalid
// inputs log a warning and the input index is returned.
func (l *raftLog) findConflictByTerm(index uint64, term uint64) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// NB: such calls should not exist, but since there is a straightfoward
// way to recover, do it.
//
// It is tempting to also check something about the first index, but
// there is odd behavior with peers that have no log, in which case
// lastIndex will return zero and firstIndex will return one, which
// leads to calls with an index of zero into this method.

// hasPendingSnapshot returns if there is pending snapshot waiting for applying.
func (l *raftLog) hasPendingSnapshot() bool { _ = "STUB: not implemented"; return false }

func (l *raftLog) snapshot() (pb.Snapshot, error) {
	_ = "STUB: not implemented"
	return *new(pb.Snapshot), nil
}

func (l *raftLog) firstIndex() uint64 { _ = "STUB: not implemented"; return 0 }

// TODO(bdarnell)

func (l *raftLog) lastIndex() uint64 { _ = "STUB: not implemented"; return 0 }

func (l *raftLog) stableLastIndex() uint64 {
	_ = "STUB: not implemented" //nolint:unused // ok
	return 0
}

// TODO(james.yin)

func (l *raftLog) compactTo(tocompact uint64) { _ = "STUB: not implemented"; return }

func (l *raftLog) appliedTo(i uint64) { _ = "STUB: not implemented"; return }

func (l *raftLog) applyingTo(i uint64) { _ = "STUB: not implemented"; return }

func (l *raftLog) localCommitTo(tocommit uint64) { _ = "STUB: not implemented"; return }

// never decrease commit

// if li := l.stableLastIndex(); li < tocommit {
// 	l.logger.Panicf("tocommit(%d) is out of range [lastIndex(%d)]. Was the raft log corrupted, truncated, or lost?", tocommit, li)
// }

// TODO(james.yin): limiting

func (l *raftLog) commitTo(tocommit uint64) {
	_ = "STUB: not implemented"
	// never decrease commit
	return
}

func (l *raftLog) persistingTo(i, t uint64) { _ = "STUB: not implemented"; return }

// if i < offset, term is matched with the snapshot
// only update the pending if term is matched with an unstable entry.

func (l *raftLog) stableTo(i, t uint64) bool { _ = "STUB: not implemented"; return false }

func (l *raftLog) stableSnapTo(i uint64) { _ = "STUB: not implemented"; return }

func (l *raftLog) lastTerm() uint64 { _ = "STUB: not implemented"; return 0 }

func (l *raftLog) term(i uint64) (uint64, error) {
	_ = "STUB: not implemented"
	// Check the unstable log first, even before computing the valid term range,
	// which may need to access stable Storage. If we find the entry's term in
	// the unstable log, we know it was in the valid range.
	return 0, nil
}

// The valid term range is [index of dummy entry, last index].

// TODO: return an error instead?

// TODO(bdarnell)

func (l *raftLog) stableTerm(i uint64) (uint64, error) {
	_ = "STUB: not implemented"
	// the valid term range is [index of dummy entry, last index]
	return 0, nil
}

// TODO: return an error instead?

// TODO(bdarnell)

func (l *raftLog) entries(i, maxsize uint64) ([]pb.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// allEntries returns all entries in the log.
func (l *raftLog) allEntries() []pb.Entry { _ = "STUB: not implemented"; return nil }

// try again if there was a racing compaction

// TODO (xiangli): handle error?

// isUpToDate determines if the given (lastIndex,term) log is more up-to-date
// by comparing the index and term of the last entries in the existing logs.
// If the logs have last entries with different terms, then the log with the
// later term is more up-to-date. If the logs end with the same term, then
// whichever log has the larger lastIndex is more up-to-date. If the logs are
// the same, the given log is up-to-date.
func (l *raftLog) isUpToDate(lasti, term uint64) bool { _ = "STUB: not implemented"; return false }

func (l *raftLog) matchTerm(i, term uint64) bool { _ = "STUB: not implemented"; return false }

func (l *raftLog) maybeCommit(maxIndex, term uint64) bool {
	_ = "STUB: not implemented"
	return
	/* && l.zeroTermOnErrCompacted(l.term(maxIndex)) == term */ false
}

func (l *raftLog) maybeCompact(i uint64) { _ = "STUB: not implemented"; return }

func (l *raftLog) restore(s pb.Snapshot) { _ = "STUB: not implemented"; return }

// NOTE: applied and compacted will be reset in raft.advance().

// slice returns a slice of log entries from lo through hi-1, inclusive.
func (l *raftLog) slice(lo, hi, maxSize uint64) ([]pb.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errorlint // it's ok

//nolint:errorlint // it's ok

// TODO(bdarnell)

// check if ents has reached the size limitation

// l.firstIndex <= lo <= hi <= l.firstIndex + len(l.entries)
func (l *raftLog) mustCheckOutOfBounds(lo, hi uint64) error { _ = "STUB: not implemented"; return nil }

// fi := l.firstIndex()
// if lo < fi {
// 	return ErrCompacted
// }

/*fi,*/

func (l *raftLog) zeroTermOnErrCompacted(t uint64, err error) uint64 {
	_ = "STUB: not implemented"
	return 0
}

//nolint:errorlint // it's ok
