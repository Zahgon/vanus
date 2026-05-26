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

package auth

import (
	"context"

	"github.com/vanus-labs/vanus/api/cluster"
	"github.com/vanus-labs/vanus/pkg/authentication"
	"github.com/vanus-labs/vanus/pkg/authorization"
)

const TokenType = "Bearer"

type Config struct {
	Disable          bool
	OpenEventbus     bool
	OpenSubscription bool
}

type Auth struct {
	config         Config
	Authentication authentication.Authentication
	Authorization  authorization.Authorization
	authorizeFunc  map[string]AuthorizeFunc
	RoleClient     authorization.RoleClient
	TokenClient    authentication.TokenClient
}

func NewAuth(config Config, cluster cluster.Cluster) *Auth { _ = "STUB: not implemented"; return nil }

func (a *Auth) Disable() bool { _ = "STUB: not implemented"; return false }

func (a *Auth) OpenEventbus() bool { _ = "STUB: not implemented"; return false }

func (a *Auth) OpenSubscription() bool { _ = "STUB: not implemented"; return false }

func (a *Auth) GetRoleClient() authorization.RoleClient {
	_ = "STUB: not implemented"
	return *new(authorization.RoleClient)
}

func (a *Auth) Authenticate(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (a *Auth) Authorize(ctx context.Context, method string, req interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// no need authorize
