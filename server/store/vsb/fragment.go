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

	// this project.
	"github.com/vanus-labs/vanus/server/store/block"
	"github.com/vanus-labs/vanus/server/store/vsb/codec"
)

const (
	OffsetSize = 8

	PayloadOffset = OffsetSize
)

type fragment struct {
	offset  int64
	entries []block.Entry
	enc     codec.EntryEncoder
	data    []byte
	sz      uint32
}

// Make sure fragment implements block.Fragment and block.FragmentMarshaler.
var (
	_ block.Fragment          = (*fragment)(nil)
	_ block.FragmentMarshaler = (*fragment)(nil)
)

func newFragment(offset int64, entries []block.Entry, enc codec.EntryEncoder) block.Fragment {
	_ = "STUB: not implemented"
	return *new(block.Fragment)
}

func (f *fragment) Payload() []byte { _ = "STUB: not implemented"; return nil }

func (f *fragment) Size() int { _ = "STUB: not implemented"; return 0 }

func (f *fragment) StartOffset() int64 { _ = "STUB: not implemented"; return 0 }

func (f *fragment) EndOffset() int64 { _ = "STUB: not implemented"; return 0 }

func (f *fragment) MarshalFragment() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *fragment) size() int { _ = "STUB: not implemented"; return 0 }

func (f *fragment) doMarshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
