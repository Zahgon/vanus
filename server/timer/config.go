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

package timer

import (
	// standard libraries.

	// first-party libraries.
	"github.com/vanus-labs/vanus/pkg/observability"

	// this project.
	"github.com/vanus-labs/vanus/server/timer/leaderelection"
	"github.com/vanus-labs/vanus/server/timer/timingwheel"
)

type Config struct {
	Observability        observability.Config `yaml:"observability"`
	Name                 string               `yaml:"name"`
	IP                   string               `yaml:"ip"`
	Port                 int                  `yaml:"port"`
	Replicas             uint                 `yaml:"replicas"`
	EtcdEndpoints        []string             `yaml:"etcd"`
	CtrlEndpoints        []string             `yaml:"controllers"`
	MetadataConfig       MetadataConfig       `yaml:"metadata"`
	LeaderElectionConfig LeaderElectionConfig `yaml:"leader_election"`
	TimingWheelConfig    TimingWheelConfig    `yaml:"timingwheel"`
}

const (
	resourceLockName = "timer"
)

func (c *Config) GetLeaderElectionConfig() *leaderelection.Config {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) GetTimingWheelConfig() *timingwheel.Config { _ = "STUB: not implemented"; return nil }

type MetadataConfig struct {
	KeyPrefix string `yaml:"key_prefix"`
}

type LeaderElectionConfig struct {
	LeaseDuration int64 `yaml:"lease_duration"`
}

type TimingWheelConfig struct {
	Tick      int64 `yaml:"tick"`
	WheelSize int64 `yaml:"wheel_size"`
	Layers    int64 `yaml:"layers"`
}

func Default(c *Config) { _ = "STUB: not implemented"; return }
