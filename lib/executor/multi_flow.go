// Copyright 2023 Linkall Inc.
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

package executor

import (
	// standard libraries.

	// this project.
	"github.com/vanus-labs/vanus/lib/container/conque/blocking"
	"github.com/vanus-labs/vanus/lib/container/conque/unbounded"
)

const defaultInvokeBatchSize = 8

type flow struct {
	q     unbounded.Queue[Task]
	state int32
	mf    *MultiFlow
}

// Make sure flow implements ExecuteCloser.
var _ ExecuteCloser = (*flow)(nil)

func (f *flow) Execute(t Task) bool { _ = "STUB: not implemented"; return false }

func (f *flow) Close() { _ = "STUB: not implemented"; return }

func (f *flow) invokeTasks(batch int) bool { _ = "STUB: not implemented"; return false }

type MultiFlow struct {
	q        blocking.Queue[*flow]
	parallel int
}

func NewMultiFlow(parallel int, handoff, startImmediately bool) *MultiFlow {
	_ = "STUB: not implemented"
	return nil
}

func (mf *MultiFlow) Init(parallel int, handoff, startImmediately bool) *MultiFlow {
	_ = "STUB: not implemented"
	return nil
}

func (mf *MultiFlow) Start() { _ = "STUB: not implemented"; return }

func (mf *MultiFlow) Close() { _ = "STUB: not implemented"; return }

func (mf *MultiFlow) NewFlow() ExecuteCloser { _ = "STUB: not implemented"; return *new(ExecuteCloser) }

func (mf *MultiFlow) run() { _ = "STUB: not implemented"; return }
