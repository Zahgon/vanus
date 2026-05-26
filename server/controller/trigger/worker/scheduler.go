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

package worker

import (
	"context"

	vanus "github.com/vanus-labs/vanus/api/vsr"

	"github.com/vanus-labs/vanus/pkg/queue"
	"github.com/vanus-labs/vanus/server/controller/trigger/subscription"
)

const (
	defaultRetryPrintLog = 5
)

type SubscriptionScheduler struct {
	normalQueue         queue.Queue
	maxRetryPrintLog    int
	policy              TriggerWorkerPolicy
	workerManager       Manager
	subscriptionManager subscription.Manager
	ctx                 context.Context
	stop                context.CancelFunc
}

func NewSubscriptionScheduler(workerManager Manager,
	subscriptionManager subscription.Manager,
) *SubscriptionScheduler {
	_ = "STUB: not implemented"
	return nil
}

func (s *SubscriptionScheduler) EnqueueSubscription(id vanus.ID) { _ = "STUB: not implemented"; return }

func (s *SubscriptionScheduler) EnqueueNormalSubscription(id vanus.ID) {
	_ = "STUB: not implemented"
	return
}

func (s *SubscriptionScheduler) Stop() { _ = "STUB: not implemented"; return }

func (s *SubscriptionScheduler) Run() { _ = "STUB: not implemented"; return }

func (s *SubscriptionScheduler) handler(ctx context.Context, subscriptionID vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}
