// Copyright 2022 Linkall Inc.
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

package controller

import (
	// first-party libraries.
	"github.com/vanus-labs/vanus/pkg/observability"

	// this project.

	"github.com/vanus-labs/vanus/server/controller/eventbus"
	"github.com/vanus-labs/vanus/server/controller/member"
	"github.com/vanus-labs/vanus/server/controller/root"
	"github.com/vanus-labs/vanus/server/controller/tenant"
	"github.com/vanus-labs/vanus/server/controller/trigger"
)

type Config struct {
	Observability        observability.Config `yaml:"observability"`
	NodeID               uint16               `yaml:"node_id"`
	Name                 string               `yaml:"name"`
	IP                   string               `yaml:"ip"`
	Port                 int                  `yaml:"port"`
	GRPCReflectionEnable bool                 `yaml:"grpc_reflection_enable"`
	MetadataConfig       MetadataConfig       `yaml:"metadata"`
	Replicas             uint                 `yaml:"replicas"`
	SecretEncryptionSalt string               `yaml:"secret_encryption_salt"`
	SegmentCapacity      int64                `yaml:"segment_capacity"`
	ClusterConfig        member.Config        `yaml:"cluster"`
	RootControllerAddr   []string             `yaml:"root_controllers"`
	NoCreateDefaultNs    bool                 `yaml:"no_create_default_namespace"`
}

func (c *Config) GetClusterConfig() member.Config {
	_ = "STUB: not implemented"
	return *new(member.Config)
}

func (c *Config) GetEventbusCtrlConfig() eventbus.Config {
	_ = "STUB: not implemented"
	return *new(eventbus.Config)
}

func (c *Config) GetSnowflakeConfig() root.Config {
	_ = "STUB: not implemented"
	return *new(root.Config)
}

func (c *Config) GetControllerAddrs() []string { _ = "STUB: not implemented"; return nil }

type MetadataConfig struct {
	KeyPrefix string `yaml:"key_prefix"`
}

func (c *Config) GetTriggerConfig() trigger.Config {
	_ = "STUB: not implemented"
	return *new(trigger.Config)
}

func (c *Config) GetTenantConfig() tenant.Config {
	_ = "STUB: not implemented"
	return *new(tenant.Config)
}
