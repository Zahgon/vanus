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

package validation

import (
	// standard libraries.
	"context"

	// third-party libraries.

	// this project.
	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	metapb "github.com/vanus-labs/vanus/api/meta"
	// this project.
)

func ValidateSubscriptionRequest(ctx context.Context, request *ctrlpb.SubscriptionRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func validateProtocol(_ context.Context, protocol metapb.Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateSinkAndProtocol(_ context.Context,
	sink string,
	protocol metapb.Protocol,
	credential *metapb.SinkCredential,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSinkCredential(ctx context.Context, sink string, credential *metapb.SinkCredential) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSubscriptionConfig(_ context.Context, cfg *metapb.SubscriptionConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateTransformer(_ context.Context, transformer *metapb.Transformer) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(james.yin): avoid to compile template

func ValidateFilterList(ctx context.Context, filters []*metapb.Filter) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateFilter(ctx context.Context, f *metapb.Filter) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCel(_ context.Context, expression string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func validateCeSQL(_ context.Context, expression string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func validateAttributeMap(attributeName string, attribute map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func hasMultipleDialects(f *metapb.Filter) bool { _ = "STUB: not implemented"; return false }
