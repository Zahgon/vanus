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
	// this project.
)

const (
	defaultDirPerm = 0o755
)

// recoverSegments rebuilds segments from specified directory.
func recoverSegments(dir string, cfg config) ([]*Segment, error) {
	_ = "STUB: not implemented"
	// Make sure the directory exists.
	return nil, nil
}

// Delete discard files.

func scanSegmentFiles(dir, ext string, segmentSize int64) (segments []*Segment, discards []*Segment, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Rebuild log stream.

// discontinuous log file

// TODO(james.yin): return error

func filterRegularFiles(entries []os.DirEntry, ext string) []os.DirEntry {
	_ = "STUB: not implemented"
	return nil
}
