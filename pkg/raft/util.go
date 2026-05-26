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
)

func (st StateType) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func min(a, b uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func max(a, b uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func IsLocalMsg(msgt pb.MessageType) bool { _ = "STUB: not implemented"; return false }

func IsResponseMsg(msgt pb.MessageType) bool { _ = "STUB: not implemented"; return false }

// voteResponseType maps vote and prevote message types to their corresponding responses.
func voteRespMsgType(msgt pb.MessageType) pb.MessageType {
	_ = "STUB: not implemented"
	return *new(pb.MessageType)
}

func DescribeHardState(hs pb.HardState) string { _ = "STUB: not implemented"; return "" }

func DescribeSoftState(ss SoftState) string { _ = "STUB: not implemented"; return "" }

func DescribeConfState(state pb.ConfState) string { _ = "STUB: not implemented"; return "" }

func DescribeSnapshot(snap pb.Snapshot) string { _ = "STUB: not implemented"; return "" }

func DescribeReady(rd Ready, f EntryFormatter) string { _ = "STUB: not implemented"; return "" }

// EntryFormatter can be implemented by the application to provide human-readable formatting
// of entry data. Nil is a valid EntryFormatter and will use a default format.
type EntryFormatter func([]byte) string

// DescribeMessage returns a concise human-readable description of a
// Message for debugging.
func DescribeMessage(m pb.Message, f EntryFormatter) string { _ = "STUB: not implemented"; return "" }

// PayloadSize is the size of the payload of this Entry. Notably, it does not
// depend on its Index or Term.
func PayloadSize(e pb.Entry) int {
	_ = "STUB: not implemented"

	// DescribeEntry returns a concise human-readable description of an
	// Entry for debugging.
	return 0
}

func DescribeEntry(e pb.Entry, f EntryFormatter) string { _ = "STUB: not implemented"; return "" }

// TODO(tbg): give the EntryFormatter a type argument so that it gets
// a chance to expose the Context.

// DescribeEntries calls DescribeEntry for each Entry, adding a newline to
// each.
func DescribeEntries(ents []pb.Entry, f EntryFormatter) string {
	_ = "STUB: not implemented"
	return ""
}

func limitSize(ents []pb.Entry, maxSize uint64) []pb.Entry { _ = "STUB: not implemented"; return nil }

func doLimitSize(ents []pb.Entry, maxSize uint64) []pb.Entry { _ = "STUB: not implemented"; return nil }

func assertConfStatesEquivalent(l Logger, cs1, cs2 pb.ConfState) { _ = "STUB: not implemented"; return }
