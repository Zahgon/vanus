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

package proxy

import (
	"context"
	stdErr "errors"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	metapb "github.com/vanus-labs/vanus/api/meta"
	proxypb "github.com/vanus-labs/vanus/api/proxy"
	vanus "github.com/vanus-labs/vanus/api/vsr"

	"github.com/vanus-labs/vanus/pkg/authorization"
)

var errMethodNotImplemented = stdErr.New("the method hasn't implemented")

func authCreateEventbus(_ context.Context, req interface{},
) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) CreateEventbus(
	ctx context.Context, req *ctrlpb.CreateEventbusRequest,
) (*metapb.Eventbus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cp *ControllerProxy) CreateSystemEventbus(
	ctx context.Context, req *ctrlpb.CreateEventbusRequest,
) (*metapb.Eventbus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authDeleteEventbus(_ context.Context, req interface{},
) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) DeleteEventbus(
	ctx context.Context, id *wrapperspb.UInt64Value,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authGetEventbus(_ context.Context, req interface{},
) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) GetEventbus(
	ctx context.Context, id *wrapperspb.UInt64Value,
) (*metapb.Eventbus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cp *ControllerProxy) ListEventbus(
	ctx context.Context, req *ctrlpb.ListEventbusRequest,
) (*ctrlpb.ListEventbusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// grant namespace all eventbus

// grant eventbus

func (cp *ControllerProxy) GetEventbusWithHumanFriendly(ctx context.Context,
	request *ctrlpb.GetEventbusWithHumanFriendlyRequest,
) (*metapb.Eventbus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cp *ControllerProxy) UpdateEventbus(
	_ context.Context, _ *ctrlpb.UpdateEventbusRequest,
) (*metapb.Eventbus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authListSegment(_ context.Context, req interface{}) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) ListSegment(
	ctx context.Context, req *ctrlpb.ListSegmentRequest,
) (*ctrlpb.ListSegmentResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cp *ControllerProxy) ValidateEventbus(
	ctx context.Context, req *proxypb.ValidateEventbusRequest,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authCreateSubscription(_ context.Context, req interface{},
) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) CreateSubscription(
	ctx context.Context, req *ctrlpb.CreateSubscriptionRequest,
) (*metapb.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authUpdateSubscription(_ context.Context, req interface{},
) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) UpdateSubscription(
	ctx context.Context, req *ctrlpb.UpdateSubscriptionRequest,
) (*metapb.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authDeleteSubscription(_ context.Context, req interface{},
) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) DeleteSubscription(
	ctx context.Context, req *ctrlpb.DeleteSubscriptionRequest,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authGetSubscription(_ context.Context, req interface{},
) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) GetSubscription(
	ctx context.Context, req *ctrlpb.GetSubscriptionRequest,
) (*metapb.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cp *ControllerProxy) ListSubscription(
	ctx context.Context, req *ctrlpb.ListSubscriptionRequest,
) (*ctrlpb.ListSubscriptionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:prealloc //ok
// grant namespace all subscription

// grant subscription

func authDisableSubscription(_ context.Context, req interface{},
) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) DisableSubscription(
	ctx context.Context, req *ctrlpb.DisableSubscriptionRequest,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authResumeSubscription(_ context.Context, req interface{},
) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) ResumeSubscription(
	ctx context.Context, req *ctrlpb.ResumeSubscriptionRequest,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authResetOffsetSubscription(_ context.Context, req interface{},
) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) ResetOffsetToTimestamp(
	ctx context.Context, req *ctrlpb.ResetOffsetToTimestampRequest,
) (*ctrlpb.ResetOffsetToTimestampResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cp *ControllerProxy) GetNamespaceWithHumanFriendly(ctx context.Context,
	value *wrapperspb.StringValue,
) (*metapb.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authCreateNamespace(_ context.Context, _ interface{},
) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) CreateNamespace(ctx context.Context,
	request *ctrlpb.CreateNamespaceRequest,
) (*metapb.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cp *ControllerProxy) ListNamespace(ctx context.Context,
	empty *emptypb.Empty,
) (*ctrlpb.ListNamespaceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authGetNamespace(_ context.Context, req interface{}) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) GetNamespace(ctx context.Context, request *ctrlpb.GetNamespaceRequest,
) (*metapb.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authDeleteNamespace(_ context.Context, req interface{},
) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) DeleteNamespace(ctx context.Context, request *ctrlpb.DeleteNamespaceRequest,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
