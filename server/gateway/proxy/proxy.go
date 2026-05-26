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
	// standard libraries.
	"context"
	"sync"
	stdtime "time"

	// third-party libraries.
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/protocol"
	cehttp "github.com/cloudevents/sdk-go/v2/protocol/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/emptypb"

	// first-party libraries.
	"github.com/vanus-labs/vanus/api/cloudevents"
	"github.com/vanus-labs/vanus/api/cluster"
	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	"github.com/vanus-labs/vanus/api/errors"
	metapb "github.com/vanus-labs/vanus/api/meta"
	proxypb "github.com/vanus-labs/vanus/api/proxy"
	vanus "github.com/vanus-labs/vanus/api/vsr"
	eb "github.com/vanus-labs/vanus/client"
	"github.com/vanus-labs/vanus/pkg/observability/tracing"

	// this project.

	"github.com/vanus-labs/vanus/pkg/authorization"
	"github.com/vanus-labs/vanus/server/gateway/auth"
)

const (
	maximumNumberPerGetRequest = 64
	eventChanCache             = 10
	readSize                   = 5
	ContentTypeProtobuf        = "application/protobuf"
	httpRequestPrefix          = "/gatewaysink"
	datacontenttype            = "datacontenttype"
	dataschema                 = "dataschema"
	subject                    = "subject"
	time                       = "time"
)

var (
	errInvalidEventbus     = errors.New("the eventbus name can't be empty")
	requestDataFromContext = cehttp.RequestDataFromContext
	zeroTime               = stdtime.Time{}
)

type Config struct {
	Endpoints              []string
	SinkPort               int
	ProxyPort              int
	CloudEventReceiverPort int
	Credentials            credentials.TransportCredentials
	GRPCReflectionEnable   bool
	AuthCfg                auth.Config
}

type ackCallback func(bool)

type message struct {
	sequenceID uint64
	event      *v2.Event
}

type subscribeCache struct {
	sequenceID      uint64
	subscriptionID  string
	subscribeStream proxypb.StoreProxy_SubscribeServer
	acks            sync.Map
	eventc          chan message
}

func newSubscribeCache(subscriptionID string, stream proxypb.StoreProxy_SubscribeServer) *subscribeCache {
	_ = "STUB: not implemented"
	return nil
}

func (s *subscribeCache) ch() chan message { _ = "STUB: not implemented"; return nil }

func (s *subscribeCache) stream() proxypb.StoreProxy_SubscribeServer {
	_ = "STUB: not implemented"
	return *new(proxypb.StoreProxy_SubscribeServer)
}

type ControllerProxy struct {
	cfg          Config
	tracer       *tracing.Tracer
	client       eb.Client
	eventbusCtrl ctrlpb.EventbusControllerClient
	eventlogCtrl ctrlpb.EventlogControllerClient
	triggerCtrl  ctrlpb.TriggerControllerClient
	nsCtrl       ctrlpb.NamespaceControllerClient
	authCtrl     ctrlpb.AuthControllerClient
	grpcSrv      *grpc.Server
	ctrl         cluster.Cluster
	writerMap    sync.Map
	cache        sync.Map
	authService  *auth.Auth
}

// Make sure ControllerProxy implements proxypb.StoreProxyServer.
var _ proxypb.StoreProxyServer = (*ControllerProxy)(nil)

func authPublish(_ context.Context, req interface{}) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) Publish(ctx context.Context, req *proxypb.PublishRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// todo  common event with delay event mixture

// validate event time

func (cp *ControllerProxy) writeEvents(
	ctx context.Context, eventbusID vanus.ID, events *cloudevents.CloudEventBatch,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Subscribe todo authentication
func (cp *ControllerProxy) Subscribe(req *proxypb.SubscribeRequest, stream proxypb.StoreProxy_SubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

// 1. modify subscription sink

// TODO add config

// 2. cache subscribe info

// 3. receive and forward events

// TODO(jiangkai): err check

func (cp *ControllerProxy) Ack(stream proxypb.StoreProxy_AckServer) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(jiangkai): err check

func ToProto(e *v2.Event) (*cloudevents.CloudEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cp *ControllerProxy) disableSubscription(
	ctx context.Context, req *proxypb.SubscribeRequest, subscriptionID uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(jiangkai): delete me after disable supports synchronization interface

func newSubscription(
	info *metapb.Subscription, subscriptionID uint64, newsink string,
) *ctrlpb.UpdateSubscriptionRequest {
	_ = "STUB: not implemented"
	return nil
}

func attributeFor(v interface{}) (*cloudevents.CloudEvent_CloudEventAttributeValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkExtension(extensions map[string]*cloudevents.CloudEvent_CloudEventAttributeValue) error {
	_ = "STUB: not implemented"
	return nil
}

// event attribute can not prefix with vanus system use

func NewControllerProxy(cfg Config) *ControllerProxy { _ = "STUB: not implemented"; return nil }

// SetClient just for test.
func (cp *ControllerProxy) SetClient(client eb.Client) { _ = "STUB: not implemented"; return }

func (cp *ControllerProxy) Start() error { _ = "STUB: not implemented"; return nil }

// for debug in developing stage

func (cp *ControllerProxy) receive(ctx context.Context, event v2.Event) (*v2.Event, protocol.Result) {
	_ = "STUB: not implemented"
	return nil, *new(protocol.Result)
}

// retry

func getSubscriptionIDFromPath(reqData *cehttp.RequestData) string {
	_ = "STUB: not implemented"
	// TODO validate
	return ""
}

func (cp *ControllerProxy) Stop() { _ = "STUB: not implemented"; return }

func (cp *ControllerProxy) ClusterInfo(_ context.Context, _ *emptypb.Empty) (*proxypb.ClusterInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authLookupOffset(_ context.Context, req interface{}) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) LookupOffset(
	ctx context.Context, req *proxypb.LookupOffsetRequest,
) (*proxypb.LookupOffsetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authGetEvent(_ context.Context, req interface{}) (authorization.ResourceKind, vanus.ID, authorization.Action) {
	_ = "STUB: not implemented"
	return *new(authorization.ResourceKind), *new(vanus.ID), *new(authorization.Action)
}

func (cp *ControllerProxy) GetEvent(
	ctx context.Context, req *proxypb.GetEventRequest,
) (*proxypb.GetEventResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cp *ControllerProxy) ValidateSubscription(
	ctx context.Context, req *proxypb.ValidateSubscriptionRequest,
) (*proxypb.ValidateSubscriptionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getByEventID why added this? can it be deleted?
func (cp *ControllerProxy) getByEventID(
	ctx context.Context, req *proxypb.GetEventRequest,
) (*proxypb.GetEventResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decodeEventID(eventID string) (uint64, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// fixed length
