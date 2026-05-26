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

package record

import (
	// standard libraries.

	"hash/crc32"
)

const (
	crcFieldSO    = 0
	crcFieldEO    = crcFieldSO + 4    // [0,4)
	lengthFieldSO = crcFieldEO        // 4
	lengthFieldEO = lengthFieldSO + 2 // [4,6)
	typeFieldSO   = lengthFieldEO     // 6
	typeFieldEO   = typeFieldSO + 1   // [6,7)
	dataFieldSO   = typeFieldEO       // [7,n)
)

const HeaderSize = dataFieldSO // 7

var crc32q = crc32.MakeTable(crc32.Castagnoli)

type Type uint8

const (
	Zero Type = iota
	Full
	First
	Middle
	Last
)

func (t Type) IsTerminal() bool { _ = "STUB: not implemented"; return false }

func (t Type) IsNonTerminal() bool { _ = "STUB: not implemented"; return false }

type Record struct {
	// CRC is crc32c of Type and Data
	CRC uint32
	// Length is len(Data). optimize?
	Length uint16
	Type   Type
	Data   []byte
}

func (r *Record) Size() int { _ = "STUB: not implemented"; return 0 }

func (r *Record) Marshal() []byte { _ = "STUB: not implemented"; return nil }

func (r *Record) MarshalTo(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// TODO(james.yin): correct error.

// calculate CRC

func Unmarshal(data []byte) (record Record, err error) {
	_ = "STUB: not implemented"
	return *new(Record), nil
}

// return empty record

// TODO(james.yin): correct error
