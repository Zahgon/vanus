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
	"time"
)

const (
	defaultBufferSize      = 1 << 10
	defaultDeliveryTimeout = 5 * time.Second
	defaultMaxWriteAttempt = 3
	defaultGoroutineSize   = 10000
	defaultMaxUACKNumber   = 10000
	defaultBatchSize       = 32
)

type Config struct {
	BufferSize        int
	MaxRetryAttempts  int32
	DeliveryTimeout   time.Duration
	RateLimit         uint32
	Controllers       []string
	MaxWriteAttempt   int
	Ordered           bool
	DisableDeadLetter bool

	GoroutineSize int
	SendBatchSize int
	PullBatchSize int
	MaxUACKNumber int
	TargetGateway *TargetGateway
}

type TargetGateway struct {
	Address          string `yaml:"address"`
	TargetHeaderName string `yaml:"header_name"`
}

func defaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

type Option func(t *trigger)

func WithBufferSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMaxRetryAttempts(attempts int32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDeliveryTimeout(timeout uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithOrdered(ordered bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRateLimit(rateLimit uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithControllers(controllers []string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDisableDeadLetter(disable bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGoroutineSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSendBatchSize(batchSize int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithPullBatchSize(batchSize int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMaxUACKNumber(maxUACKNumber int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProxy(proxy *TargetGateway) Option { _ = "STUB: not implemented"; return *new(Option) }
