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

package psync

import (
	// standard libraries.
	"os"

	// first-party libraries.
	"github.com/vanus-labs/vanus/lib/container/conque/blocking"

	// this project.
	"github.com/vanus-labs/vanus/server/store/io"
	"github.com/vanus-labs/vanus/server/store/io/engine"
	"github.com/vanus-labs/vanus/server/store/io/zone"
)

type writeTask struct {
	f   *os.File
	b   []byte
	off int64
	cb  io.WriteCallback
}

type psync struct {
	q blocking.Queue[writeTask]
}

// Make sure engine implements engine.Interface.
var _ engine.Interface = (*psync)(nil)

func New(opts ...Option) engine.Interface { _ = "STUB: not implemented"; return *new(engine.Interface) }

func (e *psync) init(cfg config) *psync {
	e.q.Init(false)
	for i := 0; i < cfg.parallel; i++ {
		go e.run()
	}
	return e
}

func (e *psync) Close() { _ = "STUB: not implemented"; return }

func (e *psync) WriteAt(z zone.Interface, b []byte, off int64, so, eo int, cb io.WriteCallback) {
	_ = "STUB: not implemented" //nolint:revive // ok
	return
	// if eo != 0 && eo != len(b) {
	// 	b = b[:eo]
	// }
	// if so != 0 {
	// 	b = b[so:]
	// 	off += int64(so)
	// }
}

func (e *psync) run() { _ = "STUB: not implemented"; return }

func (t *writeTask) invoke() {
	_ = "STUB: not implemented"
	// NOTE: data race is ok here.
	return
}
