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

package entry

import (
	// this project.
	"github.com/vanus-labs/vanus/server/store/block"
)

const (
	addedOptCount = 2
)

type entryExtWrapper struct {
	block.EntryExtWrapper
	t     uint16
	seq   int64
	stime int64
}

// Make sure entryWrapper implements block.Entry.
var _ block.EntryExt = (*entryExtWrapper)(nil)

func (w *entryExtWrapper) GetUint16(ordinal int) uint16 { _ = "STUB: not implemented"; return 0 }

func (w *entryExtWrapper) GetInt64(ordinal int) int64 { _ = "STUB: not implemented"; return 0 }

func (w *entryExtWrapper) RangeOptionalAttributes(cb block.OptionalAttributeCallback) {
	_ = "STUB: not implemented"
	return
}

func (w *entryExtWrapper) OptionalAttributeCount() int { _ = "STUB: not implemented"; return 0 }

func Wrap(e block.Entry, t uint16, seq int64, stime int64) block.Entry {
	_ = "STUB: not implemented"
	return *new(block.Entry)
}

// TODO(james.yin): entry wrapper
