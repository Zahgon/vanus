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

package io

import (
	// standard libraries.
	"os"

	// third-party libraries.
	"github.com/ncw/directio"
	// first-party libraries.
)

const (
	defaultFilePerm = 0o644
	baseFillSize    = 4 * 1024 // 4KB
	fastFillVecSize = 11       // Up to 4MB
)

var fastFillVec []fastFillTemplate

func init() { //nolint:gochecknoinits // Initialize fastFileVec.
	fastFillVec = make([]fastFillTemplate, fastFillVecSize)

	n := fastFillVecSize - 1
	for i := 0; i <= n; i++ {
		size := (1 << (n - i)) * baseFillSize
		fastFillVec[i] = fastFillTemplate{
			size: int64(size),
			data: directio.AlignedBlock(size),
		}
	}
}

type fastFillTemplate struct {
	size int64
	data []byte
}

func OpenFile(path string, flag int, sync bool, direct bool) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateFile(path string, size int64, flag int, sync bool, direct bool) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ResizeFile(f *os.File, size int64) error { _ = "STUB: not implemented"; return nil }

func doCreateFile(path string, size int64, flag int, sync bool, direct bool) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Resize file.

func doCreateFileAndWarm( //nolint:unused,nolintlint // use in linux.
	path string, size int64, flag int, sync bool, direct bool,
) (*os.File, error) {
	_ = "STUB: not implemented"
	// Create file.
	return nil, nil
}

// Resize file.

// Warm file.

func openFile(path string, flag int, sync bool, direct bool) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func makeFlag(flag int, sync bool) int { _ = "STUB: not implemented"; return 0 }

func warmFile(f *os.File, size int64) error {
	_ = "STUB: not implemented" //nolint:unused,nolintlint // use in linux.
	return nil
}
