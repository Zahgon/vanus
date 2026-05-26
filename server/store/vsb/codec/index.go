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

import ( // standard libraries.
	// this project.
	"github.com/vanus-labs/vanus/server/store/block"
)

type indexEntryEncoder struct {
	indexSize int
}

// Make sure indexEntryEncoder implements RecordDataEncoder.
var _ RecordDataEncoder = (*indexEntryEncoder)(nil)

func (e *indexEntryEncoder) Size(entry block.Entry) int { _ = "STUB: not implemented"; return 0 }

func (e *indexEntryEncoder) MarshalTo(entry block.Entry, buf []byte) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// offset
// length
// reserved
// stime

type indexEntryDecoder struct {
	indexSize int
}

// Make sure indexEntryDecoder implements RecordDataDecoder.
var _ RecordDataDecoder = (*indexEntryDecoder)(nil)

func (d *indexEntryDecoder) Unmarshal(t uint16, offset int, data []byte) (block.Entry, error) {
	_ = "STUB: not implemented" //nolint:revive // ok
	return *new(block.Entry), nil
}
