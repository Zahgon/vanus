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
	"errors"

	pb "github.com/vanus-labs/vanus/pkg/raft/raftpb"
	"github.com/vanus-labs/vanus/pkg/raft/tracker"
)

// ErrStepLocalMsg is returned when try to step a local raft message
var ErrStepLocalMsg = errors.New("raft: cannot step raft local message")

// ErrStepPeerNotFound is returned when try to step a response message
// but there is no peer found in raft.prs for that node.
var ErrStepPeerNotFound = errors.New("raft: cannot step as peer not found")

// RawNode is a thread-unsafe Node.
// The methods of this struct correspond to the methods of Node and are described
// more fully there.
type RawNode struct {
	raft *raft
}

// NewRawNode instantiates a RawNode from the given configuration.
//
// See Bootstrap() for bootstrapping an initial state; this replaces the former
// 'peers' argument to this method (with identical behavior). However, It is
// recommended that instead of calling Bootstrap, applications bootstrap their
// state manually by setting up a Storage that has a first index > 1 and which
// stores the desired ConfState as its InitialState.
func NewRawNode(config *Config) (*RawNode, error) { _ = "STUB: not implemented"; return nil, nil }

// Tick advances the internal logical clock by a single tick.
func (rn *RawNode) Tick() {
	_ = "STUB: not implemented"

	// TickQuiesced advances the internal logical clock by a single tick without
	// performing any other state machine processing. It allows the caller to avoid
	// periodic heartbeats and elections when all of the peers in a Raft group are
	// known to be at the same state. Expected usage is to periodically invoke Tick
	// or TickQuiesced depending on whether the group is "active" or "quiesced".
	//
	// WARNING: Be very careful about using this method as it subverts the Raft
	// state machine. You should probably be using Tick instead.
	return
}

func (rn *RawNode) TickQuiesced() { _ = "STUB: not implemented"; return }

// Campaign causes this RawNode to transition to candidate state.
func (rn *RawNode) Campaign() error { _ = "STUB: not implemented"; return nil }

// Propose proposes data be appended to the raft log.
func (rn *RawNode) Propose(pds ...ProposeData) { _ = "STUB: not implemented"; return }

// ProposeConfChange proposes a config change. See (Node).ProposeConfChange for
// details.
func (rn *RawNode) ProposeConfChange(cc pb.ConfChangeI) error {
	_ = "STUB: not implemented"
	return nil
}

// ApplyConfChange applies a config change to the local node. The app must call
// this when it applies a configuration change, except when it decides to reject
// the configuration change, in which case no call must take place.
func (rn *RawNode) ApplyConfChange(cc pb.ConfChangeI) *pb.ConfState {
	_ = "STUB: not implemented"
	return nil
}

func (rn *RawNode) ReportStateStatus(term uint64, vote uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (rn *RawNode) ReportLogStatus(index uint64, term uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (rn *RawNode) ReportApplyStatus(index uint64) error { _ = "STUB: not implemented"; return nil }

// Step advances the state machine using the given message.
func (rn *RawNode) Step(m pb.Message) error {
	_ = "STUB: not implemented"
	// ignore unexpected local messages receiving over network
	return nil
}

// Status returns the current status of the given group. This allocates, see
// BasicStatus and WithProgress for allocation-friendlier choices.
func (rn *RawNode) Status() Status { _ = "STUB: not implemented"; return *new(Status) }

// BasicStatus returns a BasicStatus. Notably this does not contain the
// Progress map; see WithProgress for an allocation-free way to inspect it.
func (rn *RawNode) BasicStatus() BasicStatus { _ = "STUB: not implemented"; return *new(BasicStatus) }

// ProgressType indicates the type of replica a Progress corresponds to.
type ProgressType byte

const (
	// ProgressTypePeer accompanies a Progress for a regular peer replica.
	ProgressTypePeer ProgressType = iota
	// ProgressTypeLearner accompanies a Progress for a learner replica.
	ProgressTypeLearner
)

// WithProgress is a helper to introspect the Progress for this node and its
// peers.
func (rn *RawNode) WithProgress(visitor func(id uint64, typ ProgressType, pr tracker.Progress)) {
	_ = "STUB: not implemented"
	return
}

// ReportUnreachable reports the given node is not reachable for the last send.
func (rn *RawNode) ReportUnreachable(id uint64) { _ = "STUB: not implemented"; return }

// ReportSnapshot reports the status of the sent snapshot.
func (rn *RawNode) ReportSnapshot(id uint64, status SnapshotStatus) {
	_ = "STUB: not implemented"
	return
}

// TransferLeader tries to transfer leadership to the given transferee.
func (rn *RawNode) TransferLeader(transferee uint64) { _ = "STUB: not implemented"; return }

// ReadIndex requests a read state. The read state will be set in ready.
// Read State has a read index. Once the application advances further than the read
// index, any linearizable read requests issued before the read request can be
// processed safely. The read state will have the same rctx attached.
func (rn *RawNode) ReadIndex(rctx []byte) { _ = "STUB: not implemented"; return }
