// Copyright 2022 Linkall Inc.
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

package pipeline

import (
	primitive "github.com/vanus-labs/vanus/pkg"
	"github.com/vanus-labs/vanus/pkg/transform/action"
	"github.com/vanus-labs/vanus/pkg/transform/context"
)

type Pipeline struct {
	actions []action.Action
}

func NewPipeline() *Pipeline { _ = "STUB: not implemented"; return nil }

func (p *Pipeline) Parse(actions []*primitive.Action) { _ = "STUB: not implemented"; return }

// it has check in controller so err must be nil otherwise controller check has a bug

func (p *Pipeline) Run(ceCtx *context.EventContext) error { _ = "STUB: not implemented"; return nil }
