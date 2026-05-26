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

package trigger

import (
	"time"

	primitive "github.com/vanus-labs/vanus/pkg"
	"github.com/vanus-labs/vanus/server/trigger/client"
)

type clientConfig struct {
	gateway    *TargetGateway
	sink       primitive.URI
	protocol   primitive.Protocol
	credential primitive.SinkCredential
}

func newEventClient(cfg clientConfig) client.EventClient {
	_ = "STUB: not implemented"
	return *new(client.EventClient)
}

const (
	OrderEventCode   = -1
	ErrTransformCode = 1
)

func isShouldRetry(statusCode int) (bool, string) { _ = "STUB: not implemented"; return false, "" }

func calDeliveryTime(attempts int32) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func getRetryAttempts(attempts interface{}) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
