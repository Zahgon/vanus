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

package metadata

import (
	"time"

	"github.com/vanus-labs/vanus/api/meta"
	vanus "github.com/vanus-labs/vanus/api/vsr"
)

type Eventbus struct {
	ID          vanus.ID    `json:"id"`
	Name        string      `json:"name"`
	LogNumber   int         `json:"log_number"`
	Eventlogs   []*Eventlog `json:"eventlogs"`
	Description string      `json:"description"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	NamespaceID uint64      `json:"namespace_id"`
}

func Convert2ProtoEventbus(ins ...*Eventbus) []*meta.Eventbus {
	_ = "STUB: not implemented"
	return nil
}

type Eventlog struct {
	// global unique id
	ID            vanus.ID `json:"id"`
	EventbusID    vanus.ID `json:"eventbus_id"`
	EventbusName  string   `json:"eventbus_name"`
	SegmentNumber int      `json:"segment_number"`
}

func (el *Eventlog) Eventbus() string { _ = "STUB: not implemented"; return "" }

func Convert2ProtoEventlog(ins ...*Eventlog) []*meta.Eventlog {
	_ = "STUB: not implemented"
	return nil
}

type VolumeMetadata struct {
	ID       vanus.ID          `json:"id"`
	Capacity int64             `json:"capacity"`
	Used     int64             `json:"used"`
	Blocks   map[uint64]*Block `json:"blocks"`
}

type Block struct {
	ID         vanus.ID `json:"id"`
	Capacity   int64    `json:"capacity"`
	Size       int64    `json:"size"`
	VolumeID   vanus.ID `json:"volume_id"`
	EventlogID vanus.ID `json:"eventlog_id"`
	SegmentID  vanus.ID `json:"segment_id"`
}

func (bl *Block) String() string { _ = "STUB: not implemented"; return "" }
