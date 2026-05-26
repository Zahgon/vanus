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

//go:generate mockgen -source=role.go -destination=mock_role.go -package=manager
package manager

import (
	"context"
	"sync"

	vanus "github.com/vanus-labs/vanus/api/vsr"
	"github.com/vanus-labs/vanus/pkg/kv"
	"github.com/vanus-labs/vanus/server/controller/tenant/metadata"
)

type UserRoleManager interface {
	Init(ctx context.Context) error
	AddUserRole(ctx context.Context, role *metadata.UserRole) error
	DeleteUserRole(ctx context.Context, role *metadata.UserRole) error
	GetUserRoleByUser(ctx context.Context, userIdentifier string) ([]*metadata.UserRole, error)
	GetUserRoleByResourceID(ctx context.Context, resourceID vanus.ID) ([]*metadata.UserRole, error)
	IsUserRoleExist(ctx context.Context, role *metadata.UserRole) bool
}

var _ UserRoleManager = &userRoleManager{}

type userRoleManager struct {
	lock         sync.RWMutex
	userRoles    map[string]map[string]*metadata.UserRole // key: user,roleID
	resourceRole map[vanus.ID][]*metadata.UserRole        // key: resourceID
	kvClient     kv.Client
}

func NewUserRoleManager(kvClient kv.Client) UserRoleManager {
	_ = "STUB: not implemented"
	return *new(UserRoleManager)
}

func (m *userRoleManager) Init(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// todo custom role

func (m *userRoleManager) AddUserRole(ctx context.Context, userRole *metadata.UserRole) error {
	_ = "STUB: not implemented"
	return nil
}

// todo custom role

func (m *userRoleManager) DeleteUserRole(ctx context.Context, userRole *metadata.UserRole) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *userRoleManager) IsUserRoleExist(_ context.Context, userRole *metadata.UserRole) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *userRoleManager) GetUserRoleByUser(_ context.Context, userIdentifier string) ([]*metadata.UserRole, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *userRoleManager) GetUserRoleByResourceID(_ context.Context, resourceID vanus.ID,
) ([]*metadata.UserRole, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *userRoleManager) getKVKey(userRole *metadata.UserRole) string {
	_ = "STUB: not implemented"
	return ""
}
