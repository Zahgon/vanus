// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package record

import "time"

type Eventbus struct {
	Name string
	Logs []*Eventlog
}

type Eventlog struct {
	ID   uint64
	Mode LogMode
}

func (l *Eventlog) Readable() bool { _ = "STUB: not implemented"; return false }

func (l *Eventlog) Writable() bool { _ = "STUB: not implemented"; return false }

// WritableLog return all writable logs.
func (b *Eventbus) WritableLog() []*Eventlog { _ = "STUB: not implemented"; return nil }

func allWritable(ls []*Eventlog) bool { _ = "STUB: not implemented"; return false }

// ReadableLog return all readable logs.
func (b *Eventbus) ReadableLog() []*Eventlog { _ = "STUB: not implemented"; return nil }

func allReadable(ls []*Eventlog) bool { _ = "STUB: not implemented"; return false }

type Segment struct {
	ID               uint64
	StartOffset      int64
	EndOffset        int64
	Writable         bool
	FirstEventBornAt time.Time
	LastEventBornAt  time.Time

	Blocks        map[uint64]*Block
	LeaderBlockID uint64
}

type Block struct {
	ID       uint64
	Endpoint string
}

func (s *Segment) GetLeaderEndpoint() string { _ = "STUB: not implemented"; return "" }
