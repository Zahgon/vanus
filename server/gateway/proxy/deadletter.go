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

package proxy

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/vanus-labs/vanus/api/cloudevents"
	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	proxypb "github.com/vanus-labs/vanus/api/proxy"
	vanus "github.com/vanus-labs/vanus/api/vsr"

	"github.com/vanus-labs/vanus/pkg/authorization"
)

func authGetDeadLetterEvent(_ context.Context, req interface{},
) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) GetDeadLetterEvent(
	ctx context.Context, req *proxypb.GetDeadLetterEventRequest,
) (*proxypb.GetDeadLetterEventResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// read end

// todo some error need retry read

func (cp *ControllerProxy) getDealLetterEventbusID(
	ctx context.Context, eventbusID vanus.ID,
) (vanus.ID, error) {
	_ = "STUB: not implemented"
	return *new(vanus.ID), nil
}

func authResendDeadLetterEvent(_ context.Context, req interface{},
) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) ResendDeadLetterEvent( //nolint:funlen // ok
	ctx context.Context, req *proxypb.ResendDeadLetterEventRequest,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// read end

// todo errors.ErrTryAgain maybe need retry read

// remove retry attribute

// remove dead letter attribute

func (cp *ControllerProxy) writeDeadLetterEvent(
	ctx context.Context, subscriptionID uint64, offset uint64, events []*cloudevents.CloudEvent,
) error {
	_ = "STUB: not implemented"
	return nil
}

// write to retry eventbus

// save offset

func authSetDeadLetterEventOffset(_ context.Context, req interface{},
) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) SetDeadLetterEventOffset(
	ctx context.Context, req *ctrlpb.SetDeadLetterEventOffsetRequest,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
