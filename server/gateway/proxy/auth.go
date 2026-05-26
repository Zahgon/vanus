// Copyright 2023 Linkall Inc.
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

package proxy

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	metapb "github.com/vanus-labs/vanus/api/meta"
	vanus "github.com/vanus-labs/vanus/api/vsr"

	"github.com/vanus-labs/vanus/pkg/authorization"
)

func authCreateUser(_ context.Context, _ interface{}) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) CreateUser(ctx context.Context, request *ctrlpb.CreateUserRequest) (*metapb.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authDeleteUser(_ context.Context, _ interface{}) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) DeleteUser(ctx context.Context, value *wrapperspb.StringValue) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cp *ControllerProxy) GetUser(ctx context.Context, value *wrapperspb.StringValue) (*metapb.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cp *ControllerProxy) ListUser(ctx context.Context, empty *emptypb.Empty) (*ctrlpb.ListUserResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cp *ControllerProxy) checkUserIdentifier(ctx context.Context, identifier string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// check role

func (cp *ControllerProxy) GetUserToken(ctx context.Context, request *wrapperspb.StringValue,
) (*ctrlpb.GetTokenResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authListToken(_ context.Context, _ interface{}) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) ListToken(ctx context.Context, empty *emptypb.Empty) (*ctrlpb.ListTokenResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cp *ControllerProxy) CreateToken(ctx context.Context, request *ctrlpb.CreateTokenRequest,
) (*metapb.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cp *ControllerProxy) DeleteToken(ctx context.Context, request *ctrlpb.DeleteTokenRequest,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authGrantRole(_ context.Context, req interface{}) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) GrantRole(ctx context.Context, request *ctrlpb.RoleRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authRevokeRole(_ context.Context, req interface{}) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) RevokeRole(ctx context.Context, request *ctrlpb.RoleRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cp *ControllerProxy) GetUserRole(ctx context.Context, request *ctrlpb.GetUserRoleRequest,
) (*ctrlpb.GetUserRoleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authGetResourceRole(_ context.Context, req interface{},
) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

// resource kind invalid

func (cp *ControllerProxy) GetResourceRole(ctx context.Context,
	request *ctrlpb.GetResourceRoleRequest,
) (*ctrlpb.GetResourceRoleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
