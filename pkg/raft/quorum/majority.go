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

package quorum

// MajorityConfig is a set of IDs that uses majority quorums to make decisions.
type MajorityConfig map[uint64]struct{}

func (c MajorityConfig) String() string { _ = "STUB: not implemented"; return "" }

// Describe returns a (multi-line) representation of the commit indexes for the
// given lookuper.
func (c MajorityConfig) Describe(l AckedIndexer) string { _ = "STUB: not implemented"; return "" }

// idx found?
// length of bar displayed for this tup

// Below, populate .bar so that the i-th largest commit index has bar i (we
// plot this as sort of a progress bar). The actual code is a bit more
// complicated and also makes sure that equal index => equal bar.

// Sort by index

// Populate .bar.

// Sort by ID.

// Print.

// Slice returns the MajorityConfig as a sorted slice.
func (c MajorityConfig) Slice() []uint64 { _ = "STUB: not implemented"; return nil }

func insertionSort(sl []uint64) { _ = "STUB: not implemented"; return }

// CommittedIndex computes the committed index from those supplied via the
// provided AckedIndexer (for the active config).
func (c MajorityConfig) CommittedIndex(l AckedIndexer) Index {
	_ = "STUB: not implemented"
	return *

	// This plays well with joint quorums which, when one half is the zero
	// MajorityConfig, should behave like the other half.
	new(Index)
}

// Use an on-stack slice to collect the committed indexes when n <= 7
// (otherwise we alloc). The alternative is to stash a slice on
// MajorityConfig, but this impairs usability (as is, MajorityConfig is just
// a map, and that's nice). The assumption is that running with a
// replication factor of >7 is rare, and in cases in which it happens
// performance is a lesser concern (additionally the performance
// implications of an allocation here are far from drastic).

// Fill the slice with the indexes observed. Any unused slots will be
// left as zero; these correspond to voters that may report in, but
// haven't yet. We fill from the right (since the zeroes will end up on
// the left after sorting below anyway).

// Sort by index. Use a bespoke algorithm (copied from the stdlib's sort
// package) to keep srt on the stack.

// The smallest index into the array for which the value is acked by a
// quorum. In other words, from the end of the slice, move n/2+1 to the
// left (accounting for zero-indexing).

// VoteResult takes a mapping of voters to yes/no (true/false) votes and returns
// a result indicating whether the vote is pending (i.e. neither a quorum of
// yes/no has been reached), won (a quorum of yes has been reached), or lost (a
// quorum of no has been reached).
func (c MajorityConfig) VoteResult(votes map[uint64]bool) VoteResult {
	_ = "STUB: not implemented"

	// By convention, the elections on an empty config win. This comes in
	// handy with joint quorums because it'll make a half-populated joint
	// quorum behave like a majority quorum.
	return *new(VoteResult)
}

// vote counts for no and yes, respectively
