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
	"os"
	"sync"

	// this project.

	vanus "github.com/vanus-labs/vanus/api/vsr"
	"github.com/vanus-labs/vanus/server/store/block"
	"github.com/vanus-labs/vanus/server/store/io/stream"
	"github.com/vanus-labs/vanus/server/store/io/zone"
	"github.com/vanus-labs/vanus/server/store/vsb/codec"
	"github.com/vanus-labs/vanus/server/store/vsb/index"
)

const FormatMagic = uint32(0x00627376) // ASCII of "vsb" in little endian

type meta struct {
	writeOffset int64
	// entryLength is the length of persisted entries.
	entryLength int64
	// entryNum is the number of persisted entries.
	entryNum int64
	// archived is the flag indicating Block is archived.
	archived bool
}

// vsBlock is Vanus block file.
type vsBlock struct {
	id       vanus.ID
	path     string
	capacity int64

	dataOffset int64
	indexSize  uint16

	indexOffset int64
	indexLength int

	fm      meta // flushed meta
	actx    appendContext
	indexes []index.Index
	mu      sync.RWMutex

	enc codec.EntryEncoder
	dec codec.EntryDecoder
	lis block.ArchivedListener

	f  *os.File
	z  zone.Interface
	s  stream.Stream
	wg sync.WaitGroup
}

// Make sure vsBlock implements block.Raw.
var _ block.Raw = (*vsBlock)(nil)

func (b *vsBlock) ID() vanus.ID { _ = "STUB: not implemented"; return *new(vanus.ID) }

func (b *vsBlock) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Flush metadata.

func (b *vsBlock) Delete(context.Context) error {
	_ = "STUB: not implemented"
	// FIXME(james.yin): make sure block is closed.
	return nil
}

func (b *vsBlock) Status() block.Statistics {
	_ = "STUB: not implemented"
	return *new(block.Statistics)
}

func (b *vsBlock) stat(m meta, indexes []index.Index) block.Statistics {
	_ = "STUB: not implemented"
	return *new(block.Statistics)
}

func (b *vsBlock) full() bool { _ = "STUB: not implemented"; return false }
