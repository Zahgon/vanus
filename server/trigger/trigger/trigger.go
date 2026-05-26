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

//go:generate mockgen -source=trigger.go -destination=mock_trigger.go -package=trigger
package trigger

import (
	// standard libraries.
	"context"
	"sync"

	// third-party libraries.
	ce "github.com/cloudevents/sdk-go/v2"
	"github.com/panjf2000/ants/v2"
	"go.uber.org/ratelimit"

	// first-party libraries.
	vanus "github.com/vanus-labs/vanus/api/vsr"
	eb "github.com/vanus-labs/vanus/client"
	"github.com/vanus-labs/vanus/client/pkg/api"

	// this project.
	primitive "github.com/vanus-labs/vanus/pkg"
	pInfo "github.com/vanus-labs/vanus/pkg/info"
	"github.com/vanus-labs/vanus/server/trigger/client"
	"github.com/vanus-labs/vanus/server/trigger/filter"
	"github.com/vanus-labs/vanus/server/trigger/info"
	"github.com/vanus-labs/vanus/server/trigger/offset"
	"github.com/vanus-labs/vanus/server/trigger/reader"
	"github.com/vanus-labs/vanus/server/trigger/transform"
)

type State string

const (
	TriggerCreated   State = "created"
	TriggerPending   State = "pending"
	TriggerRunning   State = "running"
	TriggerSleep     State = "sleep"
	TriggerPaused    State = "paused"
	TriggerStopped   State = "stopped"
	TriggerDestroyed State = "destroyed"
)

type Trigger interface {
	Init(ctx context.Context) error
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Change(ctx context.Context, subscription *primitive.Subscription) error
	GetOffsets(ctx context.Context) pInfo.ListOffsetInfo
}

type trigger struct {
	subscriptionIDStr string
	eventbusIDStr     string

	subscription  *primitive.Subscription
	offsetManager *offset.SubscriptionOffset
	reader        reader.Reader
	eventCh       chan info.EventRecord
	sendCh        chan *toSendEvent
	batchSendCh   chan []*toSendEvent
	eventCli      client.EventClient
	client        eb.Client
	filter        filter.Filter
	transformer   *transform.Transformer
	rateLimiter   ratelimit.Limiter
	config        Config
	batch         bool

	retryEventCh     chan info.EventRecord
	retryEventReader reader.Reader
	timerEventWriter api.BusWriter
	dlEventWriter    api.BusWriter

	state State
	stop  context.CancelFunc
	lock  sync.RWMutex
	wg    primitive.Group

	pool *ants.Pool
}

type toSendEvent struct {
	retry     bool
	record    info.EventRecord
	transform *ce.Event
}

func NewTrigger(subscription *primitive.Subscription, opts ...Option) (Trigger, error) {
	_ = "STUB: not implemented"
	return *new(Trigger), nil
}

func newTrigger(subscription *primitive.Subscription, opts ...Option) (*trigger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *trigger) applyOptions(opts ...Option) { _ = "STUB: not implemented"; return }

func (t *trigger) getConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

func (t *trigger) getClient() client.EventClient {
	_ = "STUB: not implemented"
	return *new(client.EventClient)
}

func (t *trigger) changeTarget(
	sink primitive.URI, protocol primitive.Protocol, credential primitive.SinkCredential,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *trigger) getFilter() filter.Filter { _ = "STUB: not implemented"; return *new(filter.Filter) }

func (t *trigger) changeFilter(filters []*primitive.SubscriptionFilter) {
	_ = "STUB: not implemented"
	return
}

func (t *trigger) getTransformer() *transform.Transformer { _ = "STUB: not implemented"; return nil }

func (t *trigger) changeTransformer(transformer *primitive.Transformer) {
	_ = "STUB: not implemented"
	// FIXME(james.yin): encounter error?
	return
}

func (t *trigger) changeConfig(config primitive.SubscriptionConfig) {
	_ = "STUB: not implemented"
	return
}

// eventArrived for test.
func (t *trigger) eventArrived(ctx context.Context, event info.EventRecord) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *trigger) transformEvent(record info.EventRecord, retry bool) (*toSendEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// transform will chang event which lost origin event

func (t *trigger) sendEvent(ctx context.Context, events ...*ce.Event) client.Result {
	_ = "STUB: not implemented"
	return *new(client.Result)
}

func (t *trigger) runRetryEventFilterTransform(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (t *trigger) runEventFilterTransform(ctx context.Context) { _ = "STUB: not implemented"; return }

func (t *trigger) runEventToBatch(ctx context.Context) { _ = "STUB: not implemented"; return }

////nolint:gomnd

func (t *trigger) runEventSend(ctx context.Context) { _ = "STUB: not implemented"; return }

func (t *trigger) processEvent(ctx context.Context, events ...*toSendEvent) {
	_ = "STUB: not implemented"

	// commit offset
	return
}

// todo retry util success, now ordered event no need retry direct into dead letter

func (t *trigger) writeFailEvent(ctx context.Context, e *ce.Event, code int, err error) {
	_ = "STUB: not implemented"
	return
}

// get attempts

// dead letter

// retry

func (t *trigger) writeEventToRetry(ctx context.Context, e *ce.Event, attempts int32) {
	_ = "STUB: not implemented"
	return
}

func (t *trigger) writeEventToDeadLetter(ctx context.Context, e *ce.Event, reason, errorMsg string) {
	_ = "STUB: not implemented"
	return
}

func (t *trigger) getReaderConfig() reader.Config {
	_ = "STUB: not implemented"
	return *new(reader.Config)
}

func (t *trigger) getRetryEventReaderConfig() reader.Config {
	_ = "STUB: not implemented"
	return *new(reader.Config)
}

// getOffset from subscription.
func getOffset(sub *primitive.Subscription) map[vanus.ID]uint64 {
	_ = "STUB: not implemented"
	// get offset from subscription
	return nil
}

func (t *trigger) Init(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (t *trigger) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// eb event

// retry event

func (t *trigger) Stop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (t *trigger) Change(_ context.Context, subscription *primitive.Subscription) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:contextcheck // wrong advice

// GetOffsets contains retry eventlog.
func (t *trigger) GetOffsets(_ context.Context) pInfo.ListOffsetInfo {
	_ = "STUB: not implemented"
	return *new(pInfo.ListOffsetInfo)
}
