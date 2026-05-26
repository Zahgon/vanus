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

package text

import ( // standard libraries.
	// third-party libraries.
	// this project.
	"github.com/vanus-labs/vanus/lib/bytes"
)

const escapePlan = "" +
	`................................` + // 0x00
	`................oooooooo....s.s.` + // 0x20, 0-7(0x30-37), angled brackets(0x3c,0x3e)
	`............................s...` + // 0x40, backslash(0x5c)
	"..\b...\f.......\n...\r.\tu..x......." + // 0x60
	`................................` + // 0x80
	`................................` + // 0xa0
	`................................` + // 0xc0
	`................................` //  0xe0

func parse(text string) ([]templateSegment, error) { _ = "STUB: not implemented"; return nil, nil }

// maybe unnecessary?

func expectVariable(s *bytes.MarkScanner) (templateSegment, error) {
	_ = "STUB: not implemented"
	return *new(templateSegment), nil
}
