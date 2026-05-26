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

package offset

import (
	"sync"

	"github.com/huandu/skiplist"

	vanus "github.com/vanus-labs/vanus/api/vsr"

	"github.com/vanus-labs/vanus/pkg/info"
)

func NewSubscriptionOffset(id vanus.ID, maxUACKNumber int, initOffsets info.ListOffsetInfo) *SubscriptionOffset {
	_ = "STUB: not implemented"
	return nil
}

type SubscriptionOffset struct {
	subscriptionID vanus.ID
	cond           *sync.Cond
	maxUACKNumber  int
	uACKNumber     int
	elOffsets      map[vanus.ID]*offsetTracker
	closed         bool
}

func (offset *SubscriptionOffset) Close() { _ = "STUB: not implemented"; return }

func (offset *SubscriptionOffset) EventReceive(info info.OffsetInfo) {
	_ = "STUB: not implemented"
	return
}

func (offset *SubscriptionOffset) EventCommit(info info.OffsetInfo) {
	_ = "STUB: not implemented"
	return
}

func (offset *SubscriptionOffset) GetCommit() info.ListOffsetInfo {
	_ = "STUB: not implemented"
	return *new(info.ListOffsetInfo)
}

type offsetTracker struct {
	maxOffset int64
	list      *skiplist.SkipList
}

func initOffset(initOffset uint64) *offsetTracker { _ = "STUB: not implemented"; return nil }

func (o *offsetTracker) putOffset(offset uint64) { _ = "STUB: not implemented"; return }

func (o *offsetTracker) commitOffset(offset uint64) { _ = "STUB: not implemented"; return }

func (o *offsetTracker) offsetToCommit() uint64 { _ = "STUB: not implemented"; return 0 }
