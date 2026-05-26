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

package raw

import (
	// standard libraries.
	"context"
	"fmt"

	// first-party libraries.
	vanus "github.com/vanus-labs/vanus/api/vsr"

	// this project.
	"github.com/vanus-labs/vanus/server/store/block"
)

const (
	VSB = "vsb"
)

var (
	ErrFormatRegistered = fmt.Errorf("format already registered")
	ErrNotSupported     = fmt.Errorf("not supported format")
	ErrInvalidFormat    = fmt.Errorf("invalid format")
)

type Engine interface {
	Close()

	Recover(ctx context.Context) (map[vanus.ID]block.Raw, error)

	Create(ctx context.Context, id vanus.ID, capacity int64) (block.Raw, error)
	// Open(ctx context.Context, id vanus.ID) (block.Raw, error)
}

type EngineRegistry struct {
	engines map[string]Engine
}

func NewEngineRegistry() *EngineRegistry { _ = "STUB: not implemented"; return nil }

func (er *EngineRegistry) Register(name string, engine Engine) error {
	_ = "STUB: not implemented"
	return nil
}

func (er *EngineRegistry) Resolve(engine string) (Engine, error) {
	_ = "STUB: not implemented"
	return *new(Engine), nil
}

func (er *EngineRegistry) Close() { _ = "STUB: not implemented"; return }
