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

//go:generate mockgen -source=server.go -destination=mock_server.go -package=segment
package segment

import (
	// standard libraries.
	"context"
	"net"
	"sync"
	"time"

	// third-party libraries.

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/tap"

	// first-party libraries.
	cepb "github.com/vanus-labs/vanus/api/cloudevents"
	"github.com/vanus-labs/vanus/api/cluster"
	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	metapb "github.com/vanus-labs/vanus/api/meta"
	vanus "github.com/vanus-labs/vanus/api/vsr"
	"github.com/vanus-labs/vanus/pkg/observability/tracing"

	// this project.
	primitive "github.com/vanus-labs/vanus/pkg"
	"github.com/vanus-labs/vanus/server/store/block"
	"github.com/vanus-labs/vanus/server/store/block/raw"
	raft "github.com/vanus-labs/vanus/server/store/raft/block"
)

const (
	defaultLeaderInfoBufferSize = 256
	defaultForceStopTimeout     = 30 * time.Second
)

type Server interface {
	primitive.Initializer

	Serve(lis net.Listener) error
	RegisterToController(ctx context.Context) error
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Status() primitive.ServerState

	CreateBlock(ctx context.Context, id vanus.ID, size int64) error
	RemoveBlock(ctx context.Context, id vanus.ID) error
	DescribeBlock(ctx context.Context, id vanus.ID) (*metapb.SegmentHealthInfo, error)

	ActivateSegment(ctx context.Context, logID vanus.ID, segID vanus.ID, replicas map[vanus.ID]string) error
	InactivateSegment(ctx context.Context) error

	AppendToBlock(ctx context.Context, id vanus.ID, events []*cepb.CloudEvent) ([]int64, error)
	ReadFromBlock(ctx context.Context, id vanus.ID, seq int64, num int, pollingTimeout uint32) ([]*cepb.CloudEvent, error)
	LookupOffsetInBlock(ctx context.Context, id vanus.ID, stime int64) (int64, error)
}

func NewServer(cfg Config, debug bool) (Server, error) {
	_ = "STUB: not implemented"
	return *new(Server), nil
}

// TODO(james.yin): support IPv6

// TODO(james.yin): move to close function

type leaderInfo struct {
	leader vanus.ID
	term   uint64
}

type appendResult struct {
	seqs []int64
	err  error
}

type appendFuture chan appendResult

func newAppendFuture() appendFuture { _ = "STUB: not implemented"; return *new(appendFuture) }

func (af appendFuture) onAppended(seqs []int64, err error) { _ = "STUB: not implemented"; return }

func (af appendFuture) wait() ([]int64, error) { _ = "STUB: not implemented"; return nil, nil }

type server struct {
	replicas sync.Map // <vanus.ID, Replica>

	raftEngine raft.Engine
	rawEngines *raw.EngineRegistry

	state       primitive.ServerState
	isDebugMode bool
	cfg         Config
	localAddr   string

	volumeID    uint64
	volumeIDStr string
	volumeDir   string

	ctrlAddr    []string
	credentials credentials.TransportCredentials
	ctrl        cluster.Cluster
	cc          ctrlpb.SegmentControllerClient
	leaderC     chan leaderInfo

	grpcSrv *grpc.Server
	closeC  chan struct{}

	pm     pollingManager
	tracer *tracing.Tracer
}

// Make sure server implements Server.
var _ Server = (*server)(nil)

func (s *server) Serve(lis net.Listener) error { _ = "STUB: not implemented"; return nil }

func (s *server) preGrpcStream(ctx context.Context, info *tap.Info) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (s *server) RegisterToController(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil

	// Register to controller.
}

func (s *server) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *server) startHeartbeatTask(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *server) runHeartbeat(_ context.Context) error { _ = "STUB: not implemented"; return nil }

// TODO(james.yin): move to other goroutine.

func (s *server) onLeaderChanged(blockID, leaderID vanus.ID, term uint64) {
	_ = "STUB: not implemented"
	return
}

func (s *server) Stop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// TODO(james.yin): async

// Stop grpc asynchronously.

// Force stop if timeout.

func (s *server) stop(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Close all blocks.
	return nil
}

// Stop heartbeat task, etc.

// FIXME(james.yin): reorder

func (s *server) Status() primitive.ServerState {
	_ = "STUB: not implemented"
	return *new(primitive.ServerState)
}

func (s *server) CreateBlock(ctx context.Context, id vanus.ID, size int64) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(james.yin): release resources of block.

// TODO(james.yin): open replica.

func (s *server) RemoveBlock(ctx context.Context, blockID vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(james.yin): s.host.Unregister

// FIXME(james.yin): more info.

func (s *server) DescribeBlock(_ context.Context, id vanus.ID) (*metapb.SegmentHealthInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ActivateSegment mark a block ready to using and preparing to initializing a replica group.
func (s *server) ActivateSegment(
	ctx context.Context, logID vanus.ID, segID vanus.ID, replicas map[vanus.ID]string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Register peers.

// Bootstrap raft.

// InactivateSegment mark a block ready to be removed. This method is usually used for data transfer.
func (s *server) InactivateSegment(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *server) AppendToBlock(ctx context.Context, id vanus.ID, events []*cepb.CloudEvent) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *server) processAppendError(ctx context.Context, b Replica, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *server) onEntryAppended(block vanus.ID) { _ = "STUB: not implemented"; return }

func (s *server) onBlockArchived(stat block.Statistics) { _ = "STUB: not implemented"; return }

// FIXME(james.yin): leader info.

// report to controller

// ReadFromBlock returns at most num events from seq in Block id.
func (s *server) ReadFromBlock(
	ctx context.Context, id vanus.ID, seq int64, num int, pollingTimeout uint32,
) ([]*cepb.CloudEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FIXME(james.yin) It can't read message immediately because of async apply.

func (s *server) readEvents(ctx context.Context, b Replica, seq int64, num int) ([]*cepb.CloudEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *server) processReadError(ctx context.Context, b Replica, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *server) LookupOffsetInBlock(ctx context.Context, id vanus.ID, stime int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *server) checkState() error { _ = "STUB: not implemented"; return nil }
