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

package pkg

import (
	"context"
	"sync"
	"time"
)

type Group struct {
	wg sync.WaitGroup
}

func (g *Group) Wait() {
	_ = "STUB: not implemented"

	// StartWithChannel starts f in a new goroutine in the group.
	// stopCh is passed to f as an argument. f should stop when stopCh is available.
	return
}

func (g *Group) StartWithChannel(stopCh <-chan struct{}, f func(stopCh <-chan struct{})) {
	_ = "STUB: not implemented"
	return
}

// StartWithContext starts f in a new goroutine in the group.
// ctx is passed to f as an argument. f should stop when ctx.Done() is available.
func (g *Group) StartWithContext(ctx context.Context, f func(context.Context)) {
	_ = "STUB: not implemented"
	return
}

// Start starts f in a new goroutine in the group.
func (g *Group) Start(f func()) { _ = "STUB: not implemented"; return }

func UntilWithContext(ctx context.Context, f func(context.Context), period time.Duration) {
	_ = "STUB: not implemented"
	return
}
