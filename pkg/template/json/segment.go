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
	// third-party libraries.
	"github.com/ohler55/ojg/jp"

	// this project.
	"github.com/vanus-labs/vanus/lib/bytes"
)

type templateSegment interface {
	RenderTo(w bytes.LastByteWriter, model any, variables map[string]any) error
}

type literalSegment struct {
	val []byte
}

// Make sure literalSegment implements templateSegment.
var _ templateSegment = (*literalSegment)(nil)

func (s *literalSegment) RenderTo(w bytes.LastByteWriter, _ any, _ map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

// skip leading comma

type variableSegment struct {
	name string
}

// Make sure variableSegment implements templateSegment.
var _ templateSegment = (*variableSegment)(nil)

func (s *variableSegment) RenderTo(w bytes.LastByteWriter, _ any, variables map[string]any) error {
	_ = "STUB: not implemented"
	// Variables MUST be defined. But to prevent corner cases, write a "null".
	return nil
}

type variableStringSegment struct {
	name string
}

// Make sure variableStringSegment implements templateSegment.
var _ templateSegment = (*variableStringSegment)(nil)

func (s *variableStringSegment) RenderTo(w bytes.LastByteWriter, _ any, variables map[string]any) error {
	_ = "STUB: not implemented"
	// Variables MUST be defined. But to prevent corner cases, keep it empty.
	return nil
}

type memberSegment struct {
	key   []byte
	value jp.Expr
}

// Make sure memberSegment implements templateSegment.
var _ templateSegment = (*memberSegment)(nil)

func (s *memberSegment) RenderTo(w bytes.LastByteWriter, model any, _ map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

type elementSegment struct {
	value jp.Expr
}

// Make sure elementSegment implements templateSegment.
var _ templateSegment = (*elementSegment)(nil)

func (s *elementSegment) RenderTo(w bytes.LastByteWriter, model any, _ map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

type jsonPathStringSegment struct {
	path jp.Expr
}

// Make sure jsonPathStringSegment implements templateSegment.
var _ templateSegment = (*jsonPathStringSegment)(nil)

func (s *jsonPathStringSegment) RenderTo(w bytes.LastByteWriter, model any, _ map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}
