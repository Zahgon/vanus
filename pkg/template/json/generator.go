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

import (
	// standard libraries.
	stdbytes "bytes"
	// third-party libraries.
)

type templateGenerator struct {
	stack    generatorStack
	buf      stdbytes.Buffer
	segments []templateSegment
}

func (g *templateGenerator) generate(root templateNode) []templateSegment {
	_ = "STUB: not implemented"
	return nil
}

// unreachable

func (g *templateGenerator) generateObjectNode(node *objectNode, i int) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *templateGenerator) generateArrayNode(node *arrayNode, i int) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *templateGenerator) generateDynamicStringNode(node *dynamicStringNode) {
	_ = "STUB: not implemented"
	return
}

// unreachable

func (g *templateGenerator) insertSegment(segment templateSegment) {
	_ = "STUB: not implemented"
	return
}

func (g *templateGenerator) packLiteral() { _ = "STUB: not implemented"; return }
