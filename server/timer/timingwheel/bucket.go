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

	"github.com/vanus-labs/vanus/client"
	"github.com/vanus-labs/vanus/client/pkg/api"
	primitive "github.com/vanus-labs/vanus/pkg"
	"github.com/vanus-labs/vanus/pkg/kv"
)

const (
	timerBuiltInEventbusReceivingStation    = "__Timer_RS"
	timerBuiltInEventbusDistributionStation = "__Timer_DS"
	timerBuiltInEventbus                    = "__Timer_%d_%d"
	xVanusEventbus                          = primitive.XVanusEventbus
	xVanusDeliveryTime                      = primitive.XVanusDeliveryTime
	sleepDuration                           = 100 * time.Millisecond
)

type timingMsg struct {
	expiration time.Time
	event      *ce.Event
}

func newTimingMsg(ctx context.Context, e *ce.Event) *timingMsg {
	_ = "STUB: not implemented"
	return nil
}

func (tm *timingMsg) hasExpired() bool { _ = "STUB: not implemented"; return false }

func (tm *timingMsg) getExpiration() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (tm *timingMsg) getEvent() *ce.Event { _ = "STUB: not implemented"; return nil }

type bucket struct {
	config   *Config
	tick     time.Duration
	interval time.Duration
	layer    int64
	slot     int64
	offset   int64
	eventbus string

	mu             sync.Mutex
	wg             sync.WaitGroup
	exitC          chan struct{}
	kvStore        kv.Client
	client         client.Client
	eventbusWriter api.BusWriter
	eventbusReader api.BusReader

	timingwheel *timingWheel
	element     *list.Element

	waitingForReady func(ctx context.Context, events []*ce.Event)
	eventHandler    func(ctx context.Context, event *ce.Event)
}

func newBucket(tw *timingWheel, element *list.Element, tick time.Duration, ebName string, layer, slot int64) *bucket {
	_ = "STUB: not implemented"
	return nil
}

func (b *bucket) start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *bucket) stop(ctx context.Context) { _ = "STUB: not implemented"; return }

func (b *bucket) run(ctx context.Context) { _ = "STUB: not implemented"; return }

// update offset asynchronously

// wait for all goroutines to finish before updating offset metadata

// limit the number of goroutines to no more than defaultMaxNumberOfWorkers

// batch read

// concurrent write

// block and wait here until events of the bucket ready to flow

// asynchronously update offset after the same batch of events are successfully written.

func (b *bucket) pushToDistributionStation(ctx context.Context, e *ce.Event) {
	_ = "STUB: not implemented"
	return
}

func (b *bucket) pushToPrevTimingWheel(ctx context.Context, e *ce.Event) {
	_ = "STUB: not implemented"
	return
}

func (b *bucket) waitingForExpired(ctx context.Context, events []*ce.Event) {
	_ = "STUB: not implemented"
	return
}

func (b *bucket) isReadyToDeliver(tm *timingMsg) bool { _ = "STUB: not implemented"; return false }

func (b *bucket) waitingForFlow(ctx context.Context, events []*ce.Event) {
	_ = "STUB: not implemented"
	return
}

func (b *bucket) isReadyToFlow(tm *timingMsg) bool { _ = "STUB: not implemented"; return false }

func (b *bucket) push(ctx context.Context, tm *timingMsg) bool {
	_ = "STUB: not implemented"
	return false
}

func (b *bucket) createEventbus(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *bucket) connectEventbus(ctx context.Context) { _ = "STUB: not implemented"; return }

func (b *bucket) putEvent(ctx context.Context, tm *timingMsg) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *bucket) getEvent(ctx context.Context, number int16) (events []*ce.Event, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(jiangkai): redesign here for reduce cpu overload, by jiangkai, 2022.09.16

func (b *bucket) updateOffsetMeta(ctx context.Context, offset int64) {
	_ = "STUB: not implemented"
	return
}

func (b *bucket) existsOffsetMeta(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (b *bucket) getOffsetMeta(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *bucket) deleteOffsetMeta(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *bucket) hasOnEnd(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func (b *bucket) recycle(ctx context.Context) {
	_ = "STUB: not implemented"
	// TODO(jiangkai): check for errors
	return
}

func (b *bucket) wait(_ context.Context) { _ = "STUB: not implemented"; return }

func (b *bucket) isLeader() bool { _ = "STUB: not implemented"; return false }

func (b *bucket) getEventbus() string { _ = "STUB: not implemented"; return "" }

func (b *bucket) getOffset() int64 { _ = "STUB: not implemented"; return 0 }

func (b *bucket) getTimingWheelElement() *timingWheelElement { _ = "STUB: not implemented"; return nil }

func (b *bucket) incOffset(diff int64) { _ = "STUB: not implemented"; return }
