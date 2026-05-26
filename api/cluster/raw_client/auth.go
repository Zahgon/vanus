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
	"context"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	metapb "github.com/vanus-labs/vanus/api/meta"
)

var _ io.Closer = (*authClient)(nil)

func NewAuthClient(cc *Conn) ctrlpb.AuthControllerClient {
	_ = "STUB: not implemented"
	return *new(ctrlpb.AuthControllerClient)
}

type authClient struct {
	cc *Conn
}

func (a *authClient) CreateUser(ctx context.Context, in *ctrlpb.CreateUserRequest, opts ...grpc.CallOption) (*metapb.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *authClient) DeleteUser(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *authClient) ListUser(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ctrlpb.ListUserResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *authClient) GetUser(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*metapb.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *authClient) GetUserByToken(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *authClient) ListToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ctrlpb.ListTokenResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *authClient) CreateToken(ctx context.Context, in *ctrlpb.CreateTokenRequest, opts ...grpc.CallOption) (*metapb.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *authClient) DeleteToken(ctx context.Context, in *ctrlpb.DeleteTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *authClient) GetToken(ctx context.Context, in *wrapperspb.UInt64Value, opts ...grpc.CallOption) (*metapb.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *authClient) GetUserToken(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*ctrlpb.GetTokenResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *authClient) GrantRole(ctx context.Context, in *ctrlpb.RoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *authClient) RevokeRole(ctx context.Context, in *ctrlpb.RoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *authClient) GetUserRole(ctx context.Context, in *ctrlpb.GetUserRoleRequest, opts ...grpc.CallOption) (*ctrlpb.GetUserRoleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *authClient) GetResourceRole(ctx context.Context, in *ctrlpb.GetResourceRoleRequest, opts ...grpc.CallOption) (*ctrlpb.GetResourceRoleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *authClient) Close() error { _ = "STUB: not implemented"; return nil }
