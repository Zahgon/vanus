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
	"time"

	// first-party libraries.
	vanus "github.com/vanus-labs/vanus/api/vsr"

	// this project.
	"github.com/vanus-labs/vanus/server/store/meta"
)

const (
	defaultCompactInterval = 30 * time.Second
)

var walCompactKey = []byte("wal/compact")

var ErrClosed = errors.New("WAL: closed")

// Compact discards all log entries prior to compactIndex.
// It is the application's responsibility to not attempt to compact an index greater than raftLog.applied.
func (s *Storage) Compact(ctx context.Context, i uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// FIXME(james.yin): error

// stable remaining

// Save compact information to dummy entry.

// Copy remained entries.

// NOTE: `sr` MUST NOT greater than `remaining` (sr <= remaining).

// Compact WAL.

// Reset log entries and offsets.

func (w *WAL) tryCompact(ctx context.Context, nodeID vanus.ID, offset, last, tail int64, index, term uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *WAL) markBarrier(ctx context.Context, nodeID vanus.ID, offset int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *WAL) removeBarrier(ctx context.Context, nodeID vanus.ID, offset int64) error {
	_ = "STUB: not implemented"
	return nil
}

type compactTask struct {
	nodeID             vanus.ID
	offset, last, tail int64
	info               compactInfo
}

func (t *compactTask) compact(w *WAL, cc *compactContext) {
	_ = "STUB: not implemented"
	// node is deleted.
	return
}

// Discard last barrier.

// Set new barrier.

// Set compaction info.

var emptyCompact = make([]byte, 16)

func (w *WAL) addNode(ctx context.Context, nodeID vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *WAL) removeNode(ctx context.Context, nodeID vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// Prevent compact on node.

// TODO(james.yin): handle error.

// Clean node to delete WAL.

func (w *WAL) recoverNode(nodeID vanus.ID, offset int64) { _ = "STUB: not implemented"; return }

type compactInfo struct {
	index, term uint64
}

func (ci *compactInfo) empty() bool { _ = "STUB: not implemented"; return false }

type logCompactInfos map[vanus.ID]compactInfo

// Make sure logCompactInfos implements meta.Ranger.
var _ meta.Ranger = (logCompactInfos)(nil)

func (i logCompactInfos) Range(cb meta.RangeCallback) error { _ = "STUB: not implemented"; return nil }

type compactMeta struct {
	infos  logCompactInfos
	offset int64
}

// Make sure compactMeta implements meta.Ranger.
var _ meta.Ranger = (*compactMeta)(nil)

func (m *compactMeta) Range(cb meta.RangeCallback) error { _ = "STUB: not implemented"; return nil }

type compactContext struct {
	tail      int64
	compacted int64
	toCompact int64
	infos     logCompactInfos
}

func loadCompactContext(stateStore *meta.SyncStore) *compactContext {
	_ = "STUB: not implemented"
	return nil
}

func (c *compactContext) stale() bool { _ = "STUB: not implemented"; return false }

func (c *compactContext) sync(ctx context.Context, stateStore *meta.SyncStore) bool {
	_ = "STUB: not implemented"
	return false
}

type compactFunc func(*WAL, *compactContext)

type compactJob struct {
	fn compactFunc
}

func (j *compactJob) invoke(w *WAL, cc *compactContext) {
	_ = "STUB: not implemented"

	// runCompact processes all compact jobs in a single goroutine.
	return
}

func (w *WAL) runCompact() { _ = "STUB: not implemented"; return }

func (w *WAL) doCompact(ctx context.Context, cc *compactContext) { _ = "STUB: not implemented"; return }

// Store compacted info and offset.

// Compact underlying WAL.

// reconcileBarrier scans barriers and calculates compactContext.toCompact.
func (w *WAL) reconcileBarrier(cc *compactContext) { _ = "STUB: not implemented"; return }

//  No log entry in WAL.

// Remove barrier if node is deleted.

// dispatchCompactJob dispatches a compact job to the compact goroutine.
func (w *WAL) dispatchCompactJob(ctx context.Context, job compactJob) error {
	_ = "STUB: not implemented"
	// NOTE: no panic, avoid unlocking with defer.
	return nil
}

// dispatchCompactTask dispatches a compact task to the compact goroutine.
func (w *WAL) dispatchCompactTask(ctx context.Context, task compactFunc) error {
	_ = "STUB: not implemented"
	return nil
}

type awaitableCompactFunc func(*WAL, *compactContext, chan<- error)

// invokeCompactTask invokes a compact task and waits for its completion.
func (w *WAL) invokeCompactTask(ctx context.Context, task awaitableCompactFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func CompactKey(id uint64) string { _ = "STUB: not implemented"; return "" }
