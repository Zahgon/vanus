// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

import (
	// standard libraries.
	"context"
	"time"

	// third-party libraries.

	// first-party libraries.
	"github.com/vanus-labs/vanus/api/cloudevents"
	metapb "github.com/vanus-labs/vanus/api/meta"
	"github.com/vanus-labs/vanus/pkg/observability/tracing"

	// this project.
	"github.com/vanus-labs/vanus/client/internal/net/rpc"
	"github.com/vanus-labs/vanus/client/pkg/primitive"
)

func newBlockStore(endpoint string) (*BlockStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: check error

type BlockStore struct {
	primitive.RefCount
	client rpc.Client
	tracer *tracing.Tracer
}

func (s *BlockStore) Endpoint() string { _ = "STUB: not implemented"; return "" }

func (s *BlockStore) Close() { _ = "STUB: not implemented"; return }

func (s *BlockStore) Read(
	ctx context.Context, block uint64, offset int64, size int16, pollingTimeout uint32,
) (*cloudevents.CloudEventBatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *BlockStore) LookupOffset(ctx context.Context, blockID uint64, t time.Time) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *BlockStore) Append(ctx context.Context, block uint64, events *cloudevents.CloudEventBatch) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *BlockStore) Describe(ctx context.Context, block uint64) (*metapb.SegmentHealthInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
