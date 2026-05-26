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

package convert

import (
	// standard libraries.

	// third-party libraries.

	// first-party libraries.
	ctrl "github.com/vanus-labs/vanus/api/controller"
	pb "github.com/vanus-labs/vanus/api/meta"
	triggerpb "github.com/vanus-labs/vanus/api/trigger"

	// this project.
	primitive "github.com/vanus-labs/vanus/pkg"
	"github.com/vanus-labs/vanus/pkg/info"
	"github.com/vanus-labs/vanus/server/controller/trigger/metadata"
)

func FromPbSubscriptionRequest(sub *ctrl.SubscriptionRequest) (*metadata.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fromPbProtocol(from pb.Protocol) primitive.Protocol {
	_ = "STUB: not implemented"
	return *new(primitive.Protocol)
}

func toPbProtocol(from primitive.Protocol) pb.Protocol {
	_ = "STUB: not implemented"
	return *new(pb.Protocol)
}

func fromPbProtocolSettings(from *pb.ProtocolSetting) *primitive.ProtocolSetting {
	_ = "STUB: not implemented"
	return nil
}

func toPbProtocolSettings(from *primitive.ProtocolSetting) *pb.ProtocolSetting {
	_ = "STUB: not implemented"
	return nil
}

func fromPbSinkCredentialType(from *pb.SinkCredential) *primitive.CredentialType {
	_ = "STUB: not implemented"
	return nil
}

func fromPbSinkCredential(from *pb.SinkCredential) primitive.SinkCredential {
	_ = "STUB: not implemented"
	return *new(primitive.SinkCredential)
}

func toPbSinkCredentialByType(credentialType *primitive.CredentialType) *pb.SinkCredential {
	_ = "STUB: not implemented"
	return nil
}

func toPbSinkCredential(from primitive.SinkCredential) *pb.SinkCredential {
	_ = "STUB: not implemented"
	return nil
}

func fromPbSubscriptionConfig(config *pb.SubscriptionConfig) primitive.SubscriptionConfig {
	_ = "STUB: not implemented"
	return *new(primitive.SubscriptionConfig)
}

func toPbSubscriptionConfig(config primitive.SubscriptionConfig) *pb.SubscriptionConfig {
	_ = "STUB: not implemented"
	return nil
}

func FromPbAddSubscription(req *triggerpb.AddSubscriptionRequest) (*primitive.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPbAddSubscription(sub *primitive.Subscription) *triggerpb.AddSubscriptionRequest {
	_ = "STUB: not implemented"
	return nil
}

func ToPbSubscription(sub *metadata.Subscription, offsets info.ListOffsetInfo) *pb.Subscription {
	_ = "STUB: not implemented"
	return nil
}

func fromPbFilters(filters []*pb.Filter) []*primitive.SubscriptionFilter {
	_ = "STUB: not implemented"
	return nil
}

func fromPbFilter(filter *pb.Filter) *primitive.SubscriptionFilter {
	_ = "STUB: not implemented"
	return nil
}

func toPbFilters(filters []*primitive.SubscriptionFilter) []*pb.Filter {
	_ = "STUB: not implemented"
	return nil
}

func toPbFilter(filter *primitive.SubscriptionFilter) *pb.Filter {
	_ = "STUB: not implemented"
	return nil
}

func FromPbOffsetInfos(offsets []*pb.OffsetInfo) info.ListOffsetInfo {
	_ = "STUB: not implemented"
	return *new(info.ListOffsetInfo)
}

func fromPbOffsetInfo(offset *pb.OffsetInfo) info.OffsetInfo {
	_ = "STUB: not implemented"
	return *new(info.OffsetInfo)
}

func ToPbSubscriptionInfo(sub info.SubscriptionInfo) *pb.SubscriptionInfo {
	_ = "STUB: not implemented"
	return nil
}

func ToPbOffsetInfos(offsets info.ListOffsetInfo) []*pb.OffsetInfo {
	_ = "STUB: not implemented"
	return nil
}

func toPbOffsetInfo(offset info.OffsetInfo) *pb.OffsetInfo { _ = "STUB: not implemented"; return nil }

func fromPbTransformer(transformer *pb.Transformer) (*primitive.Transformer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:nilnil // nil is valid

func fromPbActions(actions []*pb.Action) []*primitive.Action { _ = "STUB: not implemented"; return nil }

func fromPbCommand(action *pb.Action) *primitive.Action { _ = "STUB: not implemented"; return nil }

func FromPbTemplateType(t pb.TemplateType) (primitive.TemplateType, error) {
	_ = "STUB: not implemented"
	return *new(primitive.TemplateType), nil
}

// unreachable

func toPbActions(actions []*primitive.Action) []*pb.Action { _ = "STUB: not implemented"; return nil }

func toPbCommand(action *primitive.Action) *pb.Action { _ = "STUB: not implemented"; return nil }

func toPbTemplateType(t primitive.TemplateType) pb.TemplateType {
	_ = "STUB: not implemented"
	return *new(pb.TemplateType)
}

// unreachable

func ToPbTransformer(transformer *primitive.Transformer) *pb.Transformer {
	_ = "STUB: not implemented"
	return nil
}
