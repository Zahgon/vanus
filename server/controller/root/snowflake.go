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

package root

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	"github.com/vanus-labs/vanus/pkg/kv"
	"github.com/vanus-labs/vanus/server/controller/member"
)

const (
	spinInterval = 100 * time.Millisecond
)

type Config struct {
	KVEndpoints []string
	KVPrefix    string
}

func NewSnowflakeController(cfg Config, mem member.Member) *snowflake {
	_ = "STUB: not implemented" //nolint:revive // it's ok
	return nil
}

var (
	_ ctrlpb.PingServerServer = &snowflake{}
)

type snowflake struct {
	startAt  time.Time
	cfg      Config
	kvStore  kv.Client
	isLeader bool
	member   member.Member
	nodes    map[uint16]*node
	mutex    sync.RWMutex
	r        *rand.Rand
}

func (sf *snowflake) Ping(_ context.Context, _ *emptypb.Empty) (*ctrlpb.PingResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type node struct {
	StartAt time.Time
	ID      uint16
}

func (sf *snowflake) Start(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (sf *snowflake) GetClusterStartTime(_ context.Context, _ *emptypb.Empty) (*timestamppb.Timestamp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sf *snowflake) RegisterNode(ctx context.Context, in *wrapperspb.UInt32Value) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(wenfeng) find a good solution in future
// _, exist := sf.nodes[id]
//
// if exist {
//	return nil, errors.New("node has been register")
// }

func (sf *snowflake) UnregisterNode(ctx context.Context, in *wrapperspb.UInt32Value) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sf *snowflake) Stop() { _ = "STUB: not implemented"; return }

func (sf *snowflake) membershipChangedProcessor(ctx context.Context, event member.MembershipChangedEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func GetNodeIDKey(nodeID uint16) string { _ = "STUB: not implemented"; return "" }
