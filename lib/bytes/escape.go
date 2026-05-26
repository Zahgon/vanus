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

var (
	errInvalidEscapeChar  = errors.New("invalid escape character")
	errInvalidUnicodeChar = errors.New("invalid unicode character")
	errInvalidHexChar     = errors.New("invalid hexadecimal character")
	errInvalidOctChar     = errors.New("invalid octal character")
)

const (
	highSurrogateMin = 0xD800
	highSurrogateMax = 0xDBFF
	lowSurrogateMin  = 0xDC00
	lowSurrogateMax  = 0xDFFF
)

var (
	octBitmap [256]bool
	hexBitmap [256]bool
	hexToByte [256]byte
	hexToRune [256]rune
)

func init() { //nolint:gochecknoinits // init constant table
	for i := '0'; i <= '7'; i++ {
		octBitmap[i] = true
	}

	for i := '0'; i <= '9'; i++ {
		hexBitmap[i] = true
		hexToRune[i] = i - '0'
		hexToByte[i] = byte(hexToRune[i])
	}
	shift := 'a' - 'A'
	for i := 'A'; i <= 'F'; i++ {
		hexBitmap[i] = true
		hexBitmap[i+shift] = true
		hexToRune[i] = i - 'A' + 10 //nolint:gomnd // base number
		hexToRune[i+shift] = hexToRune[i]
		hexToByte[i] = byte(hexToRune[i])
		hexToByte[i+shift] = hexToByte[i]
	}
}

func ConsumeEscaped(r io.ByteReader, w io.ByteWriter, plan string) error {
	_ = "STUB: not implemented"
	return nil
}

func consumeEscapedExt(c byte, r io.ByteReader, w io.ByteWriter, plan string) error {
	_ = "STUB: not implemented"
	return nil
}

// Self

// \uNNNN

// \xNN

// \NNN

func ExpectUnicodeChar(r io.ByteReader) (rune, error) { _ = "STUB: not implemented"; return 0, nil }

// non-surrogate

// error of high-surrogate

// error of low-surrogate

func expectUnicodeSurrogate(r io.ByteReader) (rune, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ExpectHexChar(r io.ByteReader) (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func ExpectOctCharExt(b0 byte, r io.ByteReader) (byte, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
