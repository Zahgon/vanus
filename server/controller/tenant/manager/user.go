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

//go:generate mockgen -source=user.go -destination=mock_user.go -package=manager
package manager

import (
	"context"
	"sync"

	"github.com/vanus-labs/vanus/pkg/kv"
	"github.com/vanus-labs/vanus/server/controller/tenant/metadata"
)

type UserManager interface {
	Init(ctx context.Context) error
	GetUser(ctx context.Context, identifier string) *metadata.User
	AddUser(ctx context.Context, user *metadata.User) error
	DeleteUser(ctx context.Context, identifier string) error
	ListUser(ctx context.Context) []*metadata.User
}

var _ UserManager = &userManager{}

type userManager struct {
	lock     sync.RWMutex
	users    map[string]*metadata.User
	kvClient kv.Client
}

func NewUserManager(kvClient kv.Client) UserManager {
	_ = "STUB: not implemented"
	return *new(UserManager)
}

func (m *userManager) Init(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (m *userManager) ListUser(_ context.Context) []*metadata.User {
	_ = "STUB: not implemented"
	return nil
}

func (m *userManager) GetUser(_ context.Context, identifier string) *metadata.User {
	_ = "STUB: not implemented"
	return nil
}

func (m *userManager) AddUser(ctx context.Context, user *metadata.User) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *userManager) DeleteUser(ctx context.Context, identifier string) error {
	_ = "STUB: not implemented"
	return nil
}
