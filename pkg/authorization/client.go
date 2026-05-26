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

//go:generate mockgen -source=client.go -destination=mock_client.go -package=authorization
package authorization

import (
	// standard libraries.
	"context"

	// first-party libraries.
	"github.com/vanus-labs/vanus/api/cluster"
	vanus "github.com/vanus-labs/vanus/api/vsr"
)

type RoleClient interface {
	// IsClusterAdmin check the use is cluster admin.
	IsClusterAdmin(ctx context.Context, user string) (bool, error)
	// GetUserNamespaceID get grant user namespaceID.
	GetUserNamespaceID(ctx context.Context, user string) (vanus.IDList, error)
	// GetUserEventbusID get grant user eventbusID, not contains the eventbus grant namespace.
	GetUserEventbusID(ctx context.Context, user string) (vanus.IDList, error)
	// GetUserSubscriptionID get grant user SubscriptionID, not contains the subscription grant namespace.
	GetUserSubscriptionID(ctx context.Context, user string) (vanus.IDList, error)
	// GetUserRole get user role
	GetUserRole(ctx context.Context, user string) ([]*UserRole, error)
}

var _ RoleClient = &builtInClient{}

type builtInClient struct {
	cluster cluster.Cluster
}

func NewBuiltInClient(cluster cluster.Cluster) RoleClient {
	_ = "STUB: not implemented"
	return *new(RoleClient)
}

func (c *builtInClient) GetUserRole(ctx context.Context, user string) ([]*UserRole, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *builtInClient) IsClusterAdmin(ctx context.Context, user string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *builtInClient) GetUserNamespaceID(ctx context.Context, user string) (vanus.IDList, error) {
	_ = "STUB: not implemented"
	return *new(vanus.IDList), nil
}

func (c *builtInClient) GetUserEventbusID(ctx context.Context, user string) (vanus.IDList, error) {
	_ = "STUB: not implemented"
	return *new(vanus.IDList), nil
}

func (c *builtInClient) GetUserSubscriptionID(ctx context.Context, user string) (vanus.IDList, error) {
	_ = "STUB: not implemented"
	return *new(vanus.IDList), nil
}

func (c *builtInClient) getUserResourceID(ctx context.Context, user string, kind ResourceKind) (vanus.IDList, error) {
	_ = "STUB: not implemented"
	return *new(vanus.IDList), nil
}
