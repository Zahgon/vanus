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

package arg

import (
	"github.com/vanus-labs/vanus/pkg/transform/context"
)

type define struct {
	name     string
	original string
}

// newDefine name format is <var> .
func newDefine(name string) Arg { _ = "STUB: not implemented"; return *new(Arg) }

func (arg define) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (arg define) Name() string { _ = "STUB: not implemented"; return "" }

func (arg define) Original() string { _ = "STUB: not implemented"; return "" }

func (arg define) Evaluate(ceCtx *context.EventContext) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (arg define) SetValue(*context.EventContext, interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (arg define) DeleteValue(*context.EventContext) error { _ = "STUB: not implemented"; return nil }
