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

	// first-party libraries.

	vanus "github.com/vanus-labs/vanus/api/vsr"

	// this project.

	"github.com/vanus-labs/vanus/server/store/block"
)

const (
	vsbExt          = ".vsb"
	defaultFilePerm = 0o644
)

func (e *engine) Create(ctx context.Context, id vanus.ID, capacity int64) (block.Raw, error) {
	_ = "STUB: not implemented"
	return *new(block.Raw), nil
}

func processError(err error, f *os.File, path string) error { _ = "STUB: not implemented"; return nil }

func (e *engine) Open(ctx context.Context, id vanus.ID) (block.Raw, error) {
	_ = "STUB: not implemented"
	return *new(block.Raw), nil
}

func (e *engine) resolvePath(id vanus.ID) string { _ = "STUB: not implemented"; return "" }

func BlockPath(dir string, id vanus.ID) string { _ = "STUB: not implemented"; return "" }
