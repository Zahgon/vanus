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

	// first-party libraries.

	// this project.
	"github.com/vanus-labs/vanus/server/store/block"
)

// Make sure block implements block.Reader.
var _ block.Reader = (*vsBlock)(nil)

// Read date from file.
func (b *vsBlock) Read(ctx context.Context, seq int64, num int) ([]block.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *vsBlock) entryRange(start, num int) (int64, int64, int, error) {
	_ = "STUB: not implemented"
	// TODO(james.yin): optimize lock.
	return 0, 0, 0, nil
}
