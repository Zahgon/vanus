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

	"time"

	// third-party libraries.

	// first-party libraries.
	cepb "github.com/vanus-labs/vanus/api/cloudevents"

	// this project.
	"github.com/vanus-labs/vanus/server/store/block"
)

type ceWrapper struct {
	e block.Entry
}

func (w *ceWrapper) ID() string { _ = "STUB: not implemented"; return "" }

func (w *ceWrapper) Source() string { _ = "STUB: not implemented"; return "" }

func (w *ceWrapper) SpecVersion() string { _ = "STUB: not implemented"; return "" }

func (w *ceWrapper) Type() string { _ = "STUB: not implemented"; return "" }

func (w *ceWrapper) DataContentType() string { _ = "STUB: not implemented"; return "" }

func (w *ceWrapper) DataSchema() string { _ = "STUB: not implemented"; return "" }

func (w *ceWrapper) Subject() string { _ = "STUB: not implemented"; return "" }

func (w *ceWrapper) Time() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (w *ceWrapper) Extension(ext []byte) []byte { _ = "STUB: not implemented"; return nil }

func (w *ceWrapper) Data() []byte { _ = "STUB: not implemented"; return nil }

func ToPb(e block.Entry) *cepb.CloudEvent { _ = "STUB: not implemented"; return nil }

// Overwrite XVanusBlockOffset and XVanusStime if exists.

func toPbValue(val block.Value) *cepb.CloudEvent_CloudEventAttributeValue {
	_ = "STUB: not implemented"
	return nil
}

// unreachable

// attrTypeNone
