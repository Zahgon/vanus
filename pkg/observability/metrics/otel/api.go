// Copyright 2023 Linkall Inc.
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

package otel

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel/attribute"
)

var (
	countMap          = make(map[string]ICounter, 0)
	gaugeMap          = make(map[string]IGauge, 0)
	histogramMap      = make(map[string]IHistogram, 0)
	metricCreateMutex = sync.Mutex{}
	emptyCount        = &promCounter{}
	emptyGauge        = &promGauge{}
	emptyHistogram    = &promHistogram{}
)

type ICounter interface {
	IncrInt(int64, ...attribute.KeyValue)
	IncrFloat(float64, ...attribute.KeyValue)
	Async(func(context.Context, ICounter))
}

type unit string

const (
	UnitMillisecond   = unit("ms")
	UnitByte          = unit("byte")
	UnitDimensionless = unit("1")
)

type metricKey struct {
	name        string
	description string
	unit        unit
}

func NewMetricKey(name string, u unit, desc string) *metricKey {
	_ = "STUB: not implemented"
	return nil
}

func newCounter(k *metricKey) ICounter { _ = "STUB: not implemented"; return *new(ICounter) }

type IGauge interface {
	IncrInt(int64, ...attribute.KeyValue)
	IncrFloat(float64, ...attribute.KeyValue)
	Async(func(context.Context, IGauge))
}

func newGauge(k *metricKey) IGauge { _ = "STUB: not implemented"; return *new(IGauge) }

type IHistogram interface {
	RecordInt(int64, ...attribute.KeyValue)
	RecordFloat(float64, ...attribute.KeyValue)
	Async(func(context.Context, IHistogram))
}

func newHistogram(k *metricKey) IHistogram { _ = "STUB: not implemented"; return *new(IHistogram) }

func GetCounter(key *metricKey) ICounter { _ = "STUB: not implemented"; return *new(ICounter) }

func GetGauge(key *metricKey) IGauge { _ = "STUB: not implemented"; return *new(IGauge) }

func GetHistogram(key *metricKey) IHistogram { _ = "STUB: not implemented"; return *new(IHistogram) }

func isValidKey(k *metricKey) bool { _ = "STUB: not implemented"; return false }
