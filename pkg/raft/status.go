// Copyright 2015 The etcd Authors
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

package raft

import (
	pb "github.com/vanus-labs/vanus/pkg/raft/raftpb"
	"github.com/vanus-labs/vanus/pkg/raft/tracker"
)

// Status contains information about this Raft peer and its view of the system.
// The Progress is only populated on the leader.
type Status struct {
	BasicStatus
	Config   tracker.Config
	Progress map[uint64]tracker.Progress
}

// BasicStatus contains basic information about the Raft peer. It does not allocate.
type BasicStatus struct {
	ID uint64

	pb.HardState
	SoftState

	Applied uint64

	LeadTransferee uint64
}

func getProgressCopy(r *raft) map[uint64]tracker.Progress { _ = "STUB: not implemented"; return nil }

func getBasicStatus(r *raft) BasicStatus { _ = "STUB: not implemented"; return *new(BasicStatus) }

// getStatus gets a copy of the current raft status.
func getStatus(r *raft) Status { _ = "STUB: not implemented"; return *new(Status) }

// MarshalJSON translates the raft status into JSON.
// TODO: try to simplify this by introducing ID type into raft
func (s Status) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// remove the trailing ","

func (s Status) String() string { _ = "STUB: not implemented"; return "" }
