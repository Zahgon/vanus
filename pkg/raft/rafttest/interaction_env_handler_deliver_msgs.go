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

func (env *InteractionEnv) handleDeliverMsgs(t *testing.T, d datadriven.TestData) error {
	_ = "STUB: not implemented"
	return nil
}

type Recipient struct {
	ID   uint64
	Drop bool
}

// DeliverMsgs goes through env.Messages and, depending on the Drop flag,
// delivers or drops messages to the specified Recipients. Returns the
// number of messages handled (i.e. delivered or dropped). A handled message
// is removed from env.Messages.
func (env *InteractionEnv) DeliverMsgs(rs ...Recipient) int { _ = "STUB: not implemented"; return 0 }

// NB: it's allowed to drop messages to nodes that haven't been instantiated yet,
// we haven't used msg.To yet.
