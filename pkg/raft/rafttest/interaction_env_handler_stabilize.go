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

package rafttest

import (
	"testing"

	"github.com/cockroachdb/datadriven"

	"github.com/vanus-labs/vanus/pkg/raft/raftpb"
)

func (env *InteractionEnv) handleStabilize(t *testing.T, d datadriven.TestData) error {
	_ = "STUB: not implemented"
	return nil
}

// Stabilize repeatedly runs Ready handling on and message delivery to the set
// of nodes specified via the idxs slice until reaching a fixed point.
func (env *InteractionEnv) Stabilize(idxs ...int) error { _ = "STUB: not implemented"; return nil }

// NB: we grab the messages just to see whether to print the header.
// DeliverMsgs will do it again.

func splitMsgs(msgs []raftpb.Message, to uint64) (toMsgs []raftpb.Message, rmdr []raftpb.Message) {
	_ = "STUB: not implemented"
	// NB: this method does not reorder messages.
	return nil, nil
}
