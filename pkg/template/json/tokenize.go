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
	"io"

	// this project.
	"github.com/vanus-labs/vanus/lib/bytes"
)

var (
	errInvalidDynamicString = errors.New("invalid dynamic string")
	errInvalidString        = errors.New("invalid JSON string")
	errInvalidNumber        = errors.New("invalid JSON number")
	errInvalidNull          = errors.New("invalid JSON null")
	errInvalidTrue          = errors.New("invalid JSON true")
	errInvalidFalse         = errors.New("invalid JSON false")
)

const (
	// The lowest and highest control characters.
	locc = 0x00
	hicc = 0x1F
)

const dynamicStringPlan = "" +
	`................................` + // 0x00
	`..s............s............s.s.` + // 0x20, double quote(0x22), slash(0x2f), angled brackets(0x3c,0x3e)
	`............................s...` + // 0x40, backslash(0x5c)
	"..\b...\f.......\n...\r.\tu.........." + // 0x60
	`................................` + // 0x80
	`................................` + // 0xa0
	`................................` + // 0xc0
	`................................` //  0xe0

func skipWhitespace(r io.ByteReader) (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func consumeDynamicString(r io.ByteReader, w io.ByteWriter) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// quotation mark, end of string

// open angled bracket, begin of variable

// reverse solidus

// control characters

func unescapeBracket(bs []byte) []byte { _ = "STUB: not implemented"; return nil }

type number struct {
	negative, negExp            bool
	integer, fraction, exponent []byte
}

func expectNumberExt(c byte, s *bytes.MarkScanner) (number, error) {
	_ = "STUB: not implemented"
	return *new(number), nil
}

// at least one digit

// fraction

// exponent

func consumeDigits(r io.ByteReader, w io.ByteWriter) (int, bool, byte, error) {
	_ = "STUB: not implemented"
	return 0, false, 0, nil
}

func exceptNullExt(r io.ByteReader) error { _ = "STUB: not implemented"; return nil }

func exceptTrueExt(r io.ByteReader) error { _ = "STUB: not implemented"; return nil }

func exceptFalseExt(r io.ByteReader) error { _ = "STUB: not implemented"; return nil }
