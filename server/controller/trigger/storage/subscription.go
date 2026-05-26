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

//go:generate mockgen -source=subscription.go -destination=mock_subscription.go -package=storage
package storage

import (
	"context"

	vanus "github.com/vanus-labs/vanus/api/vsr"

	"github.com/vanus-labs/vanus/pkg/kv"
	"github.com/vanus-labs/vanus/server/controller/trigger/metadata"
)

type SubscriptionStorage interface {
	CreateSubscription(ctx context.Context, sub *metadata.Subscription) error
	UpdateSubscription(ctx context.Context, sub *metadata.Subscription) error
	DeleteSubscription(ctx context.Context, id vanus.ID) error
	GetSubscription(ctx context.Context, id vanus.ID) (*metadata.Subscription, error)
	ListSubscription(ctx context.Context) ([]*metadata.Subscription, error)
}

type subscriptionStorage struct {
	client kv.Client
}

func NewSubscriptionStorage(client kv.Client) SubscriptionStorage {
	_ = "STUB: not implemented"
	return *new(SubscriptionStorage)
}

func (s *subscriptionStorage) getKey(subID vanus.ID) string { _ = "STUB: not implemented"; return "" }

func (s *subscriptionStorage) CreateSubscription(ctx context.Context, sub *metadata.Subscription) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *subscriptionStorage) UpdateSubscription(ctx context.Context, sub *metadata.Subscription) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *subscriptionStorage) DeleteSubscription(ctx context.Context, id vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *subscriptionStorage) GetSubscription(ctx context.Context, id vanus.ID) (*metadata.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *subscriptionStorage) ListSubscription(ctx context.Context) ([]*metadata.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
