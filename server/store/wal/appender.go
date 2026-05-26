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
	"io"

	// this project.
	"github.com/vanus-labs/vanus/server/store/wal/record"
)

func (w *WAL) newAppender(ctx context.Context, entries [][]byte, direct bool, callback AppendCallback) *appender {
	_ = "STUB: not implemented"
	return nil
}

type appender struct {
	w       *WAL
	entries [][]byte
	records []record.Record
	padding int
	i, j    int

	ranges []Range

	ctx      context.Context
	direct   bool
	callback AppendCallback
}

// Make sure Data implements io.Reader.
var _ io.Reader = (*appender)(nil)

func (a *appender) invoke() { _ = "STUB: not implemented"; return }

// metrics.WALEntryWriteCounter.Add(float64(len(entries)))
// metrics.WALEntryWriteSizeCounter.Add(float64(entrySize))
// metrics.WALRecordWriteCounter.Add(float64(recordCount))
// metrics.WALRecordWriteSizeCounter.Add(float64(recordSize))

func (a *appender) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Release memory.

func (a *appender) onAppended(_ int, err error) { _ = "STUB: not implemented"; return }
