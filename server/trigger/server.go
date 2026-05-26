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
	// standard libraries.
	"context"
	"time"

	// first-party libraries.

	triggerpb "github.com/vanus-labs/vanus/api/trigger"

	// this project.
	primitive "github.com/vanus-labs/vanus/pkg"
)

var _ triggerpb.TriggerWorkerServer = &server{}

type server struct {
	worker    Worker
	config    Config
	state     primitive.ServerState
	startTime time.Time
}

func NewTriggerServer(config Config) triggerpb.TriggerWorkerServer {
	_ = "STUB: not implemented"
	return *new(triggerpb.TriggerWorkerServer)
}

func (s *server) Start(ctx context.Context,
	_ *triggerpb.StartTriggerWorkerRequest,
) (*triggerpb.StartTriggerWorkerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *server) Stop(ctx context.Context,
	_ *triggerpb.StopTriggerWorkerRequest,
) (*triggerpb.StopTriggerWorkerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *server) AddSubscription(ctx context.Context,
	request *triggerpb.AddSubscriptionRequest,
) (*triggerpb.AddSubscriptionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *server) RemoveSubscription(ctx context.Context,
	request *triggerpb.RemoveSubscriptionRequest,
) (*triggerpb.RemoveSubscriptionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *server) PauseSubscription(ctx context.Context,
	request *triggerpb.PauseSubscriptionRequest,
) (*triggerpb.PauseSubscriptionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *server) ResumeSubscription(ctx context.Context,
	request *triggerpb.ResumeSubscriptionRequest,
) (*triggerpb.ResumeSubscriptionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *server) Initialize(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *server) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *server) stop(ctx context.Context, sendUnregister bool) { _ = "STUB: not implemented"; return }
