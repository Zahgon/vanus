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
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/vanus-labs/vanus/pkg/raft/quorum"
	pb "github.com/vanus-labs/vanus/pkg/raft/raftpb"
	"github.com/vanus-labs/vanus/pkg/raft/tracker"
)

// None is a placeholder node ID used when there is no leader.
const (
	None    uint64 = 0
	noLimit        = math.MaxUint64
)

// Possible values for StateType.
const (
	StateFollower StateType = iota
	StateCandidate
	StateLeader
	StatePreCandidate
	numStates
)

type ReadOnlyOption int

const (
	// ReadOnlySafe guarantees the linearizability of the read only request by
	// communicating with the quorum. It is the default and suggested option.
	ReadOnlySafe ReadOnlyOption = iota
	// ReadOnlyLeaseBased ensures linearizability of the read only request by
	// relying on the leader lease. It can be affected by clock drift.
	// If the clock drift is unbounded, leader might keep the lease longer than it
	// should (clock can move backward/pause without any bound). ReadIndex is not safe
	// in that case.
	ReadOnlyLeaseBased
)

// Possible values for CampaignType
const (
	// campaignPreElection represents the first phase of a normal election when
	// Config.PreVote is true.
	campaignPreElection CampaignType = "CampaignPreElection"
	// campaignElection represents a normal (time-based) election (the second phase
	// of the election when Config.PreVote is true).
	campaignElection CampaignType = "CampaignElection"
	// campaignTransfer represents the type of leader transfer
	campaignTransfer CampaignType = "CampaignTransfer"
)

// ErrProposalDropped is returned when the proposal is ignored by some cases,
// so that the proposer can be notified and fail fast.
var ErrProposalDropped = errors.New("raft proposal dropped")

// lockedRand is a small wrapper around rand.Rand to provide
// synchronization among multiple raft groups. Only the methods needed
// by the code are exposed (e.g. Intn).
type lockedRand struct {
	mu   sync.Mutex
	rand *rand.Rand
}

func (r *lockedRand) Intn(n int) int { _ = "STUB: not implemented"; return 0 }

var globalRand = &lockedRand{
	rand: rand.New(rand.NewSource(time.Now().UnixNano())),
}

// CampaignType represents the type of campaigning
// the reason we use the type of string instead of uint64
// is because it's simpler to compare and fill in raft entries
type CampaignType string

// StateType represents the role of a node in a cluster.
type StateType uint64

var stmap = [...]string{
	"StateFollower",
	"StateCandidate",
	"StateLeader",
	"StatePreCandidate",
}

func (st StateType) String() string { _ = "STUB: not implemented"; return "" }

// Config contains the parameters to start a raft.
type Config struct {
	// ID is the identity of the local raft. ID cannot be 0.
	ID uint64

	// ElectionTick is the number of Node.Tick invocations that must pass between
	// elections. That is, if a follower does not receive any message from the
	// leader of current term before ElectionTick has elapsed, it will become
	// candidate and start an election. ElectionTick must be greater than
	// HeartbeatTick. We suggest ElectionTick = 10 * HeartbeatTick to avoid
	// unnecessary leader switching.
	ElectionTick int
	// HeartbeatTick is the number of Node.Tick invocations that must pass between
	// heartbeats. That is, a leader sends heartbeat messages to maintain its
	// leadership every HeartbeatTick ticks.
	HeartbeatTick int

	// Storage is the storage for raft. raft generates entries and states to be
	// stored in storage. raft reads the persisted entries and states out of
	// Storage when it needs. raft reads out the previous state and configuration
	// out of storage when restarting.
	Storage Storage
	Keeper  Keeper
	// Applied is the last applied index. It should only be set when restarting
	// raft. raft will not return entries to the application smaller or equal to
	// Applied. If Applied is unset when restarting, raft might return previous
	// applied entries. This is a very application dependent configuration.
	Applied uint64

	Compacted uint64

	// MaxSizePerMsg limits the max byte size of each append message. Smaller
	// value lowers the raft recovery cost(initial probing and message lost
	// during normal operation). On the other side, it might affect the
	// throughput during normal replication. Note: math.MaxUint64 for unlimited,
	// 0 for at most one entry per message.
	MaxSizePerMsg uint64
	// MaxCommittedSizePerReady limits the size of the committed entries which
	// can be applied.
	MaxCommittedSizePerReady uint64
	// MaxUncommittedEntriesSize limits the aggregate byte size of the
	// uncommitted entries that may be appended to a leader's log. Once this
	// limit is exceeded, proposals will begin to return ErrProposalDropped
	// errors. Note: 0 for no limit.
	MaxUncommittedEntriesSize uint64
	// MaxInflightMsgs limits the max number of in-flight append messages during
	// optimistic replication phase. The application transportation layer usually
	// has its own sending buffer over TCP/UDP. Setting MaxInflightMsgs to avoid
	// overflowing that sending buffer. TODO (xiangli): feedback to application to
	// limit the proposal rate?
	MaxInflightMsgs int

	// CheckQuorum specifies if the leader should check quorum activity. Leader
	// steps down when quorum is not active for an electionTimeout.
	CheckQuorum bool

	// PreVote enables the Pre-Vote algorithm described in raft thesis section
	// 9.6. This prevents disruption when a node that has been partitioned away
	// rejoins the cluster.
	PreVote bool

	// ReadOnlyOption specifies how the read only request is processed.
	//
	// ReadOnlySafe guarantees the linearizability of the read only request by
	// communicating with the quorum. It is the default and suggested option.
	//
	// ReadOnlyLeaseBased ensures linearizability of the read only request by
	// relying on the leader lease. It can be affected by clock drift.
	// If the clock drift is unbounded, leader might keep the lease longer than it
	// should (clock can move backward/pause without any bound). ReadIndex is not safe
	// in that case.
	// CheckQuorum MUST be enabled if ReadOnlyOption is ReadOnlyLeaseBased.
	ReadOnlyOption ReadOnlyOption

	// Logger is the logger used for raft log. For multinode which can host
	// multiple raft group, each raft group can have its own logger
	Logger Logger

	// DisableProposalForwarding set to true means that followers will drop
	// proposals, rather than forwarding them to the leader. One use case for
	// this feature would be in a situation where the Raft leader is used to
	// compute the data of a proposal, for example, adding a timestamp from a
	// hybrid logical clock to data in a monotonically increasing way. Forwarding
	// should be disabled to prevent a follower with an inaccurate hybrid
	// logical clock from assigning the timestamp and then forwarding the data
	// to the leader.
	DisableProposalForwarding bool
}

func (c *Config) validate() error { _ = "STUB: not implemented"; return nil }

// default MaxCommittedSizePerReady to MaxSizePerMsg because they were
// previously the same parameter.

type raft struct {
	id uint64

	Term uint64
	Vote uint64

	termStart uint64
	receiving bool
	voted     bool

	readStates []ReadState

	// the log
	raftLog *raftLog

	maxMsgSize         uint64
	maxUncommittedSize uint64
	// TODO(tbg): rename to trk.
	prs tracker.ProgressTracker

	state StateType

	// isLearner is true if the local raft node is a learner.
	isLearner bool

	// the leader id
	lead uint64
	// leadTransferee is id of the leader transfer target when its value is not zero.
	// Follow the procedure defined in raft thesis 3.10.
	leadTransferee uint64
	// Only one conf change may be pending (in the log, but not yet
	// applied) at a time. This is enforced via pendingConfIndex, which
	// is set to a value >= the log index of the latest pending
	// configuration change (if any). Config changes are only allowed to
	// be proposed if the leader's applied index is greater than this
	// value.
	pendingConfIndex uint64
	// an estimate of the size of the uncommitted tail of the Raft log. Used to
	// prevent unbounded log growth. Only maintained by the leader. Reset on
	// term changes.
	uncommittedSize uint64

	readOnly *readOnly

	// number of ticks since it reached last electionTimeout when it is leader
	// or candidate.
	// number of ticks since it reached last electionTimeout or received a
	// valid message from current leader when it is a follower.
	electionElapsed int

	// number of ticks since it reached last heartbeatTimeout.
	// only leader keeps heartbeatElapsed.
	heartbeatElapsed int

	checkQuorum bool
	preVote     bool

	heartbeatTimeout int
	electionTimeout  int
	// randomizedElectionTimeout is a random number between
	// [electiontimeout, 2 * electiontimeout - 1]. It gets reset
	// when raft changes its state to follower or candidate.
	randomizedElectionTimeout int
	disableProposalForwarding bool

	tick    func()
	step    stepFunc
	propose proposeFunc

	logger Logger

	// pendingReadIndexMessages is used to store messages of type MsgReadIndex
	// that can't be answered as new leader didn't committed any log in
	// current term. Those will be handled as fast as first log is committed in
	// current term.
	pendingReadIndexMessages []pb.Message
}

func newRaft(c *Config) *raft { _ = "STUB: not implemented"; return nil }

// TODO(bdarnell)

// FIXME(james.yin): fix term

func (r *raft) hasLeader() bool { _ = "STUB: not implemented"; return false }

func (r *raft) softState() SoftState { _ = "STUB: not implemented"; return *new(SoftState) }

func (r *raft) hardState() pb.HardState { _ = "STUB: not implemented"; return *new(pb.HardState) }

// send schedules persisting state to a stable storage and AFTER that
// sending the message (as part of next Ready message processing).
func (r *raft) send(m pb.Message) { _ = "STUB: not implemented"; return }

// All {pre-,}campaign messages need to have the term set when
// sending.
// - MsgVote: m.Term is the term the node is campaigning for,
//   non-zero as we increment the term when campaigning.
// - MsgVoteResp: m.Term is the new r.Term if the MsgVote was
//   granted, non-zero for the same reason MsgVote is
// - MsgPreVote: m.Term is the term the node will campaign,
//   non-zero as we use m.Term to indicate the next term we'll be
//   campaigning for
// - MsgPreVoteResp: m.Term is the term received in the original
//   MsgPreVote if the pre-vote was granted, non-zero for the
//   same reasons MsgPreVote is

// do not attach term to MsgProp, MsgReadIndex
// proposals are a way to forward to the leader and
// should be treated as local message.
// MsgReadIndex is also forwarded to leader.

// sendAppend sends an append RPC with new entries (if any) and the
// current commit index to the given peer.
func (r *raft) sendAppend(to uint64) { _ = "STUB: not implemented"; return }

// maybeSendAppend sends an append RPC with new entries to the given peer,
// if necessary. Returns true if a message was sent. The sendIfEmpty
// argument controls whether messages with no entries will be sent
// ("empty" messages are useful to convey updated Commit indexes, but
// are undesirable when we're sending multiple messages in a batch).
func (r *raft) maybeSendAppend(to uint64, sendIfEmpty bool) bool {
	_ = "STUB: not implemented"
	return false
}

// send snapshot if we failed to get term or entries

// TODO(bdarnell)

// optimistically increase the next when in StateReplicate

// sendHeartbeat sends a heartbeat RPC to the given peer.
func (r *raft) sendHeartbeat(to uint64, ctx []byte) {
	_ = "STUB: not implemented"
	// Attach the commit as min(to.matched, r.committed).
	// When the leader sends out heartbeat message,
	// the receiver(follower) might not be matched with the leader
	// or it might not have all the committed entries.
	// The leader MUST NOT forward the follower's commit to
	// an unmatched index.
	return
}

// bcastAppend sends RPC, with entries to all peers that are not up-to-date
// according to the progress recorded in r.prs.
func (r *raft) bcastAppend() { _ = "STUB: not implemented"; return }

// bcastHeartbeat sends RPC, without entries to all the peers.
func (r *raft) bcastHeartbeat() { _ = "STUB: not implemented"; return }

func (r *raft) bcastHeartbeatWithCtx(ctx []byte) { _ = "STUB: not implemented"; return }

func (r *raft) advance(rd Ready) { _ = "STUB: not implemented"; return }

// If entries were applied (or a snapshot), update our cursor for
// the next Ready. Note that if the current HardState contains a
// new Commit index, this does not mean that we're also applying
// all of the new entries due to commit pagination by size.

// If the current (and most recent, at least for this leader's term)
// configuration should be auto-left, initiate that now. We use a
// nil Data which unmarshals into an empty ConfChangeV2 and has the
// benefit that appendEntry can never refuse it based on its size
// (which registers as zero).

// There's no way in which this proposal should be able to be rejected.

// maybeCommit attempts to advance the commit index. Returns true if
// the commit index changed (in which case the caller should call
// r.bcastAppend).
func (r *raft) maybeCommit() bool { _ = "STUB: not implemented"; return false }

func (r *raft) maybeCompact() { _ = "STUB: not implemented"; return }

func (r *raft) reset(term uint64) { _ = "STUB: not implemented"; return }

func (r *raft) appendEntry(es ...pb.Entry) (accepted bool) { _ = "STUB: not implemented"; return false }

// Track the size of this uncommitted proposal.

// Drop the proposal.

// tickElection is run by followers and candidates after r.electionTimeout.
func (r *raft) tickElection() { _ = "STUB: not implemented"; return }

// tickHeartbeat is run by leaders to send a MsgBeat after r.heartbeatTimeout.
func (r *raft) tickHeartbeat() { _ = "STUB: not implemented"; return }

// If current leader cannot transfer leadership in electionTimeout, it becomes leader again.

func (r *raft) becomeFollower(term uint64, lead uint64) { _ = "STUB: not implemented"; return }

func (r *raft) becomeFollowerExt(term uint64, lead uint64) { _ = "STUB: not implemented"; return }

func (r *raft) becomeCandidate() {
	_ = "STUB: not implemented"
	// TODO(xiangli) remove the panic when the raft implementation is stable
	return
}

func (r *raft) becomeCandidateExt() { _ = "STUB: not implemented"; return }

func (r *raft) becomePreCandidate() {
	_ = "STUB: not implemented"
	// TODO(xiangli) remove the panic when the raft implementation is stable
	return
}

// Becoming a pre-candidate changes our step functions and state,
// but doesn't change anything else. In particular it does not increase
// r.Term or change r.Vote.

func (r *raft) becomeLeader() {
	_ = "STUB: not implemented"
	// TODO(xiangli) remove the panic when the raft implementation is stable
	return
}

// Followers enter replicate mode when they've been successfully probed
// (perhaps after having received a snapshot as a result). The leader is
// trivially in this state. Note that r.reset() has initialized this
// progress with the last index already.

// Conservatively set the pendingConfIndex to the last index in the
// log. There may or may not be a pending config change, but it's
// safe to delay any future proposals until we commit all our
// pending log entries, and scanning the entire tail of the log
// could be expensive.

// This won't happen because we just called reset() above.

// As a special case, don't count the initial empty entry towards the
// uncommitted log quota. This is because we want to preserve the
// behavior of allowing one entry larger than quota if the current
// usage is zero.

func (r *raft) hup(t CampaignType) { _ = "STUB: not implemented"; return }

// campaign transitions the raft instance to candidate state. This must only be
// called after verifying that this is a legitimate transition.
func (r *raft) campaign(t CampaignType) { _ = "STUB: not implemented"; return }

// This path should not be hit (callers are supposed to check), but
// better safe than sorry.

// We won the election after voting for ourselves (which must mean that
// this is a single-node cluster). Advance to the next state.

// PreVote RPCs are sent for the next term before we've incremented r.Term.

func (r *raft) poll(id uint64, t pb.MessageType, v bool) (granted int, rejected int, result quorum.VoteResult) {
	_ = "STUB: not implemented"
	return 0, 0, *new(quorum.VoteResult)
}

func (r *raft) Propose(pds ...ProposeData) { _ = "STUB: not implemented"; return }

type proposeFunc func(r *raft, m pb.Message) error

func proposeLeader(r *raft, m pb.Message) error { _ = "STUB: not implemented"; return nil }

// If we are not currently a member of the range (i.e. this node
// was removed from the configuration while serving as leader),
// drop any new proposals.

// e.Type == pb.EntryConfChangeV2

func proposeCandidate(r *raft, m pb.Message) error { _ = "STUB: not implemented"; return nil }

func proposeFollower(r *raft, m pb.Message) error { _ = "STUB: not implemented"; return nil }

func (r *raft) Step(m pb.Message) error {
	_ = "STUB: not implemented"
	// Handle the message term, which may result in our stepping down to a follower.
	return nil
}

// local message

// We have received messages from a leader at a lower term. It is possible
// that these messages were simply delayed in the network, but this could
// also mean that this node has advanced its term number during a network
// partition, and it is now unable to either win an election or to rejoin
// the majority on the old term. If checkQuorum is false, this will be
// handled by incrementing term numbers in response to MsgVote with a
// higher term, but if checkQuorum is true we may not advance the term on
// MsgVote and must generate other messages to advance the term. The net
// result of these two features is to minimize the disruption caused by
// nodes that have been removed from the cluster's configuration: a
// removed node will send MsgVotes (or MsgPreVotes) which will be ignored,
// but it will not receive MsgApp or MsgHeartbeat, so it will not create
// disruptive term increases, by notifying leader of this node's activeness.
// The above comments also true for Pre-Vote
//
// When follower gets isolated, it soon starts an election ending
// up with a higher term than leader, although it won't receive enough
// votes to win the election. When it regains connectivity, this response
// with "pb.MsgAppResp" of higher term would force leader to step down.
// However, this disruption is inevitable to free this stuck node with
// fresh election. This can be prevented with Pre-Vote phase.

// Before Pre-Vote enable, there may have candidate with higher term,
// but less log. After update to Pre-Vote, the cluster may deadlock if
// we drop messages with a lower term.

// ignore other cases

// If a server receives a RequestVote request within the minimum election timeout
// of hearing from a current leader, it does not update its term or grant its vote

// We send pre-vote requests with a term in our future. If the
// pre-vote is granted, we will increment our term when we get a
// quorum. If it is not, the term comes from the node that
// rejected our vote so we should become a follower at the new
// term.

type stepFunc func(r *raft, m pb.Message) error

func stepLeader(r *raft, m pb.Message) error {
	_ = "STUB: not implemented"
	// These message types do not require any progress for m.From.
	return nil
}

// The leader should always see itself as active. As a precaution, handle
// the case in which the leader isn't in the configuration any more (for
// example if it just removed itself).
//
// TODO(tbg): I added a TODO in removeNode, it doesn't seem that the
// leader steps down when removing itself. I might be missing something.

// Mark everyone (but ourselves) as inactive in preparation for the next
// CheckQuorum.

// TODO(james.yin): reduce uncommitted size

// Config change

// If the current (and most recent, at least for this leader's term)
// configuration should be auto-left, initiate that now. We use a
// nil Data which unmarshals into an empty ConfChangeV2 and has the
// benefit that appendEntry can never refuse it based on its size
// (which registers as zero).

// There's no way in which this proposal should be able to be rejected.

// only one voting member (the leader) in the cluster

// Postpone read only request when this leader has not committed
// any log entry at its term.

// All other message types require a progress for m.From (pr).

// RejectHint is the suggested next base entry for appending (i.e.
// we try to append entry RejectHint+1 next), and LogTerm is the
// term that the follower has at index RejectHint. Older versions
// of this library did not populate LogTerm for rejections and it
// is zero for followers with an empty log.
//
// Under normal circumstances, the leader's log is longer than the
// follower's and the follower's log is a prefix of the leader's
// (i.e. there is no divergent uncommitted suffix of the log on the
// follower). In that case, the first probe reveals where the
// follower's log ends (RejectHint=follower's last index) and the
// subsequent probe succeeds.
//
// However, when networks are partitioned or systems overloaded,
// large divergent log tails can occur. The naive attempt, probing
// entry by entry in decreasing order, will be the product of the
// length of the diverging tails and the network round-trip latency,
// which can easily result in hours of time spent probing and can
// even cause outright outages. The probes are thus optimized as
// described below.

// If the follower has an uncommitted log tail, we would end up
// probing one by one until we hit the common prefix.
//
// For example, if the leader has:
//
//   idx        1 2 3 4 5 6 7 8 9
//              -----------------
//   term (L)   1 3 3 3 5 5 5 5 5
//   term (F)   1 1 1 1 2 2
//
// Then, after sending an append anchored at (idx=9,term=5) we
// would receive a RejectHint of 6 and LogTerm of 2. Without the
// code below, we would try an append at index 6, which would
// fail again.
//
// However, looking only at what the leader knows about its own
// log and the rejection hint, it is clear that a probe at index
// 6, 5, 4, 3, and 2 must fail as well:
//
// For all of these indexes, the leader's log term is larger than
// the rejection's log term. If a probe at one of these indexes
// succeeded, its log term at that index would match the leader's,
// i.e. 3 or 5 in this example. But the follower already told the
// leader that it is still at term 2 at index 6, and since the
// log term only ever goes up (within a log), this is a contradiction.
//
// At index 1, however, the leader can draw no such conclusion,
// as its term 1 is not larger than the term 2 from the
// follower's rejection. We thus probe at 1, which will succeed
// in this example. In general, with this approach we probe at
// most once per term found in the leader's log.
//
// There is a similar mechanism on the follower (implemented in
// handleAppendEntries via a call to findConflictByTerm) that is
// useful if the follower has a large divergent uncommitted log
// tail[1], as in this example:
//
//   idx        1 2 3 4 5 6 7 8 9
//              -----------------
//   term (L)   1 3 3 3 3 3 3 3 7
//   term (F)   1 3 3 4 4 5 5 5 6
//
// Naively, the leader would probe at idx=9, receive a rejection
// revealing the log term of 6 at the follower. Since the leader's
// term at the previous index is already smaller than 6, the leader-
// side optimization discussed above is ineffective. The leader thus
// probes at index 8 and, naively, receives a rejection for the same
// index and log term 5. Again, the leader optimization does not improve
// over linear probing as term 5 is above the leader's term 3 for that
// and many preceding indexes; the leader would have to probe linearly
// until it would finally hit index 3, where the probe would succeed.
//
// Instead, we apply a similar optimization on the follower. When the
// follower receives the probe at index 8 (log term 3), it concludes
// that all of the leader's log preceding that index has log terms of
// 3 or below. The largest index in the follower's log with a log term
// of 3 or below is index 3. The follower will thus return a rejection
// for index=3, log term=3 instead. The leader's next probe will then
// succeed at that index.
//
// [1]: more precisely, if the log terms in the large uncommitted
// tail on the follower are larger than the leader's. At first,
// it may seem unintuitive that a follower could even have such
// a large tail, but it can happen:
//
// 1. Leader appends (but does not commit) entries 2 and 3, crashes.
//   idx        1 2 3 4 5 6 7 8 9
//              -----------------
//   term (L)   1 2 2     [crashes]
//   term (F)   1
//   term (F)   1
//
// 2. a follower becomes leader and appends entries at term 3.
//              -----------------
//   term (x)   1 2 2     [down]
//   term (F)   1 3 3 3 3
//   term (F)   1
//
// 3. term 3 leader goes down, term 2 leader returns as term 4
//    leader. It commits the log & entries at term 4.
//
//              -----------------
//   term (L)   1 2 2 2
//   term (x)   1 3 3 3 3 [down]
//   term (F)   1
//              -----------------
//   term (L)   1 2 2 2 4 4 4
//   term (F)   1 3 3 3 3 [gets probed]
//   term (F)   1 2 2 2 4 4 4
//
// 4. the leader will now probe the returning follower at index
//    7, the rejection points it at the end of the follower's log
//    which is at a higher log term than the actually committed
//    log.

// TODO(tbg): we should also enter this branch if a snapshot is
// received that is below pr.PendingSnapshot but which makes it
// possible to use the log again.

// Transition back to replicating state via probing state
// (which takes the snapshot into account). If we didn't
// move to replicating state, that would only happen with
// the next round of appends (but there may not be a next
// round for a while, exposing an inconsistent RaftStatus).

// committed index has progressed for the term, so it is safe
// to respond to pending read index requests

// If we were paused before, this node may be missing the
// latest commit index, so send it.

// We've updated flow control information above, which may
// allow us to send multiple (size-limited) in-flight messages
// at once (such as when transitioning from probe to
// replicate, or when freeTo() covers multiple messages). If
// we have more entries to send, send as many messages as we
// can (without sending empty messages for the commit index)

// Transfer leadership is in progress.

// free one slot for the full inflights window to allow progress.

// TODO(tbg): this code is very similar to the snapshot handling in
// MsgAppResp above. In fact, the code there is more correct than the
// code here and should likely be updated to match (or even better, the
// logic pulled into a newly created Progress state machine handler).

// NB: the order here matters or we'll be probing erroneously from
// the snapshot index, but the snapshot never applied.

// If snapshot finish, wait for the MsgAppResp from the remote node before sending
// out the next MsgApp.
// If snapshot failure, wait for a heartbeat interval before next try

// During optimistic replication, if the remote becomes unreachable,
// there is huge probability that a MsgApp is lost.

// Transfer leadership to third party.

// Transfer leadership should be finished in one electionTimeout, so reset r.electionElapsed.

// stepCandidate is shared by StateCandidate and StatePreCandidate; the difference is
// whether they respond to MsgVoteResp or MsgPreVoteResp.
func stepCandidate(r *raft, m pb.Message) error {
	_ = "STUB: not implemented"
	// Only handle vote responses corresponding to our candidacy (while in
	// StateCandidate, we may get stale MsgPreVoteResp messages in this term from
	// our pre-candidate state).
	return nil
}

// always m.Term == r.Term

// TODO(james.yin):

// always m.Term == r.Term

// always m.Term == r.Term

// pb.MsgPreVoteResp contains future term of pre-candidate
// m.Term > r.Term; reuse r.Term

func stepFollower(r *raft, m pb.Message) error { _ = "STUB: not implemented"; return nil }

// TODO(james.yin):

// Leadership transfers never use pre-vote even if r.preVote is true; we
// know we are not recovering from a partition so there is no need for the
// extra round trip.

func (r *raft) handlePreVote(from, term, index, logTerm uint64) { _ = "STUB: not implemented"; return }

// TODO(james.yin): not received heartbeats from a valid leader for at least a baseline election timeout
// We can vote for a future term...

// ...we can vote if this is a repeat of a vote we've already cast.

// ...we haven't voted and we don't think there's a leader yet in this term...

// ...and we believe the candidate is up to date.

func (r *raft) handleVote(from, term, index, logTerm uint64) { _ = "STUB: not implemented"; return }

// We can vote if this is a repeat of a vote we've already cast.

// m.Term > r.Term

// We haven't voted and we don't think there's a leader yet in this term...

// ...and we believe the candidate is up to date.

// Note: it turns out that that learners must be allowed to cast votes.
// This seems counter- intuitive but is necessary in the situation in which
// a learner has been promoted (i.e. is now a voter) but has not learned
// about this yet.
// For example, consider a group in which id=1 is a learner and id=2 and
// id=3 are voters. A configuration change promoting 1 can be committed on
// the quorum `{2,3}` without the config change being appended to the
// learner's log. If the leader (say 2) fails, there are de facto two
// voters remaining. Only 3 can win an election (due to its log containing
// all committed entries), but to do so it will need 1 to vote. But 1
// considers itself a learner and will continue to do so until 3 has
// stepped up as leader, replicates the conf change to 1, and 1 applies it.
// Ultimately, by receiving a request to vote, the learner realizes that
// the candidate believes it to be a voter, and that it should act
// accordingly. The candidate's config may be stale, too; but in that case
// it won't win the election, at least in the absence of the bug discussed
// in:
// https://github.com/etcd-io/etcd/issues/7625#issuecomment-488798263.

// When responding to Msg{Pre,}Vote messages we include the term
// from the message, not the local term. To see why, consider the
// case where a single node was previously partitioned away and
// it's local term is now out of date. If we include the local term
// (recall that for pre-votes we don't update the local term), the
// (pre-)campaigning node on the other end will proceed to ignore
// the message (it ignores all out of date messages).
// The term in the original message and current local term are the
// same in the case of regular votes, but different for pre-votes.
// Only record real votes.

// TODO(james.yin): wait hard state is persisted?

func (r *raft) handleStatePersisted(term, vote uint64) { _ = "STUB: not implemented"; return }

// Vote to ourself in current term.

// Vote to other peer in current term.

func (r *raft) handleAppendEntries(m pb.Message) { _ = "STUB: not implemented"; return }

// Return a hint to the leader about the maximum index and term that the
// two logs could be divergent at. Do this by searching through the
// follower's log for the maximum (index, term) pair with a term <= the
// MsgApp's LogTerm and an index <= the MsgApp's Index. This can help
// skip all indexes in the follower's uncommitted tail with terms
// greater than the MsgApp's LogTerm.
//
// See the other caller for findConflictByTerm (in stepLeader) for a much
// more detailed explanation of this mechanism.

func (r *raft) handleHeartbeat(m pb.Message) { _ = "STUB: not implemented"; return }

func (r *raft) handleSnapshot(m pb.Message) { _ = "STUB: not implemented"; return }

// restore recovers the state machine from a snapshot. It restores the log and the
// configuration of state machine. If this method returns false, the snapshot was
// ignored, either because it was obsolete or because of an error.
func (r *raft) restore(s pb.Snapshot) bool { _ = "STUB: not implemented"; return false }

// This is defense-in-depth: if the leader somehow ended up applying a
// snapshot, it could move into a new term without moving into a
// follower state. This should never fire, but if it did, we'd have
// prevented damage by returning early, so log only a loud warning.
//
// At the time of writing, the instance is guaranteed to be in follower
// state when this method is called.

// More defense-in-depth: throw away snapshot if recipient is not in the
// config. This shouldn't ever happen (at the time of writing) but lots of
// code here and there assumes that r.id is in the progress tracker.

// `LearnersNext` doesn't need to be checked. According to the rules, if a peer in
// `LearnersNext`, it has to be in `VotersOutgoing`.

// Now go ahead and actually restore.

// Reset the configuration and add the (potentially updated) peers in anew.

// This should never happen. Either there's a bug in our config change
// handling or the client corrupted the conf change.

// TODO(tbg): this is untested and likely unneeded

// promotable indicates whether state machine can be promoted to leader,
// which is true when its own id is in progress list.
func (r *raft) promotable() bool { _ = "STUB: not implemented"; return false }

func (r *raft) applyConfChange(cc pb.ConfChangeV2) pb.ConfState {
	_ = "STUB: not implemented"
	return *new(pb.ConfState)
}

// TODO(tbg): return the error to the caller.

// switchToConfig reconfigures this node to use the provided configuration. It
// updates the in-memory state and, when necessary, carries out additional
// actions such as reacting to the removal of nodes or changed quorum
// requirements.
//
// The inputs usually result from restoring a ConfState or applying a ConfChange.
func (r *raft) switchToConfig(cfg tracker.Config, prs tracker.ProgressMap) pb.ConfState {
	_ = "STUB: not implemented"
	return *new(pb.ConfState)
}

// Update whether the node itself is a learner, resetting to false when the
// node is removed.

// This node is leader and was removed or demoted. We prevent demotions
// at the time writing but hypothetically we handle them the same way as
// removing the leader: stepping down into the next Term.
//
// TODO(tbg): step down (for sanity) and ask follower with largest Match
// to TimeoutNow (to avoid interruption). This might still drop some
// proposals but it's better than nothing.
//
// TODO(tbg): test this branch. It is untested at the time of writing.

// The remaining steps only make sense if this node is the leader and there
// are other nodes.

// If the configuration change means that more entries are committed now,
// broadcast/append to everyone in the updated config.

// Otherwise, still probe the newly added replicas; there's no reason to
// let them wait out a heartbeat interval (or the next incoming
// proposal).

/* sendIfEmpty */

// If the the leadTransferee was removed or demoted, abort the leadership transfer.

func (r *raft) loadState(state pb.HardState) { _ = "STUB: not implemented"; return }

// pastElectionTimeout returns true iff r.electionElapsed is greater
// than or equal to the randomized election timeout in
// [electiontimeout, 2 * electiontimeout - 1].
func (r *raft) pastElectionTimeout() bool { _ = "STUB: not implemented"; return false }

func (r *raft) resetRandomizedElectionTimeout() { _ = "STUB: not implemented"; return }

func (r *raft) sendTimeoutNow(to uint64) { _ = "STUB: not implemented"; return }

func (r *raft) abortLeaderTransfer() { _ = "STUB: not implemented"; return }

// committedEntryInCurrentTerm return true if the peer has committed an entry in its term.
func (r *raft) committedEntryInCurrentTerm() bool { _ = "STUB: not implemented"; return false }

// responseToReadIndexReq constructs a response for `req`. If `req` comes from the peer
// itself, a blank value will be returned.
func (r *raft) responseToReadIndexReq(req pb.Message, readIndex uint64) pb.Message {
	_ = "STUB: not implemented"
	return *new(pb.Message)
}

// increaseUncommittedSize computes the size of the proposed entries and
// determines whether they would push leader over its maxUncommittedSize limit.
// If the new entries would exceed the limit, the method returns false. If not,
// the increase in uncommitted entry size is recorded and the method returns
// true.
//
// Empty payloads are never refused. This is used both for appending an empty
// entry at a new leader's term, as well as leaving a joint configuration.
func (r *raft) increaseUncommittedSize(ents []pb.Entry) bool {
	_ = "STUB: not implemented"
	return false
}

// If the uncommitted tail of the Raft log is empty, allow any size
// proposal. Otherwise, limit the size of the uncommitted tail of the
// log and drop any proposal that would push the size over the limit.
// Note the added requirement s>0 which is used to make sure that
// appending single empty entries to the log always succeeds, used both
// for replicating a new leader's initial empty entry, and for
// auto-leaving joint configurations.

// reduceUncommittedSize accounts for the newly committed entries by decreasing
// the uncommitted entry size limit.
func (r *raft) reduceUncommittedSize(ents []pb.Entry) { _ = "STUB: not implemented"; return }

// Fast-path for followers, who do not track or enforce the limit.

// uncommittedSize may underestimate the size of the uncommitted Raft
// log tail but will never overestimate it. Saturate at 0 instead of
// allowing overflow.

func numOfPendingConf(ents []pb.Entry) int { _ = "STUB: not implemented"; return 0 }

func releasePendingReadIndexMessages(r *raft) { _ = "STUB: not implemented"; return }

func sendMsgReadIndexResponse(r *raft, m pb.Message) {
	_ = "STUB: not implemented"
	// thinking: use an internally defined context instead of the user given context.
	// We can express this in terms of the term and index instead of a user-supplied value.
	// This would allow multiple reads to piggyback on the same message.
	return
}

// If more than the local vote is needed, go through a full broadcast.

// The local node automatically acks the request.
