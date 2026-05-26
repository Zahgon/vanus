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

package eventlog

import (
	"context"
	"time"

	metapb "github.com/vanus-labs/vanus/api/meta"
	vanus "github.com/vanus-labs/vanus/api/vsr"

	"github.com/vanus-labs/vanus/server/controller/eventbus/metadata"
)

type SegmentState string

const (
	StateCreated  = SegmentState("created")
	StateWorking  = SegmentState("working")
	StateFrozen   = SegmentState("frozen")
	StateArchived = SegmentState("archived")
	StateExpired  = SegmentState("expired")
)

type Segment struct {
	ID                 vanus.ID      `json:"id,omitempty"`
	Capacity           int64         `json:"capacity,omitempty"`
	EventlogID         vanus.ID      `json:"event_log_id,omitempty"`
	PreviousSegmentID  vanus.ID      `json:"previous_segment_id,omitempty"`
	NextSegmentID      vanus.ID      `json:"next_segment_id,omitempty"`
	StartOffsetInLog   int64         `json:"start_offset_in_log,omitempty"`
	Replicas           *ReplicaGroup `json:"replicas,omitempty"`
	State              SegmentState  `json:"state,omitempty"`
	Size               int64         `json:"size,omitempty"`
	Number             int32         `json:"number,omitempty"`
	FirstEventBornTime time.Time     `json:"first_event_born_time"`
	LastEventBornTime  time.Time     `json:"last_event_born_time"`
}

func (seg *Segment) IsAppendable() bool { _ = "STUB: not implemented"; return false }

func (seg *Segment) GetLeaderBlock() *metadata.Block { _ = "STUB: not implemented"; return nil }

func (seg *Segment) String() string { _ = "STUB: not implemented"; return "" }

// TODO Don't update field in here
func (seg *Segment) isNeedUpdate(newSeg Segment) bool { _ = "STUB: not implemented"; return false }

// TODO(wenfeng): follow state shift

func (seg *Segment) isFull() bool { _ = "STUB: not implemented"; return false }

func (seg *Segment) isReady() bool { _ = "STUB: not implemented"; return false }

func (seg *Segment) Copy() Segment { _ = "STUB: not implemented"; return *new(Segment) }

type ReplicaGroup struct {
	ID vanus.ID `json:"id"`
	// the id of LeaderBlock
	Leader uint64 `json:"leader"`
	// blockID *metadata.Block
	Peers     map[uint64]*metadata.Block `json:"blocks"`
	Term      uint64                     `json:"term"`
	CreateAt  time.Time                  `json:"create_at"`
	DestroyAt time.Time                  `json:"destroy_at"`
}

func Convert2ProtoSegment(ctx context.Context, ins ...Segment) []*metapb.Segment {
	_ = "STUB: not implemented"
	return nil
}
