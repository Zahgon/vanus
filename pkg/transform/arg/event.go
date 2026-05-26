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

type eventAttribute struct {
	attr     string
	original string
}

// newEventAttribute name format is $.source .
func newEventAttribute(name string) (Arg, error) { _ = "STUB: not implemented"; return *new(Arg), nil }

func (arg eventAttribute) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (arg eventAttribute) Name() string { _ = "STUB: not implemented"; return "" }

func (arg eventAttribute) Original() string { _ = "STUB: not implemented"; return "" }

func (arg eventAttribute) Evaluate(ceCtx *context.EventContext) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (arg eventAttribute) SetValue(ceCtx *context.EventContext, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (arg eventAttribute) DeleteValue(ceCtx *context.EventContext) error {
	_ = "STUB: not implemented"
	return nil
}

type eventData struct {
	path     string
	original string
}

// newEventData name format is $.data.key .
func newEventData(name string) Arg { _ = "STUB: not implemented"; return *new(Arg) }

func (arg eventData) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (arg eventData) Name() string { _ = "STUB: not implemented"; return "" }

func (arg eventData) Original() string { _ = "STUB: not implemented"; return "" }

func (arg eventData) Evaluate(ceCtx *context.EventContext) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (arg eventData) SetValue(ceCtx *context.EventContext, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (arg eventData) DeleteValue(ceCtx *context.EventContext) error {
	_ = "STUB: not implemented"
	return nil
}

type eventDataAll struct {
	eventData
}

func (arg eventDataAll) Evaluate(ceCtx *context.EventContext) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (arg eventDataAll) SetValue(ceCtx *context.EventContext, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (arg eventDataAll) DeleteValue(ceCtx *context.EventContext) error {
	_ = "STUB: not implemented"
	return nil
}
