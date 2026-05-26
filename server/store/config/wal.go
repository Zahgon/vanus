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

package config

import (
	// standard libraries.

	"time"

	// this project.
	"github.com/vanus-labs/vanus/server/store/wal"
)

const (
	baseKB             = 1024
	baseMB             = 1024 * baseKB
	baseWALBlockSize   = 4 * baseKB
	minWALFlushTimeout = 200 * time.Microsecond
)

type WAL struct {
	BlockSize    int    `yaml:"block_size"`
	FileSize     uint64 `yaml:"file_size"`
	FlushTimeout string `yaml:"flush_timeout"`
	IO           `yaml:"io"`
}

func (c *WAL) Validate(minFileSize uint64) error { _ = "STUB: not implemented"; return nil }

func (c *WAL) Options() (opts []wal.Option) { _ = "STUB: not implemented"; return nil }

// unreachable
