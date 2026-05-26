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

package raw_client

import (
	// standard libraries.
	"context"
	"io"

	// third-party libraries.
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	// first-party libraries.
	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	// this project.
)

var (
	_ io.Closer = (*segmentClient)(nil)
	_ Heartbeat = (*segmentClient)(nil)
)

type segmentClient struct {
	cc              *Conn
	heartBeatClient ctrlpb.SegmentController_SegmentHeartbeatClient
}

func NewSegmentClient(cc *Conn) ctrlpb.SegmentControllerClient {
	_ = "STUB: not implemented"
	return *new(ctrlpb.SegmentControllerClient)
}

func (sc *segmentClient) Beat(ctx context.Context, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO panic

func (sc *segmentClient) Close() error { _ = "STUB: not implemented"; return nil }

func (sc *segmentClient) RegisterSegmentServer(
	ctx context.Context, in *ctrlpb.RegisterSegmentServerRequest, opts ...grpc.CallOption,
) (*ctrlpb.RegisterSegmentServerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sc *segmentClient) UnregisterSegmentServer(
	ctx context.Context, in *ctrlpb.UnregisterSegmentServerRequest, opts ...grpc.CallOption,
) (*ctrlpb.UnregisterSegmentServerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sc *segmentClient) ReportSegmentBlockIsFull(
	ctx context.Context, in *ctrlpb.SegmentHeartbeatRequest, opts ...grpc.CallOption,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sc *segmentClient) ReportSegmentLeader(
	ctx context.Context, in *ctrlpb.ReportSegmentLeaderRequest, opts ...grpc.CallOption,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sc *segmentClient) QuerySegmentRouteInfo(
	ctx context.Context, in *ctrlpb.QuerySegmentRouteInfoRequest, opts ...grpc.CallOption,
) (*ctrlpb.QuerySegmentRouteInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sc *segmentClient) SegmentHeartbeat(_ context.Context, _ ...grpc.CallOption) (ctrlpb.SegmentController_SegmentHeartbeatClient, error) {
	_ = "STUB: not implemented"
	// TODO implement me
	return *new(ctrlpb.SegmentController_SegmentHeartbeatClient), nil
}
