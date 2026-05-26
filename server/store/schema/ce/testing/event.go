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
	"time"

	// third-party libraries.
	ce "github.com/cloudevents/sdk-go/v2"
	. "github.com/smartystreets/goconvey/convey"

	// first-party libraries.
	cepb "github.com/vanus-labs/vanus/api/cloudevents"
)

const (
	seq0              int64 = 0
	seq1              int64 = 1
	seq2              int64 = 2
	ceID0                   = "ce-id0"
	ceID1                   = "ce-id1"
	ceSource                = "ce-source"
	ceSpecVersion           = "1.0"
	ceType                  = "ce-type"
	ceDataContentType       = ce.TextPlain
	ceSubject               = "ce-subject"
)

var (
	Stime  int64 = 0x182D2E76BF3
	ceData       = []byte("ce-data")
	ceTime       = time.Unix(0x6306E32E, 0x04030201)
)

func MakeEvent0() *cepb.CloudEvent { _ = "STUB: not implemented"; return nil }

func MakeEvent1() *cepb.CloudEvent { _ = "STUB: not implemented"; return nil }

func CheckEvent0(event *cepb.CloudEvent) { _ = "STUB: not implemented"; return }

func CheckEvent1(event *cepb.CloudEvent) { _ = "STUB: not implemented"; return }
