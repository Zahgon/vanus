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

//go:build linux
// +build linux

package uring

import (
	// third-party libraries.
	"github.com/iceber/iouring-go"

	// first-party libraries.

	// this project.
	"github.com/vanus-labs/vanus/server/store/io"
	"github.com/vanus-labs/vanus/server/store/io/engine"
	"github.com/vanus-labs/vanus/server/store/io/zone"
)

const (
	defaultResultBufferSize = 64
)

type uRing struct {
	ring    *iouring.IOURing
	resultC chan iouring.Result
}

// Make sure uRing implements engine.Interface.
var _ engine.Interface = (*uRing)(nil)

func New() engine.Interface { _ = "STUB: not implemented"; return *new(engine.Interface) }

func (e *uRing) Close() { _ = "STUB: not implemented"; return }

func (e *uRing) runCallback() { _ = "STUB: not implemented"; return }

func (e *uRing) WriteAt(z zone.Interface, b []byte, off int64, so, eo int, cb io.WriteCallback) {
	_ = "STUB: not implemented" //nolint:revive // ok
	return
}
