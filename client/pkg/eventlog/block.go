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

package eventlog

import (
	// standard libraries.
	"context"
	"time"

	// first-party libraries.
	"github.com/vanus-labs/vanus/api/cloudevents"
	metapb "github.com/vanus-labs/vanus/api/meta"

	// this project.
	"github.com/vanus-labs/vanus/client/internal/store"
	"github.com/vanus-labs/vanus/client/pkg/record"
)

func newBlock(ctx context.Context, r *record.Block) (*block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type block struct {
	id    uint64
	store *store.BlockStore
}

func (b *block) Close(ctx context.Context) { _ = "STUB: not implemented"; return }

func (b *block) LookupOffset(ctx context.Context, t time.Time) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *block) Append(ctx context.Context, event *cloudevents.CloudEventBatch) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *block) Read(ctx context.Context, offset int64, size int16, pollingTimeout uint32) (*cloudevents.CloudEventBatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// doRead

func (b *block) Describe(ctx context.Context) (*metapb.SegmentHealthInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
