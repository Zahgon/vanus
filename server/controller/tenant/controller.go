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

package tenant

import (
	"context"
	"sync"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/vanus-labs/vanus/api/cluster"
	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	metapb "github.com/vanus-labs/vanus/api/meta"

	"github.com/vanus-labs/vanus/pkg/kv"
	"github.com/vanus-labs/vanus/server/controller/member"
	"github.com/vanus-labs/vanus/server/controller/tenant/manager"
	"github.com/vanus-labs/vanus/server/controller/tenant/metadata"
)

var (
	_               ctrlpb.NamespaceControllerServer = &controller{}
	_               ctrlpb.AuthControllerServer      = &controller{}
	tokenRandLength                                  = 32
)

func NewController(config Config, mem member.Member) *controller {
	_ = "STUB: not implemented"
	return nil
}

type controller struct {
	config           Config
	member           member.Member
	membershipMutex  sync.Mutex
	isLeader         bool
	kvClient         kv.Client
	namespaceManager manager.NamespaceManager
	userManager      manager.UserManager
	tokenManager     manager.TokenManager
	userRoleManager  manager.UserRoleManager
	cluster          cluster.Cluster
}

func (ctrl *controller) CreateUser(ctx context.Context, request *ctrlpb.CreateUserRequest) (*metapb.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) DeleteUser(ctx context.Context, value *wrapperspb.StringValue) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// todo delete token and role

func (ctrl *controller) ListUser(ctx context.Context, _ *emptypb.Empty) (*ctrlpb.ListUserResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) GetUser(ctx context.Context, value *wrapperspb.StringValue) (*metapb.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) GetUserByToken(ctx context.Context,
	token *wrapperspb.StringValue,
) (*wrapperspb.StringValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) CreateToken(ctx context.Context,
	request *ctrlpb.CreateTokenRequest,
) (*metapb.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) DeleteToken(ctx context.Context,
	request *ctrlpb.DeleteTokenRequest,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) GetToken(ctx context.Context, request *wrapperspb.UInt64Value) (*metapb.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) GetUserToken(ctx context.Context,
	request *wrapperspb.StringValue,
) (*ctrlpb.GetTokenResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) ListToken(ctx context.Context, _ *emptypb.Empty,
) (*ctrlpb.ListTokenResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) GrantRole(ctx context.Context, request *ctrlpb.RoleRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// todo support custom define role

func (ctrl *controller) RevokeRole(ctx context.Context, request *ctrlpb.RoleRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) GetUserRole(ctx context.Context,
	request *ctrlpb.GetUserRoleRequest,
) (*ctrlpb.GetUserRoleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) GetResourceRole(ctx context.Context,
	request *ctrlpb.GetResourceRoleRequest,
) (*ctrlpb.GetResourceRoleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// todo check resourceID exist

func (ctrl *controller) CreateNamespace(ctx context.Context,
	request *ctrlpb.CreateNamespaceRequest,
) (*metapb.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) createNamespace(ctx context.Context, ns *metadata.Namespace) error {
	_ = "STUB: not implemented"
	return nil
}

func (ctrl *controller) ListNamespace(ctx context.Context,
	_ *emptypb.Empty,
) (*ctrlpb.ListNamespaceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) GetNamespace(ctx context.Context,
	request *ctrlpb.GetNamespaceRequest,
) (*metapb.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) GetNamespaceWithHumanFriendly(ctx context.Context,
	name *wrapperspb.StringValue,
) (*metapb.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) DeleteNamespace(ctx context.Context,
	request *ctrlpb.DeleteNamespaceRequest,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctrl *controller) Start() error { _ = "STUB: not implemented"; return nil }

func (ctrl *controller) init(ctx context.Context) error {
	err := ctrl.createSystemNamespace(ctx)
	if err != nil {
		return err
	}
	err = ctrl.createDefaultUserAndRole(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (ctrl *controller) createDefaultUserAndRole(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// create default user

// create default user token

// create default user role

func (ctrl *controller) createSystemNamespace(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// create default namespace

// create system namespace

func (ctrl *controller) membershipChangedProcessor(ctx context.Context,
	event member.MembershipChangedEvent,
) error {
	_ = "STUB: not implemented"
	return nil
}
