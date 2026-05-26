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

//go:generate mockgen -source=offset.go -destination=mock_offset.go -package=storage
package storage

import (
	"context"

	vanus "github.com/vanus-labs/vanus/api/vsr"

	"github.com/vanus-labs/vanus/pkg/info"
	"github.com/vanus-labs/vanus/pkg/kv"
)

type OffsetStorage interface {
	CreateOffset(ctx context.Context, subscriptionID vanus.ID, info info.OffsetInfo) error
	UpdateOffset(ctx context.Context, subscriptionID vanus.ID, info info.OffsetInfo) error
	GetOffsets(ctx context.Context, subscriptionID vanus.ID) (info.ListOffsetInfo, error)
	DeleteOffset(ctx context.Context, subscriptionID vanus.ID) error
}

var (
	base    = 10
	bitSize = 64
)

type offsetStorage struct {
	client kv.Client
}

func NewOffsetStorage(client kv.Client) OffsetStorage {
	_ = "STUB: not implemented"
	return *new(OffsetStorage)
}

func (s *offsetStorage) getKey(subscriptionID, eventlogID vanus.ID) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *offsetStorage) getSubKey(subscriptionID vanus.ID) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *offsetStorage) int64ToByteArr(v uint64) []byte {
	_ = "STUB: not implemented"
	/*
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, vanus.ID(v))
		return b
	*/return nil
}

func (s *offsetStorage) byteArrToUint64(b []byte) uint64 {
	_ = "STUB: not implemented"
	/*
		v := binary.LittleEndian.vanus.ID(b)
		return int64(v)
	*/return 0
}

func (s *offsetStorage) CreateOffset(ctx context.Context, subscriptionID vanus.ID, info info.OffsetInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *offsetStorage) UpdateOffset(ctx context.Context, subscriptionID vanus.ID, info info.OffsetInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *offsetStorage) GetOffsets(ctx context.Context, subscriptionID vanus.ID) (info.ListOffsetInfo, error) {
	_ = "STUB: not implemented"
	return *new(info.ListOffsetInfo), nil
}

func (s *offsetStorage) DeleteOffset(ctx context.Context, subscriptionID vanus.ID) error {
	_ = "STUB: not implemented"
	return nil
}
