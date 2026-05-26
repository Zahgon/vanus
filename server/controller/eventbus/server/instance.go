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

//go:generate mockgen -source=instance.go -destination=mock_instance.go -package=server
package server

import (
	"context"
	"sync"

	vanus "github.com/vanus-labs/vanus/api/vsr"

	"github.com/vanus-labs/vanus/server/controller/eventbus/metadata"
)

type Instance interface {
	ID() vanus.ID
	Address() string
	Close() error
	GetMeta() *metadata.VolumeMetadata
	CreateBlock(context.Context, int64) (*metadata.Block, error)
	DeleteBlock(context.Context, vanus.ID) error
	GetServer() Server
	SetServer(Server)
}

func NewInstance(md *metadata.VolumeMetadata) Instance {
	_ = "STUB: not implemented"
	return *new(Instance)
}

type volumeInstance struct {
	md        *metadata.VolumeMetadata
	metaMutex sync.Mutex
	srv       Server
	rwMutex   sync.RWMutex
}

func (ins *volumeInstance) GetMeta() *metadata.VolumeMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (ins *volumeInstance) CreateBlock(ctx context.Context, capacity int64) (*metadata.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ins *volumeInstance) DeleteBlock(ctx context.Context, id vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (ins *volumeInstance) ID() vanus.ID { _ = "STUB: not implemented"; return *new(vanus.ID) }

func (ins *volumeInstance) Address() string { _ = "STUB: not implemented"; return "" }

func (ins *volumeInstance) Close() error { _ = "STUB: not implemented"; return nil }

func (ins *volumeInstance) SetServer(srv Server) { _ = "STUB: not implemented"; return }

func (ins *volumeInstance) GetServer() Server { _ = "STUB: not implemented"; return *new(Server) }
