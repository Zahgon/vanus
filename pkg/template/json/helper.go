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

package json

import (
	// standard libraries.
	"io"
	"sync"
	"time"
	"unsafe"

	// third-party libraries.
	"github.com/ohler55/ojg/oj"
	// first-party libraries.
)

var writerPool = sync.Pool{
	New: func() any {
		opts := oj.DefaultOptions
		opts.TimeFormat = time.RFC3339
		return &oj.Writer{Options: opts}
	},
}

func writeJSON(w io.Writer, v any) error { _ = "STUB: not implemented"; return nil }

func writeJSONInJSONString(w io.Writer, v any) error {
	_ = "STUB: not implemented"
	// TODO(james.yin): optimize it.
	return nil
}

func writeInJSONString(w io.Writer, v any) error { _ = "STUB: not implemented"; return nil }

// underlying array of byte buffer in stack.

func appendInJSONString(dst []byte, v any) []byte { _ = "STUB: not implemented"; return nil }

// TODO

func ignoreCount(_ int, err error) error {
	_ = "STUB: not implemented"

	// noescape hides a pointer from escape analysis.  noescape is
	// the identity function but escape analysis doesn't think the
	// output depends on the input.  noescape is inlined and currently
	// compiles down to zero instructions.
	// USE CAREFULLY!
	//
	//go:nosplit
	return nil
}

func noescape(p unsafe.Pointer) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

//nolint:staticcheck // copy from go source code.
