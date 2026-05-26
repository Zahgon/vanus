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

package transport

import (
	// standard libraries.
	"sync"
)

type Resolver interface {
	Resolve(node uint64) string
}

type SimpleResolver struct {
	sync.RWMutex
	nodes map[uint64]string
}

func NewSimpleResolver() *SimpleResolver { _ = "STUB: not implemented"; return nil }

func (r *SimpleResolver) Resolve(node uint64) string { _ = "STUB: not implemented"; return "" }

func (r *SimpleResolver) Register(node uint64, endpoint string) { _ = "STUB: not implemented"; return }

func (r *SimpleResolver) Unregister(node uint64) { _ = "STUB: not implemented"; return }
