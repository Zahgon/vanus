// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package primitive

import (
	"context"
	"sync"
	"time"
)

func NewWatcher(period time.Duration, lookupFunc func(), cleanFuncs ...func()) *Watcher {
	_ = "STUB: not implemented"
	return nil
}

// TODO: no buffer

type Watcher struct {
	period     time.Duration
	lookupFunc func()
	cleanFuncs []func()

	ch chan interface{}
	wg *sync.WaitGroup
	mu sync.RWMutex
}

func (w *Watcher) Close() { _ = "STUB: not implemented"; return }

func (w *Watcher) Run() { _ = "STUB: not implemented"; return }

// do first lookup immediately

// reset timer

func (w *Watcher) Refresh(ctx context.Context) error {
	_ = "STUB: not implemented"
	// batch multi-refresh into a group
	return nil
}

// double check

// TODO: non-blocking

func (w *Watcher) Wakeup() { _ = "STUB: not implemented"; return }

// clear
