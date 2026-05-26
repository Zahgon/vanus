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
	stderr "errors"

	// first-party libraries.

	// this project.

	"github.com/vanus-labs/vanus/server/store/vsb/codec"
	"github.com/vanus-labs/vanus/server/store/vsb/index"
)

var (
	errCorrupted  = stderr.New("corrupted vsb")
	errIncomplete = stderr.New("incomplete vsb")
)

func (b *vsBlock) Open(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// TODO(james.yin): use direct IO

func (b *vsBlock) init(ctx context.Context) error {
	if err := b.loadHeader(ctx); err != nil {
		return err
	}

	b.enc = codec.NewEncoder()
	if dec, err := codec.NewDecoder(false, int(b.indexSize)); err == nil {
		b.dec = dec
	} else {
		return err
	}

	if err := b.repairMeta(); err != nil {
		return err
	}

	return b.validate(ctx)
}

func (b *vsBlock) repairMeta() error { _ = "STUB: not implemented"; return nil }

// Scan entries.

// Note: use math.MaxInt64-off to avoid overflow.

func (b *vsBlock) rebuildIndexes(num int, tail []index.Index) error {
	_ = "STUB: not implemented"
	return nil
}

// Scan entries.

func (b *vsBlock) validate(_ context.Context) error { _ = "STUB: not implemented"; return nil }
