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

package template

import (
	// standard libraries.
	"errors"
	"io"

	// this project.
	"github.com/vanus-labs/vanus/lib/bytes"
	jp "github.com/vanus-labs/vanus/lib/json/path"
)

var errVariable = errors.New("invalid variable")

func ExpectVariable(s *bytes.MarkScanner) ([]byte, jp.Path, error) {
	_ = "STUB: not implemented"
	// exclude '<'
	return nil, *new(jp.Path), nil
}

// JSON path, begin with the root identifier '$'

func consumeIdentifierExt(c byte, r io.ByteReader, w io.ByteWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// close angled bracket, end of variable
