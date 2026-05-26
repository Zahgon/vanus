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

//go:generate mockgen -source=worker.go -destination=mock_worker.go -package=worker
package worker

import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/vanus-labs/vanus/api/errors"
	"github.com/vanus-labs/vanus/api/trigger"
	vanus "github.com/vanus-labs/vanus/api/vsr"

	primitive "github.com/vanus-labs/vanus/pkg"
	"github.com/vanus-labs/vanus/pkg/queue"
	"github.com/vanus-labs/vanus/server/controller/trigger/metadata"
	"github.com/vanus-labs/vanus/server/controller/trigger/subscription"
)

type TriggerWorker interface {
	Start(ctx context.Context) error
	RemoteStart(ctx context.Context) error
	RemoteStop(ctx context.Context) error
	Close() error
	IsActive() bool
	Reset()
	GetInfo() metadata.TriggerWorkerInfo
	GetAddr() string
	SetPhase(metadata.TriggerWorkerPhase)
	GetPhase() metadata.TriggerWorkerPhase
	GetPendingTime() time.Time
	GetHeartbeatTime() time.Time
	Polish()
	AssignSubscription(id vanus.ID)
	UnAssignSubscription(id vanus.ID) error
	GetAssignedSubscriptions() []vanus.ID
}

// triggerWorker send subscription to trigger worker server.
type triggerWorker struct {
	info                  *metadata.TriggerWorkerInfo
	cc                    *grpc.ClientConn
	client                trigger.TriggerWorkerClient
	lock                  sync.RWMutex
	assignSubscriptionIDs sync.Map
	pendingTime           time.Time
	heartbeatTime         time.Time
	ctx                   context.Context
	stop                  context.CancelFunc
	subscriptionManager   subscription.Manager
	subscriptionQueue     queue.Queue
}

var newTriggerWorker = NewTriggerWorker

func NewTriggerWorkerByAddr(addr string, subscriptionManager subscription.Manager) TriggerWorker {
	_ = "STUB: not implemented"
	return *new(TriggerWorker)
}

func NewTriggerWorker(twInfo *metadata.TriggerWorkerInfo, subscriptionManager subscription.Manager) TriggerWorker {
	_ = "STUB: not implemented"
	return *new(TriggerWorker)
}

func (tw *triggerWorker) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (tw *triggerWorker) handler(ctx context.Context, subscriptionID vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// no assign to this trigger worker,remove subscription

// modify phase to stopped.

// modify subscription to running

func (tw *triggerWorker) IsActive() bool { _ = "STUB: not implemented"; return false }

// Reset when trigger worker restart and re-connect.
func (tw *triggerWorker) Reset() { _ = "STUB: not implemented"; return }

func (tw *triggerWorker) GetInfo() metadata.TriggerWorkerInfo {
	_ = "STUB: not implemented"
	return *new(metadata.TriggerWorkerInfo)
}

func (tw *triggerWorker) GetAddr() string { _ = "STUB: not implemented"; return "" }

func (tw *triggerWorker) SetPhase(phase metadata.TriggerWorkerPhase) {
	_ = "STUB: not implemented"
	return
}

func (tw *triggerWorker) GetPhase() metadata.TriggerWorkerPhase {
	_ = "STUB: not implemented"
	return *new(metadata.TriggerWorkerPhase)
}

func (tw *triggerWorker) Polish() { _ = "STUB: not implemented"; return }

func (tw *triggerWorker) AssignSubscription(id vanus.ID) { _ = "STUB: not implemented"; return }

func (tw *triggerWorker) UnAssignSubscription(id vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (tw *triggerWorker) GetAssignedSubscriptions() []vanus.ID {
	_ = "STUB: not implemented"
	return nil
}

func (tw *triggerWorker) GetPendingTime() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (tw *triggerWorker) GetHeartbeatTime() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (tw *triggerWorker) init(ctx context.Context) error {
	if tw.cc != nil {
		return nil
	}
	var err error
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	tw.cc, err = grpc.DialContext(ctx, tw.info.Addr, opts...)
	if err != nil {
		return errors.ErrTriggerWorker.WithMessage("grpc dial error").Wrap(err)
	}
	tw.client = trigger.NewTriggerWorkerClient(tw.cc)
	return nil
}

func (tw *triggerWorker) Close() error { _ = "STUB: not implemented"; return nil }

func (tw *triggerWorker) RemoteStop(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (tw *triggerWorker) RemoteStart(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (tw *triggerWorker) addSubscription(ctx context.Context, sub *primitive.Subscription) error {
	_ = "STUB: not implemented"
	return nil
}

func (tw *triggerWorker) removeSubscription(ctx context.Context, id vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}
