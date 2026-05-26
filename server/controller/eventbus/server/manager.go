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

//go:generate mockgen -source=manager.go -destination=mock_manager.go -package=server
package server

import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	segpb "github.com/vanus-labs/vanus/api/segment"
)

type Manager interface {
	AddServer(ctx context.Context, srv Server) error
	RemoveServer(ctx context.Context, srv Server) error
	GetServerByAddress(addr string) Server
	GetServerByVolumeID(id uint64) Server
	Run(ctx context.Context) error
	Stop(ctx context.Context)
	CanCreateEventbus(ctx context.Context, replicaNum int) bool
}

const (
	serverStateRunning = "running"
)

func NewServerManager() Manager { _ = "STUB: not implemented"; return *new(Manager) }

type segmentServerManager struct {
	segmentServerCredentials credentials.TransportCredentials
	mutex                    sync.Mutex
	// map[string]Server
	segmentServerMapByIP sync.Map
	// map[string]Server
	segmentServerMapByVolumeID sync.Map
	cancelCtx                  context.Context
	cancel                     func()
	ticker                     *time.Ticker
	onlineServerNumber         int64
}

func (mgr *segmentServerManager) AddServer(ctx context.Context, srv Server) error {
	_ = "STUB: not implemented"
	return nil
}

func (mgr *segmentServerManager) RemoveServer(ctx context.Context, srv Server) error {
	_ = "STUB: not implemented"
	return nil
}

func (mgr *segmentServerManager) GetServerByAddress(addr string) Server {
	_ = "STUB: not implemented"
	return *new(Server)
}

func (mgr *segmentServerManager) GetServerByVolumeID(id uint64) Server {
	_ = "STUB: not implemented"
	return *new(Server)
}

func (mgr *segmentServerManager) Run(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (mgr *segmentServerManager) Stop(ctx context.Context) { _ = "STUB: not implemented"; return }

func (mgr *segmentServerManager) CanCreateEventbus(ctx context.Context, replicaNum int) bool {
	_ = "STUB: not implemented"
	return false
}

type Server interface {
	RemoteStart(ctx context.Context) error
	RemoteStop(ctx context.Context)
	GetClient() segpb.SegmentServerClient
	VolumeID() uint64
	Address() string
	Close() error
	Polish()
	IsActive(ctx context.Context) bool
	Uptime() time.Time
}

type segmentServer struct {
	volumeID          uint64
	addr              string
	grpcConn          *grpc.ClientConn
	client            segpb.SegmentServerClient
	uptime            time.Time
	lastHeartbeatTime time.Time
}

type Getter func(volumeID uint64, addr string) (Server, error)

var (
	getter Getter = newSegmentServer
	mutex  sync.Mutex
)

func NewSegmentServerWithVolumeID(volumeID uint64, addr string) (Server, error) {
	_ = "STUB: not implemented"
	return *new(Server), nil
}

func MockServerGetter(gt Getter) { _ = "STUB: not implemented"; return }

func MockReset() { _ = "STUB: not implemented"; return }

func NewSegmentServer(volumeID uint64, addr string) (Server, error) {
	_ = "STUB: not implemented"
	return *new(Server), nil
}

func newSegmentServer(volumeID uint64, addr string) (Server, error) {
	_ = "STUB: not implemented"
	return *new(Server), nil
}

func (ss *segmentServer) RemoteStart(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (ss *segmentServer) RemoteStop(ctx context.Context) { _ = "STUB: not implemented"; return }

func (ss *segmentServer) GetClient() segpb.SegmentServerClient {
	_ = "STUB: not implemented"
	return *new(segpb.SegmentServerClient)
}

func (ss *segmentServer) VolumeID() uint64 { _ = "STUB: not implemented"; return 0 }

func (ss *segmentServer) Address() string { _ = "STUB: not implemented"; return "" }

func (ss *segmentServer) Close() error { _ = "STUB: not implemented"; return nil }

func (ss *segmentServer) Polish() { _ = "STUB: not implemented"; return }

func (ss *segmentServer) IsActive(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// maximum heartbeat interval is 1 minute
// return time.Now().Sub(ss.lastHeartbeatTime) > time.Minute
// TODO optimize here

func (ss *segmentServer) Uptime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
