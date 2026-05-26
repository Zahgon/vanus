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

//go:generate mockgen -source=manager.go -destination=mock_manager.go -package=worker
package worker

import (
	"context"
	"fmt"
	"sync"
	"time"

	vanus "github.com/vanus-labs/vanus/api/vsr"

	"github.com/vanus-labs/vanus/server/controller/trigger/metadata"
	"github.com/vanus-labs/vanus/server/controller/trigger/storage"
	"github.com/vanus-labs/vanus/server/controller/trigger/subscription"
)

const (
	defaultCheckInterval       = 5 * time.Second
	defaultLostHeartbeatTime   = 30 * time.Second
	defaultHeartbeatTimeout    = 60 * time.Second
	defaultDisconnectCleanTime = 120 * time.Second
	defaultWaitRunningTimeout  = 30 * time.Second
	defaultStartWorkerDuration = 10 * time.Second
)

type Manager interface {
	AddTriggerWorker(ctx context.Context, addr string) error
	GetTriggerWorker(addr string) TriggerWorker
	RemoveTriggerWorker(ctx context.Context, addr string)
	UpdateTriggerWorkerInfo(ctx context.Context, addr string) error
	GetActiveRunningTriggerWorker() []metadata.TriggerWorkerInfo
	Init(ctx context.Context) error
	Start()
	Stop()
}

var ErrTriggerWorkerNotFound = fmt.Errorf("trigger worker not found")

type OnTriggerWorkerRemoveSubscription func(ctx context.Context, subId vanus.ID, addr string) error

type Config struct {
	CheckInterval       time.Duration
	LostHeartbeatTime   time.Duration
	HeartbeatTimeout    time.Duration
	DisconnectCleanTime time.Duration
	WaitRunningTimeout  time.Duration

	StartWorkerDuration       time.Duration
	StartSubscriptionDuration time.Duration
}

func (c *Config) init() {
	if c.CheckInterval <= 0 {
		c.CheckInterval = defaultCheckInterval
	}
	if c.LostHeartbeatTime <= 0 {
		c.LostHeartbeatTime = defaultLostHeartbeatTime
	}
	if c.HeartbeatTimeout <= 0 {
		c.HeartbeatTimeout = defaultHeartbeatTimeout
	}
	if c.DisconnectCleanTime <= 0 {
		c.DisconnectCleanTime = defaultDisconnectCleanTime
	}
	if c.WaitRunningTimeout <= 0 {
		c.WaitRunningTimeout = defaultWaitRunningTimeout
	}
	if c.StartWorkerDuration <= 0 {
		c.StartWorkerDuration = defaultStartWorkerDuration
	}
}

type manager struct {
	config               Config
	triggerWorkers       map[string]TriggerWorker
	storage              storage.TriggerWorkerStorage
	subscriptionManager  subscription.Manager
	lock                 sync.RWMutex
	onRemoveSubscription OnTriggerWorkerRemoveSubscription
	ctx                  context.Context
	stop                 context.CancelFunc
}

func NewTriggerWorkerManager(config Config,
	storage storage.TriggerWorkerStorage,
	subscriptionManager subscription.Manager,
	handler OnTriggerWorkerRemoveSubscription,
) Manager {
	_ = "STUB: not implemented"
	return *new(Manager)
}

func (m *manager) GetTriggerWorker(addr string) TriggerWorker {
	_ = "STUB: not implemented"
	return *new(TriggerWorker)
}

func (m *manager) AddTriggerWorker(ctx context.Context, addr string) error {
	_ = "STUB: not implemented"
	return nil
}

// wait clean

func (m *manager) RemoveTriggerWorker(ctx context.Context, addr string) {
	_ = "STUB: not implemented"
	return
}

func (m *manager) UpdateTriggerWorkerInfo(ctx context.Context, addr string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) startTriggerWorker(ctx context.Context, tWorker TriggerWorker) {
	_ = "STUB: not implemented"
	return
}

// trigger worker restart need assign to trigger worker again

func (m *manager) cleanTriggerWorker(ctx context.Context, tWorker TriggerWorker) {
	_ = "STUB: not implemented"
	return
}

func (m *manager) doTriggerWorkerLeave(ctx context.Context, tWorker TriggerWorker) bool {
	_ = "STUB: not implemented"
	return false
}

// reallocate subscription

func (m *manager) GetActiveRunningTriggerWorker() []metadata.TriggerWorkerInfo {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) getTriggerWorkers() []TriggerWorker { _ = "STUB: not implemented"; return nil }

func (m *manager) deleteTriggerWorker(addr string) { _ = "STUB: not implemented"; return }

func (m *manager) Init(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (m *manager) Stop() { _ = "STUB: not implemented"; return }

func (m *manager) Start() { _ = "STUB: not implemented"; return }

func (m *manager) check(ctx context.Context) {
	_ = "STUB: not implemented"
	// log.Debug(ctx, "trigger worker check begin", nil)
	return
}

func (m *manager) pendingTriggerWorkerHandler(ctx context.Context, tWorker TriggerWorker) {
	_ = "STUB: not implemented"
	return
}

func (m *manager) runningTriggerWorkerHandler(ctx context.Context, tWorker TriggerWorker) {
	_ = "STUB: not implemented"
	return
}
