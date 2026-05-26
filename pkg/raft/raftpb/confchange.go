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

package raftpb

// ConfChangeI abstracts over ConfChangeV2 and (legacy) ConfChange to allow
// treating them in a unified manner.
type ConfChangeI interface {
	AsV2() ConfChangeV2
	AsV1() (ConfChange, bool)
}

// MarshalConfChange calls Marshal on the underlying ConfChange or ConfChangeV2
// and returns the result along with the corresponding EntryType.
func MarshalConfChange(c ConfChangeI) (EntryType, []byte, error) {
	_ = "STUB: not implemented"
	return *new(EntryType), nil, nil
}

// AsV2 returns a V2 configuration change carrying out the same operation.
func (c ConfChange) AsV2() ConfChangeV2 { _ = "STUB: not implemented"; return *new(ConfChangeV2) }

// AsV1 returns the ConfChange and true.
func (c ConfChange) AsV1() (ConfChange, bool) {
	_ = "STUB: not implemented"

	// AsV2 is the identity.
	return *new(ConfChange), false
}

func (c ConfChangeV2) AsV2() ConfChangeV2 {
	_ = "STUB: not implemented"

	// AsV1 returns ConfChange{} and false.
	return *new(ConfChangeV2)
}

func (c ConfChangeV2) AsV1() (ConfChange, bool) {
	_ = "STUB: not implemented"
	return *new(ConfChange), false
}

// EnterJoint returns two bools. The second bool is true if and only if this
// config change will use Joint Consensus, which is the case if it contains more
// than one change or if the use of Joint Consensus was requested explicitly.
// The first bool can only be true if second one is, and indicates whether the
// Joint State will be left automatically.
func (c ConfChangeV2) EnterJoint() (autoLeave bool, ok bool) {
	_ = "STUB: not implemented"
	// NB: in theory, more config changes could qualify for the "simple"
	// protocol but it depends on the config on top of which the changes apply.
	// For example, adding two learners is not OK if both nodes are part of the
	// base config (i.e. two voters are turned into learners in the process of
	// applying the conf change). In practice, these distinctions should not
	// matter, so we keep it simple and use Joint Consensus liberally.
	return false, false
}

// Use Joint Consensus.

// LeaveJoint is true if the configuration change leaves a joint configuration.
// This is the case if the ConfChangeV2 is zero, with the possible exception of
// the Context field.
func (c ConfChangeV2) LeaveJoint() bool {
	_ = "STUB: not implemented"
	// NB: c is already a copy.
	return false
}

// ConfChangesFromString parses a Space-delimited sequence of operations into a
// slice of ConfChangeSingle. The supported operations are:
// - vn: make n a voter,
// - ln: make n a learner,
// - rn: remove n, and
// - un: update n.
func ConfChangesFromString(s string) ([]ConfChangeSingle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConfChangesToString is the inverse to ConfChangesFromString.
func ConfChangesToString(ccs []ConfChangeSingle) string { _ = "STUB: not implemented"; return "" }
