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

package option

import (
	"time"

	"github.com/vanus-labs/vanus/client/pkg/api"
)

func WithWritePolicy(policy api.WritePolicy) api.WriteOption {
	_ = "STUB: not implemented"
	return *new(api.WriteOption)
}

func WithOneway() api.WriteOption { _ = "STUB: not implemented"; return *new(api.WriteOption) }

func WithBatchSize(size int) api.ReadOption { _ = "STUB: not implemented"; return *new(api.ReadOption) }

func WithPollingTimeout(d time.Duration) api.ReadOption {
	_ = "STUB: not implemented"
	return *new(api.ReadOption)
}

func WithDisablePolling() api.ReadOption { _ = "STUB: not implemented"; return *new(api.ReadOption) }

func WithReadPolicy(policy api.ReadPolicy) api.ReadOption {
	_ = "STUB: not implemented"
	return *new(api.ReadOption)
}

func WithLogPolicy(policy api.LogPolicy) api.LogOption {
	_ = "STUB: not implemented"
	return *new(api.LogOption)
}
