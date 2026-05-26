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

//go:generate mockgen -source=snapshot_storage.go -destination=mock_snapshot_storage.go -package=storage
package storage

import (
	// standard libraries.
	"context"

	// third-party libraries.

	// first-party libraries.

	"github.com/vanus-labs/vanus/pkg/raft/raftpb"
	// this project.
)

type SnapshotOperator interface {
	GetSnapshot(index uint64) ([]byte, error)
	ApplySnapshot(data []byte) error
}

type snapshotStorage struct {
	snapOp SnapshotOperator
}

func (ss *snapshotStorage) SetSnapshotOperator(op SnapshotOperator) {
	_ = "STUB: not implemented"

	// Snapshot returns the most recent snapshot.
	// If snapshot is temporarily unavailable, it should return ErrSnapshotTemporarilyUnavailable,
	// so raft state machine could know that Storage needs some time to prepare
	// snapshot and call Snapshot later.
	return
}

func (s *Storage) Snapshot() (raftpb.Snapshot, error) {
	_ = "STUB: not implemented"
	return *new(raftpb.Snapshot), nil
}

// ApplySnapshot overwrites the contents of this Storage object with
// those of the given snapshot.
func (s *Storage) ApplySnapshot(ctx context.Context, snap raftpb.Snapshot) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle check for old snapshot being applied.
