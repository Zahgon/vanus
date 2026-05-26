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

//go:generate mockgen -source=replica.go -destination=mock_replica.go -package=segment
package segment

import (
	// standard libraries.
	"context"

	// first-party libraries.
	metapb "github.com/vanus-labs/vanus/api/meta"
	vanus "github.com/vanus-labs/vanus/api/vsr"

	// this project.

	"github.com/vanus-labs/vanus/server/store/block"
	raft "github.com/vanus-labs/vanus/server/store/raft/block"
)

type Replica interface {
	block.Block

	IDStr() string
	Bootstrap(ctx context.Context, blocks []raft.Peer) error
	Close(ctx context.Context) error
	Delete(ctx context.Context) error
	Status() *metapb.SegmentHealthInfo
}

type replica struct {
	id       vanus.ID
	idStr    string
	raw      block.Raw
	appender raft.Appender
}

var _ Replica = (*replica)(nil)

func (r *replica) ID() vanus.ID { _ = "STUB: not implemented"; return *new(vanus.ID) }

func (r *replica) IDStr() string { _ = "STUB: not implemented"; return "" }

func (r *replica) Bootstrap(ctx context.Context, peers []raft.Peer) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *replica) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *replica) Delete(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *replica) Seek(ctx context.Context, index int64, key block.Entry, flag block.SeekKeyFlag) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *replica) Read(ctx context.Context, seq int64, num int) ([]block.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *replica) Append(ctx context.Context, entries []block.Entry, cb block.AppendCallback) {
	_ = "STUB: not implemented"
	return
}

func (r *replica) Status() *metapb.SegmentHealthInfo { _ = "STUB: not implemented"; return nil }

// TODO(james.yin): fill EntLogId and SerializationVersion.

func (s *server) createBlock(ctx context.Context, id vanus.ID, size int64) (Replica, error) {
	_ = "STUB: not implemented"
	// Create block.
	return *new(Replica), nil
}

// Create raft appender.
