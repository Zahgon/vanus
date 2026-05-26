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

package stream

import (
	// standard libraries.
	stdio "io"
	"sync"

	// this project.
	"github.com/vanus-labs/vanus/lib/executor"
	"github.com/vanus-labs/vanus/server/store/io"
	"github.com/vanus-labs/vanus/server/store/io/block"
	"github.com/vanus-labs/vanus/server/store/io/zone"
)

type Stream interface {
	// Zone() zone.Interface
	WriteOffset() int64

	Append(r stdio.Reader, cb io.WriteCallback)
	Sync()
}

type flushTask struct {
	ready bool
	off   int
	cbs   []io.WriteCallback
}

type stream struct {
	s *scheduler
	z zone.Interface

	mu  sync.Mutex
	buf *block.Buffer
	// off is the base offset of Buffer buf.
	off int64
	// dirty is a flag to indicate whether the Buffer buf is dirty.
	dirty   bool
	waiting []io.WriteCallback

	timer PendingID

	pending map[int64]*flushTask

	callbackExecutor executor.ExecuteCloser
}

// Make sure handle implements Stream and io.WriterAt.
var (
	_ Stream      = (*stream)(nil)
	_ io.WriterAt = (*stream)(nil)
	_ PendingTask = (*stream)(nil)
)

func (s *stream) Zone() zone.Interface { _ = "STUB: not implemented"; return *new(zone.Interface) }

func (s *stream) WriteOffset() int64 { _ = "STUB: not implemented"; return 0 }

func (s *stream) Sync() { _ = "STUB: not implemented"; return }

func (s *stream) Append(r stdio.Reader, cb io.WriteCallback) { _ = "STUB: not implemented"; return }

//nolint:errorlint // compare to EOF is ok

func (s *stream) startFlushTimer() { _ = "STUB: not implemented"; return }

func (s *stream) cancelFlushTimer() { _ = "STUB: not implemented"; return }

func (s *stream) OnTimeout(pid PendingID) { _ = "STUB: not implemented"; return }

func (s *stream) flushBuffer(b *block.Buffer, cbs []io.WriteCallback) {
	_ = "STUB: not implemented"
	return
}

//nolint:errorlint // compare to ErrAlreadyFlushed is ok

func (s *stream) flushBlock(b block.Interface, cbs []io.WriteCallback) {
	_ = "STUB: not implemented"
	return
}

//nolint:errorlint // compare to ErrAlreadyFlushed is ok

func (s *stream) onFlushed(base int64, off int, cbs []io.WriteCallback) {
	_ = "STUB: not implemented"
	return
}

// Wait previous block flushed.

// FIXME(james.yin): pass n

// Partial flush.

// Check next block.

// FIXME(james.yin): pass n

func invokeCallbacks(cbs []io.WriteCallback, n int, err error) { _ = "STUB: not implemented"; return }

func (s *stream) WriteAt(b []byte, off int64, so, eo int, cb io.WriteCallback) {
	_ = "STUB: not implemented"
	return
}
