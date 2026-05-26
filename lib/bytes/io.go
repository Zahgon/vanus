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

package bytes

import (
	"io"
)

var DummyWriter = &nopWriter{}

type MarkScanner struct {
	Buf []byte
	off int
}

// Make sure MarkScanner implements io.ByteScanner.
var _ io.ByteScanner = (*MarkScanner)(nil)

func NewMarkScanner(b []byte) *MarkScanner { _ = "STUB: not implemented"; return nil }

func (s *MarkScanner) empty() bool { _ = "STUB: not implemented"; return false }

func (s *MarkScanner) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *MarkScanner) UnreadByte() error { _ = "STUB: not implemented"; return nil }

func (s *MarkScanner) Mark(off int) int { _ = "STUB: not implemented"; return 0 }

func (s *MarkScanner) Since(mark int, off int) []byte { _ = "STUB: not implemented"; return nil }

func (s *MarkScanner) From(mark int) []byte { _ = "STUB: not implemented"; return nil }

func (s *MarkScanner) Resume(mark int) error { _ = "STUB: not implemented"; return nil }

func ScannedBytes(s *MarkScanner, mark int, eof bool) []byte { _ = "STUB: not implemented"; return nil }

type nopWriter struct{}

// Make sure nopWriter implements io.Write and io.ByteWriter.
var (
	_ io.Writer     = (*nopWriter)(nil)
	_ io.ByteWriter = (*nopWriter)(nil)
)

func (w *nopWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *nopWriter) WriteByte(_ byte) error { _ = "STUB: not implemented"; return nil }

type CopyOnDiffWriter struct {
	Buf []byte
	new []byte
	off int
}

// Make sure CopyOnDiffWriter implements io.ByteWriter.
var _ io.ByteWriter = (*CopyOnDiffWriter)(nil)

func (w *CopyOnDiffWriter) WriteByte(c byte) error { _ = "STUB: not implemented"; return nil }

func (w *CopyOnDiffWriter) Bytes() []byte { _ = "STUB: not implemented"; return nil }

type LastByteWriter interface {
	io.Writer

	LastByte() (byte, bool)
	TruncateLastByte()
}
