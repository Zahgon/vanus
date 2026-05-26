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

package queue

import (
	"time"

	"k8s.io/client-go/util/workqueue"

	vanus "github.com/vanus-labs/vanus/api/vsr"
)

type Queue interface {
	Add(key vanus.ID)
	Len() int
	Get() (value vanus.ID, shutdown bool)
	Done(key vanus.ID)
	ShutDown()
	IsShutDown() bool
	ReAdd(key vanus.ID)
	GetFailNum(key vanus.ID) int
	ClearFailNum(key vanus.ID)
}

const (
	limitSize      = 10
	burst          = 100
	limitBaseDelay = 500 * time.Millisecond
	limitMaxDelay  = 10 * time.Second
)

type queue struct {
	queue workqueue.RateLimitingInterface
}

func New() Queue { _ = "STUB: not implemented"; return *new(Queue) }

func DefaultControllerRateLimiter() workqueue.RateLimiter {
	_ = "STUB: not implemented"
	return *new(workqueue.RateLimiter)
}

// 10 qps, 100 bucket size.  This is only for retry speed and its only the overall factor (not per item)

func (q *queue) Add(key vanus.ID) { _ = "STUB: not implemented"; return }

func (q *queue) Len() int { _ = "STUB: not implemented"; return 0 }

func (q *queue) Get() (value vanus.ID, shutdown bool) {
	_ = "STUB: not implemented"
	return *new(vanus.ID), false
}

func (q *queue) Done(key vanus.ID) { _ = "STUB: not implemented"; return }

func (q *queue) ShutDown() { _ = "STUB: not implemented"; return }

func (q *queue) IsShutDown() bool { _ = "STUB: not implemented"; return false }

func (q *queue) ReAdd(key vanus.ID) { _ = "STUB: not implemented"; return }

func (q *queue) GetFailNum(key vanus.ID) int { _ = "STUB: not implemented"; return 0 }

func (q *queue) ClearFailNum(key vanus.ID) { _ = "STUB: not implemented"; return }
