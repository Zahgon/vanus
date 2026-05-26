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

package strings

import (
	"regexp"
	"sync"

	"github.com/vanus-labs/vanus/pkg/transform/action"
	"github.com/vanus-labs/vanus/pkg/transform/arg"
	"github.com/vanus-labs/vanus/pkg/transform/context"
)

type replaceWithRegexAction struct {
	action.CommonAction
	pattern *regexp.Regexp
	expr    string
	lock    sync.RWMutex
}

// NewReplaceWithRegexAction ["replace_with_regex", "key", "pattern", "value"].
func NewReplaceWithRegexAction() action.Action {
	_ = "STUB: not implemented"
	return *new(action.Action)
}

func (a *replaceWithRegexAction) Init(args []arg.Arg) error { _ = "STUB: not implemented"; return nil }

func (a *replaceWithRegexAction) Execute(ceCtx *context.EventContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *replaceWithRegexAction) setPattern(expr string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *replaceWithRegexAction) getPattern() *regexp.Regexp { _ = "STUB: not implemented"; return nil }
