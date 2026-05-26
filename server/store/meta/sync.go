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

package meta

import (
	// standard libraries.
	"context"
	"time"

	// third-party libraries.
	"github.com/huandu/skiplist"

	// this project.
	walog "github.com/vanus-labs/vanus/server/store/wal"
)

const (
	runSnapshotInterval = 30 * time.Second
)

type SyncStore struct {
	store

	snapshotC chan struct{}
	doneC     chan struct{}
}

func newSyncStore(wal *walog.WAL, committed *skiplist.SkipList, version, snapshot int64) *SyncStore {
	_ = "STUB: not implemented"
	return nil
}

func (s *SyncStore) Close(_ context.Context) {
	_ = "STUB: not implemented"
	// Close WAL.
	return
}

// NOTE: Can not close the snapshotC before close the WAL,
// because write to snapshotC in callback of WAL append.

func (s *SyncStore) Load(key []byte) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *SyncStore) Range(begin, end []byte, cb RangeCallback) error {
	_ = "STUB: not implemented"
	return nil
}

type StoreCallback = func(error)

func (s *SyncStore) Store(ctx context.Context, key []byte, value interface{}, cb StoreCallback) {
	_ = "STUB: not implemented"
	return
}

func (s *SyncStore) BatchStore(ctx context.Context, kvs Ranger, cb StoreCallback) {
	_ = "STUB: not implemented"
	return
}

func (s *SyncStore) Delete(ctx context.Context, key []byte, cb StoreCallback) {
	_ = "STUB: not implemented"
	return
}

func (s *SyncStore) BatchDelete(ctx context.Context, keys [][]byte, cb StoreCallback) {
	_ = "STUB: not implemented"
	return
}

func (s *SyncStore) set(ctx context.Context, kvs Ranger, cb StoreCallback) {
	_ = "STUB: not implemented"
	return
}

// Use callbacks for ordering guarantees.

// Convert ErrClosed.

// Update state.

func (s *SyncStore) runSnapshot() { _ = "STUB: not implemented"; return }

func RecoverSyncStore(ctx context.Context, dir string, opts ...walog.Option) (*SyncStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type storeFuture chan error

func newStoreFuture() storeFuture { _ = "STUB: not implemented"; return *new(storeFuture) }

func (sf storeFuture) onStored(err error) { _ = "STUB: not implemented"; return }

func (sf storeFuture) wait() error { _ = "STUB: not implemented"; return nil }

func Store(ctx context.Context, s *SyncStore, key []byte, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func BatchStore(ctx context.Context, s *SyncStore, kvs Ranger) error {
	_ = "STUB: not implemented"
	return nil
}

func Delete(ctx context.Context, s *SyncStore, key []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func BatchDelete(ctx context.Context, s *SyncStore, keys [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}
