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

//go:generate mockgen -source=namespace.go -destination=mock_namespace.go -package=manager
package manager

import (
	"context"
	"sync"

	vanus "github.com/vanus-labs/vanus/api/vsr"

	"github.com/vanus-labs/vanus/pkg/kv"
	"github.com/vanus-labs/vanus/server/controller/tenant/metadata"
)

type NamespaceManager interface {
	Init(ctx context.Context) error
	AddNamespace(ctx context.Context, ns *metadata.Namespace) error
	DeleteNamespace(ctx context.Context, id vanus.ID) error
	GetNamespace(ctx context.Context, id vanus.ID) *metadata.Namespace
	GetNamespaceByName(ctx context.Context, name string) *metadata.Namespace
	ListNamespace(ctx context.Context) []*metadata.Namespace
}

var _ NamespaceManager = &namespaceManager{}

type namespaceManager struct {
	lock       sync.RWMutex
	namespaces map[vanus.ID]*metadata.Namespace
	kvClient   kv.Client
}

func NewNamespaceManager(client kv.Client) NamespaceManager {
	_ = "STUB: not implemented"
	return *new(NamespaceManager)
}

func (m *namespaceManager) Init(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (m *namespaceManager) AddNamespace(ctx context.Context, ns *metadata.Namespace) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *namespaceManager) DeleteNamespace(ctx context.Context, id vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *namespaceManager) GetNamespace(_ context.Context, id vanus.ID) *metadata.Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (m *namespaceManager) GetNamespaceByName(_ context.Context, name string) *metadata.Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (m *namespaceManager) ListNamespace(_ context.Context) []*metadata.Namespace {
	_ = "STUB: not implemented"
	return nil
}
