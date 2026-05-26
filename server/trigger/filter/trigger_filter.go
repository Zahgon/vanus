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

package filter

import (
	ce "github.com/cloudevents/sdk-go/v2"

	primitive "github.com/vanus-labs/vanus/pkg"
)

func extractFilter(subscriptionFilter *primitive.SubscriptionFilter) Filter {
	_ = "STUB: not implemented"
	return *new(Filter)
}

func extractFilters(subscriptionFilters []*primitive.SubscriptionFilter) []Filter {
	_ = "STUB: not implemented"
	return nil
}

func GetFilter(subscriptionFilters []*primitive.SubscriptionFilter) Filter {
	_ = "STUB: not implemented"
	return *new(Filter)
}

func Run(f Filter, event ce.Event) Result { _ = "STUB: not implemented"; return *new(Result) }
