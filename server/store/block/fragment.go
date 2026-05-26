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

//go:generate mockgen -source=fragment.go -destination=testing/mock_fragment.go -package=testing
package block

// Fragment is a fragment of a block.
//
// The layout of `Fragment` is:
//
//	┌───────────────────────────────────┬───────────────────────────────────┐
//	│             Offset(8)             │              Payload ...          │
//	└───────────────────────────────────┴───────────────────────────────────┘
//
// All values little-endian
//
//	+00 8B Offset
//	+08    Payload
type Fragment interface {
	Payload() []byte
	Size() int
	StartOffset() int64
	EndOffset() int64
}

type FragmentMarshaler interface {
	MarshalFragment() ([]byte, error)
}

func MarshalFragment(frag Fragment) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type fragment struct {
	data []byte
}

// Make sure fragment implements Fragment and FragmentMarshaler.
var (
	_ Fragment          = (*fragment)(nil)
	_ FragmentMarshaler = (*fragment)(nil)
)

func NewFragment(data []byte) Fragment { _ = "STUB: not implemented"; return *new(Fragment) }

func (f *fragment) Payload() []byte { _ = "STUB: not implemented"; return nil }

func (f *fragment) Size() int { _ = "STUB: not implemented"; return 0 }

func (f *fragment) StartOffset() int64 { _ = "STUB: not implemented"; return 0 }

func (f *fragment) EndOffset() int64 { _ = "STUB: not implemented"; return 0 }

func (f *fragment) MarshalFragment() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
