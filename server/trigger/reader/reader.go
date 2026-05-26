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

//go:generate mockgen -source=reader.go  -destination=mock_reader.go -package=reader
package reader

import (
	"context"
	"sync"
	"time"

	ce "github.com/cloudevents/sdk-go/v2"

	vanus "github.com/vanus-labs/vanus/api/vsr"
	eb "github.com/vanus-labs/vanus/client"
	"github.com/vanus-labs/vanus/client/pkg/api"
	"github.com/vanus-labs/vanus/server/trigger/info"
)

const (
	lookupReadableLogsTimeout = 5 * time.Second
	readEventTimeout          = 5 * time.Second
	readErrSleepTime          = 2 * time.Second
	checkEventlogInterval     = 2 * time.Minute
	logFrequencyMini          = 10
)

type Config struct {
	EventbusID        vanus.ID
	Client            eb.Client
	SubscriptionID    vanus.ID
	SubscriptionIDStr string
	EventbusIDStr     string
	Offset            EventlogOffset
	BatchSize         int
}
type EventlogOffset map[vanus.ID]uint64

type Reader interface {
	Start() error
	Close()
}

type reader struct {
	config      Config
	events      chan<- info.EventRecord
	stop        context.CancelFunc
	wg          sync.WaitGroup
	eventlogMap map[uint64]*eventlogReader
}

func NewReader(config Config, events chan<- info.EventRecord) Reader {
	_ = "STUB: not implemented"
	return *new(Reader)
}

func (r *reader) Close() { _ = "STUB: not implemented"; return }

func (r *reader) findEventlog(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *reader) Start() error { _ = "STUB: not implemented"; return nil }

func (r *reader) startEventlog(ctx context.Context, l api.Eventlog) {
	_ = "STUB: not implemented"
	return
}

func (r *reader) getOffset(eventlogID vanus.ID) uint64 { _ = "STUB: not implemented"; return 0 }

type eventlogReader struct {
	config        Config
	eventlogID    vanus.ID
	eventlogIDStr string
	policy        api.ReadPolicy
	events        chan<- info.EventRecord
	offset        uint64
	cancel        context.CancelFunc
}

func (elReader *eventlogReader) stop() { _ = "STUB: not implemented"; return }

// getOffset get earliest offset.
func (elReader *eventlogReader) getOffset(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (elReader *eventlogReader) run(parentCtx context.Context) { _ = "STUB: not implemented"; return }

func (elReader *eventlogReader) loop(ctx context.Context, lr api.BusReader) error {
	_ = "STUB: not implemented"
	return nil
}

func (elReader *eventlogReader) putEvent(ctx context.Context, event info.EventRecord) error {
	_ = "STUB: not implemented"
	return nil
}

func readEvents(ctx context.Context, lr api.BusReader) ([]*ce.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
