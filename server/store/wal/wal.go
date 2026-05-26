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

package wal

import (
	// standard libraries.
	"context"
	"errors"
	"sync"

	// first-party libraries.

	// this project.
	"github.com/vanus-labs/vanus/lib/container/conque/blocking"
	"github.com/vanus-labs/vanus/server/store/io/engine"
	"github.com/vanus-labs/vanus/server/store/io/stream"
	"github.com/vanus-labs/vanus/server/store/io/zone/segmentedfile"
)

var (
	ErrClosed          = errors.New("wal: closed")
	ErrNotFoundLogFile = errors.New("wal: not found log file")
)

type Range struct {
	SO int64
	EO int64
}

type AppendOneCallback = func(Range, error)

type AppendCallback = func([]Range, error)

// WAL is write-ahead log.
// All append tasks be processed in WAL.runAppend by a single goroutine.
type WAL struct {
	sf *segmentedfile.SegmentedFile
	s  stream.Stream

	engine    engine.Interface
	scheduler stream.Scheduler

	blockSize int

	appendQ blocking.Queue[*appender]

	appendWg sync.WaitGroup

	doneC chan struct{}
}

func Open(ctx context.Context, dir string, opts ...Option) (*WAL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func open(ctx context.Context, dir string, cfg config) (*WAL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check wal entries from pos.

// Skip padding.

func (w *WAL) Dir() string { _ = "STUB: not implemented"; return "" }

func (w *WAL) Close() { _ = "STUB: not implemented"; return }

func (w *WAL) doClose() { _ = "STUB: not implemented"; return }

func (w *WAL) Wait() { _ = "STUB: not implemented"; return }

func (w *WAL) AppendOne(ctx context.Context, entry []byte, cb AppendOneCallback) {
	_ = "STUB: not implemented"
	return
}

// Append appends entries to WAL.
func (w *WAL) Append(ctx context.Context, entries [][]byte, cb AppendCallback) {
	_ = "STUB: not implemented"
	return
}

func (w *WAL) append(ctx context.Context, entries [][]byte, direct bool, cb AppendCallback) {
	_ = "STUB: not implemented"
	// Check entries.
	return
}

// TODO(james.yin): invoke callback in another goroutine.

func (w *WAL) runAppend() { _ = "STUB: not implemented"; return }

// Invoke remained tasks in w.appendQ.

func (w *WAL) Compact(_ context.Context, off int64) error { _ = "STUB: not implemented"; return nil }

type appendResult struct {
	ranges []Range
	err    error
}

type appendFuture chan appendResult

func newAppendFuture() appendFuture { _ = "STUB: not implemented"; return *new(appendFuture) }

func (af appendFuture) onAppended(ranges []Range, err error) { _ = "STUB: not implemented"; return }

func (af appendFuture) wait() ([]Range, error) { _ = "STUB: not implemented"; return nil, nil }

func Append(ctx context.Context, w *WAL, entries [][]byte) ([]Range, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DirectAppend(ctx context.Context, w *WAL, entries [][]byte) ([]Range, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AppendOne(ctx context.Context, w *WAL, entry []byte) (Range, error) {
	_ = "STUB: not implemented"
	return *new(Range), nil
}

func DirectAppendOne(ctx context.Context, w *WAL, entry []byte) (Range, error) {
	_ = "STUB: not implemented"
	return *new(Range), nil
}
