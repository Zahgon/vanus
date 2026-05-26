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

//go:generate mockgen -source=authorization.go -destination=mock_authorization.go -package=authorization
package authorization

import (
	"context"

	"github.com/vanus-labs/vanus/api/cluster"
	vanus "github.com/vanus-labs/vanus/api/vsr"
)

type Authorization interface {
	Authorize(ctx context.Context, user string, attributes Attributes) (bool, error)
}

var _ Authorization = &authorization{}

type authorization struct {
	client  RoleClient
	cluster cluster.Cluster
}

func NewAuthorization(client RoleClient, cluster cluster.Cluster) Authorization {
	_ = "STUB: not implemented"
	return *new(Authorization)
}

func (a *authorization) Authorize(ctx context.Context, user string, attributes Attributes) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// hasPermission is loop all user role to check has permission
// todo optimize use role compare with resource action role
func hasPermission(roles []*UserRole, attributes Attributes, resourceID vanus.ID) bool {
	_ = "STUB: not implemented"
	return false
}

// todo custom role
