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

package convert

import (
	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	metapb "github.com/vanus-labs/vanus/api/meta"

	"github.com/vanus-labs/vanus/server/controller/tenant/metadata"
)

func FromPbCreateNamespace(ns *ctrlpb.CreateNamespaceRequest) *metadata.Namespace {
	_ = "STUB: not implemented"
	return nil
}

func ToPbNamespace(ns *metadata.Namespace) *metapb.Namespace { _ = "STUB: not implemented"; return nil }

func ToPbUser(from *metadata.User) *metapb.User { _ = "STUB: not implemented"; return nil }

func ToPbToken(token *metadata.Token) *metapb.Token { _ = "STUB: not implemented"; return nil }

func FromPbRoleRequest(from *ctrlpb.RoleRequest) *metadata.UserRole {
	_ = "STUB: not implemented"
	return nil
}

func ToPbUserRole(from *metadata.UserRole) *metapb.UserRole { _ = "STUB: not implemented"; return nil }

func ToPbResourceRole(from *metadata.UserRole) *metapb.ResourceRole {
	_ = "STUB: not implemented"
	return nil
}
