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

package path

import (
	// standard libraries.
	"errors"

	// this project.
	"github.com/vanus-labs/vanus/lib/bytes"
)

var errInvalidJSONPath = errors.New("invalid JSON path")

func Parse(text string) (Path, error) { _ = "STUB: not implemented"; return *new(Path), nil }

// root identifier

func ConsumeExt(c byte, s *bytes.MarkScanner) (Path, error) {
	_ = "STUB: not implemented"
	return *new(Path), nil
}

func ConsumeSegments(s *bytes.MarkScanner) ([]Segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func consumeSegment(s *bytes.MarkScanner) (Segment, error) {
	_ = "STUB: not implemented"
	return *new(Segment), nil
}

func consumeBracketedSelection(s *bytes.MarkScanner) (Segment, error) {
	_ = "STUB: not implemented"
	return *new(Segment), nil
}

// end of bracketed selection

// continue

func consumeSelectorExt(c byte, s *bytes.MarkScanner) (Selector, error) {
	_ = "STUB: not implemented"
	return *

	// name selector
	new(Selector), nil
}

// name selector

// wildcard selector

// array slice selector

// filter selector
// TODO(james.yin)

func consumeIndexOrSliceSelector(c byte, s *bytes.MarkScanner) (Selector, error) {
	_ = "STUB: not implemented"
	return *new(Selector), nil
}

// array slice selector

func consumeSliceSelector(start *int, s *bytes.MarkScanner) (*arraySliceSelector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func consumeDotSegment(s *bytes.MarkScanner) (Segment, error) {
	_ = "STUB: not implemented"
	return *new(Segment), nil
}

// wildcard selector

// descendant segment
// TODO(james.yin)

// member-name-shorthand

func consumeMemberNameShorthandExt(r rune, s *bytes.MarkScanner) (Segment, error) {
	_ = "STUB: not implemented"
	return *new(Segment), nil
}
