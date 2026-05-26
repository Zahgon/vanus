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

//go:generate mockgen -source=subscription.go -destination=mock_subscription.go -package=subscription
package subscription

import (
	"context"
	"sync"
	"time"

	"github.com/vanus-labs/vanus/api/cluster"
	vanus "github.com/vanus-labs/vanus/api/vsr"
	eb "github.com/vanus-labs/vanus/client"

	"github.com/vanus-labs/vanus/pkg/info"
	"github.com/vanus-labs/vanus/server/controller/trigger/metadata"
	"github.com/vanus-labs/vanus/server/controller/trigger/secret"
	"github.com/vanus-labs/vanus/server/controller/trigger/storage"
	"github.com/vanus-labs/vanus/server/controller/trigger/subscription/offset"
)

type Manager interface {
	SaveOffset(ctx context.Context, id vanus.ID, offsets info.ListOffsetInfo, commit bool) error
	// GetOrSaveOffset get offset only from etcd, if it isn't exist will get from cli and save to etcd,
	// and it contains retry eb offset
	GetOrSaveOffset(ctx context.Context, id vanus.ID) (info.ListOffsetInfo, error)
	// GetOffset get offset only from etcd, it doesn't contain retry eb offset
	GetOffset(ctx context.Context, id vanus.ID) (info.ListOffsetInfo, error)
	GetDeadLetterOffset(ctx context.Context, id vanus.ID) (uint64, error)
	SaveDeadLetterOffset(ctx context.Context, id vanus.ID, offset uint64) error
	ResetOffsetByTimestamp(ctx context.Context, id vanus.ID, timestamp uint64) (info.ListOffsetInfo, error)
	ListSubscription(ctx context.Context) []*metadata.Subscription
	GetSubscription(ctx context.Context, id vanus.ID) *metadata.Subscription
	GetSubscriptionByName(ctx context.Context, namespaceID vanus.ID, name string) *metadata.Subscription
	AddSubscription(ctx context.Context, subscription *metadata.Subscription) error
	UpdateSubscription(ctx context.Context, subscription *metadata.Subscription) error
	Heartbeat(ctx context.Context, id vanus.ID, addr string, time time.Time) error
	DeleteSubscription(ctx context.Context, id vanus.ID) error
	Init(ctx context.Context) error
	Start()
	Stop()
}

const (
	defaultCommitInterval = time.Second
)

type manager struct {
	cl              cluster.Cluster
	ebCli           eb.Client
	secretStorage   secret.Storage
	storage         storage.Storage
	offsetManager   offset.Manager
	lock            sync.RWMutex
	subscriptionMap map[vanus.ID]*metadata.Subscription
	// key: eventbusID, value: deadLetterEventbusID
	deadLetterEventbusMap map[vanus.ID]vanus.ID
	// key: deadLetterEventbusID, value: eventlogID
	deadLetterEventlogMap map[vanus.ID]vanus.ID
	retryEventlogID       vanus.ID
	retryEventbusID       vanus.ID
	timerEventbusID       vanus.ID
}

func NewSubscriptionManager(storage storage.Storage, secretStorage secret.Storage,
	ebCli eb.Client, cl cluster.Cluster,
) Manager {
	_ = "STUB: not implemented"
	return *new(Manager)
}

func (m *manager) ListSubscription(_ context.Context) []*metadata.Subscription {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) GetSubscriptionByName(_ context.Context, namespaceID vanus.ID, name string) *metadata.Subscription {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) GetSubscription(_ context.Context, id vanus.ID) *metadata.Subscription {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) getSystemEventbusAndEventlog(ctx context.Context, name string) (vanus.ID, vanus.ID, error) {
	_ = "STUB: not implemented"
	return *new(vanus.ID), *new(vanus.ID), nil
}

func (m *manager) initDeadLetterEventbus(ctx context.Context, eventbusID vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO dlq eb belongs to system?

func (m *manager) initRetryEventbus(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) initTimerEventbus(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) AddSubscription(ctx context.Context, subscription *metadata.Subscription) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) UpdateSubscription(ctx context.Context, sub *metadata.Subscription) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteSubscription will do
// 1.delete offset
// 2.delete subscription .
func (m *manager) DeleteSubscription(ctx context.Context, id vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) Heartbeat(ctx context.Context, id vanus.ID, addr string, time time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// data is not consistent, record

func (m *manager) Stop() { _ = "STUB: not implemented"; return }

func (m *manager) Init(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (m *manager) Start() { _ = "STUB: not implemented"; return }
