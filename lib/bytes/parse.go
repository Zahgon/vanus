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

package bytes

import (
	// standard libraries.
	"errors"
	"io"
)

var errUnexpectedChar = errors.New("unexpected character")

func ExpectChar(r io.ByteReader, c byte) error { _ = "STUB: not implemented"; return nil }

func ConsumeUntil(r io.ByteReader, w io.ByteWriter, stop func(byte) bool) (int, byte, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func Skip(r io.ByteReader, expect func(byte) bool) (int, byte, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func IgnoreCount(_ int, c byte, err error) (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func AcceptEOF(count int, c byte, err error) (int, bool, byte, error) {
	_ = "STUB: not implemented"
	return 0, false, 0, nil
}

//nolint:errorlint // io.EOF is not an error

func Unread(r io.ByteScanner, eof bool, err error) error { _ = "STUB: not implemented"; return nil }
