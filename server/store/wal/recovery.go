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
	"bytes"
	"errors"

	// third-party project.

	// this project.

	"github.com/vanus-labs/vanus/server/store/io/zone/segmentedfile"
	"github.com/vanus-labs/vanus/server/store/wal/record"
)

type OnEntryCallback = func(entry []byte, r Range) error

var (
	ErrOutOfRange = errors.New("WAL: out of range")
	errEndOfLog   = errors.New("WAL: end of log")
)

func scanLogEntries(sf *segmentedfile.SegmentedFile, blockSize int, from int64, cb OnEntryCallback) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TODO(james.yin): has empty log file(s).
// if i != len(s.stream)-1 {
// 	panic("has empty log file")
// }

// TODO(james.yin): Has incomplete entry, truncate it.

type scanner struct {
	blockSize int64
	buf       []byte
	buffer    *bytes.Buffer
	last      record.Type
	eo        int64 // end offset of entry
	from      int64
	cb        OnEntryCallback
}

func (sc *scanner) scanSegmentFile(s *segmentedfile.Segment) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// TODO(james.yin): handle parse error

// no new record

// TODO(james.yin): check crc

func (sc *scanner) firstBlockOffset(s *segmentedfile.Segment) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (sc *scanner) firstRecordOffset(so int64) int64 { _ = "STUB: not implemented"; return 0 }

func onRecord(ctx *scanner, r record.Record, eo int64) error { _ = "STUB: not implemented"; return nil }

// TODO(james.yin): unexpected state

// TODO(james.yin): unexpected state

// TODO(james.yin): unexpected state

// TODO(james.yin): unexpected state

func noopOnEntry(_ []byte, r Range) error {
	_ = "STUB: not implemented" //nolint:revive // ok
	return nil
}
