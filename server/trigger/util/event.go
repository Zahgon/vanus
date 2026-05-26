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

package util

import (
	"fmt"

	ce "github.com/cloudevents/sdk-go/v2"
)

// LookupAttribute lookup event attribute value by attribute name.
func LookupAttribute(event ce.Event, attr string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// LookupData lookup event data value by JSON path.
func LookupData(data interface{}, path string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	errValueIsNil           = fmt.Errorf("value is nil")
	errAttributeValue       = fmt.Errorf("attribute value is invalid")
	errDeleteAttrNotSupport = fmt.Errorf("delete attribute not support")

	specAttributes = map[string]struct{}{
		"id":          {},
		"source":      {},
		"type":        {},
		"specversion": {},
	}
)

func SetAttribute(e *ce.Event, attr string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func DeleteAttribute(e *ce.Event, attr string) error { _ = "STUB: not implemented"; return nil }

// SetData set value to data path, now data must is map, not support array.
func SetData(data interface{}, path string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// todo ,now not support

func setData(data interface{}, paths []string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// key .

// arr[2] .

// arr[2].

// todo multidimensional array

type pathType string

const (
	pathMap   pathType = "map"
	pathArray pathType = "array"
)

func getPathIndex(path string) (pathType, string, int, error) {
	_ = "STUB: not implemented"
	return *new(pathType), "", 0, nil
}

// todo map or array

func DeleteData(data interface{}, path string) error { _ = "STUB: not implemented"; return nil }

// todo ,now not support

func deleteData(data interface{}, paths []string) { _ = "STUB: not implemented"; return }

// todo ,now not support
