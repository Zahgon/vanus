// Copyright 2019 The etcd Authors
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

package rafttest

import (
	"testing"

	"github.com/cockroachdb/datadriven"

	"github.com/vanus-labs/vanus/pkg/raft"
	pb "github.com/vanus-labs/vanus/pkg/raft/raftpb"
)

func (env *InteractionEnv) handleAddNodes(t *testing.T, d datadriven.TestData) error {
	_ = "STUB: not implemented"
	return nil
}

type snapOverrideStorage struct {
	Storage
	snapshotOverride func() (pb.Snapshot, error)
}

func (s snapOverrideStorage) Snapshot() (pb.Snapshot, error) {
	_ = "STUB: not implemented"
	return *new(pb.Snapshot), nil
}

var _ raft.Storage = snapOverrideStorage{}

// AddNodes adds n new nodes initializes from the given snapshot (which may be
// empty). They will be assigned consecutive IDs.
func (env *InteractionEnv) AddNodes(n int, snap pb.Snapshot) error {
	_ = "STUB: not implemented"
	return nil
}

// When you ask for a snapshot, you get the most recent snapshot.
//
// TODO(tbg): this is sort of clunky, but MemoryStorage itself will
// give you some fixed snapshot and also the snapshot changes
// whenever you compact the logs and vice versa, so it's all a bit
// awkward to use.

// NB: we could make this work with 1, but MemoryStorage just
// doesn't play well with that and it's not a loss of generality.

// At the time of writing and for *MemoryStorage, applying a
// snapshot also truncates appropriately, but this would change with
// other storage engines potentially.

// This could be supported but then we need to do more work
// translating back and forth -- not worth it.

// TODO(tbg): allow a more general Storage, as long as it also allows
// us to apply snapshots, append entries, and update the HardState.
