// Copyright 2023 Linkall Inc.
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

// vanus resource.
package vsr

import (
	"errors"
)

const (
	emptyID = ID(0)
	base    = 16
	bitSize = 64
)

var ErrEmptyID = errors.New("id: empty")

type ID uint64

func EmptyID() ID { _ = "STUB: not implemented"; return *new(ID) }

func NewIDFromUint64(id uint64) ID { _ = "STUB: not implemented"; return *new(ID) }

func NewIDFromString(id string) (ID, error) { _ = "STUB: not implemented"; return *new(ID), nil }

func (id ID) String() string { _ = "STUB: not implemented"; return "" }

func (id ID) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }

func (id ID) Key() string { _ = "STUB: not implemented"; return "" }

func (id ID) Equals(cID ID) bool { _ = "STUB: not implemented"; return false }

type IDList []ID

func (l IDList) Contains(id ID) bool { _ = "STUB: not implemented"; return false }
