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

package confchange

import (
	"github.com/vanus-labs/vanus/pkg/raft/quorum"
	pb "github.com/vanus-labs/vanus/pkg/raft/raftpb"
	"github.com/vanus-labs/vanus/pkg/raft/tracker"
)

// Changer facilitates configuration changes. It exposes methods to handle
// simple and joint consensus while performing the proper validation that allows
// refusing invalid configuration changes before they affect the active
// configuration.
type Changer struct {
	Tracker   tracker.ProgressTracker
	LastIndex uint64
}

// EnterJoint verifies that the outgoing (=right) majority config of the joint
// config is empty and initializes it with a copy of the incoming (=left)
// majority config. That is, it transitions from
//
//	(1 2 3)&&()
//
// to
//
//	(1 2 3)&&(1 2 3).
//
// The supplied changes are then applied to the incoming majority config,
// resulting in a joint configuration that in terms of the Raft thesis[1]
// (Section 4.3) corresponds to `C_{new,old}`.
//
// [1]: https://github.com/ongardie/dissertation/blob/master/online-trim.pdf
func (c Changer) EnterJoint(autoLeave bool, ccs ...pb.ConfChangeSingle) (tracker.Config, tracker.ProgressMap, error) {
	_ = "STUB: not implemented"
	return *new(tracker.Config), *new(tracker.ProgressMap), nil
}

// We allow adding nodes to an empty config for convenience (testing and
// bootstrap), but you can't enter a joint state.

// Clear the outgoing config.

// Copy incoming to outgoing.

// LeaveJoint transitions out of a joint configuration. It is an error to call
// this method if the configuration is not joint, i.e. if the outgoing majority
// config Voters[1] is empty.
//
// The outgoing majority config of the joint configuration will be removed,
// that is, the incoming config is promoted as the sole decision maker. In the
// notation of the Raft thesis[1] (Section 4.3), this method transitions from
// `C_{new,old}` into `C_new`.
//
// At the same time, any staged learners (LearnersNext) the addition of which
// was held back by an overlapping voter in the former outgoing config will be
// inserted into Learners.
//
// [1]: https://github.com/ongardie/dissertation/blob/master/online-trim.pdf
func (c Changer) LeaveJoint() (tracker.Config, tracker.ProgressMap, error) {
	_ = "STUB: not implemented"
	return *new(tracker.Config), *new(tracker.ProgressMap), nil
}

// Simple carries out a series of configuration changes that (in aggregate)
// mutates the incoming majority config Voters[0] by at most one. This method
// will return an error if that is not the case, if the resulting quorum is
// zero, or if the configuration is in a joint state (i.e. if there is an
// outgoing configuration).
func (c Changer) Simple(ccs ...pb.ConfChangeSingle) (tracker.Config, tracker.ProgressMap, error) {
	_ = "STUB: not implemented"
	return *new(tracker.Config), *new(tracker.ProgressMap), nil
}

// apply a change to the configuration. By convention, changes to voters are
// always made to the incoming majority config Voters[0]. Voters[1] is either
// empty or preserves the outgoing majority configuration while in a joint state.
func (c Changer) apply(cfg *tracker.Config, prs tracker.ProgressMap, ccs ...pb.ConfChangeSingle) error {
	_ = "STUB: not implemented"
	return nil
}

// etcd replaces the NodeID with zero if it decides (downstream of
// raft) to not apply a change, so we have to have explicit code
// here to ignore these.

// makeVoter adds or promotes the given ID to be a voter in the incoming
// majority config.
func (c Changer) makeVoter(cfg *tracker.Config, prs tracker.ProgressMap, id uint64) {
	_ = "STUB: not implemented"
	return
}

/* isLearner */

// makeLearner makes the given ID a learner or stages it to be a learner once
// an active joint configuration is exited.
//
// The former happens when the peer is not a part of the outgoing config, in
// which case we either add a new learner or demote a voter in the incoming
// config.
//
// The latter case occurs when the configuration is joint and the peer is a
// voter in the outgoing config. In that case, we do not want to add the peer
// as a learner because then we'd have to track a peer as a voter and learner
// simultaneously. Instead, we add the learner to LearnersNext, so that it will
// be added to Learners the moment the outgoing config is removed by
// LeaveJoint().
func (c Changer) makeLearner(cfg *tracker.Config, prs tracker.ProgressMap, id uint64) {
	_ = "STUB: not implemented"
	return
}

/* isLearner */

// Remove any existing voter in the incoming config...

// ... but save the Progress.

// Use LearnersNext if we can't add the learner to Learners directly, i.e.
// if the peer is still tracked as a voter in the outgoing config. It will
// be turned into a learner in LeaveJoint().
//
// Otherwise, add a regular learner right away.

// remove this peer as a voter or learner from the incoming config.
func (c Changer) remove(cfg *tracker.Config, prs tracker.ProgressMap, id uint64) {
	_ = "STUB: not implemented"
	return
}

// If the peer is still a voter in the outgoing config, keep the Progress.

// initProgress initializes a new progress for the given node or learner.
func (c Changer) initProgress(cfg *tracker.Config, prs tracker.ProgressMap, id uint64, isLearner bool) {
	_ = "STUB: not implemented"
	return
}

// Initializing the Progress with the last index means that the follower
// can be probed (with the last index).
//
// TODO(tbg): seems awfully optimistic. Using the first index would be
// better. The general expectation here is that the follower has no log
// at all (and will thus likely need a snapshot), though the app may
// have applied a snapshot out of band before adding the replica (thus
// making the first index the better choice).

// When a node is first added, we should mark it as recently active.
// Otherwise, CheckQuorum may cause us to step down if it is invoked
// before the added node has had a chance to communicate with us.

// checkInvariants makes sure that the config and progress are compatible with
// each other. This is used to check both what the Changer is initialized with,
// as well as what it returns.
func checkInvariants(cfg tracker.Config, prs tracker.ProgressMap) error {
	_ = "STUB: not implemented"
	// NB: intentionally allow the empty config. In production we'll never see a
	// non-empty config (we prevent it from being created) but we will need to
	// be able to *create* an initial config, for example during bootstrap (or
	// during tests). Instead of having to hand-code this, we allow
	// transitioning from an empty config into any other legal and non-empty
	// config.
	return nil
}

// Any staged learner was staged because it could not be directly added due
// to a conflicting voter in the outgoing config.

// Conversely Learners and Voters doesn't intersect at all.

// We enforce that empty maps are nil instead of zero.

// checkAndCopy copies the tracker's config and progress map (deeply enough for
// the purposes of the Changer) and returns those copies. It returns an error
// if checkInvariants does.
func (c Changer) checkAndCopy() (tracker.Config, tracker.ProgressMap, error) {
	_ = "STUB: not implemented"
	return *new(tracker.Config), *new(tracker.ProgressMap), nil
}

// A shallow copy is enough because we only mutate the Learner field.

// checkAndReturn calls checkInvariants on the input and returns either the
// resulting error or the input.
func checkAndReturn(cfg tracker.Config, prs tracker.ProgressMap) (tracker.Config, tracker.ProgressMap, error) {
	_ = "STUB: not implemented"
	return *new(tracker.Config), *new(tracker.ProgressMap), nil
}

// err returns zero values and an error.
func (c Changer) err(err error) (tracker.Config, tracker.ProgressMap, error) {
	_ = "STUB: not implemented"
	return *new(tracker.Config), *new(tracker.ProgressMap), nil
}

// nilAwareAdd populates a map entry, creating the map if necessary.
func nilAwareAdd(m *map[uint64]struct{}, id uint64) { _ = "STUB: not implemented"; return }

// nilAwareDelete deletes from a map, nil'ing the map itself if it is empty after.
func nilAwareDelete(m *map[uint64]struct{}, id uint64) { _ = "STUB: not implemented"; return }

// symdiff returns the count of the symmetric difference between the sets of
// uint64s, i.e. len( (l - r) \union (r - l)).
func symdiff(l, r map[uint64]struct{}) int { _ = "STUB: not implemented"; return 0 }

// count elems in l but not in r
// count elems in r but not in l

func joint(cfg tracker.Config) bool { _ = "STUB: not implemented"; return false }

func incoming(voters quorum.JointConfig) quorum.MajorityConfig {
	_ = "STUB: not implemented"
	return *new(quorum.MajorityConfig)
}
func outgoing(voters quorum.JointConfig) quorum.MajorityConfig {
	_ = "STUB: not implemented"
	return *new(quorum.MajorityConfig)
}
func outgoingPtr(voters *quorum.JointConfig) *quorum.MajorityConfig {
	_ = "STUB: not implemented"

	// Describe prints the type and NodeID of the configuration changes as a
	// space-delimited string.
	return nil
}

func Describe(ccs ...pb.ConfChangeSingle) string { _ = "STUB: not implemented"; return "" }
