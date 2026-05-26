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

	// third-party libraries.

	// this project.
	"github.com/vanus-labs/vanus/server/store/block"
	"github.com/vanus-labs/vanus/server/store/vsb/index"
)

// Make sure block implements block.Reader.
var _ block.Seeker = (*vsBlock)(nil)

func (b *vsBlock) Seek(ctx context.Context, index int64, key block.Entry, flag block.SeekKeyFlag) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *vsBlock) seekKeyExact(_ context.Context, idx int64, key block.Entry, indexes []index.Index) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *vsBlock) seekKeyOrNext(_ context.Context, idx int64, key block.Entry, indexes []index.Index) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *vsBlock) seekKeyOrPrev(_ context.Context, idx int64, key block.Entry, indexes []index.Index) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *vsBlock) seekAfterKey(_ context.Context, idx int64, key block.Entry, indexes []index.Index) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *vsBlock) seekBeforeKey(_ context.Context, idx int64, key block.Entry, indexes []index.Index) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *vsBlock) selectComparer(_ int64, key block.Entry) func(index.Index) int {
	_ = "STUB: not implemented"
	// TODO(james.yin): support non-stime index.
	return nil
}

// v > val

func searchGE(indexes []index.Index, cmp func(index.Index) int) int64 {
	_ = "STUB: not implemented"
	return 0
}

func searchGT(indexes []index.Index, cmp func(index.Index) int) int64 {
	_ = "STUB: not implemented"
	return 0
}
