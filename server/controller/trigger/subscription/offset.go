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

package subscription

import (
	"context"

	vanus "github.com/vanus-labs/vanus/api/vsr"

	primitive "github.com/vanus-labs/vanus/pkg"
	"github.com/vanus-labs/vanus/pkg/info"
)

func (m *manager) SaveOffset(ctx context.Context, id vanus.ID, offsets info.ListOffsetInfo, commit bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) ResetOffsetByTimestamp(ctx context.Context, id vanus.ID,
	timestamp uint64,
) (info.ListOffsetInfo, error) {
	_ = "STUB: not implemented"
	return *new(info.ListOffsetInfo), nil
}

func (m *manager) GetOffset(ctx context.Context, id vanus.ID) (info.ListOffsetInfo, error) {
	_ = "STUB: not implemented"
	return *new(info.ListOffsetInfo), nil
}

// todo filter retry and deadLetter eventlog

func (m *manager) GetOrSaveOffset(ctx context.Context, id vanus.ID) (info.ListOffsetInfo, error) {
	_ = "STUB: not implemented"
	return *new(info.ListOffsetInfo), nil
}

// get retry eb offset.

func (m *manager) GetDeadLetterOffset(ctx context.Context, id vanus.ID) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// storage offsets no exist

func (m *manager) SaveDeadLetterOffset(ctx context.Context, id vanus.ID, offset uint64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) getOffsetFromCli(ctx context.Context, eventbusID vanus.ID,
	config primitive.SubscriptionConfig,
) (info.ListOffsetInfo, error) {
	_ = "STUB: not implemented"
	return *new(info.ListOffsetInfo), nil
}

// fix offset is negative which convert to uint64 is big.
