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

//go:generate mockgen -source=offset.go -destination=mock_offset.go -package=offset
package offset

import (
	"context"
	"sync"
	"time"

	vanus "github.com/vanus-labs/vanus/api/vsr"
	"github.com/vanus-labs/vanus/pkg/info"
	"github.com/vanus-labs/vanus/server/controller/trigger/storage"
)

type Manager interface {
	GetOffset(ctx context.Context, subscriptionID vanus.ID) (info.ListOffsetInfo, error)
	Offset(ctx context.Context, subscriptionID vanus.ID, offsets info.ListOffsetInfo, commit bool) error
	RemoveRegisterSubscription(ctx context.Context, id vanus.ID) error
	Start()
	Stop()
}

const (
	defaultCommitInterval = time.Second
	defaultCloseWaitTime  = 2 * time.Second
)

type manager struct {
	subscriptionOffset sync.Map
	storage            storage.OffsetStorage
	commitInterval     time.Duration
	closeWaitTimeout   time.Duration
	ctx                context.Context
	stop               context.CancelFunc
	wg                 sync.WaitGroup
}

func NewOffsetManager(storage storage.OffsetStorage, commitInterval time.Duration) Manager {
	_ = "STUB: not implemented"
	return *new(Manager)
}

func (m *manager) GetOffset(ctx context.Context, subscriptionID vanus.ID) (info.ListOffsetInfo, error) {
	_ = "STUB: not implemented"
	return *new(info.ListOffsetInfo), nil
}

func (m *manager) Offset(ctx context.Context, subscriptionID vanus.ID, offsets info.ListOffsetInfo, commit bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) getSubscriptionOffset(ctx context.Context, id vanus.ID) (*subscriptionOffset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *manager) RemoveRegisterSubscription(ctx context.Context, id vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// stop commit

func (m *manager) Stop() { _ = "STUB: not implemented"; return }

func (m *manager) Start() { _ = "STUB: not implemented"; return }

func (m *manager) commit(ctx context.Context) { _ = "STUB: not implemented"; return }

type subscriptionOffset struct {
	subscriptionID vanus.ID
	offsets        sync.Map
	stopped        bool
	lock           sync.Mutex
}

func initSubscriptionOffset(ctx context.Context,
	storage storage.OffsetStorage,
	subscriptionID vanus.ID,
) (*subscriptionOffset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getEventlogOffset if not exist create.
func (o *subscriptionOffset) getEventlogOffset(info info.OffsetInfo) *eventlogOffset {
	_ = "STUB: not implemented"
	return nil
}

func (o *subscriptionOffset) offset(infos info.ListOffsetInfo) { _ = "STUB: not implemented"; return }

func (o *subscriptionOffset) getOffsets() info.ListOffsetInfo {
	_ = "STUB: not implemented"
	return *new(info.ListOffsetInfo)
}

func (o *subscriptionOffset) stop() { _ = "STUB: not implemented"; return }

func (o *subscriptionOffset) commitOffset(ctx context.Context, storage storage.OffsetStorage) {
	_ = "STUB: not implemented"
	return
}

type eventlogOffset struct {
	subscriptionID vanus.ID
	eventlogID     vanus.ID
	offset         uint64
	commit         uint64
	checkExist     bool
}

func (o *eventlogOffset) setOffset(offset uint64) { _ = "STUB: not implemented"; return }

func (o *eventlogOffset) commitOffset(ctx context.Context, storage storage.OffsetStorage) error {
	_ = "STUB: not implemented"
	return nil
}
