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

//go:generate mockgen -source=polling.go -destination=mock_polling.go -package=segment
package segment

import (
	"context"
	"sync"

	vanus "github.com/vanus-labs/vanus/api/vsr"
)

var _ pollingManager = (*pollingMgr)(nil)

type pollingManager interface {
	Add(ctx context.Context, blockID vanus.ID) <-chan struct{}
	NewMessageArrived(blockID vanus.ID)
	Destroy()
}

type pollingMgr struct {
	// vanus.ID, *blockPolling
	blockPollingMap sync.Map
}

func (p *pollingMgr) Destroy() { _ = "STUB: not implemented"; return }

func (p *pollingMgr) Add(ctx context.Context, blockID vanus.ID) <-chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

func (p *pollingMgr) NewMessageArrived(blockID vanus.ID) { _ = "STUB: not implemented"; return }

type blockPolling struct {
	mutex sync.RWMutex
	ch    chan struct{}
}

func newBlockPolling() *blockPolling { _ = "STUB: not implemented"; return nil }

func (bp *blockPolling) add(ctx context.Context) <-chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

func (bp *blockPolling) messageArrived() { _ = "STUB: not implemented"; return }

func (bp *blockPolling) destroy() { _ = "STUB: not implemented"; return }
