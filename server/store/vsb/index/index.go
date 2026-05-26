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

//go:generate mockgen -source=index.go -destination=testing/mock_index.go -package=testing
package index

import (
	// this project.
	"github.com/vanus-labs/vanus/server/store/block"
)

type Index interface {
	StartOffset() int64
	EndOffset() int64
	Length() int32
	Stime() int64
}

type Option func(*index)

func WithEntry(entry block.Entry) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithStime(stime int64) Option { _ = "STUB: not implemented"; return *new(Option) }

func NewIndex(offset int64, length int32, opts ...Option) Index {
	_ = "STUB: not implemented"
	return *new(Index)
}

type index struct {
	offset int64
	length int32
	stime  int64
}

// Make sure index implements Index.
var _ Index = (*index)(nil)

func (i *index) StartOffset() int64 { _ = "STUB: not implemented"; return 0 }

func (i *index) EndOffset() int64 { _ = "STUB: not implemented"; return 0 }

func (i *index) Length() int32 { _ = "STUB: not implemented"; return 0 }

func (i *index) Stime() int64 { _ = "STUB: not implemented"; return 0 }
