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

const (
	defaultExt         = ""
	defaultSegmentSize = int64(64 * 1024 * 1024)
)

type config struct {
	ext         string
	segmentSize int64
}

func defaultConfig() config { _ = "STUB: not implemented"; return *new(config) }

type Option func(*config)

func makeConfig(opts ...Option) config { _ = "STUB: not implemented"; return *new(config) }

func WithSegmentSize(size int64) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithExtension(ext string) Option { _ = "STUB: not implemented"; return *new(Option) }
