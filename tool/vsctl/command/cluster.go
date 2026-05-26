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

package command

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

const (
	DefaultInitialVersion    = "v0.9.0"
	DefaultImagePullPolicy   = "Always"
	DefaultResourceLimitsCPU = "500m"
	DefaultResourceLimitsMem = "1Gi"
)

var clusterVersionList = []string{DefaultInitialVersion}

type ClusterCreate struct {
	Version     string            `json:"version,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

type ClusterDelete struct {
	Force *bool `json:"force,omitempty"`
}

type ClusterPatch struct {
	Version     string            `json:"version,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

type ClusterInfo struct {
	Status  string `json:"status,omitempty"`
	Version string `json:"version,omitempty"`
}

type ClusterOKBody struct {
	Code    *int32       `json:"code"`
	Data    *ClusterInfo `json:"data"`
	Message *string      `json:"message"`
}

type ClusterSpec struct {
	Version         *string     `yaml:"version"`
	ImagePullPolicy *string     `yaml:"image_pull_policy"`
	Etcd            *Etcd       `yaml:"etcd"`
	Controller      *Controller `yaml:"controller"`
	Store           *Store      `yaml:"store"`
	Gateway         *Gateway    `yaml:"gateway"`
	Trigger         *Trigger    `yaml:"trigger"`
	Timer           *Timer      `yaml:"timer"`
}

type Etcd struct {
	Ports        *EtcdPorts `yaml:"ports"`
	Replicas     *int32     `yaml:"replicas"`
	StorageSize  *string    `yaml:"storage_size"`
	StorageClass *string    `yaml:"storage_class"`
	Resources    *Resources `yaml:"resources"`
}

type Controller struct {
	Ports           *ControllerPorts `yaml:"ports"`
	Replicas        *int32           `yaml:"replicas"`
	SegmentCapacity *string          `yaml:"segment_capacity"`
	Resources       *Resources       `yaml:"resources"`
}

type Store struct {
	Replicas     *int32     `yaml:"replicas"`
	StorageSize  *string    `yaml:"storage_size"`
	StorageClass *string    `yaml:"storage_class"`
	Resources    *Resources `yaml:"resources"`
}

type Gateway struct {
	Ports     *GatewayPorts     `yaml:"ports"`
	NodePorts *GatewayNodePorts `yaml:"nodeports"`
	Replicas  *int32            `yaml:"replicas"`
	Resources *Resources        `yaml:"resources"`
}

type Trigger struct {
	Replicas  *int32     `yaml:"replicas"`
	Resources *Resources `yaml:"resources"`
}

type Timer struct {
	Replicas    *int32       `yaml:"replicas"`
	TimingWheel *TimingWheel `yaml:"timingwheel"`
	Resources   *Resources   `yaml:"resources"`
}

type EtcdPorts struct {
	Client *int32 `yaml:"client"`
	Peer   *int32 `yaml:"peer"`
}

type ControllerPorts struct {
	Controller     *int32 `yaml:"controller"`
	RootController *int32 `yaml:"root_controller"`
}

type GatewayPorts struct {
	Proxy       *int32 `yaml:"proxy"`
	CloudEvents *int32 `yaml:"cloudevents"`
	SinkProxy   *int32 `yaml:"sink_proxy"`
}

type GatewayNodePorts struct {
	Proxy       *int32 `yaml:"proxy"`
	CloudEvents *int32 `yaml:"cloudevents"`
}

type TimingWheel struct {
	Tick   *int32 `yaml:"tick"`
	Size   *int32 `yaml:"wheel_size"`
	Layers *int32 `yaml:"layers"`
}

type Resources struct {
	LimitsCPU *string `yaml:"limits_cpu"`
	LimitsMem *string `yaml:"limits_mem"`
}

func NewClusterCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func createClusterCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func deleteClusterCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func upgradeClusterCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func scaleClusterCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func scaleStoreReplicas() *cobra.Command { _ = "STUB: not implemented"; return nil }

func scaleTriggerReplicas() *cobra.Command { _ = "STUB: not implemented"; return nil }

func getClusterCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func genClusterCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func getCluster(cmd *cobra.Command, endpoint string) (*ClusterOKBody, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func clusterColConfigs() []table.ColumnConfig { _ = "STUB: not implemented"; return nil }

func getUpgradableVersionList(curVersion string) []string { _ = "STUB: not implemented"; return nil }

func clusterIsVaild(c *ClusterSpec) bool { _ = "STUB: not implemented"; return false }

func genClusterAnnotations(c *ClusterSpec) map[string]string { _ = "STUB: not implemented"; return nil }

// Etcd

// Controller

// Store

// Gateway

// Trigger

// Timer
