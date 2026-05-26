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

package meta

import (
	// standard libraries.

	"reflect"
	// third-party libraries.
)

var (
	DeletedMark  deletedMarkType
	defaultCodec codec
)

type Marshaler interface {
	Marshal(data Ranger) ([]byte, error)
}

type Unmarshaler interface {
	Unmarshal(data []byte, cb RangeCallback) error
}

type deletedMarkType struct{}

type Kind uint8

const (
	Invalid Kind = iota
	Deleted
	True
	False
	Bytes
	String
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
)

type codec struct{}

// Make sure codec implements Marshaler and Unmarshaler.
var (
	_ Marshaler   = (*codec)(nil)
	_ Unmarshaler = (*codec)(nil)
)

func (codec) Marshal(data Ranger) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Encode key.

// Encode value.

func appendKey(buf []byte, key, last []byte) []byte { _ = "STUB: not implemented"; return nil }

func encodeKey(key, last []byte) (int, []byte) { _ = "STUB: not implemented"; return 0, nil }

func appendValue(buf []byte, value interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(james.yin): validate type

func toKind(k reflect.Kind) Kind { _ = "STUB: not implemented"; return *new(Kind) }

func (codec) Unmarshal(data []byte, cb RangeCallback) error { _ = "STUB: not implemented"; return nil }

// Decode key.

// Decode value.

func consumeKey(buf []byte, last []byte) ([]byte, int) { _ = "STUB: not implemented"; return nil, 0 }

func decodeKey(private, last []byte, shared int) []byte { _ = "STUB: not implemented"; return nil }

func consumeValue(buf []byte) (interface{}, int) { _ = "STUB: not implemented"; return nil, 0 }
