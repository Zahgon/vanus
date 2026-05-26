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

package timingwheel

import (
	"container/list"
	"context"
	"sync"
	"time"

	ce "github.com/cloudevents/sdk-go/v2"

	"github.com/vanus-labs/vanus/api/cluster"
	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	"github.com/vanus-labs/vanus/client"

	"github.com/vanus-labs/vanus/pkg/kv"
	"github.com/vanus-labs/vanus/pkg/kv/etcd"
)

const (
	// check waiting period every 1/defaultCheckWaitingPeriodRatio tick time by default.
	defaultCheckWaitingPeriodRatio = 10

	// frequent check waiting period every 1/defaultFrequentCheckWaitingPeriodRatio tick time by default.
	defaultFrequentCheckWaitingPeriodRatio = 100

	// number of tick flow in advance by default.
	defaultNumberOfTickFlowInAdvance = 1

	// number of events read each time by default.
	defaultNumberOfEventsRead = 10

	// the max number of workers by default.
	defaultMaxNumberOfWorkers = 1000

	recycleInterval = 60 * time.Second
)

var newEtcdClientV3 = etcd.NewEtcdClientV3

type Manager interface {
	Init(ctx context.Context) error
	Start(ctx context.Context) error
	Push(ctx context.Context, e *ce.Event) bool
	SetLeader(isleader bool)
	IsLeader() bool
	IsDeployed(ctx context.Context) bool
	Recover(ctx context.Context) error
	StopNotify() <-chan struct{}
	Stop(ctx context.Context)
}

// timingWheel timewheel contains multiple layers.
type timingWheel struct {
	config  *Config
	kvStore kv.Client
	ctrlCli ctrlpb.EventbusControllerClient
	client  client.Client
	cache   sync.Map
	twList  *list.List // element: *timingWheelElement

	ctrl cluster.Cluster

	receivingStation    *bucket
	distributionStation *bucket

	leader bool
	exitC  chan struct{}
	wg     sync.WaitGroup
}

func NewTimingWheel(c *Config) Manager { _ = "STUB: not implemented"; return *new(Manager) }

// Init the current timing wheel.
func (tw *timingWheel) Init(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Init Hierarchical Timing Wheels.

// Start starts the current timing wheel.
func (tw *timingWheel) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// here is to wait for the leader to complete the creation of all eventbus

// start distribution station for scheduled events distributing

// start all bucket of each layer

// start receiving station for scheduled events receiving

// start bucket recycling

func (tw *timingWheel) StopNotify() <-chan struct{} {
	_ = "STUB: not implemented"

	// Stop stops the current timing wheel.
	return nil
}

func (tw *timingWheel) Stop(ctx context.Context) { _ = "STUB: not implemented"; return }

// wait for all goroutine to end

func (tw *timingWheel) SetLeader(isLeader bool) { _ = "STUB: not implemented"; return }

func (tw *timingWheel) IsLeader() bool { _ = "STUB: not implemented"; return false }

func (tw *timingWheel) IsDeployed(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (tw *timingWheel) Recover(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// no offset metadata, no recovery required

// Push the scheduled event to the timingwheel.
func (tw *timingWheel) Push(ctx context.Context, e *ce.Event) bool {
	_ = "STUB: not implemented"
	return false
}

// Already expired

func (tw *timingWheel) getReceivingStation() *bucket { _ = "STUB: not implemented"; return nil }

func (tw *timingWheel) getDistributionStation() *bucket { _ = "STUB: not implemented"; return nil }

func (tw *timingWheel) startRecycling(ctx context.Context) { _ = "STUB: not implemented"; return }

func (tw *timingWheel) startReceivingStation(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

const receiveGoroutineNum = 2

// runReceivingStation as the unified entrance of scheduled events and pushed to the timingwheel.
func (tw *timingWheel) runReceivingStation(ctx context.Context) { _ = "STUB: not implemented"; return }

// update offset asynchronously

// wait for all goroutines to finish before updating offset metadata

// limit the number of goroutines to no more than defaultMaxNumberOfWorkers

// batch read

// concurrent write

// TODO(jiangkai): check event dst eventbus is vaild.

// asynchronously update offset after the same batch of events are successfully written

func (tw *timingWheel) startDistributionStation(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// runDistributionStation as the unified exit of scheduled events and popped to the timingwheel.
func (tw *timingWheel) runDistributionStation(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// update offset asynchronously

// wait for all goroutines to finish before updating offset metadata

// limit the number of goroutines to no more than defaultMaxNumberOfWorkers

// batch read

// concurrent write

// asynchronously update offset after the same batch of events are successfully written

func (tw *timingWheel) deliver(ctx context.Context, e *ce.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// timingWheelElement timingwheelelement has N number of buckets, every bucket is an eventbus.
type timingWheelElement struct {
	config   *Config
	kvStore  kv.Client
	ctrlCli  ctrlpb.EventbusControllerClient
	tick     time.Duration
	layer    int64
	interval time.Duration
	buckets  map[int64]*bucket

	exitC chan struct{}
	mu    sync.RWMutex
	wg    sync.WaitGroup

	timingwheel *timingWheel
	element     *list.Element

	pushHandler func(ctx context.Context, tm *timingMsg) bool
}

// newTimingWheel is an internal helper function that really creates an instance of TimingWheel.
func newTimingWheelElement(tw *timingWheel, tick time.Duration, layer int64) *timingWheelElement {
	_ = "STUB: not implemented"
	return nil
}

func (twe *timingWheelElement) push(ctx context.Context, tm *timingMsg) bool {
	_ = "STUB: not implemented"
	return false
}

// Put it into its own bucket

// Out of the interval. Put it into the overflow wheel

func (twe *timingWheelElement) pushBack(ctx context.Context, tm *timingMsg) bool {
	_ = "STUB: not implemented"
	return false
}

// Put it into its own bucket

func (twe *timingWheelElement) allowPush(tm *timingMsg) bool {
	_ = "STUB: not implemented"
	return false
}

func (twe *timingWheelElement) flow(ctx context.Context, tm *timingMsg) bool {
	_ = "STUB: not implemented"
	return false
}

// Put it into its own bucket

func (twe *timingWheelElement) calculateIndex(tm *timingMsg) int64 {
	_ = "STUB: not implemented"
	// the timing message comes from the timingwheel of the upper layer
	return 0
}

// Put it into its buffer bucket

// Put it into its own bucket

func (twe *timingWheelElement) makeSureBucketExist(ctx context.Context, index int64) error {
	_ = "STUB: not implemented"
	// TODO(jiangkai): redesign locks if here is a performance bottleneck in the future, by jiangkai, 2022.09.16
	// the segmented lock may solve the problem.
	return nil
}

func (twe *timingWheelElement) recycling(ctx context.Context) { _ = "STUB: not implemented"; return }

func (twe *timingWheelElement) wait(_ context.Context) { _ = "STUB: not implemented"; return }

func (twe *timingWheelElement) getBuckets() map[int64]*bucket {
	_ = "STUB: not implemented"
	return nil
}

func (twe *timingWheelElement) setElement(element *list.Element) { _ = "STUB: not implemented"; return }

func (twe *timingWheelElement) prev() *timingWheelElement { _ = "STUB: not implemented"; return nil }

func (twe *timingWheelElement) next() *timingWheelElement { _ = "STUB: not implemented"; return nil }
