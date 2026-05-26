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

package cluster

import (
	"context"
	"sync"

	"github.com/vanus-labs/vanus/api/cluster/raw_client"
	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	metapb "github.com/vanus-labs/vanus/api/meta"
)

const (
	systemNamespace  = "vanus-system"
	defaultNamespace = "default"
)

type namespaceService struct {
	client ctrlpb.NamespaceControllerClient
	cache  sync.Map
}

func (ns *namespaceService) GetNamespace(ctx context.Context, id uint64) (*metapb.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ns.cache.Store(id, n) unmask when dirty cache is resolved

func (ns *namespaceService) GetSystemNamespace(ctx context.Context) (*metapb.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ns *namespaceService) GetDefaultNamespace(ctx context.Context) (*metapb.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ns *namespaceService) GetNamespaceByName(ctx context.Context, name string) (*metapb.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ns.cache.Store(name, n) unmask when dirty cache is resolved

func newNamespaceService(cc *raw_client.Conn) NamespaceService {
	_ = "STUB: not implemented"
	return *new(NamespaceService)
}

func (ns *namespaceService) RawClient() ctrlpb.NamespaceControllerClient {
	_ = "STUB: not implemented"
	return *new(ctrlpb.NamespaceControllerClient)
}
