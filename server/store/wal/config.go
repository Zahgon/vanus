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

package wal

import (
	// standard libraries.
	"time"

	// this project.
	ioengine "github.com/vanus-labs/vanus/server/store/io/engine"
	"github.com/vanus-labs/vanus/server/store/io/stream"
	"github.com/vanus-labs/vanus/server/store/io/zone/segmentedfile"
)

const (
	logFileExt              = ".log"
	defaultBlockSize        = 16 * 1024
	defaultFileSize         = 128 * 1024 * 1024
	defaultAppendBufferSize = 64
)

type config struct {
	pos            int64
	cb             OnEntryCallback
	blockSize      int
	fileSize       int64
	flushDelayTime time.Duration // default: 3 * time.Millisecond
	engine         ioengine.Interface
	readOnly       bool
}

func (cfg *config) segmentedFileOptions() []segmentedfile.Option {
	_ = "STUB: not implemented"
	return nil
}

func (cfg *config) streamSchedulerOptions() []stream.Option { _ = "STUB: not implemented"; return nil }

func defaultConfig() config { _ = "STUB: not implemented"; return *new(config) }

type Option func(*config)

func makeConfig(opts ...Option) config { _ = "STUB: not implemented"; return *new(config) }

func FromPosition(pos int64) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRecoveryCallback(cb OnEntryCallback) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithBlockSize(blockSize int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithFileSize(fileSize int64) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithFlushDelayTime(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithIOEngine(engine ioengine.Interface) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithReadOnly() Option { _ = "STUB: not implemented"; return *new(Option) }
