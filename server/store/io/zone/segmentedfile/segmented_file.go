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

package segmentedfile

import (
	"os"
	"sync"

	// first-party libraries.

	// this project.

	"github.com/vanus-labs/vanus/server/store/io/zone"
)

type SegmentedFile struct {
	segments []*Segment
	mu       sync.RWMutex

	dir         string
	ext         string
	segmentSize int64
}

// Make sure file implements zone.Interface.
var _ zone.Interface = (*SegmentedFile)(nil)

func Open(dir string, opts ...Option) (*SegmentedFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sf *SegmentedFile) Close() { _ = "STUB: not implemented"; return }

func (sf *SegmentedFile) Raw(off int64) (*os.File, int64) { _ = "STUB: not implemented"; return nil, 0 }

func (sf *SegmentedFile) SelectSegment(offset int64, autoCreate bool) *Segment {
	_ = "STUB: not implemented"
	return nil
}

// Fast return for append.

func (sf *SegmentedFile) createNextSegment(last *Segment) *Segment {
	_ = "STUB: not implemented"
	return nil
}

func createSegment(dir, ext string, so, size int64) (*Segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sf *SegmentedFile) firstSegment() *Segment { _ = "STUB: not implemented"; return nil }

func (sf *SegmentedFile) lastSegment() *Segment { _ = "STUB: not implemented"; return nil }

func (sf *SegmentedFile) Dir() string { _ = "STUB: not implemented"; return "" }

func (sf *SegmentedFile) Len() int { _ = "STUB: not implemented"; return 0 }
