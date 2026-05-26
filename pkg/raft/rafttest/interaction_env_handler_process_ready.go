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

func (env *InteractionEnv) handleProcessReady(t *testing.T, d datadriven.TestData) error {
	_ = "STUB: not implemented"
	return nil
}

// ProcessReady runs Ready handling on the node with the given index.
func (env *InteractionEnv) ProcessReady(idx int) error {
	_ = "STUB: not implemented"
	// TODO(tbg): Allow simulating crashes here.
	return nil
}

// TODO(tbg): the order of operations here is not necessarily safe. See:
// https://github.com/etcd-io/etcd/pull/10861

// Record the new state by starting with the current state and applying
// the command.

// NB: this hard-codes an "appender" state machine.
