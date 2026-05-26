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

package cel

import (
	"fmt"

	ce "github.com/cloudevents/sdk-go/v2"
	"github.com/google/cel-go/cel"
)

var ErrInvalidExpression = fmt.Errorf("expression is invalid,format is: $json_path.(type)")

type Expression struct {
	program   cel.Program
	variables map[string]Variable
}

type Variable struct {
	Name string
	Path string
	Type string
}

func Parse(expression string) (*Expression, error) { _ = "STUB: not implemented"; return nil, nil }

// parseExpressionString breaks inline expression string into Google Expression expression
// and a set of variable definitions, e.g.:
// '$foo.(string) == "bar"' becomes
// expr: foo == "bar", vars: ["foo": string]
func parseExpression(expression string) (string, map[string]Variable, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// integer as the pos name first symbol causes issue with matching
// var types. String prefix ensures that we don't have first integer symbol.

func newCelProgram(expr string, vars map[string]Variable) (cel.Program, error) {
	_ = "STUB: not implemented"
	return *new(cel.Program), nil
}

func (e *Expression) Eval(event ce.Event) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func stringValue(value interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

func intValue(value interface{}) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func uintValue(value interface{}) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func boolValue(value interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func floatValue(value interface{}) (float64, error) { _ = "STUB: not implemented"; return 0, nil }
