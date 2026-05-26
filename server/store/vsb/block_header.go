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

package vsb

import (
	// standard libraries.
	"context"
	"hash/crc32"
	"os"
	// this project.
)

const (
	headerBlockSize = 4 * 1024
	headerSize      = 4 + 4 + 4 + 4 + 4 + 1 + 1 + 2 + 8 + 8 + 4 + 2

	magicOffset       = 0
	crcOffset         = 4
	flagsOffset       = 8
	breakFlagsOffset  = 12
	dataOffsetOffset  = 16
	stateOffset       = 20
	indexSizeOffset   = 22
	capacityOffset    = 24
	entryLengthOffset = 32
	entryNumOffset    = 40
	indexOffsetOffset = 44
)

var (
	crc32q      = crc32.MakeTable(crc32.Castagnoli)
	emptyHeader = make([]byte, headerBlockSize)
)

type Header struct {
	Magic       uint32
	Crc         uint32
	Flags       uint32
	BreakFlags  uint32
	DataOffset  uint32
	State       uint8
	_pad        uint8 //nolint:unused // padding
	IndexSize   uint16
	Capacity    uint64
	EntryLength uint64
	EntryNum    uint32
	IndexOffset uint16
}

func LoadHeader(f *os.File) (hdr Header, err error) {
	_ = "STUB: not implemented"
	return *new(Header), nil
}

func (b *vsBlock) persistHeader(_ context.Context, m meta) error {
	_ = "STUB: not implemented"
	return nil
}

// magic
// flags
// break flags
// data offset
// state

// index size
// capacity
// entry length
// entry number
// index offset

// crc

func (b *vsBlock) loadHeader(_ context.Context) error { _ = "STUB: not implemented"; return nil }
