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

package authorization

import (
	vanus "github.com/vanus-labs/vanus/api/vsr"
)

type Attributes interface {
	GetResourceKind() ResourceKind
	GetResourceID() vanus.ID
	GetAction() Action
}

var _ Attributes = &defaultAttributes{}

type defaultAttributes struct {
	resourceKind ResourceKind
	resourceID   vanus.ID
	action       Action
}

func NewDefaultAttributes(resourceKind ResourceKind, resourceID vanus.ID, action Action) Attributes {
	_ = "STUB: not implemented"
	return *new(Attributes)
}

func (d *defaultAttributes) GetResourceKind() ResourceKind {
	_ = "STUB: not implemented"
	return *new(ResourceKind)
}

func (d *defaultAttributes) GetResourceID() vanus.ID {
	_ = "STUB: not implemented"
	return *new(vanus.ID)
}

func (d *defaultAttributes) GetAction() Action { _ = "STUB: not implemented"; return *new(Action) }
