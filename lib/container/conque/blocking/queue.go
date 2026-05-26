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

package blocking

import (
	// standard libraries.
	stdsync "sync"

	// this project.
	"github.com/vanus-labs/vanus/lib/container/conque/unbounded"
	"github.com/vanus-labs/vanus/lib/sync"
)

type Queue[T any] struct {
	q     unbounded.Queue[T]
	sem   sync.Semaphore
	mu    stdsync.RWMutex
	state int32
}

func New[T any](handoff bool) *Queue[T] { _ = "STUB: not implemented"; return nil }

func (q *Queue[T]) Init(handoff bool) *Queue[T] { _ = "STUB: not implemented"; return nil }

func (q *Queue[T]) Close() { _ = "STUB: not implemented"; return }

// Wait ensures that all incoming Pushes observe that the queue is closed.
func (q *Queue[T]) Wait() {
	_ = "STUB: not implemented"
	// Make sure no inflight Push.
	return
}

// no op

func (q *Queue[T]) Push(v T) bool {
	_ = "STUB: not implemented"
	// NOTE: no panic, avoid unlocking with defer.
	return false
}

// TODO: maybe atomic is unnecessary.

func (q *Queue[T]) SharedPop() (T, bool) {
	_ = "STUB: not implemented"

	// Check close.
	return *new(T), false
}

func (q *Queue[T]) UniquePop() (T, bool) {
	_ = "STUB: not implemented"

	// Check close.
	return *new(T), false
}

func (q *Queue[T]) RawPop() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }
