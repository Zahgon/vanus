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

package action

import (
	"fmt"

	"github.com/vanus-labs/vanus/pkg/transform/arg"
	"github.com/vanus-labs/vanus/pkg/transform/common"
	"github.com/vanus-labs/vanus/pkg/transform/context"
	"github.com/vanus-labs/vanus/pkg/transform/function"
)

type Action interface {
	// Name func name
	Name() string
	// Arity arg number
	Arity() int
	// ArgType arg type
	ArgType(index int) arg.TypeList
	// IsVariadic is exist variadic
	IsVariadic() bool
	Init(args []arg.Arg) error
	Execute(ceCtx *context.EventContext) error
}

type CommonAction struct {
	ActionName  string
	FixedArgs   []arg.TypeList
	VariadicArg arg.TypeList
	Fn          function.Function

	Args      []arg.Arg
	ArgTypes  []common.Type
	TargetArg arg.Arg
}

func (a *CommonAction) Name() string { _ = "STUB: not implemented"; return "" }

func (a *CommonAction) Arity() int { _ = "STUB: not implemented"; return 0 }

func (a *CommonAction) ArgType(index int) arg.TypeList {
	_ = "STUB: not implemented"
	return *new(arg.TypeList)
}

func (a *CommonAction) IsVariadic() bool { _ = "STUB: not implemented"; return false }

func (a *CommonAction) RunArgs(ceCtx *context.EventContext) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type FunctionAction struct {
	CommonAction
}

func (a *FunctionAction) Init(args []arg.Arg) error { _ = "STUB: not implemented"; return nil }

func (a *FunctionAction) setArgTypes() error { _ = "STUB: not implemented"; return nil }

func (a *FunctionAction) Execute(ceCtx *context.EventContext) error {
	_ = "STUB: not implemented"
	return nil
}

type SourceTargetSameAction struct {
	FunctionAction
}

func (a *SourceTargetSameAction) Init(args []arg.Arg) error { _ = "STUB: not implemented"; return nil }

var (
	ErrExist     = fmt.Errorf("action have exist")
	ErrArgNumber = fmt.Errorf("action arg number invalid")
)

type NestAction interface {
	Action
	InitAction(actions []Action) error
}

type NestActionImpl struct {
	CommonAction
	Actions []Action
}

func (c *NestActionImpl) InitAction(actions []Action) error { _ = "STUB: not implemented"; return nil }
