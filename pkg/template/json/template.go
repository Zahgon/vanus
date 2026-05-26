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

package json

import (
	// this project.
	"github.com/vanus-labs/vanus/lib/bytes"
	"github.com/vanus-labs/vanus/pkg/template"
)

func Compile(text string) (template.Template, error) {
	_ = "STUB: not implemented"
	return *new(template.Template), nil
}

// TODO(james.yin): check segments

type jsonTemplate struct {
	segments []templateSegment
}

// Make sure jsonTemplate implements template.Template.
var _ template.Template = (*jsonTemplate)(nil)

func (t *jsonTemplate) ContentType() string { _ = "STUB: not implemented"; return "" }

func (t *jsonTemplate) Execute(model interface{}, variables map[string]interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type executeBuffer struct {
	buf []byte
}

// Make sure executeBuffer implements bytes.LastByteWriter.
var _ bytes.LastByteWriter = (*executeBuffer)(nil)

func (b *executeBuffer) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *executeBuffer) LastByte() (byte, bool) { _ = "STUB: not implemented"; return 0, false }

func (b *executeBuffer) TruncateLastByte() { _ = "STUB: not implemented"; return }
