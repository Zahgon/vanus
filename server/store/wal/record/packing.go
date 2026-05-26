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

func Pack(entry []byte, firstSize, otherSize int) ([]Record, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

// first packet

// middle packet(s)

// last packet

func calPacketNum(entry []byte, firstSize, otherSize int) int { _ = "STUB: not implemented"; return 0 }

// 1 + ((payload-(firstSize-HeaderSize))+((otherSize-HeaderSize)-1))/(otherSize-HeaderSize)

func makePacket(t Type, payload []byte) Record { _ = "STUB: not implemented"; return *new(Record) }
