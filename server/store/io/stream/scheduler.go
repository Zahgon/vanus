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

package stream

import (
	// this project.
	"github.com/vanus-labs/vanus/lib/executor"
	"github.com/vanus-labs/vanus/server/store/io"
	"github.com/vanus-labs/vanus/server/store/io/block"
	"github.com/vanus-labs/vanus/server/store/io/engine"
	"github.com/vanus-labs/vanus/server/store/io/zone"
)

type Scheduler interface {
	Close()

	Register(z zone.Interface, wo int64, direct bool) Stream
	Unregister(s Stream)
}

type scheduler struct {
	e  engine.Interface
	bp *block.BufferPool
	pq pendingQueue

	callbackExecutor executor.MultiFlow
}

// Make sure scheduler implements Scheduler.
var _ Scheduler = (*scheduler)(nil)

func NewScheduler(e engine.Interface, opts ...Option) Scheduler {
	_ = "STUB: not implemented"
	return *new(Scheduler)
}

func (s *scheduler) init(e engine.Interface, cfg config) *scheduler {
	s.e = e
	s.bp = block.NewBufferPool(cfg.flushBatchSize)
	s.pq.init(cfg.flushDelayTime)
	s.callbackExecutor.Init(cfg.callbackParallel, false, true)
	return s
}

func (s *scheduler) Close() { _ = "STUB: not implemented"; return }

func (s *scheduler) Register(z zone.Interface, wo int64, direct bool) Stream {
	_ = "STUB: not implemented"
	return *new(Stream)
}

// TODO(james.yin)

func (s *scheduler) Unregister(ss Stream) { _ = "STUB: not implemented"; return }

func (s *scheduler) writeAt(z zone.Interface, b []byte, off int64, so, eo int, cb io.WriteCallback) {
	_ = "STUB: not implemented"
	return
}

func (s *scheduler) bufferSize() int { _ = "STUB: not implemented"; return 0 }

func (s *scheduler) getBuffer(base int64) *block.Buffer { _ = "STUB: not implemented"; return nil }

func (s *scheduler) putBuffer(b *block.Buffer) { _ = "STUB: not implemented"; return }

func (s *scheduler) delayFlush(ss *stream) PendingID {
	_ = "STUB: not implemented"
	return *new(PendingID)
}

func (s *scheduler) cancelFlushTask(pid PendingID) { _ = "STUB: not implemented"; return }
