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

//go:generate mockgen -source=authentication.go -destination=mock_authentication.go -package=authentication
package authentication

import (
	"context"
	"sync"
	"time"
)

type Authentication interface {
	// Authenticate check token valid and return user identifier
	Authenticate(ctx context.Context, token string) (string, error)
}

var _ Authentication = &authentication{}

const checkExpireTime = 30 * time.Second

type authentication struct {
	client     TokenClient
	tokens     sync.Map
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewAuthentication(client TokenClient) Authentication {
	_ = "STUB: not implemented"
	return *new(Authentication)
}

func (a *authentication) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (a *authentication) Stop(ctx context.Context) error {
	_ = "STUB: not implemented" //nolint:revive // ignore
	return nil
}

func (a *authentication) checkTokenExpired() { _ = "STUB: not implemented"; return }

func (a *authentication) Authenticate(ctx context.Context, token string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// todo breakdown cache
