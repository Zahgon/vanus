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

package convert

import (
	// standard libraries.

	"sync/atomic"
	"time"

	// third-party libraries.

	// first-party libraries.
	cepb "github.com/vanus-labs/vanus/api/cloudevents"

	// this project.
	"github.com/vanus-labs/vanus/server/store/block"
)

const (
	dataContentTypeAttr = "datacontenttype"
	dataSchemaAttr      = "dataschema"
	subjectAttr         = "subject"
	timeAttr            = "time"
)

var emptyAttrs = make([]string, 0)

type ceEntry struct {
	block.EmptyEntry
	ce       *cepb.CloudEvent
	extAttrs atomic.Value
}

// Make sure ceEntry implements block.EntryExt.
var _ block.EntryExt = (*ceEntry)(nil)

func (e *ceEntry) GetBytes(ordinal int) []byte { _ = "STUB: not implemented"; return nil }

func (e *ceEntry) GetString(ordinal int) string { _ = "STUB: not implemented"; return "" }

func (e *ceEntry) GetTime(ordinal int) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (e *ceEntry) RangeOptionalAttributes(cb block.OptionalAttributeCallback) {
	_ = "STUB: not implemented"
	// id, source, specversion, type, datacontenttype, dataschema, subject, time
	return
}

// TODO(james.yin): TypeUrl

func (e *ceEntry) OptionalAttributeCount() int { _ = "STUB: not implemented"; return 0 }

func (e *ceEntry) GetExtensionAttribute(attr []byte) []byte { _ = "STUB: not implemented"; return nil }

func (e *ceEntry) RangeExtensionAttributes(cb block.ExtensionAttributeCallback) {
	_ = "STUB: not implemented"
	return
}

// Make sure the order of attributes.

func (e *ceEntry) ExtensionAttributeCount() int { _ = "STUB: not implemented"; return 0 }

func attrValue(v *cepb.CloudEvent_CloudEventAttributeValue) block.ValueMarshaler {
	_ = "STUB: not implemented"
	return *new(block.ValueMarshaler)
}

func ToEntry(event *cepb.CloudEvent) block.EntryExt {
	_ = "STUB: not implemented"
	return *new(block.EntryExt)
}
