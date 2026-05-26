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

package tracing

import (
	"context"

	"go.opentelemetry.io/otel/sdk/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
)

const (
	vanusVersion     = "v0.9.0"
	environmentKey   = "environment"
	environmentValue = "local"
)

type Config struct {
	ServerName    string `yaml:"-"`
	Enable        bool   `yaml:"enable"`
	OtelCollector string `yaml:"otel_collector"`
}

var tp *tracerProvider

func Init(cfg Config) { _ = "STUB: not implemented"; return }

// if otel_collector is empty, switch to noop tracer

// Test input two num, return sum.
func Test() { _ = "STUB: not implemented"; return }

func Start(ctx context.Context, pkgName, methodName string) (context.Context, oteltrace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(oteltrace.Span)
}

type tracerProvider struct {
	p          oteltrace.TracerProvider
	serverName string
}

type Tracer struct {
	tracer     oteltrace.Tracer
	kind       oteltrace.SpanKind
	moduleName string
}

func (t *Tracer) Start(ctx context.Context, methodName string, opts ...oteltrace.SpanStartOption) (context.Context, oteltrace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(oteltrace.Span)
}

func NewTracer(moduleName string, kind oteltrace.SpanKind) *Tracer {
	_ = "STUB: not implemented"
	return nil
}

func newTracerProvider(serviceName string, collectorEndpoint string) (*trace.TracerProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set up a trace exporter

// Register the trace exporter with a TracerProvider, using a batch
// span processor to aggregate spans before export.
