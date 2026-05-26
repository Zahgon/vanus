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

package eventbus

import (
	// standard libraries.
	"context"
	"sync"

	// third-party libraries.
	"github.com/scylladb/go-set/u64set"

	// first-party libraries.
	"github.com/vanus-labs/vanus/api/cloudevents"
	"github.com/vanus-labs/vanus/pkg/observability/tracing"

	// this project.
	eb "github.com/vanus-labs/vanus/client/internal/eventbus"
	"github.com/vanus-labs/vanus/client/pkg/api"
	"github.com/vanus-labs/vanus/client/pkg/eventlog"
)

func NewEventbus(cfg *eb.Config) *eventbus { _ = "STUB: not implemented"; return nil }

type eventbus struct {
	cfg         *eb.Config
	nameService *eb.NameService

	writableWatcher *WritableLogsWatcher
	writableLogSet  *u64set.Set
	writableLogs    map[uint64]eventlog.Eventlog
	writableMu      sync.RWMutex
	writableState   error

	readableWatcher *ReadableLogsWatcher
	readableLogSet  *u64set.Set
	readableLogs    map[uint64]eventlog.Eventlog
	readableMu      sync.RWMutex
	readableState   error

	tracer *tracing.Tracer
}

// make sure eventbus implements api.Eventbus.
var _ api.Eventbus = (*eventbus)(nil)

func (b *eventbus) defaultWriteOptions() *api.WriteOptions { _ = "STUB: not implemented"; return nil }

func (b *eventbus) defaultReadOptions() *api.ReadOptions { _ = "STUB: not implemented"; return nil }

func (b *eventbus) Writer(opts ...api.WriteOption) api.BusWriter {
	_ = "STUB: not implemented"
	return *new(api.BusWriter)
}

func (b *eventbus) Reader(opts ...api.ReadOption) api.BusReader {
	_ = "STUB: not implemented"
	return *new(api.BusReader)
}

func (b *eventbus) GetLog(ctx context.Context, logID uint64, opts ...api.LogOption) (api.Eventlog, error) {
	_ = "STUB: not implemented"
	return *new(api.Eventlog), nil
}

func (b *eventbus) ListLog(ctx context.Context, opts ...api.LogOption) ([]api.Eventlog, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *eventbus) CheckHealth(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *eventbus) ID() uint64 { _ = "STUB: not implemented"; return 0 }

func (b *eventbus) Close(ctx context.Context) { _ = "STUB: not implemented"; return }

func (b *eventbus) getWritableState() error { _ = "STUB: not implemented"; return nil }

func (b *eventbus) setWritableState(err error) { _ = "STUB: not implemented"; return }

func (b *eventbus) isNeedUpdateWritableLogs(err error) bool {
	_ = "STUB: not implemented"
	return false
}

func (b *eventbus) updateWritableLogs(ctx context.Context, re *WritableLogsResult) {
	_ = "STUB: not implemented"
	return
}

func (b *eventbus) setWritableLogs(s *u64set.Set, lws map[uint64]eventlog.Eventlog) {
	_ = "STUB: not implemented"
	return
}

func (b *eventbus) getWritableLog(ctx context.Context, logID uint64) eventlog.Eventlog {
	_ = "STUB: not implemented"
	return *new(eventlog.Eventlog)
}

func (b *eventbus) refreshWritableLogs(ctx context.Context) { _ = "STUB: not implemented"; return }

func (b *eventbus) getReadableState() error { _ = "STUB: not implemented"; return nil }

func (b *eventbus) setReadableState(err error) { _ = "STUB: not implemented"; return }

func (b *eventbus) isNeedUpdateReadableLogs(err error) bool {
	_ = "STUB: not implemented"
	return false
}

func (b *eventbus) updateReadableLogs(ctx context.Context, re *ReadableLogsResult) {
	_ = "STUB: not implemented"
	return
}

func (b *eventbus) setReadableLogs(s *u64set.Set, lws map[uint64]eventlog.Eventlog) {
	_ = "STUB: not implemented"
	return
}

func (b *eventbus) getReadableLog(ctx context.Context, logID uint64) eventlog.Eventlog {
	_ = "STUB: not implemented"
	return *new(eventlog.Eventlog)
}

func (b *eventbus) refreshReadableLogs(ctx context.Context) { _ = "STUB: not implemented"; return }

type busWriter struct {
	ebus   *eventbus
	opts   *api.WriteOptions
	tracer *tracing.Tracer
}

var _ api.BusWriter = (*busWriter)(nil)

func (w *busWriter) Append(ctx context.Context, events *cloudevents.CloudEventBatch, opts ...api.WriteOption) (eids []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 1. pick a writer of eventlog

// 2. append the event to the eventlog

func (w *busWriter) Bus() api.Eventbus { _ = "STUB: not implemented"; return *new(api.Eventbus) }

func (w *busWriter) pickWritableLog(ctx context.Context, opts *api.WriteOptions) (eventlog.LogWriter, error) {
	_ = "STUB: not implemented"
	return *new(eventlog.LogWriter), nil
}

func genEventID(logID uint64, off int64) string { _ = "STUB: not implemented"; return "" }

type busReader struct {
	ebus   *eventbus
	opts   *api.ReadOptions
	tracer *tracing.Tracer
}

var _ api.BusReader = (*busReader)(nil)

func (r *busReader) Read(ctx context.Context, opts ...api.ReadOption) (events *cloudevents.CloudEventBatch, off int64, logid uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, nil
}

// 1. pick a reader of eventlog

// TODO(jiangkai): refactor eventlog interface to avoid seek every time, by jiangkai, 2022.10.24

// 2. read the event to the eventlog

func (r *busReader) Bus() api.Eventbus { _ = "STUB: not implemented"; return *new(api.Eventbus) }

func (r *busReader) pickReadableLog(ctx context.Context, opts *api.ReadOptions) (eventlog.LogReader, error) {
	_ = "STUB: not implemented"
	return *new(eventlog.LogReader), nil
}
