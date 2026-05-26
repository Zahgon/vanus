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

type Type uint8

const (
	Constant Type = iota
	EventAttribute
	EventData
	Define
	Any
)

func (t Type) String() string { _ = "STUB: not implemented"; return "" }

type TypeList []Type

func (list TypeList) Contains(arg Arg) bool { _ = "STUB: not implemented"; return false }

var (
	EventList = []Type{EventAttribute, EventData}
	All       = []Type{EventAttribute, EventData, Constant, Define}
)

type Arg interface {
	Type() Type
	Name() string
	Original() string
	// Evaluate arg value
	Evaluate(ceCtx *context.EventContext) (interface{}, error)
	SetValue(ceCtx *context.EventContext, v interface{}) error
	DeleteValue(ceCtx *context.EventContext) error
}

func NewArg(arg interface{}) (Arg, error) { _ = "STUB: not implemented"; return *new(Arg), nil }

const (
	EventArgPrefix        = "$."
	EventDataArgPrefix    = EventArgPrefix + "data"
	EventDataSubArgPrefix = "@."
)
