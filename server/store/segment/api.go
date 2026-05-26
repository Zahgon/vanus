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

package segment

import (
	// standard libraries.
	"context"

	// third-party libraries.
	"google.golang.org/protobuf/types/known/emptypb"

	// first-party libraries.

	segpb "github.com/vanus-labs/vanus/api/segment"
	// this project.
)

type segmentServer struct {
	srv Server
}

// Make sure segmentServer implements segpb.SegmentServerServer.
var _ segpb.SegmentServerServer = (*segmentServer)(nil)

func (s *segmentServer) Start(
	ctx context.Context, _ *segpb.StartSegmentServerRequest,
) (*segpb.StartSegmentServerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *segmentServer) Stop(
	ctx context.Context, _ *segpb.StopSegmentServerRequest,
) (*segpb.StopSegmentServerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *segmentServer) Status(_ context.Context, _ *emptypb.Empty) (*segpb.StatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *segmentServer) CreateBlock(ctx context.Context, req *segpb.CreateBlockRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *segmentServer) RemoveBlock(ctx context.Context, req *segpb.RemoveBlockRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *segmentServer) DescribeBlock(
	ctx context.Context, req *segpb.DescribeBlockRequest,
) (*segpb.DescribeBlockResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *segmentServer) ActivateSegment(
	ctx context.Context, req *segpb.ActivateSegmentRequest,
) (*segpb.ActivateSegmentResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *segmentServer) InactivateSegment(
	_ context.Context, _ *segpb.InactivateSegmentRequest,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *segmentServer) AppendToBlock(
	ctx context.Context, req *segpb.AppendToBlockRequest,
) (*segpb.AppendToBlockResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *segmentServer) ReadFromBlock(
	ctx context.Context, req *segpb.ReadFromBlockRequest,
) (*segpb.ReadFromBlockResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *segmentServer) LookupOffsetInBlock(
	ctx context.Context, req *segpb.LookupOffsetInBlockRequest,
) (*segpb.LookupOffsetInBlockResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
