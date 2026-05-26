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

package meta

import "github.com/huandu/skiplist"

type RangeCallback func(key []byte, value interface{}) error

type Ranger interface {
	Range(cb RangeCallback) error
}

type skiplistRange struct {
	l *skiplist.SkipList
}

var _ Ranger = (*skiplistRange)(nil)

func SkiplistRange(l *skiplist.SkipList) Ranger { _ = "STUB: not implemented"; return *new(Ranger) }

func (r *skiplistRange) Range(cb RangeCallback) error { _ = "STUB: not implemented"; return nil }

type kvRange struct {
	key   []byte
	value interface{}
}

var _ Ranger = (*kvRange)(nil)

func KVRange(key []byte, value interface{}) Ranger { _ = "STUB: not implemented"; return *new(Ranger) }

func (r *kvRange) Range(cb RangeCallback) error { _ = "STUB: not implemented"; return nil }

type deleteRange struct {
	keys [][]byte
}

var _ Ranger = (*deleteRange)(nil)

func (r *deleteRange) Range(cb RangeCallback) error { _ = "STUB: not implemented"; return nil }
