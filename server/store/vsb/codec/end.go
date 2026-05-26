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

const (
	endEntrySize = 8 + 8 + 8

	endBitmap uint64 = 0b11
)

type endEntryEncoder struct{}

// Make sure endEntryEncoder implements RecordDataEncoder.
var _ RecordDataEncoder = (*endEntryEncoder)(nil)

func (e *endEntryEncoder) Size(_ block.Entry) int { _ = "STUB: not implemented"; return 0 }

func (e *endEntryEncoder) MarshalTo(entry block.Entry, buf []byte) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// bitmap
// seq num
// stime
