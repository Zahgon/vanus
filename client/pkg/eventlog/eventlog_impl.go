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

	// third-party libraries.

	// first-party libraries.
	"github.com/vanus-labs/vanus/api/cloudevents"
	"github.com/vanus-labs/vanus/pkg/observability/tracing"

	// this project.
	el "github.com/vanus-labs/vanus/client/internal/eventlog"
	"github.com/vanus-labs/vanus/client/pkg/record"
)

const (
	defaultRetryTimes = 10
	pollingThreshold  = 200 // in milliseconds.
	pollingPostSpan   = 100 // in milliseconds.
)

func NewEventlog(cfg *el.Config) Eventlog { _ = "STUB: not implemented"; return *new(Eventlog) }

type eventlog struct {
	cfg         *el.Config
	nameService *el.NameService

	writableWatcher *WritableSegmentWatcher
	writableSegment *segment
	writableMu      sync.RWMutex

	readableWatcher  *ReadableSegmentsWatcher
	readableSegments []*segment
	readableMu       sync.RWMutex
	tracer           *tracing.Tracer
}

// make sure eventlog implements Eventlog.
var _ Eventlog = (*eventlog)(nil)

func (l *eventlog) ID() uint64 { _ = "STUB: not implemented"; return 0 }

func (l *eventlog) Close(ctx context.Context) { _ = "STUB: not implemented"; return }

func (l *eventlog) Writer() LogWriter { _ = "STUB: not implemented"; return *new(LogWriter) }

func (l *eventlog) Reader(cfg ReaderConfig) LogReader {
	_ = "STUB: not implemented"
	return *new(LogReader)
}

func (l *eventlog) EarliestOffset(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (l *eventlog) LatestOffset(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (l *eventlog) Length(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	// TODO(kai.jiangkai)
	return 0, nil
}

func (l *eventlog) QueryOffsetByTime(ctx context.Context, timestamp int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil

	// get all segments
}

// the target offset maybe in newer segment, refresh immediately

func (l *eventlog) CheckHealth(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (l *eventlog) updateWritableSegment(ctx context.Context, r *record.Segment) {
	_ = "STUB: not implemented"
	return
}

func (l *eventlog) selectWritableSegment(ctx context.Context) (*segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *eventlog) fetchWritableSegment(ctx context.Context) *segment {
	_ = "STUB: not implemented"
	return nil
}

// refresh

func (l *eventlog) refreshWritableSegment(ctx context.Context) { _ = "STUB: not implemented"; return }

func (l *eventlog) updateReadableSegments(ctx context.Context, rs []*record.Segment) {
	_ = "STUB: not implemented"
	return
}

// TODO: find

// FIXME: create or update segment failed

func (l *eventlog) selectReadableSegment(ctx context.Context, offset int64) (*segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: make sure the segments are in order.

func (l *eventlog) fetchReadableSegments(ctx context.Context) []*segment {
	_ = "STUB: not implemented"
	return nil
}

// refresh

func (l *eventlog) refreshReadableSegments(ctx context.Context) { _ = "STUB: not implemented"; return }

var _ LogWriter = &logWriter{}

// logWriter is the writer of eventlog.
//
// Append is thread-safety.
type logWriter struct {
	elog *eventlog
	cur  *segment
	mu   sync.RWMutex
}

func (w *logWriter) Append(ctx context.Context, events *cloudevents.CloudEventBatch) (offs []int64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *logWriter) Log() Eventlog { _ = "STUB: not implemented"; return *new(Eventlog) }

func (w *logWriter) Close(ctx context.Context) {
	_ = "STUB: not implemented"
	// TODO: by jiangkai, 2022.10.19
	return
}

func (w *logWriter) doAppend(ctx context.Context, event *cloudevents.CloudEventBatch) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *logWriter) selectWritableSegment(ctx context.Context) (*segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// double check

type logReader struct {
	elog *eventlog
	pos  int64
	cur  *segment
	cfg  ReaderConfig
}

func (r *logReader) Log() Eventlog { _ = "STUB: not implemented"; return *new(Eventlog) }

func (r *logReader) Close(ctx context.Context) {
	_ = "STUB: not implemented"
	// TODO: by jiangkai, 2022.10.19
	return
}

func (r *logReader) Read(ctx context.Context, size int16) (*cloudevents.CloudEventBatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *logReader) pollingTimeout(ctx context.Context) int64 { _ = "STUB: not implemented"; return 0 }

func (r *logReader) switchSegment(ctx context.Context) bool {
	_ = "STUB: not implemented"
	// switch to next segment
	return false
}

func (r *logReader) Seek(ctx context.Context, offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	// TODO
	return 0, nil
}
