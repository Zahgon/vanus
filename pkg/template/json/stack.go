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

package json

type parserContext struct {
	node  templateNode
	state parserState
}

type parserStack struct {
	stack []parserContext
}

func (ps *parserStack) push(n templateNode, s parserState) { _ = "STUB: not implemented"; return }

func (ps *parserStack) pop() (templateNode, parserState) {
	_ = "STUB: not implemented"
	return *new(templateNode), *new(parserState)
}

func (ps *parserStack) peek() (templateNode, parserState) {
	_ = "STUB: not implemented"
	return *new(templateNode), *new(parserState)
}

type generatorContext struct {
	node templateNode
	iter int
}

type generatorStack struct {
	stack []generatorContext
}

func (gs *generatorStack) push(n templateNode) { _ = "STUB: not implemented"; return }

func (gs *generatorStack) advance() { _ = "STUB: not implemented"; return }

func (gs *generatorStack) advanceThenPush(n templateNode) { _ = "STUB: not implemented"; return }

func (gs *generatorStack) pop() bool { _ = "STUB: not implemented"; return false }

func (gs *generatorStack) peek() (templateNode, int) {
	_ = "STUB: not implemented"
	return *new(templateNode), 0
}
