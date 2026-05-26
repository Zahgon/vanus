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

package trigger

import (
	// standard libraries.
	"context"
	"sync"
	"time"

	// third-party libraries.

	"google.golang.org/protobuf/types/known/emptypb"

	// first-party libraries.
	"github.com/vanus-labs/vanus/api/cluster"
	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	metapb "github.com/vanus-labs/vanus/api/meta"
	vanus "github.com/vanus-labs/vanus/api/vsr"
	eb "github.com/vanus-labs/vanus/client"

	// this project.
	primitive "github.com/vanus-labs/vanus/pkg"
	"github.com/vanus-labs/vanus/server/controller/member"
	"github.com/vanus-labs/vanus/server/controller/trigger/metadata"
	"github.com/vanus-labs/vanus/server/controller/trigger/secret"
	"github.com/vanus-labs/vanus/server/controller/trigger/storage"
	"github.com/vanus-labs/vanus/server/controller/trigger/subscription"
	"github.com/vanus-labs/vanus/server/controller/trigger/worker"
)

var _ ctrlpb.TriggerControllerServer = &controller{}

const (
	defaultGcSubscriptionInterval = time.Second * 10
)

func NewController(config Config, mem member.Member) *controller {
	_ = "STUB: not implemented"
	return nil
}

type controller struct {
	config                Config
	member                member.Member
	storage               storage.Storage
	secretStorage         secret.Storage
	subscriptionManager   subscription.Manager
	workerManager         worker.Manager
	scheduler             *worker.SubscriptionScheduler
	needCleanSubscription map[vanus.ID]string
	lock                  sync.Mutex
	membershipMutex       sync.Mutex
	isLeader              bool
	ctx                   context.Context
	stopFunc              context.CancelFunc
	state                 primitive.ServerState
	cl                    cluster.Cluster
	ebClient              eb.Client
}

func (ctrl *controller) SetDeadLetterEventOffset(
	ctx context.Context, request *ctrlpb.SetDeadLetterEventOffsetRequest,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) GetDeadLetterEventOffset(
	ctx context.Context, request *ctrlpb.GetDeadLetterEventOffsetRequest,
) (*ctrlpb.GetDeadLetterEventOffsetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) CommitOffset(
	ctx context.Context, request *ctrlpb.CommitOffsetRequest,
) (*ctrlpb.CommitOffsetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) ResetOffsetToTimestamp(
	ctx context.Context, request *ctrlpb.ResetOffsetToTimestampRequest,
) (*ctrlpb.ResetOffsetToTimestampResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) CreateSubscription(
	ctx context.Context, request *ctrlpb.CreateSubscriptionRequest,
) (*metapb.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) UpdateSubscription(
	ctx context.Context, request *ctrlpb.UpdateSubscriptionRequest,
) (*metapb.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) DeleteSubscription(
	ctx context.Context, request *ctrlpb.DeleteSubscriptionRequest,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) DisableSubscription(
	ctx context.Context, request *ctrlpb.DisableSubscriptionRequest,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) ResumeSubscription(
	ctx context.Context, request *ctrlpb.ResumeSubscriptionRequest,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) GetSubscription(
	ctx context.Context, request *ctrlpb.GetSubscriptionRequest,
) (*metapb.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) TriggerWorkerHeartbeat(
	heartbeat ctrlpb.TriggerController_TriggerWorkerHeartbeatServer,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (ctrl *controller) triggerWorkerHeartbeatRequest(
	ctx context.Context, req *ctrlpb.TriggerWorkerHeartbeatRequest,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (ctrl *controller) RegisterTriggerWorker(
	ctx context.Context, request *ctrlpb.RegisterTriggerWorkerRequest,
) (*ctrlpb.RegisterTriggerWorkerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) UnregisterTriggerWorker(
	ctx context.Context, request *ctrlpb.UnregisterTriggerWorkerRequest,
) (*ctrlpb.UnregisterTriggerWorkerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) ListSubscription(
	ctx context.Context, request *ctrlpb.ListSubscriptionRequest,
) (*ctrlpb.ListSubscriptionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// gcSubscription before delete subscription,need
//
// 1.trigger worker remove subscription
// 2.delete offset
// 3.delete subscription .
func (ctrl *controller) gcSubscription(ctx context.Context, id vanus.ID, addr string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ctrl *controller) gcSubscriptions(ctx context.Context) { _ = "STUB: not implemented"; return }

func (ctrl *controller) requeueSubscription(ctx context.Context, id vanus.ID, addr string) error {
	_ = "STUB: not implemented"
	return nil
}

// data is not consistent, record

func (ctrl *controller) init(ctx context.Context) error {
	ctrl.initTriggerSystemEventbus()
	err := ctrl.subscriptionManager.Init(ctx)
	if err != nil {
		return err
	}
	err = ctrl.workerManager.Init(ctx)
	if err != nil {
		return err
	}
	// restart,need reschedule
	for _, sub := range ctrl.subscriptionManager.ListSubscription(ctx) {
		switch sub.Phase {
		case metadata.SubscriptionPhaseCreated:
			ctrl.scheduler.EnqueueNormalSubscription(sub.ID)
		case metadata.SubscriptionPhasePending, metadata.SubscriptionPhaseStopping:
			ctrl.scheduler.EnqueueSubscription(sub.ID)
		case metadata.SubscriptionPhaseToDelete:
			ctrl.needCleanSubscription[sub.ID] = sub.TriggerWorker
		}
	}
	return nil
}

func (ctrl *controller) membershipChangedProcessor(
	ctx context.Context, event member.MembershipChangedEvent,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (ctrl *controller) stop(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (ctrl *controller) Start() error { _ = "STUB: not implemented"; return nil }

func (ctrl *controller) Stop(ctx context.Context) { _ = "STUB: not implemented"; return }

func (ctrl *controller) initTriggerSystemEventbus() {
	_ = "STUB: not implemented"
	// avoid blocking starting
	return
}
