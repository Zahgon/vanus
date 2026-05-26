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
	// third-party libraries.
	. "github.com/smartystreets/goconvey/convey"
	. "go.uber.org/mock/gomock"

	// this project.

	"github.com/vanus-labs/vanus/server/store/vsb/index"
)

func MakeIndex0(ctrl *Controller) index.Index { _ = "STUB: not implemented"; return *new(index.Index) }

func MakeIndex1(ctrl *Controller) index.Index { _ = "STUB: not implemented"; return *new(index.Index) }

func CheckIndex0(i index.Index, ignoreStime bool) { _ = "STUB: not implemented"; return }

func CheckIndex1(i index.Index, ignoreStime bool) { _ = "STUB: not implemented"; return }
