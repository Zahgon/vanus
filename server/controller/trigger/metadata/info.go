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

package metadata

import (
	"time"

	vanus "github.com/vanus-labs/vanus/api/vsr"

	primitive "github.com/vanus-labs/vanus/pkg"
)

type TriggerWorkerPhase string

const (
	TriggerWorkerPhasePending    TriggerWorkerPhase = "pending"
	TriggerWorkerPhaseRunning    TriggerWorkerPhase = "running"
	TriggerWorkerPhasePaused     TriggerWorkerPhase = "paused"
	TriggerWorkerPhaseDisconnect TriggerWorkerPhase = "disconnect"
)

type TriggerWorkerInfo struct {
	ID    string             `json:"-"`
	Addr  string             `json:"addr"`
	Phase TriggerWorkerPhase `json:"phase"`
}

func NewTriggerWorkerInfo(addr string) *TriggerWorkerInfo { _ = "STUB: not implemented"; return nil }

func (tw *TriggerWorkerInfo) String() string { _ = "STUB: not implemented"; return "" }

type SubscriptionPhase string

const (
	SubscriptionPhaseCreated  SubscriptionPhase = "created"
	SubscriptionPhasePending  SubscriptionPhase = "pending"
	SubscriptionPhaseRunning  SubscriptionPhase = "running"
	SubscriptionPhaseStopping SubscriptionPhase = "stopping"
	SubscriptionPhaseStopped  SubscriptionPhase = "stopped"
	SubscriptionPhaseToDelete SubscriptionPhase = "toDelete"
)

type Subscription struct {
	ID                 vanus.ID                        `json:"id"`
	Source             string                          `json:"source,omitempty"`
	Types              []string                        `json:"types,omitempty"`
	Config             primitive.SubscriptionConfig    `json:"config,omitempty"`
	Filters            []*primitive.SubscriptionFilter `json:"filters,omitempty"`
	Sink               primitive.URI                   `json:"sink,omitempty"`
	SinkCredentialType *primitive.CredentialType       `json:"sink_credential_type,omitempty"`
	SinkCredential     primitive.SinkCredential        `json:"-"`
	Protocol           primitive.Protocol              `json:"protocol,omitempty"`
	ProtocolSetting    *primitive.ProtocolSetting      `json:"protocol_settings,omitempty"`
	EventbusID         vanus.ID                        `json:"eventbus_id"`
	NamespaceID        vanus.ID                        `json:"namespace_id"`
	Transformer        *primitive.Transformer          `json:"transformer,omitempty"`
	Name               string                          `json:"name"`
	Description        string                          `json:"description"`
	CreatedAt          time.Time                       `json:"created_at"`
	UpdatedAt          time.Time                       `json:"updated_at"`

	// not from api
	DeadLetterEventbusID vanus.ID          `json:"dead_letter_eventbus_id"`
	RetryEventbusID      vanus.ID          `json:"retry_eventbus_id"`
	TimerEventbusID      vanus.ID          `json:"timer_eventbus_id"`
	Phase                SubscriptionPhase `json:"phase"`
	TriggerWorker        string            `json:"trigger_worker,omitempty"`
	HeartbeatTime        time.Time         `json:"-"`
}

// Update property change from api .
func (s *Subscription) Update(update *Subscription) bool { _ = "STUB: not implemented"; return false }
