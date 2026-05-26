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

	// this project.
	"github.com/vanus-labs/vanus/server/store/block"
	"github.com/vanus-labs/vanus/server/store/vsb/index"
)

// Make sure block implements block.Snapshoter.
var _ block.Snapshoter = (*vsBlock)(nil)

func (b *vsBlock) makeSnapshot() (meta, []index.Index) {
	_ = "STUB: not implemented"
	return *new(meta), nil
}

func makeSnapshot(actx appendContext, indexes []index.Index) (meta, []index.Index) {
	_ = "STUB: not implemented"
	return *new(meta), nil
}

func (b *vsBlock) Snapshot(_ context.Context) (block.Fragment, error) {
	_ = "STUB: not implemented"
	return *new(block.Fragment), nil
}

func (b *vsBlock) ApplySnapshot(_ context.Context, snap block.Fragment) error {
	_ = "STUB: not implemented"
	return nil
}

// Build indexes from data.
