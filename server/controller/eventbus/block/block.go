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

//go:generate mockgen -source=block.go -destination=mock_block.go -package=block
package block

import (
	"context"
	"sync"
	"time"

	vanus "github.com/vanus-labs/vanus/api/vsr"

	"github.com/vanus-labs/vanus/pkg/kv"
	"github.com/vanus-labs/vanus/server/controller/eventbus/metadata"
	"github.com/vanus-labs/vanus/server/controller/eventbus/server"
)

const (
	defaultBlockSize                = int64(64 * 1024 * 1024)
	minimumBlockSize                = int64(4 * 1024 * 1024)
	defaultBlockBufferSizePerVolume = 8
)

type Allocator interface {
	Run(ctx context.Context, kvCli kv.Client, dynamicAllocate bool) error
	Pick(ctx context.Context, num int) ([]*metadata.Block, error)
	PickByVolumes(ctx context.Context, volumes []vanus.ID) ([]*metadata.Block, error)
	Stop()
}

func NewAllocator(defaultBlockCapacity int64, selector VolumeSelector) Allocator {
	_ = "STUB: not implemented"
	return *new(Allocator)
}

type allocator struct {
	selector VolumeSelector
	// key: volumeID, value: SkipList of *metadata.Block
	volumeBlockBuffer sync.Map
	kvClient          kv.Client
	mutex             sync.Mutex
	allocateTicker    *time.Ticker
	blockCapacity     int64
}

func (al *allocator) PickByVolumes(ctx context.Context, volumes []vanus.ID) ([]*metadata.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (al *allocator) Run(ctx context.Context, kvCli kv.Client, startDynamicAllocate bool) error {
	_ = "STUB: not implemented" //nolint:revive // ok
	return nil
}

// if startDynamicAllocate {
// disable by wenfeng on 3.10
// al.cancelCtx, al.cancel = context.WithCancel(context.Background())
// go al.dynamicAllocateBlockTask(al.cancelCtx)
// }

func (al *allocator) Pick(ctx context.Context, num int) ([]*metadata.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (al *allocator) pick(ctx context.Context, volumes []server.Instance) ([]*metadata.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (al *allocator) Stop() { _ = "STUB: not implemented"; return }

func (al *allocator) dynamicAllocateBlockTask(ctx context.Context) {
	_ = "STUB: not implemented" //nolint:unused // ok
	return
}

func (al *allocator) updateBlockInKV(ctx context.Context, block *metadata.Block) error {
	_ = "STUB: not implemented"
	return nil
}
