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

package filter

import (
	ce "github.com/cloudevents/sdk-go/v2"
)

type commonFilter struct {
	attribute     map[string]string
	dataValue     string
	data          map[string]string
	meetCondition meetCondition
}

type meetCondition func(value, compareValue string) bool

func newCommonFilter(value map[string]string, meetCondition meetCondition) *commonFilter {
	_ = "STUB: not implemented"
	return nil
}

// event attribute.

func (filter *commonFilter) Filter(event ce.Event) Result {
	_ = "STUB: not implemented"
	return *new(Result)
}

func dataValue2String(value interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

func attrValue2String(value interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }
