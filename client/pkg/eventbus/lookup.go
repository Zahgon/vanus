// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eventbus

import (
	// standard libraries.

	// first-party libraries.

	// this project.
	"github.com/vanus-labs/vanus/client/pkg/primitive"
	"github.com/vanus-labs/vanus/client/pkg/record"
)

type WritableLogsResult struct {
	Eventlogs []*record.Eventlog
	Err       error
}

type WritableLogsWatcher struct {
	*primitive.Watcher
	ch chan *WritableLogsResult
}

func (w *WritableLogsWatcher) Chan() <-chan *WritableLogsResult {
	_ = "STUB: not implemented"
	return nil
}

func (w *WritableLogsWatcher) Start() { _ = "STUB: not implemented"; return }

func WatchWritableLogs(bus *eventbus) *WritableLogsWatcher { _ = "STUB: not implemented"; return nil }

type ReadableLogsResult struct {
	Eventlogs []*record.Eventlog
	Err       error
}

type ReadableLogsWatcher struct {
	*primitive.Watcher
	ch chan *ReadableLogsResult
}

func (w *ReadableLogsWatcher) Chan() <-chan *ReadableLogsResult {
	_ = "STUB: not implemented"
	return nil
}

func (w *ReadableLogsWatcher) Start() { _ = "STUB: not implemented"; return }

func WatchReadableLogs(bus *eventbus) *ReadableLogsWatcher { _ = "STUB: not implemented"; return nil }
