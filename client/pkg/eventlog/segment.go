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
	"sync"
	"sync/atomic"
	"time"

	// third-party libraries.

	// first-party libraries.
	"github.com/vanus-labs/vanus/api/cloudevents"
	"github.com/vanus-labs/vanus/pkg/observability/tracing"

	// this project.
	"github.com/vanus-labs/vanus/client/pkg/record"
)

func newSegment(ctx context.Context, r *record.Segment, towrite bool) (*segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newBlockExt(ctx context.Context, r *record.Segment, leaderOnly bool) (*block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type segment struct {
	id               uint64
	startOffset      int64
	endOffset        atomic.Int64
	writable         atomic.Bool
	firstEventBornAt time.Time
	lastEventBornAt  time.Time

	prefer *block
	mu     sync.RWMutex
	tracer *tracing.Tracer
}

func (s *segment) ID() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *segment) StartOffset() int64 { _ = "STUB: not implemented"; return 0 }

func (s *segment) EndOffset() int64 { _ = "STUB: not implemented"; return 0 }

func (s *segment) Writable() bool { _ = "STUB: not implemented"; return false }

func (s *segment) SetNotWritable() { _ = "STUB: not implemented"; return }

func (s *segment) Close(ctx context.Context) { _ = "STUB: not implemented"; return }

func (s *segment) Update(ctx context.Context, r *record.Segment, towrite bool) error {
	_ = "STUB: not implemented"
	// When a segment become read-only, the end offset needs to be set to the real value.
	// TODO(wenfeng) data race?
	return nil
}

func (s *segment) Append(ctx context.Context, event *cloudevents.CloudEventBatch) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *segment) Read(ctx context.Context, from int64, size int16, pollingTimeout uint32) (*cloudevents.CloudEventBatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: cached read

func (s *segment) preferSegmentBlock() *block { _ = "STUB: not implemented"; return nil }

func (s *segment) setPreferSegmentBlock(prefer *block) { _ = "STUB: not implemented"; return }

func (s *segment) LookupOffset(ctx context.Context, t time.Time) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *segment) CheckHealth(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// FIXME: no leader

// TODO: maybe corrupted metadata
