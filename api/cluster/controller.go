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

//go:generate mockgen -source=controller.go -destination=mock_controller.go -package=cluster
package cluster

import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc/credentials"

	"github.com/vanus-labs/vanus/api/cluster/raw_client"
	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	metapb "github.com/vanus-labs/vanus/api/meta"
)

var defaultClusterStartTimeout = 3 * time.Minute

type Topology struct {
	ControllerLeader string
	ControllerURLs   []string
	Uptime           time.Time
}

type Cluster interface {
	WaitForControllerReady(createEventbus bool) error
	Status() Topology
	IsReady(createEventbus bool) bool
	NamespaceService() NamespaceService
	EventbusService() EventbusService
	SegmentService() SegmentService
	EventlogService() EventlogService
	TriggerService() TriggerService
	IDService() IDService
	AuthService() AuthService
}

type NamespaceService interface {
	RawClient() ctrlpb.NamespaceControllerClient
	GetSystemNamespace(ctx context.Context) (*metapb.Namespace, error)
	GetDefaultNamespace(ctx context.Context) (*metapb.Namespace, error)
	GetNamespace(ctx context.Context, id uint64) (*metapb.Namespace, error)
	GetNamespaceByName(ctx context.Context, name string) (*metapb.Namespace, error)
}

type EventbusService interface {
	CreateSystemEventbusIfNotExist(ctx context.Context, name string, desc string) (*metapb.Eventbus, error)
	Delete(ctx context.Context, id uint64) error
	GetSystemEventbusByName(ctx context.Context, name string) (*metapb.Eventbus, error)
	GetEventbus(ctx context.Context, id uint64) (*metapb.Eventbus, error)
	GetEventbusByName(ctx context.Context, ns, name string) (*metapb.Eventbus, error)
	RawClient() ctrlpb.EventbusControllerClient
}

type EventlogService interface {
	RawClient() ctrlpb.EventlogControllerClient
}

type TriggerService interface {
	RawClient() ctrlpb.TriggerControllerClient
	GetSubscription(ctx context.Context, id uint64) (*metapb.Subscription, error)
	RegisterHeartbeat(ctx context.Context, interval time.Duration, reqFunc func() interface{}) error
}

type IDService interface {
	RawClient() ctrlpb.SnowflakeControllerClient
}

type SegmentService interface {
	RegisterHeartbeat(ctx context.Context, interval time.Duration, reqFunc func() interface{}) error
	RawClient() ctrlpb.SegmentControllerClient
}

type AuthService interface {
	RawClient() ctrlpb.AuthControllerClient
	GetUserByToken(ctx context.Context, token string) (string, error)
	GetUserRole(ctx context.Context, user string) ([]*metapb.UserRole, error)
}

var (
	connCache = map[string]*raw_client.Conn{}
	mutex     sync.Mutex
)

func NewClusterController(endpoints []string, credentials credentials.TransportCredentials) Cluster {
	_ = "STUB: not implemented"
	return *new(Cluster)
}

// single instance

type cluster struct {
	controllerAddress []string
	cc                *raw_client.Conn
	nsSvc             NamespaceService
	ebSvc             EventbusService
	elSvc             EventlogService
	triggerSvc        TriggerService
	idSvc             IDService
	segmentSvc        SegmentService
	ping              ctrlpb.PingServerClient
	authSvc           AuthService
}

func (c *cluster) WaitForControllerReady(createEventbus bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *cluster) IsReady(createEventbus bool) bool { _ = "STUB: not implemented"; return false }

func (c *cluster) Status() Topology {
	_ = "STUB: not implemented"
	// TODO(wenfeng)
	return *new(Topology)
}

func (c *cluster) NamespaceService() NamespaceService {
	_ = "STUB: not implemented"
	return *new(NamespaceService)
}

func (c *cluster) EventbusService() EventbusService {
	_ = "STUB: not implemented"
	return *new(EventbusService)
}

func (c *cluster) SegmentService() SegmentService {
	_ = "STUB: not implemented"
	return *new(SegmentService)
}

func (c *cluster) EventlogService() EventlogService {
	_ = "STUB: not implemented"
	return *new(EventlogService)
}

func (c *cluster) TriggerService() TriggerService {
	_ = "STUB: not implemented"
	return *new(TriggerService)
}

func (c *cluster) IDService() IDService { _ = "STUB: not implemented"; return *new(IDService) }

func (c *cluster) AuthService() AuthService { _ = "STUB: not implemented"; return *new(AuthService) }
