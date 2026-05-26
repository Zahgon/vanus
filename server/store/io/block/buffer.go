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

package block

import (
	// standard libraries.
	"errors"
	stdio "io"
	"os"
	"unsafe"

	// this project.
	"github.com/vanus-labs/vanus/server/store/io"
)

var ErrAlreadyFlushed = errors.New("already flushed")

type flushTask struct {
	b      *Buffer
	writer io.WriterAt
	off    int
	cb     FlushCallback
	next   *flushTask
}

// Buffer is an append-only buffer for block IO.
//
// NOTE: calling Append and Flush is not thread-safe.
type Buffer struct {
	base int64
	buf  []byte
	// wp is write pointer
	wp int
	// fp is flush pointer
	fp int
	// cp is commit pointer
	cp int
	// nf is next flush task
	nf unsafe.Pointer
}

// Make sure Buffer implements Interface.
var _ Interface = (*Buffer)(nil)

func (b *Buffer) Base() int64 { _ = "STUB: not implemented"; return 0 }

func (b *Buffer) Capacity() int { _ = "STUB: not implemented"; return 0 }

func (b *Buffer) Size() int { _ = "STUB: not implemented"; return 0 }

func (b *Buffer) Committed() int { _ = "STUB: not implemented"; return 0 }

func (b *Buffer) Remaining() int { _ = "STUB: not implemented"; return 0 }

func (b *Buffer) remaining(offset int) int { _ = "STUB: not implemented"; return 0 }

func (b *Buffer) Full() bool { _ = "STUB: not implemented"; return false }

func (b *Buffer) Empty() bool { _ = "STUB: not implemented"; return false }

func (b *Buffer) Append(r stdio.Reader) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:errorlint // compare to EOF is ok.

// Flush flushes data in the buffer to storage by writer.
// Invoking callbacks for multiple flushes on the same Buffer is sequence.
func (b *Buffer) Flush(writer io.WriterAt, cb FlushCallback) {
	_ = "STUB: not implemented"
	// TODO(james.yin): Synchronization in concurrency.
	return
}

// Already flush, skip.

// Shortcut if it is final flush.

// partial flush

func (ft *flushTask) invoke(so int) { _ = "STUB: not implemented"; return }

func (ft *flushTask) onWrite(_ int, err error) { _ = "STUB: not implemented"; return }

// NOTE: If it is final flush, DO NOT use b after invoke callback.

// reload

// truncate task list

// TODO(james.yin): optimize goroutine

func (b *Buffer) relocateFlushTask(ft *flushTask) (unsafe.Pointer, *flushTask) {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer), nil
}

func (ft *flushTask) invokeCallback(err error) { _ = "STUB: not implemented"; return }

func (b *Buffer) RecoverFromFile(f *os.File, at int64, committed int, direct bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Fill zero.
