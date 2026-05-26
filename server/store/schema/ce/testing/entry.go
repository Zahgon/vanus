// Copyright 2022 Linkall Inc.
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

package testing

import (
	// standard libraries.

	// third-party libraries.
	. "github.com/smartystreets/goconvey/convey"
	. "go.uber.org/mock/gomock"

	// this project.
	"github.com/vanus-labs/vanus/server/store/block"
	blktest "github.com/vanus-labs/vanus/server/store/block/testing"
)

var (
	value0 = []byte{0x01}
	value1 = []byte{0x02}
	value2 = []byte{0x78, 0x56, 0x34, 0x12, 0x03}
	value3 = []byte("value3\x04")
	value4 = []byte("value4\x05")
	value5 = []byte("value5\x06")
	value6 = []byte("value6\x07")
	value7 = []byte{
		0x2E, 0xE3, 0x06, 0x63, 0x00, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x08,
	}
)

func MakeEntry0(ctrl *Controller) block.EntryExt {
	_ = "STUB: not implemented"
	return *new(block.EntryExt)
}

func MakeEntry1(ctrl *Controller) block.EntryExt {
	_ = "STUB: not implemented"
	return *new(block.EntryExt)
}

func setEntry1Extension(entry *blktest.MockEntryExt, raw bool) { _ = "STUB: not implemented"; return }

func MakeStoredEntry0(ctrl *Controller) block.EntryExt {
	_ = "STUB: not implemented"
	return *new(block.EntryExt)
}

func MakeStoredEntry1(ctrl *Controller, raw bool) block.EntryExt {
	_ = "STUB: not implemented"
	return *new(block.EntryExt)
}

func MakeStoredEndEntry(ctrl *Controller) block.EntryExt {
	_ = "STUB: not implemented"
	return *new(block.EntryExt)
}

func CheckEntry0(entry block.Entry, ignoreSeq, ignoreStime bool) { _ = "STUB: not implemented"; return }

func CheckEntryExt0(entry block.EntryExt) { _ = "STUB: not implemented"; return }

func CheckEntry1(entry block.Entry, ignoreSeq, ignoreStime bool) { _ = "STUB: not implemented"; return }

func CheckEntryExt1(entry block.EntryExt) { _ = "STUB: not implemented"; return }

func CheckEndEntry(entry block.Entry, ignoreStime bool) { _ = "STUB: not implemented"; return }
