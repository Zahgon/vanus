// Copyright 2023 Linkall Inc.
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

package text

import (
	// standard libraries.
	"io"
	"sync"

	// third-party libraries.
	"github.com/ohler55/ojg/oj"
	// first-party libraries.
)

var writerPool = sync.Pool{
	New: func() any {
		return &oj.Writer{Options: oj.DefaultOptions}
	},
}

func writeJSON(w io.Writer, v any) error { _ = "STUB: not implemented"; return nil }

func write(w io.Writer, v any) error { _ = "STUB: not implemented"; return nil }

func ignoreCount(_ int, err error) error { _ = "STUB: not implemented"; return nil }
