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

//go:generate mockgen -source=server.go -destination=mock_server.go -package=segment
package segment

import (
	// standard libraries.
	"context"

	// first-party libraries.

	metapb "github.com/vanus-labs/vanus/api/meta"

	// this project.

	"github.com/vanus-labs/vanus/server/store/config"
)

func (s *server) Initialize(ctx context.Context) error {
	_ = "STUB: not implemented"
	// TODO(james.yin): how to organize block engine?
	return nil
}

// Recover replicas.

// Fetch block information in volume from controller, and make state up to date.

func (s *server) loadVSBEngine(_ context.Context, cfg config.VSB) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *server) initRaftEngine(ctx context.Context, cfg config.Raft) error {
	_ = "STUB: not implemented"
	// TODO(james.yin): move metaStore and offsetStore to raftEngine?
	return nil
}

// recover recovers replicas.
func (s *server) recover(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Recover replicas.

// TODO(james.yin): remain this block?

func (s *server) reconcileBlocks(_ context.Context) error {
	_ = "STUB: not implemented"
	// TODO(james.yin): Fetch block information in volume from controller, and make state up to date.
	return nil
}

func (s *server) registerSelf(ctx context.Context) error {
	_ = "STUB: not implemented"
	// TODO(james.yin): pass information of blocks.
	return nil
}

// FIXME(james.yin): some blocks may not be bound to segment.

// No block in the volume of this server.

func (s *server) reconcileSegments(ctx context.Context, segments map[uint64]*metapb.Segment) {
	_ = "STUB: not implemented"
	return
}

// Don't use address to compare.

// FIXME(james.yin): multiple blocks of same segment in this server.

// TODO(james.yin): no my block

func (s *server) registerReplicas(ctx context.Context, segment *metapb.Segment) {
	_ = "STUB: not implemented"
	return
}
