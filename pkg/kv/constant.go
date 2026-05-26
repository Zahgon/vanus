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

package kv

import (
	vanus "github.com/vanus-labs/vanus/api/vsr"
)

const (
	ResourceEventbus          = "/vanus/core/eventbus_controller/eventbus"
	ResourceEventlog          = "/vanus/core/eventbus_controller/eventlog"
	ResourceSegment           = "/vanus/core/eventbus_controller/segment"
	ResourceSegmentOfEventlog = "/vanus/core/eventbus_controller/segs_of_eventlog"
	ResourceVolumeMetadata    = "/vanus/core/eventbus_controller/volume/metadata"
	ResourceVolumeBlock       = "/vanus/core/eventbus_controller/volume/block"
	ResourceVolumeInstance    = "/vanus/core/eventbus_controller/volume/instance"
	ResourceSubscription      = "/vanus/core/trigger_controller/subscriptions"
	MetadataSecret            = "/vanus/core/trigger_controller/secrets" //nolint:gosec // ok
	MetadataOffset            = "/vanus/core/trigger_controller/offsets"
	TriggerWorker             = "/vanus/core/trigger_controller/trigger_workers"
	ClusterNode               = "/vanus/core/cluster/nodes"
	ClusterStart              = "/vanus/core/cluster/start_at"

	leaderLock = "/vanus/core/cluster/resource_lock"
	leaderInfo = "/vanus/core/cluster/leader_info"

	namespace = "/vanus/core/tenant/namespaces"
	user      = "/vanus/core/tenant/users"
	userRole  = "/vanus/core/tenants/user_role"
	userToken = "/vanus/core/tenant/tokens" //nolint:gosec // ok
)

func DistributedLockKey(component string) string { _ = "STUB: not implemented"; return "" }

func ComponentLeaderKey(component string) string { _ = "STUB: not implemented"; return "" }

func NamespaceAllKey() string { _ = "STUB: not implemented"; return "" }

func NamespaceKey(id vanus.ID) string { _ = "STUB: not implemented"; return "" }

func UserAllKey() string { _ = "STUB: not implemented"; return "" }

func UserKey(identifier string) string { _ = "STUB: not implemented"; return "" }

func UserTokenAllKey() string { _ = "STUB: not implemented"; return "" }

func UserTokenKey(id vanus.ID) string { _ = "STUB: not implemented"; return "" }

func UserRoleAllKey() string { _ = "STUB: not implemented"; return "" }

func UserRoleKey(identifier, role string) string { _ = "STUB: not implemented"; return "" }
