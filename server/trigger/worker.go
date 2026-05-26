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

//go:generate mockgen -source=worker.go -destination=mock_worker.go -package=trigger
package trigger

import (
	"context"
	"sync"
	"time"

	"github.com/vanus-labs/vanus/api/cluster"
	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	metapb "github.com/vanus-labs/vanus/api/meta"
	vanus "github.com/vanus-labs/vanus/api/vsr"
	primitive "github.com/vanus-labs/vanus/pkg"

	"github.com/vanus-labs/vanus/server/trigger/trigger"
)

type Worker interface {
	Init(ctx context.Context) error
	Register(ctx context.Context) error
	Unregister(ctx context.Context) error
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	AddSubscription(ctx context.Context, subscription *primitive.Subscription) error
	RemoveSubscription(ctx context.Context, id vanus.ID) error
	PauseSubscription(ctx context.Context, id vanus.ID) error
	StartSubscription(ctx context.Context, id vanus.ID) error
}

const (
	defaultHeartbeatInterval = 2 * time.Second
)

type newTrigger func(subscription *primitive.Subscription, options ...trigger.Option) (trigger.Trigger, error)

type worker struct {
	triggerMap map[vanus.ID]trigger.Trigger
	ctx        context.Context
	stop       context.CancelFunc
	config     Config
	newTrigger newTrigger
	wg         sync.WaitGroup
	lock       sync.RWMutex
	tgLock     sync.RWMutex
	client     ctrlpb.TriggerControllerClient
	ctrl       cluster.Cluster
}

func NewWorker(config Config) Worker { _ = "STUB: not implemented"; return *new(Worker) }

func (w *worker) getTrigger(id vanus.ID) (trigger.Trigger, bool) {
	_ = "STUB: not implemented"
	return *new(trigger.Trigger), false
}

func (w *worker) addTrigger(id vanus.ID, t trigger.Trigger) { _ = "STUB: not implemented"; return }

func (w *worker) deleteTrigger(id vanus.ID) { _ = "STUB: not implemented"; return }

func (w *worker) Init(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (w *worker) Register(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (w *worker) Unregister(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (w *worker) Start(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (w *worker) Stop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// stop subscription

// commit offset

// stop heartbeat

// clean trigger

func (w *worker) AddSubscription(ctx context.Context, subscription *primitive.Subscription) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *worker) RemoveSubscription(ctx context.Context, id vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *worker) PauseSubscription(ctx context.Context, id vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *worker) StartSubscription(ctx context.Context, id vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *worker) startHeartbeat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (w *worker) stopSubscription(ctx context.Context, id vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *worker) startSubscription(ctx context.Context, id vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *worker) commitOffsets(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (w *worker) getAllSubscriptionInfo(ctx context.Context) []*metapb.SubscriptionInfo {
	_ = "STUB: not implemented"
	return nil
}

func (w *worker) getTriggerOptions(subscription *primitive.Subscription) []trigger.Option {
	_ = "STUB: not implemented"
	return nil
}

// todo use subscription config replace global config
