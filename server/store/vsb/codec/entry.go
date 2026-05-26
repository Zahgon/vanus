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

	"time"

	// this project.
	"github.com/vanus-labs/vanus/server/store/block"
)

const (
	offsetOffset        = 4 * 8
	sizeMask     uint64 = 0x00000000FFFFFFFF

	extAttrCountSize = 2
	bitmapSize       = 6
	entryHeaderSize  = extAttrCountSize + bitmapSize

	bitmapOffset = extAttrCountSize * 8

	refSize = 8

	optAttrSize      = refSize
	extAttrKeySize   = refSize
	extAttrValueSize = refSize
	extAttrPairSize  = extAttrKeySize + extAttrValueSize

	extAttrKeyOffset   = 0
	extAttrValueOffset = extAttrKeyOffset + extAttrKeySize
)

type entry struct {
	t    uint16
	data []byte
}

// Make sure entry implements block.Entry.
var _ block.Entry = (*entry)(nil)

func (e *entry) Get(_ int) interface{} { _ = "STUB: not implemented"; return nil }

func (e *entry) GetBytes(ordinal int) []byte { _ = "STUB: not implemented"; return nil }

func (e *entry) GetString(ordinal int) string { _ = "STUB: not implemented"; return "" }

func (e *entry) GetUint16(ordinal int) uint16 { _ = "STUB: not implemented"; return 0 }

func (e *entry) GetUint64(ordinal int) uint64 { _ = "STUB: not implemented"; return 0 }

func (e *entry) GetInt64(ordinal int) int64 { _ = "STUB: not implemented"; return 0 }

func (e *entry) GetTime(ordinal int) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (e *entry) ExtensionAttributeCount() int { _ = "STUB: not implemented"; return 0 }

func (e *entry) GetExtensionAttribute(attr []byte) []byte { _ = "STUB: not implemented"; return nil }

func (e *entry) RangeExtensionAttributes(cb block.ExtensionAttributeCallback) {
	_ = "STUB: not implemented"
	return
}

func (e *entry) bitmap() uint64 { _ = "STUB: not implemented"; return 0 }

func (e *entry) valueCount() int { _ = "STUB: not implemented"; return 0 }

func (e *entry) valueIndex(ordinal int) int { _ = "STUB: not implemented"; return 0 }

func (e *entry) extCount() int { _ = "STUB: not implemented"; return 0 }

func (e *entry) extVecBase() int { _ = "STUB: not implemented"; return 0 }

func (e *entry) deref(base int) []byte { _ = "STUB: not implemented"; return nil }

func makeRef(offset, length int) uint64 { _ = "STUB: not implemented"; return 0 }

func offsetAndLength(data []byte) (uint32, uint32) { _ = "STUB: not implemented"; return 0, 0 }

func valueIndex(bitmap uint64, mask uint64) int { _ = "STUB: not implemented"; return 0 }

func doValueIndex(bitmap uint64, mask uint64) int { _ = "STUB: not implemented"; return 0 }

func valueOffset(idx int) int { _ = "STUB: not implemented"; return 0 }

func attrKeyOffset(base int, idx int) int { _ = "STUB: not implemented"; return 0 }

func attrValueOffset(base int, idx int) int { _ = "STUB: not implemented"; return 0 }

type entryEncoder struct {
	ceEnc    ceEntryEncoder
	endEnc   endEntryEncoder
	indexEnc indexEntryEncoder
}

// Make sure entryEncoder implements RecordDataEncoder.
var _ RecordDataEncoder = (*entryEncoder)(nil)

func (e *entryEncoder) Size(entry block.Entry) int { _ = "STUB: not implemented"; return 0 }

func (e *entryEncoder) MarshalTo(entry block.Entry, buf []byte) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

type entryDecoder struct {
	indexDec indexEntryDecoder
}

// Make sure entryDecoder implements RecordDataDecoder.
var _ RecordDataDecoder = (*entryDecoder)(nil)

func (d *entryDecoder) Unmarshal(t uint16, offset int, data []byte) (block.Entry, error) {
	_ = "STUB: not implemented"
	return *new(block.Entry), nil
}
