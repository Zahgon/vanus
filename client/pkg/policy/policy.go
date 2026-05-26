// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package policy

import (
	"context"

	"github.com/vanus-labs/vanus/client/pkg/api"
)

var _ api.WritePolicy = (*roundRobinWritePolicy)(nil)

func NewRoundRobinWritePolicy(eb api.Eventbus) api.WritePolicy {
	_ = "STUB: not implemented"
	return *new(api.WritePolicy)
}

type roundRobinWritePolicy struct {
	bus api.Eventbus
	idx uint64
}

func (w *roundRobinWritePolicy) Type() api.PolicyType {
	_ = "STUB: not implemented"
	return *new(api.PolicyType)
}

func (w *roundRobinWritePolicy) NextLog(ctx context.Context) (api.Eventlog, error) {
	_ = "STUB: not implemented"
	return *new(api.Eventlog), nil
}

var _ api.ReadPolicy = (*roundRobinReadPolicy)(nil)

func NewRoundRobinReadPolicy(eb api.Eventbus, fromWhere api.ConsumeFromWhere) *roundRobinReadPolicy {
	_ = "STUB: not implemented"
	return nil
}

type roundRobinReadPolicy struct {
	bus    api.Eventbus
	idx    uint64
	offset int64
}

func (r *roundRobinReadPolicy) Type() api.PolicyType {
	_ = "STUB: not implemented"
	return *new(api.PolicyType)
}

func (r *roundRobinReadPolicy) NextLog(ctx context.Context) (api.Eventlog, error) {
	_ = "STUB: not implemented"
	return *new(api.Eventlog), nil
}

func (r *roundRobinReadPolicy) Offset() int64 { _ = "STUB: not implemented"; return 0 }

func (r *roundRobinReadPolicy) Forward(diff int) { _ = "STUB: not implemented"; return }

var _ api.ReadPolicy = (*manuallyReadPolicy)(nil)

func NewManuallyReadPolicy(log api.Eventlog, offset int64) *manuallyReadPolicy {
	_ = "STUB: not implemented"
	return nil
}

type manuallyReadPolicy struct {
	log    api.Eventlog
	offset int64
}

func (r manuallyReadPolicy) Type() api.PolicyType {
	_ = "STUB: not implemented"
	return *new(api.PolicyType)
}

func (r manuallyReadPolicy) NextLog(ctx context.Context) (api.Eventlog, error) {
	_ = "STUB: not implemented"
	return *new(api.Eventlog), nil
}

func (r manuallyReadPolicy) Offset() int64 { _ = "STUB: not implemented"; return 0 }

func (r *manuallyReadPolicy) Forward(diff int) { _ = "STUB: not implemented"; return }

var _ api.LogPolicy = (*readOnlyPolicy)(nil)

func NewReadOnlyPolicy() api.LogPolicy { _ = "STUB: not implemented"; return *new(api.LogPolicy) }

type readOnlyPolicy struct{}

func (w *readOnlyPolicy) AccessMode() api.PolicyType {
	_ = "STUB: not implemented"
	return *new(api.PolicyType)
}

var _ api.LogPolicy = (*readWritePolicy)(nil)

func NewReadWritePolicy() api.LogPolicy { _ = "STUB: not implemented"; return *new(api.LogPolicy) }

type readWritePolicy struct{}

func (w *readWritePolicy) AccessMode() api.PolicyType {
	_ = "STUB: not implemented"
	return *new(api.PolicyType)
}
