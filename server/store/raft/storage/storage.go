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
	"sync"

	// first-party libraries.
	vanus "github.com/vanus-labs/vanus/api/vsr"
	"github.com/vanus-labs/vanus/pkg/raft"

	// this project.
	"github.com/vanus-labs/vanus/lib/executor"
	"github.com/vanus-labs/vanus/server/store/meta"
)

type Storage struct {
	// Protects access to raft states.
	mu sync.RWMutex

	nodeID vanus.ID

	stateStorage
	logStorage
	snapshotStorage

	// AppendExecutor is the Executor that executes Append, postAppend, and Compact.
	AppendExecutor executor.Executor
}

// Make sure Storage implements raft.Storage.
var _ raft.Storage = (*Storage)(nil)

// NewStorage creates an empty Storage.
func NewStorage(
	ctx context.Context, nodeID vanus.ID, wal *WAL, stateStore *meta.SyncStore, hintStore *meta.AsyncStore,
	snapOp SnapshotOperator,
) (*Storage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newStorage(
	nodeID vanus.ID, wal *WAL, stateStore *meta.SyncStore, hintStore *meta.AsyncStore, snapOp SnapshotOperator,
) *Storage {
	_ = "STUB: not implemented"
	return nil
}

// When starting from scratch populate the list with a dummy entry at term zero.

// Delete discards all data of Storage.
// NOTE: waiting for inflight append calls is the responsibility of the caller.
func (s *Storage) Delete(ctx context.Context) { _ = "STUB: not implemented"; return }

// TODO(james.yin): handle error.

// Clean metadata in stateStore and hintStore.
