// Copyright 2019 The etcd Authors
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

package tracker

import (
	"github.com/vanus-labs/vanus/pkg/raft/quorum"
	pb "github.com/vanus-labs/vanus/pkg/raft/raftpb"
)

// Config reflects the configuration tracked in a ProgressTracker.
type Config struct {
	Voters quorum.JointConfig
	// AutoLeave is true if the configuration is joint and a transition to the
	// incoming configuration should be carried out automatically by Raft when
	// this is possible. If false, the configuration will be joint until the
	// application initiates the transition manually.
	AutoLeave bool
	// Learners is a set of IDs corresponding to the learners active in the
	// current configuration.
	//
	// Invariant: Learners and Voters does not intersect, i.e. if a peer is in
	// either half of the joint config, it can't be a learner; if it is a
	// learner it can't be in either half of the joint config. This invariant
	// simplifies the implementation since it allows peers to have clarity about
	// its current role without taking into account joint consensus.
	Learners map[uint64]struct{}
	// When we turn a voter into a learner during a joint consensus transition,
	// we cannot add the learner directly when entering the joint state. This is
	// because this would violate the invariant that the intersection of
	// voters and learners is empty. For example, assume a Voter is removed and
	// immediately re-added as a learner (or in other words, it is demoted):
	//
	// Initially, the configuration will be
	//
	//   voters:   {1 2 3}
	//   learners: {}
	//
	// and we want to demote 3. Entering the joint configuration, we naively get
	//
	//   voters:   {1 2} & {1 2 3}
	//   learners: {3}
	//
	// but this violates the invariant (3 is both voter and learner). Instead,
	// we get
	//
	//   voters:   {1 2} & {1 2 3}
	//   learners: {}
	//   next_learners: {3}
	//
	// Where 3 is now still purely a voter, but we are remembering the intention
	// to make it a learner upon transitioning into the final configuration:
	//
	//   voters:   {1 2}
	//   learners: {3}
	//   next_learners: {}
	//
	// Note that next_learners is not used while adding a learner that is not
	// also a voter in the joint config. In this case, the learner is added
	// right away when entering the joint configuration, so that it is caught up
	// as soon as possible.
	LearnersNext map[uint64]struct{}
}

func (c Config) String() string { _ = "STUB: not implemented"; return "" }

// Clone returns a copy of the Config that shares no memory with the original.
func (c *Config) Clone() Config { _ = "STUB: not implemented"; return *new(Config) }

// ProgressTracker tracks the currently active configuration and the information
// known about the nodes and learners in it. In particular, it tracks the match
// index for each peer which in turn allows reasoning about the committed index.
type ProgressTracker struct {
	Config

	Progress ProgressMap

	Votes map[uint64]bool

	MaxInflight int
}

// MakeProgressTracker initializes a ProgressTracker.
func MakeProgressTracker(maxInflight int) ProgressTracker {
	_ = "STUB: not implemented"
	return *new(ProgressTracker)
}

// only populated when used

// only populated when used
// only populated when used

// ConfState returns a ConfState representing the active configuration.
func (p *ProgressTracker) ConfState() pb.ConfState {
	_ = "STUB: not implemented"
	return *new(pb.ConfState)
}

// IsSingleton returns true if (and only if) there is only one voting member
// (i.e. the leader) in the current configuration.
func (p *ProgressTracker) IsSingleton() bool { _ = "STUB: not implemented"; return false }

type matchAckIndexer map[uint64]*Progress

var _ quorum.AckedIndexer = matchAckIndexer(nil)

// AckedIndex implements IndexLookuper.
func (l matchAckIndexer) AckedIndex(id uint64) (quorum.Index, bool) {
	_ = "STUB: not implemented"
	return *new(quorum.Index), false
}

// Committed returns the largest log index known to be committed based on what
// the voting members of the group have acknowledged.
func (p *ProgressTracker) Committed() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *ProgressTracker) AllAcked() uint64 { _ = "STUB: not implemented"; return 0 }

// TODO(james.yin): tolerate offline node.

func insertionSort(sl []uint64) { _ = "STUB: not implemented"; return }

// Visit invokes the supplied closure for all tracked progresses in stable order.
func (p *ProgressTracker) Visit(f func(id uint64, pr *Progress)) {
	_ = "STUB: not implemented"

	// We need to sort the IDs and don't want to allocate since this is hot code.
	// The optimization here mirrors that in `(MajorityConfig).CommittedIndex`,
	// see there for details.
	return
}

// QuorumActive returns true if the quorum is active from the view of the local
// raft state machine. Otherwise, it returns false.
func (p *ProgressTracker) QuorumActive() bool { _ = "STUB: not implemented"; return false }

// VoterNodes returns a sorted slice of voters.
func (p *ProgressTracker) VoterNodes() []uint64 { _ = "STUB: not implemented"; return nil }

// LearnerNodes returns a sorted slice of learners.
func (p *ProgressTracker) LearnerNodes() []uint64 { _ = "STUB: not implemented"; return nil }

// ResetVotes prepares for a new round of vote counting via recordVote.
func (p *ProgressTracker) ResetVotes() { _ = "STUB: not implemented"; return }

// RecordVote records that the node with the given id voted for this Raft
// instance if v == true (and declined it otherwise).
func (p *ProgressTracker) RecordVote(id uint64, v bool) { _ = "STUB: not implemented"; return }

// TallyVotes returns the number of granted and rejected Votes, and whether the
// election outcome is known.
func (p *ProgressTracker) TallyVotes() (granted int, rejected int, _ quorum.VoteResult) {
	_ = "STUB: not implemented"
	// Make sure to populate granted/rejected correctly even if the Votes slice
	// contains members no longer part of the configuration. This doesn't really
	// matter in the way the numbers are used (they're informational), but might
	// as well get it right.
	return 0, 0, *new(quorum.VoteResult)
}
