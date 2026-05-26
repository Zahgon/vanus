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

//go:generate mockgen -source=manager.go -destination=mock_manager.go -package=volume
package volume

import (
	"context"
	"sync"

	vanus "github.com/vanus-labs/vanus/api/vsr"

	"github.com/vanus-labs/vanus/pkg/kv"
	"github.com/vanus-labs/vanus/server/controller/eventbus/metadata"
	"github.com/vanus-labs/vanus/server/controller/eventbus/server"
)

type Manager interface {
	Init(ctx context.Context, kvClient kv.Client) error
	GetAllActiveVolumes() []server.Instance
	RegisterVolume(ctx context.Context, md *metadata.VolumeMetadata) (server.Instance, error)
	UpdateRouting(ctx context.Context, ins server.Instance, srv server.Server)
	GetVolumeInstanceByID(id vanus.ID) server.Instance
	LookupVolumeByID(id uint64) server.Instance
	GetBlocksOfVolume(ctx context.Context, instance server.Instance) (map[uint64]*metadata.Block, error)
}

var mgr = &volumeMgr{}

func NewVolumeManager(serverMgr server.Manager) Manager {
	_ = "STUB: not implemented"
	return *new(Manager)
}

type volumeMgr struct {
	// volumeID(vanus.ID) server.Instance
	volInstanceMap sync.Map
	// volumeID(uint64) server.Instance
	volInstanceByPhysicalVolumeID sync.Map
	kvCli                         kv.Client
	serverMgr                     server.Manager
}

func (mgr *volumeMgr) RegisterVolume(ctx context.Context, md *metadata.VolumeMetadata) (server.Instance, error) {
	_ = "STUB: not implemented"
	return *new(server.Instance), nil
}

func (mgr *volumeMgr) Init(ctx context.Context, kvClient kv.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// load block meta in this volume

func (mgr *volumeMgr) GetVolumeInstanceByID(id vanus.ID) server.Instance {
	_ = "STUB: not implemented"
	return *new(server.Instance)
}

func (mgr *volumeMgr) LookupVolumeByID(id uint64) server.Instance {
	_ = "STUB: not implemented"
	return *new(server.Instance)
}

func (mgr *volumeMgr) GetAllActiveVolumes() []server.Instance {
	_ = "STUB: not implemented"
	return nil
}

func (mgr *volumeMgr) UpdateRouting(ctx context.Context, ins server.Instance, srv server.Server) {
	_ = "STUB: not implemented"
	return
}

func (mgr *volumeMgr) GetBlocksOfVolume(ctx context.Context,
	instance server.Instance,
) (map[uint64]*metadata.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
