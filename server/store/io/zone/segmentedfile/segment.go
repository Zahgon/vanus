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
	// standard libraries.
	"os"
)

type Segment struct {
	// so is the start offset of the log file.
	so int64
	// eo is the end offset of the log file.
	eo int64
	f  *os.File

	size int64
	path string
}

func newSegment(path string, so int64, size int64, f *os.File) *Segment {
	_ = "STUB: not implemented"
	return nil
}

func (s *Segment) Close() error { _ = "STUB: not implemented"; return nil }

func (s *Segment) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (s *Segment) SO() int64 { _ = "STUB: not implemented"; return 0 }

func (s *Segment) EO() int64 { _ = "STUB: not implemented"; return 0 }

func (s *Segment) File() *os.File { _ = "STUB: not implemented"; return nil }
