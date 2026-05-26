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

//go:generate mockgen -source=trigger_worker.go -destination=mock_trigger_worker.go -package=storage
package storage

import (
	"context"

	"github.com/vanus-labs/vanus/pkg/kv"
	"github.com/vanus-labs/vanus/server/controller/trigger/metadata"
)

type TriggerWorkerStorage interface {
	SaveTriggerWorker(context.Context, metadata.TriggerWorkerInfo) error
	GetTriggerWorker(ctx context.Context, id string) (*metadata.TriggerWorkerInfo, error)
	DeleteTriggerWorker(ctx context.Context, id string) error
	ListTriggerWorker(ctx context.Context) ([]*metadata.TriggerWorkerInfo, error)
}

type triggerWorkerStorage struct {
	client kv.Client
}

func NewTriggerWorkerStorage(client kv.Client) TriggerWorkerStorage {
	_ = "STUB: not implemented"
	return *new(TriggerWorkerStorage)
}

func (s *triggerWorkerStorage) getKey(id string) string { _ = "STUB: not implemented"; return "" }

func (s *triggerWorkerStorage) SaveTriggerWorker(ctx context.Context, info metadata.TriggerWorkerInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *triggerWorkerStorage) GetTriggerWorker(ctx context.Context, id string) (*metadata.TriggerWorkerInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *triggerWorkerStorage) DeleteTriggerWorker(ctx context.Context, id string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *triggerWorkerStorage) ListTriggerWorker(ctx context.Context) ([]*metadata.TriggerWorkerInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
