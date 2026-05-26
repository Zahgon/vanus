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

package codec

import (
	// standard libraries.

	"hash/crc32"
	"io"

	// first-party libraries.

	"github.com/vanus-labs/vanus/pkg/observability/tracing"

	// this project.
	"github.com/vanus-labs/vanus/server/store/block"
)

const (
	packetLengthSize = 4
	packetCRCSize    = 4

	packetFooterSize = packetLengthSize + packetCRCSize
	packetMetaSize   = packetLengthSize + packetFooterSize

	packetLengthOffset  = 0
	packetPayloadOffset = packetLengthOffset + packetLengthSize
	packetCRCOffset     = packetLengthSize
)

var crc32q = crc32.MakeTable(crc32.Castagnoli)

type PacketDataEncoder interface {
	Size(entry block.Entry) int
	MarshalTo(entry block.Entry, buf []byte) (int, error)
}

type packetEncoder struct {
	pde    PacketDataEncoder
	tracer *tracing.Tracer
}

// Make sure packetEncoder implements EntryEncoder.
var _ EntryEncoder = (*packetEncoder)(nil)

func (e *packetEncoder) Size(entry block.Entry) int { _ = "STUB: not implemented"; return 0 }

func (e *packetEncoder) MarshalTo(entry block.Entry, buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type PacketDataDecoder interface {
	Unmarshal(data []byte) (block.Entry, error)
}

type packetDecoder struct {
	pdd      PacketDataDecoder
	checkCRC bool
}

var _ EntryDecoder = (*packetDecoder)(nil)

func (pd *packetDecoder) Unmarshal(data []byte) (int, block.Entry, error) {
	_ = "STUB: not implemented"
	return 0, *new(block.Entry), nil
}

func (pd *packetDecoder) UnmarshalLast(data []byte) (int, block.Entry, error) {
	_ = "STUB: not implemented"
	return 0, *new(block.Entry), nil
}

func (pd *packetDecoder) UnmarshalReader(r io.ReadSeeker) (int, block.Entry, error) {
	_ = "STUB: not implemented"
	return 0, *new(block.Entry), nil
}

// Seek to footer.

// Read length in footer, and check if it matches.

// Seek to start.

// FIXME(james.yin): persisted data is partial.

func (pd *packetDecoder) doUnmarshal(length int, data []byte) (int, block.Entry, error) {
	_ = "STUB: not implemented"
	return 0, *new(block.Entry), nil
}
