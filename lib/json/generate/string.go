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

// Inspired by https://github.com/ohler55/ojg/blob/9a6023ee37455ae9a61af3368f0635a2fabd2021/string.go

package generate

// standard libraries.

const (
	hex = "0123456789abcdef"

	hioff  = 4
	himask = 0xf0
	lomask = 0x0f
)

const basicPlan = "" +
	`........btn.fr..................` + // 0x00
	`oo"ooooooooooooooooooooooooooooo` + // 0x20
	`oooooooooooooooooooooooooooo\ooo` + // 0x40
	`ooooooooooooooooooooooooooooooo.` + // 0x60
	`88888888888888888888888888888888` + // 0x80
	`88888888888888888888888888888888` + // 0xa0
	`88888888888888888888888888888888` + // 0xc0
	`88888888888888888888888888888888` //  0xe0

const htmlSafePlan = "" + //nolint:unused // reserved for future use.
	`........btn.fr..................` + // 0x00
	`oo"ooo.ooooooooooooooooooooo.o.o` + // 0x20
	`oooooooooooooooooooooooooooo\ooo` + // 0x40
	`ooooooooooooooooooooooooooooooo.` + // 0x60
	`88888888888888888888888888888888` + // 0x80
	`88888888888888888888888888888888` + // 0xa0
	`88888888888888888888888888888888` + // 0xc0
	`88888888888888888888888888888888` //  0xe0

// AppendString appends a JSON encoding of a string to the provided byte slice.
func AppendString(dst []byte, s string) []byte { _ = "STUB: not implemented"; return nil }

func AppendRawString(dst []byte, s string) []byte { _ = "STUB: not implemented"; return nil }

func appendRawString(dst []byte, bs []byte, plan string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// maybe unnecessary?

func AppendByteAsUnicode(dst []byte, b byte) []byte { _ = "STUB: not implemented"; return nil }
