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

package unbounded

import (
	// standard libraries.
	stdrt "runtime"
	"unsafe"
	// this project.
)

const (
	enableReloadPtr1 = false

	activeSpin    = 4
	activeSpinCnt = 30
	passiveSpin   = 1
)

var (
	ncpu = stdrt.NumCPU()
	lock = unsafe.Pointer(&struct{}{})
)

type node[T any] struct {
	_    [7]uint64 // padding to fill cache line
	next unsafe.Pointer
	v    T
}

// Queue is a concurrency safety non-blocking queue,
// used in scenarios with multiple Producers and single Consumers.
type Queue[T any] struct {
	_    [7]uint64 // padding to fill cache line
	tail unsafe.Pointer
	_    [7]uint64 // padding to fill cache line
	head unsafe.Pointer
	_    [7]uint64 // padding to fill cache line
}

func New[T any]() *Queue[T] { _ = "STUB: not implemented"; return nil }

func (q *Queue[T]) Push(v T) bool { _ = "STUB: not implemented"; return false }

func (q *Queue[T]) push(n *node[T]) bool {
	_ = "STUB: not implemented"
	// n.next = nil
	return false
}

func (q *Queue[T]) SharedPop() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (q *Queue[T]) sharedPop() (*node[T], bool) { _ = "STUB: not implemented"; return nil, false }

// No node.

// Only one element, lock.

// Push-pop conflict, spin.

func (q *Queue[T]) UniquePop() (T, bool, bool) {
	_ = "STUB: not implemented"
	return *new(T), false, false
}

func (q *Queue[T]) uniquePop() (*node[T], bool, bool) {
	_ = "STUB: not implemented"
	return nil, false, false

	// No node.
}

// Only one element.

// Push-pop conflict, spin.

func (q *Queue[T]) Peek() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (q *Queue[T]) peek() (*node[T], bool) { _ = "STUB: not implemented"; return nil, false }

func waitUnlock(addr *unsafe.Pointer) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

func waitStable(addr *unsafe.Pointer) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

func reloadPtr(addr *unsafe.Pointer, unexpected unsafe.Pointer) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

func reloadPtr1(addr *unsafe.Pointer, unexpected unsafe.Pointer) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

// TODO(james.yin): use synchronization primitive?

func reloadPtrN(addr *unsafe.Pointer, unexpected unsafe.Pointer) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

// TODO(james.yin): use synchronization primitive?
