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

package storage

import (
	"context"

	vanus "github.com/vanus-labs/vanus/api/vsr"
	pInfo "github.com/vanus-labs/vanus/pkg/info"
	"github.com/vanus-labs/vanus/server/controller/trigger/metadata"
)

type fake struct {
	subs     map[vanus.ID]*metadata.Subscription
	offset   map[vanus.ID]map[vanus.ID]pInfo.OffsetInfo
	tWorkers map[string]*metadata.TriggerWorkerInfo
}

func NewFakeStorage() Storage { _ = "STUB: not implemented"; return *new(Storage) }

func (f *fake) Close() { _ = "STUB: not implemented"; return }

func (f *fake) CreateSubscription(_ context.Context, sub *metadata.Subscription) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *fake) UpdateSubscription(_ context.Context, sub *metadata.Subscription) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *fake) DeleteSubscription(_ context.Context, id vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *fake) GetSubscription(_ context.Context, id vanus.ID) (*metadata.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *fake) ListSubscription(_ context.Context) ([]*metadata.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *fake) CreateOffset(_ context.Context, subscriptionID vanus.ID, info pInfo.OffsetInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *fake) UpdateOffset(_ context.Context, subscriptionID vanus.ID, info pInfo.OffsetInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *fake) GetOffsets(_ context.Context, subscriptionID vanus.ID) (pInfo.ListOffsetInfo, error) {
	_ = "STUB: not implemented"
	return *new(pInfo.ListOffsetInfo), nil
}

func (f *fake) DeleteOffset(_ context.Context, subscriptionID vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *fake) SaveTriggerWorker(_ context.Context, info metadata.TriggerWorkerInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *fake) GetTriggerWorker(_ context.Context, id string) (*metadata.TriggerWorkerInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *fake) DeleteTriggerWorker(_ context.Context, id string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *fake) ListTriggerWorker(_ context.Context) ([]*metadata.TriggerWorkerInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
