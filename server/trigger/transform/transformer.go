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

package transform

import (
	// standard libraries.

	// third-party libraries.
	ce "github.com/cloudevents/sdk-go/v2"

	// first-party project.
	"github.com/vanus-labs/vanus/pkg/template"

	// this project.
	primitive "github.com/vanus-labs/vanus/pkg"
	"github.com/vanus-labs/vanus/server/trigger/transform/define"
	"github.com/vanus-labs/vanus/server/trigger/transform/pipeline"
)

type Transformer struct {
	define   *define.Define
	pipeline *pipeline.Pipeline
	template template.Template
}

func NewTransformer(transformer *primitive.Transformer) (*Transformer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:nilnil // nil is valid

func (tf *Transformer) Execute(event *ce.Event) (err error) { _ = "STUB: not implemented"; return nil }

func buildTemplateModel(event *ce.Event, data any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func CompileTemplate(tc primitive.TemplateConfig) (t template.Template, err error) {
	_ = "STUB: not implemented"
	return *new(template.Template), nil
}
