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
)

func (env *InteractionEnv) handleRaftLog(t *testing.T, d datadriven.TestData) error {
	_ = "STUB: not implemented"
	return nil
}

// RaftLog pretty prints the raft log to the output buffer.
func (env *InteractionEnv) RaftLog(idx int) error { _ = "STUB: not implemented"; return nil }

// TODO(tbg): this is what MemoryStorage returns, but unclear if it's
// the "correct" thing to do.
