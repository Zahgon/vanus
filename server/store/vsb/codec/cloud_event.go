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
	byteAligned   int = 8
	alignAddition     = byteAligned - 1
	alignMask         = -byteAligned
	baseAttrSize      = refSize
	timeAttrSize      = refSize + 8
)

type ceEntryEncoder struct{}

// Make sure ceEntryEncoder implements RecordDataEncoder.
var _ RecordDataEncoder = (*ceEntryEncoder)(nil)

type sizeOptAttrCallback struct {
	size int
}

// Make sure sizeOptAttrCallback implements block.OptionalAttributeCallback.
var _ block.OptionalAttributeCallback = (*sizeOptAttrCallback)(nil)

func (cb *sizeOptAttrCallback) OnBytes(_ int, val []byte) { _ = "STUB: not implemented"; return }

func (cb *sizeOptAttrCallback) OnString(_ int, val string) { _ = "STUB: not implemented"; return }

func (cb *sizeOptAttrCallback) OnUint16(_ int, _ uint16) { _ = "STUB: not implemented"; return }

func (cb *sizeOptAttrCallback) OnUint64(_ int, _ uint64) { _ = "STUB: not implemented"; return }

func (cb *sizeOptAttrCallback) OnInt64(_ int, _ int64) { _ = "STUB: not implemented"; return }

func (cb *sizeOptAttrCallback) OnTime(_ int, _ time.Time) { _ = "STUB: not implemented"; return }

func (cb *sizeOptAttrCallback) OnAttribute(_ int, val interface{}) {
	_ = "STUB: not implemented"
	return
}

type sizeExtAttrCallback struct {
	size int
}

// Make sure sizeExtAttrCallback implements block.ExtensionAttributesCallback.
var _ block.ExtensionAttributeCallback = (*sizeExtAttrCallback)(nil)

func (cb *sizeExtAttrCallback) OnAttribute(attr []byte, val block.Value) {
	_ = "STUB: not implemented"
	return
}

func (e *ceEntryEncoder) Size(entry block.Entry) int { _ = "STUB: not implemented"; return 0 }

// ext count + non-null bitmap

type optAttrMarshaler struct {
	buf       []byte
	bitmap    uint64
	nextAlloc int
}

// Make sure marshalOptAttrCallback implements block.OptionalAttributeCallback.
var _ block.OptionalAttributeCallback = (*optAttrMarshaler)(nil)

func newOptAttrMarshaler(buf []byte, optCnt, extCnt int) *optAttrMarshaler {
	_ = "STUB: not implemented"
	return nil
}

func (oam *optAttrMarshaler) marshal(e block.EntryExt) (uint64, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (oam *optAttrMarshaler) OnBytes(ordinal int, _ []byte) { _ = "STUB: not implemented"; return }

// TODO(james.yin):

func (oam *optAttrMarshaler) OnString(ordinal int, val string) { _ = "STUB: not implemented"; return }

func (oam *optAttrMarshaler) OnUint16(ordinal int, _ uint16) { _ = "STUB: not implemented"; return }

// TODO(james.yin):

func (oam *optAttrMarshaler) OnUint64(ordinal int, _ uint64) { _ = "STUB: not implemented"; return }

// TODO(james.yin):

func (oam *optAttrMarshaler) OnInt64(ordinal int, val int64) { _ = "STUB: not implemented"; return }

func (oam *optAttrMarshaler) OnTime(ordinal int, val time.Time) { _ = "STUB: not implemented"; return }

func (oam *optAttrMarshaler) OnAttribute(ordinal int, _ interface{}) {
	_ = "STUB: not implemented"
	return
}

// TODO(james.yin):

type extAttrMarshaler struct {
	buf        []byte
	valCache   []block.Value
	i          int
	baseOffset int
	nextAlloc  int
}

// Make sure marshalExtAttrCallback implements block.ExtensionAttributesCallback.
var _ block.ExtensionAttributeCallback = (*extAttrMarshaler)(nil)

func newExtAttrMarshaler(buf []byte, optCnt, extCnt, nextAlloc int) *extAttrMarshaler {
	_ = "STUB: not implemented"
	return nil
}

func (eam *extAttrMarshaler) marshal(e block.EntryExt) int {
	_ = "STUB: not implemented"
	// marshal attr key
	return 0
}

// marshal attr value

func (eam *extAttrMarshaler) OnAttribute(attr []byte, val block.Value) {
	_ = "STUB: not implemented"
	// cache attr value
	return
}

// marshal attr key

func (eam *extAttrMarshaler) marshalAttrValue() { _ = "STUB: not implemented"; return }

func (e *ceEntryEncoder) MarshalTo(entry block.Entry, buf []byte) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// fill opt attr (exclude data)

// fill ext count and non-null attributes bitmap

// fill ext attr

// fill data

func alignment(n int) int { _ = "STUB: not implemented"; return 0 }

// vlvrOffset returns the offset of Variable Length Values Region.
func vlvrOffset(valCnt, attrCnt int) int { _ = "STUB: not implemented"; return 0 }
