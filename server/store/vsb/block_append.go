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
	stdio "io"

	// third-party libraries.

	// this project.

	"github.com/vanus-labs/vanus/server/store/block"
	"github.com/vanus-labs/vanus/server/store/io"
	"github.com/vanus-labs/vanus/server/store/vsb/index"
)

var (
	errCorruptedFragment = stderr.New("vsb: corrupted fragment")
	dummyReader          = stdio.LimitReader(nil, 0)
)

type appendContext struct {
	seq      int64
	offset   int64
	archived uint32
}

// Make sure appendContext implements block.AppendContext.
var _ block.AppendContext = (*appendContext)(nil)

func (c *appendContext) size(dataOffset int64) int64 { _ = "STUB: not implemented"; return 0 }

func (c *appendContext) WriteOffset() int64 { _ = "STUB: not implemented"; return 0 }

func (c *appendContext) Archived() bool { _ = "STUB: not implemented"; return false }

// Make sure vsBlock implements block.TwoPCAppender.
var _ block.TwoPCAppender = (*vsBlock)(nil)

func (b *vsBlock) NewAppendContext(last block.Fragment) block.AppendContext {
	_ = "STUB: not implemented"
	return *new(block.AppendContext)
}

// Copy append context.

func (b *vsBlock) PrepareAppend(
	_ context.Context, appendCtx block.AppendContext, entries ...block.Entry,
) ([]int64, block.Fragment, bool, error) {
	_ = "STUB: not implemented"
	return nil, *new(block.Fragment), false, nil
}

// TODO(james.yin): fill auto fields in a general way.

func (b *vsBlock) PrepareArchive(ctx context.Context, appendCtx block.AppendContext) (block.Fragment, error) {
	_ = "STUB: not implemented"
	return *new(block.Fragment), nil
}

func (b *vsBlock) CommitAppend(ctx context.Context, frag block.Fragment, cb block.CommitAppendCallback) {
	_ = "STUB: not implemented"
	return
}

// TODO(james.yin): get offset from Stream or AppendContext?
// b.actx.offset

// TODO(james.yin): use new method.

// TODO(james.yin): use new method.

// always false currently.

// NOTE: must update archived flag after append all indexes.

// No more data, so call Sync() to avoid waiting.

func (b *vsBlock) buildIndexes(
	_ context.Context, expected int64, frag block.Fragment,
) ([]index.Index, int64, bool, error) {
	_ = "STUB: not implemented"
	return nil, 0, false, nil
}

// End entry must be the last.

func (b *vsBlock) appendIndexEntry(_ context.Context, indexes []index.Index, cb io.WriteCallback) {
	_ = "STUB: not implemented"
	return
}

// return sz, nil
