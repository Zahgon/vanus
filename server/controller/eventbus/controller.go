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

package eventbus

import (
	// standard libraries.
	"context"
	"sync"

	// third-party libraries.

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	// first-party libraries.
	"github.com/vanus-labs/vanus/api/cluster"
	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	metapb "github.com/vanus-labs/vanus/api/meta"
	vanus "github.com/vanus-labs/vanus/api/vsr"

	// this project.

	"github.com/vanus-labs/vanus/pkg/kv"
	"github.com/vanus-labs/vanus/server/controller/eventbus/eventlog"
	"github.com/vanus-labs/vanus/server/controller/eventbus/metadata"
	"github.com/vanus-labs/vanus/server/controller/eventbus/server"
	"github.com/vanus-labs/vanus/server/controller/eventbus/volume"
	"github.com/vanus-labs/vanus/server/controller/member"
)

var (
	_ ctrlpb.EventbusControllerServer = &controller{}
	_ ctrlpb.EventlogControllerServer = &controller{}
	_ ctrlpb.SegmentControllerServer  = &controller{}
	_ ctrlpb.PingServerServer         = &controller{}
)

const (
	maximumEventlogNum = 64
	mappingKey         = "@%d_%s@" // @{namespace_id}_{eventbus}@
)

func NewController(cfg Config, mem member.Member) *controller {
	_ = "STUB: not implemented"
	return nil
}

type controller struct {
	cfg                      *Config
	kvStore                  kv.Client
	volumeMgr                volume.Manager
	eventlogMgr              eventlog.Manager
	ssMgr                    server.Manager
	eventbusMap              map[vanus.ID]*metadata.Eventbus
	eventbusNamespaceMapping sync.Map // string, *metadata.Eventbus
	member                   member.Member
	cancelCtx                context.Context
	cancelFunc               context.CancelFunc
	membershipMutex          sync.Mutex
	isLeader                 bool
	readyNotify              chan error
	stopNotify               chan error
	mutex                    sync.Mutex
	eventbusUpdatedCount     int64
	eventbusDeletedCount     int64
	clusterCli               cluster.Cluster
}

func (ctrl *controller) Start(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (ctrl *controller) Stop() { _ = "STUB: not implemented"; return }

func (ctrl *controller) ReadyNotify() <-chan error { _ = "STUB: not implemented"; return nil }

func (ctrl *controller) StopNotify() <-chan error { _ = "STUB: not implemented"; return nil }

func (ctrl *controller) CreateEventbus(
	ctx context.Context, req *ctrlpb.CreateEventbusRequest,
) (*metapb.Eventbus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO async create
// create dead letter eventbus

func isValidEventbusName(name string) error { _ = "STUB: not implemented"; return nil }

func (ctrl *controller) GetEventbusWithHumanFriendly(_ context.Context,
	request *ctrlpb.GetEventbusWithHumanFriendlyRequest,
) (*metapb.Eventbus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetMappingKey(namespace uint64, name string) string { _ = "STUB: not implemented"; return "" }

func (ctrl *controller) CreateSystemEventbus(
	ctx context.Context, req *ctrlpb.CreateEventbusRequest,
) (*metapb.Eventbus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) isEventbusExist(ctx context.Context, req *ctrlpb.CreateEventbusRequest) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// check id exist

// check name exist

func (ctrl *controller) createEventbus(
	ctx context.Context, req *ctrlpb.CreateEventbusRequest,
) (*metapb.Eventbus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) getDeadLetterEventbusID(_ context.Context, id vanus.ID) vanus.ID {
	_ = "STUB: not implemented"
	return *new(vanus.ID)
}

func (ctrl *controller) DeleteEventbus(ctx context.Context, eb *wrapperspb.UInt64Value) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO async delete
// delete dead letter eventbus

func (ctrl *controller) deleteEventbus(ctx context.Context, id vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// todo user can't delete system eventbus, but timer need to delete system eventbus

// TODO(wenfeng.wang) notify gateway to cut flow

func (ctrl *controller) GetEventbus(_ context.Context, eb *wrapperspb.UInt64Value) (*metapb.Eventbus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) getEventbus(id vanus.ID) (*metapb.Eventbus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) ListEventbus(_ context.Context,
	req *ctrlpb.ListEventbusRequest,
) (*ctrlpb.ListEventbusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) UpdateEventbus(
	_ context.Context, _ *ctrlpb.UpdateEventbusRequest,
) (*metapb.Eventbus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) ListSegment(
	ctx context.Context, req *ctrlpb.ListSegmentRequest,
) (*ctrlpb.ListSegmentResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) RegisterSegmentServer(
	ctx context.Context, req *ctrlpb.RegisterSegmentServerRequest,
) (*ctrlpb.RegisterSegmentServerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// need to compare metadata if existed?

func (ctrl *controller) UnregisterSegmentServer(ctx context.Context,
	req *ctrlpb.UnregisterSegmentServerRequest,
) (*ctrlpb.UnregisterSegmentServerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) QuerySegmentRouteInfo(_ context.Context,
	_ *ctrlpb.QuerySegmentRouteInfoRequest,
) (*ctrlpb.QuerySegmentRouteInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) SegmentHeartbeat(srv ctrlpb.SegmentController_SegmentHeartbeatServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (ctrl *controller) processHeartbeat(ctx context.Context, req *ctrlpb.SegmentHeartbeatRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO

func (ctrl *controller) GetAppendableSegment(
	ctx context.Context, req *ctrlpb.GetAppendableSegmentRequest,
) (*ctrlpb.GetAppendableSegmentResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) ReportSegmentBlockIsFull(
	ctx context.Context, req *ctrlpb.SegmentHeartbeatRequest,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) Ping(ctx context.Context, _ *emptypb.Empty) (*ctrlpb.PingResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) isReady(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func (ctrl *controller) ReportSegmentLeader(
	ctx context.Context, req *ctrlpb.ReportSegmentLeaderRequest,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) recordMetrics() { _ = "STUB: not implemented"; return }

func (ctrl *controller) membershipChangedProcessor(ctx context.Context, event member.MembershipChangedEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (ctrl *controller) loadEventbus(ctx context.Context) error {
	_ = "STUB: not implemented"
	// load eventbus metadata
	return nil
}

func (ctrl *controller) stop(ctx context.Context, err error) { _ = "STUB: not implemented"; return }
