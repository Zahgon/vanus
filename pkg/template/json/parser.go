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
	"errors"

	// this project.
	"github.com/vanus-labs/vanus/lib/bytes"
)

var errParseJSONTemplate = errors.New("cannot parse JSON template")

type parserState int

const (
	waitValue parserState = iota
	waitObjectFirstKey
	waitObjectKey
	waitObjectColon
	waitObjectValue
	waitObjectComma
	waitArrayFirstElement
	waitArrayElement
	waitArrayComma
	waitEOF
)

type templateParser struct {
	stack parserStack
	state parserState
}

func (p *templateParser) parse(text string) (templateNode, error) {
	_ = "STUB: not implemented"
	return *new(templateNode), nil
}

func (p *templateParser) doParse(s *bytes.MarkScanner) error { _ = "STUB: not implemented"; return nil }

//nolint:errorlint // io.EOF is not an error

// empty array

// empty object

func (p *templateParser) expectValue(c byte, s *bytes.MarkScanner) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *templateParser) expectVariable(s *bytes.MarkScanner) (templateNode, error) {
	_ = "STUB: not implemented"
	return *new(templateNode), nil
}

func (p *templateParser) expectDynamicString(s *bytes.MarkScanner) (templateNode, error) {
	_ = "STUB: not implemented"
	return *new(templateNode), nil
}

// string

// variable

func (p *templateParser) expectString(s *bytes.MarkScanner) (*stringNode, error) {
	_ = "STUB: not implemented"
	// exclude '"'
	return nil, nil
}

func (p *templateParser) expectNumber(c byte, s *bytes.MarkScanner) (*numberNode, error) {
	_ = "STUB: not implemented"
	// include c
	return nil, nil
}

func (p *templateParser) reduce() { _ = "STUB: not implemented"; return }

// resume last state

// then do reduce

func (p *templateParser) reduceExt(n templateNode) { _ = "STUB: not implemented"; return }

func (p *templateParser) reduceObjectMember(n templateNode) { _ = "STUB: not implemented"; return }

// expect next member

func (p *templateParser) reduceArrayElement(n templateNode) { _ = "STUB: not implemented"; return }

// expect next element
