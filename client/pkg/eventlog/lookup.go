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

package eventlog

import (
	// standard libraries.

	"time"

	// this project.
	"github.com/vanus-labs/vanus/client/pkg/primitive"
	"github.com/vanus-labs/vanus/client/pkg/record"
)

const (
	defaultWatchInterval = 30 * time.Second
)

type WritableSegmentWatcher struct {
	*primitive.Watcher
	ch chan *record.Segment
}

func (w *WritableSegmentWatcher) Chan() <-chan *record.Segment {
	_ = "STUB: not implemented"
	return nil
}

func (w *WritableSegmentWatcher) Start() { _ = "STUB: not implemented"; return }

func WatchWritableSegment(l *eventlog) *WritableSegmentWatcher {
	_ = "STUB: not implemented"
	return nil
}

type ReadableSegmentsWatcher struct {
	*primitive.Watcher
	ch chan []*record.Segment
}

func (w *ReadableSegmentsWatcher) Chan() <-chan []*record.Segment {
	_ = "STUB: not implemented"
	return nil
}

func (w *ReadableSegmentsWatcher) Start() { _ = "STUB: not implemented"; return }

func WatchReadableSegments(l *eventlog) *ReadableSegmentsWatcher {
	_ = "STUB: not implemented"
	return nil
}
