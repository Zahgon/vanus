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
	"github.com/spf13/cobra"
)

var supportedConnectors = []ConnectorSpec{
	{
		Kind:    "source",
		Type:    "http",
		Version: "latest",
	},
	{
		Kind:    "sink",
		Type:    "feishu",
		Version: "latest",
	},
}

type ConnectorSpec struct {
	Kind    string
	Type    string
	Version string
}

type ConnectorCreate struct {
	Kind        string                 `json:"kind,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Type        string                 `json:"type,omitempty"`
	Version     string                 `json:"version,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
	Annotations map[string]string      `json:"annotations,omitempty"`
}

type ConnectorDelete struct {
	Force *bool `json:"force,omitempty"`
}

type ConnectorInfo struct {
	Kind    string `json:"kind,omitempty"`
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
	Version string `json:"version,omitempty"`
	Status  string `json:"status,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

type ConnectorOKBody struct {
	Code    *int32         `json:"code"`
	Data    *ConnectorInfo `json:"data"`
	Message *string        `json:"message"`
}

type ListConnectorOKBody struct {
	Code    *int32           `json:"code"`
	Data    []*ConnectorInfo `json:"data"`
	Message *string          `json:"message"`
}

func NewConnectorCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func installConnectorCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func uninstallConnectorCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func listConnectorCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func getConnectorCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func getConfig(file string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func isUnsupported(kind, ctype, version string) bool { _ = "STUB: not implemented"; return false }
