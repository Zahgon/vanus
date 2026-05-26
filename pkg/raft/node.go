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
	// standard libraries.
	"context"
	"errors"

	// third-party libraries.

	// this project.
	pb "github.com/vanus-labs/vanus/pkg/raft/raftpb"
)

type SnapshotStatus int

const (
	SnapshotFinish  SnapshotStatus = 1
	SnapshotFailure SnapshotStatus = 2
)

var (
	emptyState = pb.HardState{}

	// ErrStopped is returned by methods on Nodes that have been stopped.
	ErrStopped = errors.New("raft: stopped")
)

// SoftState provides state that is useful for logging and debugging.
// The state is volatile and does not need to be persisted to the WAL.
type SoftState struct {
	Lead      uint64 // must use atomic operations to access; keep 64-bit aligned.
	RaftState StateType
}

func (a *SoftState) equal(b *SoftState) bool { _ = "STUB: not implemented"; return false }

// Ready encapsulates the entries and messages that are ready to read,
// be saved to stable storage, committed or sent to other peers.
// All fields in Ready are read-only.
type Ready struct {
	// The current volatile state of a Node.
	// SoftState will be nil if there is no update.
	// It is not required to consume or store SoftState.
	*SoftState

	// The current state of a Node to be saved to stable storage BEFORE
	// Messages are sent.
	// HardState will be equal to empty state if there is no update.
	pb.HardState

	// ReadStates can be used for node to serve linearizable read requests locally
	// when its applied index is greater than the index in ReadState.
	// Note that the readState will be returned when raft receives msgReadIndex.
	// The returned is only valid for the request that requested to read.
	ReadStates []ReadState

	// Entries specifies entries to be saved to stable storage BEFORE
	// Messages are sent.
	Entries []pb.Entry

	// Snapshot specifies the snapshot to be saved to stable storage.
	Snapshot pb.Snapshot

	// CommittedEntries specifies entries to be committed to a
	// store/state-machine. These have previously been committed to stable
	// store.
	CommittedEntries []pb.Entry

	// Messages specifies outbound messages to be sent AFTER Entries are
	// committed to stable storage.
	// If it contains a MsgSnap message, the application MUST report back to raft
	// when the snapshot has been received or has failed by calling ReportSnapshot.
	Messages []pb.Message

	// MustSync indicates whether the HardState and Entries must be synchronously
	// written to disk or if an asynchronous write is permissible.
	MustSync bool

	Compact uint64
}

func isHardStateEqual(a, b pb.HardState) bool { _ = "STUB: not implemented"; return false }

// IsEmptyHardState returns true if the given HardState is empty.
func IsEmptyHardState(st pb.HardState) bool { _ = "STUB: not implemented"; return false }

// IsEmptySnap returns true if the given Snapshot is empty.
func IsEmptySnap(sp pb.Snapshot) bool { _ = "STUB: not implemented"; return false }

func (rd Ready) containsUpdates() bool { _ = "STUB: not implemented"; return false }

// appliedCursor extracts from the Ready the highest index the client has
// applied (once the Ready is confirmed via Advance). If no information is
// contained in the Ready, returns zero.
func (rd Ready) appliedCursor() uint64 { _ = "STUB: not implemented"; return 0 }

// Node represents a node in a raft cluster.
type Node interface {
	// Tick increments the internal logical clock for the Node by a single tick. Election
	// timeouts and heartbeat timeouts are in units of ticks.
	Tick()
	// Campaign causes the Node to transition to candidate state and start campaigning to become leader.
	Campaign(ctx context.Context) error
	// Propose proposes that data be appended to the log.
	Propose(ctx context.Context, pds ...ProposeData)
	// ProposeConfChange proposes a configuration change. Like any proposal, the
	// configuration change may be dropped with or without an error being
	// returned. In particular, configuration changes are dropped unless the
	// leader has certainty that there is no prior unapplied configuration
	// change in its log.
	//
	// The method accepts either a pb.ConfChange (deprecated) or pb.ConfChangeV2
	// message. The latter allows arbitrary configuration changes via joint
	// consensus, notably including replacing a voter. Passing a ConfChangeV2
	// message is only allowed if all Nodes participating in the cluster run a
	// version of this library aware of the V2 API. See pb.ConfChangeV2 for
	// usage details and semantics.
	ProposeConfChange(ctx context.Context, cc pb.ConfChangeI) error

	// Step advances the state machine using the given message. ctx.Err() will be returned, if any.
	Step(ctx context.Context, msg pb.Message) error

	// ApplyConfChange applies a config change (previously passed to
	// ProposeConfChange) to the node. This must be called whenever a config
	// change is observed in Ready.CommittedEntries, except when the app decides
	// to reject the configuration change (i.e. treats it as a noop instead), in
	// which case it must not be called.
	//
	// Returns an opaque non-nil ConfState protobuf which must be recorded in
	// snapshots.
	ApplyConfChange(cc pb.ConfChangeI) *pb.ConfState

	ReportStateStatus(ctx context.Context, term uint64, vote uint64) error
	ReportLogStatus(ctx context.Context, index uint64, term uint64) error
	ReportApplyStatus(ctx context.Context, index uint64) error

	// TransferLeadership attempts to transfer leadership to the given transferee.
	TransferLeadership(ctx context.Context, lead, transferee uint64)

	// ReadIndex request a read state. The read state will be set in the ready.
	// Read state has a read index. Once the application advances further than the read
	// index, any linearizable read requests issued before the read request can be
	// processed safely. The read state will have the same rctx attached.
	// Note that request can be lost without notice, therefore it is user's job
	// to ensure read index retries.
	ReadIndex(ctx context.Context, rctx []byte) error

	// Status returns the current status of the raft state machine.
	Status() Status
	// ReportUnreachable reports the given node is not reachable for the last send.
	ReportUnreachable(id uint64)
	// ReportSnapshot reports the status of the sent snapshot. The id is the raft ID of the follower
	// who is meant to receive the snapshot, and the status is SnapshotFinish or SnapshotFailure.
	// Calling ReportSnapshot with SnapshotFinish is a no-op. But, any failure in applying a
	// snapshot (for e.g., while streaming it from leader to follower), should be reported to the
	// leader with SnapshotFailure. When leader sends a snapshot to a follower, it pauses any raft
	// log probes until the follower can apply the snapshot and advance its state. If the follower
	// can't do that, for e.g., due to a crash, it could end up in a limbo, never getting any
	// updates from the leader. Therefore, it is crucial that the application ensures that any
	// failure in snapshot sending is caught and reported back to the leader; so it can resume raft
	// log probing in the follower.
	ReportSnapshot(id uint64, status SnapshotStatus)
	// Stop performs any necessary termination of the Node.
	Stop()
	Bootstrap(peers []Peer) error
}

type Peer struct {
	ID      uint64
	Context []byte
}

// StartNode returns a new Node given configuration and a list of raft peers.
// It appends a ConfChangeAddNode entry for each given peer to the initial log.
//
// Peers must not be zero length; call RestartNode in that case.
func StartNode(c *Config, peers []Peer) Node { _ = "STUB: not implemented"; return *new(Node) }

// RestartNode is similar to StartNode but does not take a list of peers.
// The current membership of the cluster will be restored from the Storage.
// If the caller has an existing state machine, pass in the last log index that
// has been applied to it; otherwise use zero.
func RestartNode(c *Config) Node { _ = "STUB: not implemented"; return *new(Node) }

type peersWithResult struct {
	peers  []Peer
	result chan error
}

// node is the canonical implementation of the Node interface
type node struct {
	propc      chan []ProposeData
	recvc      chan pb.Message
	confc      chan pb.ConfChangeV2
	confstatec chan pb.ConfState
	bootstrapc chan peersWithResult
	tickc      chan struct{}
	done       chan struct{}
	stop       chan struct{}
	status     chan chan Status

	rn *RawNode
}

func newNode(rn *RawNode) node { _ = "STUB: not implemented"; return *new(node) }

// make tickc a buffered chan, so raft node can buffer some ticks when the node
// is busy processing raft messages. Raft node will resume process buffered
// ticks when it becomes idle.

func (n *node) Stop() { _ = "STUB: not implemented"; return }

// Not already stopped, so trigger it

// Node has already been stopped - no need to do anything

// Block until the stop has been acknowledged by run()

func (n *node) Bootstrap(peers []Peer) error {
	_ = "STUB: not implemented"
	// TODO
	return nil
}

func (n *node) run() {
	_ = "STUB: not implemented" //nolint:funlen // ok
	return
}

// TODO: maybe buffer the config propose if there exists one (the way
// described in raft dissertation)
// Currently it is dropped in Step silently.

// Filter out response message from unknown From.

// If the node was removed, block incoming proposals. Note that we
// only do this if the node was in the config before. Nodes may be
// a member of the group without knowing this (when they're catching
// up on the log and don't have the latest config) and we don't want
// to block the proposal channel in that case.
//
// NB: propc is reset when the leader changes, which, if we learn
// about it, sort of implies that we got readded, maybe? This isn't
// very sound and likely has bugs.

// Tick increments the internal logical clock for this Node. Election timeouts
// and heartbeat timeouts are in units of ticks.
func (n *node) Tick() { _ = "STUB: not implemented"; return }

func (n *node) Campaign(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (n *node) Propose(ctx context.Context, pds ...ProposeData) { _ = "STUB: not implemented"; return }

func (n *node) Step(ctx context.Context, m pb.Message) error { _ = "STUB: not implemented"; return nil }

// ignore unexpected local messages receiving over network

// TODO: return an error?

func confChangeToMsg(c pb.ConfChangeI) (pb.Message, error) {
	_ = "STUB: not implemented"
	return *new(pb.Message), nil
}

func (n *node) ProposeConfChange(ctx context.Context, cc pb.ConfChangeI) error {
	_ = "STUB: not implemented"
	return nil
}

// Step advances the state machine using msgs. The ctx.Err() will be returned,
// if any.
func (n *node) step(ctx context.Context, m pb.Message) error { _ = "STUB: not implemented"; return nil }

func (n *node) ApplyConfChange(cc pb.ConfChangeI) *pb.ConfState {
	_ = "STUB: not implemented"
	return nil
}

func (n *node) ReportStateStatus(ctx context.Context, term uint64, vote uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *node) ReportLogStatus(ctx context.Context, index uint64, term uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *node) ReportApplyStatus(ctx context.Context, index uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *node) Status() Status { _ = "STUB: not implemented"; return *new(Status) }

func (n *node) ReportUnreachable(id uint64) { _ = "STUB: not implemented"; return }

func (n *node) ReportSnapshot(id uint64, status SnapshotStatus) { _ = "STUB: not implemented"; return }

func (n *node) TransferLeadership(ctx context.Context, lead, transferee uint64) {
	_ = "STUB: not implemented"

	// manually set 'from' and 'to', so that leader can voluntarily transfers its leadership
	return
}

func (n *node) ReadIndex(ctx context.Context, rctx []byte) error {
	_ = "STUB: not implemented"
	return nil
}
