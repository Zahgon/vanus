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
	runCommitInterval = 3 * time.Second
)

type AsyncStore struct {
	store

	pending *skiplist.SkipList

	commitC chan struct{}
	closeC  chan struct{}
	doneC   chan struct{}
}

func newAsyncStore(wal *walog.WAL, committed *skiplist.SkipList, version, snapshot int64) *AsyncStore {
	_ = "STUB: not implemented"
	return nil
}

func (s *AsyncStore) Close() { _ = "STUB: not implemented"; return }

// Close WAL.

func (s *AsyncStore) Load(key []byte) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *AsyncStore) Store(_ context.Context, key []byte, value interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *AsyncStore) BatchStore(_ context.Context, kvs Ranger) { _ = "STUB: not implemented"; return }

func (s *AsyncStore) Delete(key []byte) { _ = "STUB: not implemented"; return }

func (s *AsyncStore) BatchDelete(keys [][]byte) { _ = "STUB: not implemented"; return }

func (s *AsyncStore) set(kvs Ranger) error { _ = "STUB: not implemented"; return nil }

func (s *AsyncStore) tryCommit() { _ = "STUB: not implemented"; return }

func (s *AsyncStore) needCommit() bool {
	_ = "STUB: not implemented"
	// TODO(james.yin): commit condition
	return false
}

func (s *AsyncStore) runCommit() { _ = "STUB: not implemented"; return }

func (s *AsyncStore) commit() { _ = "STUB: not implemented"; return }

// Marshal changed data.

// Update state.

// Write WAL.

func RecoverAsyncStore(ctx context.Context, dir string, opts ...walog.Option) (*AsyncStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
