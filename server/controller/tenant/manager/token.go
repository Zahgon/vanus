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

//go:generate mockgen -source=token.go -destination=mock_token.go -package=manager
package manager

import (
	"context"
	"sync"

	vanus "github.com/vanus-labs/vanus/api/vsr"
	"github.com/vanus-labs/vanus/pkg/kv"
	"github.com/vanus-labs/vanus/server/controller/tenant/metadata"
)

type TokenManager interface {
	Init(ctx context.Context) error
	GetUser(ctx context.Context, token string) (string, error)
	AddToken(ctx context.Context, user *metadata.Token) error
	DeleteToken(ctx context.Context, id vanus.ID) error
	GetToken(ctx context.Context, id vanus.ID) (*metadata.Token, error)
	GetUserToken(ctx context.Context, identifier string) []*metadata.Token
	ListToken(ctx context.Context) []*metadata.Token
}

var _ TokenManager = &tokenManager{}

type tokenManager struct {
	lock       sync.RWMutex
	tokens     map[vanus.ID]*metadata.Token
	users      map[string]map[vanus.ID]struct{}
	tokenUsers map[string]string
	kvClient   kv.Client
}

func NewTokenManager(kvClient kv.Client) TokenManager {
	_ = "STUB: not implemented"
	return *new(TokenManager)
}

func (m *tokenManager) Init(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (m *tokenManager) GetUser(_ context.Context, token string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *tokenManager) AddToken(ctx context.Context, token *metadata.Token) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *tokenManager) DeleteToken(ctx context.Context, id vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *tokenManager) GetToken(_ context.Context, id vanus.ID) (*metadata.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *tokenManager) GetUserToken(_ context.Context, identifier string) []*metadata.Token {
	_ = "STUB: not implemented"
	return nil
}

func (m *tokenManager) ListToken(_ context.Context) []*metadata.Token {
	_ = "STUB: not implemented"
	return nil
}
