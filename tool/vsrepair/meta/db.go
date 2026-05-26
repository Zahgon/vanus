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

	"errors"

	// first-party libraries.
	"github.com/vanus-labs/vanus/pkg/raft/raftpb"
	"github.com/vanus-labs/vanus/server/store/meta"
)

var ErrNotFound = errors.New("not found")

type DB struct {
	metaStore   *meta.SyncStore
	offsetStore *meta.AsyncStore
}

type config struct {
	skipMetaStore bool
	readOnly      bool
}

type Option func(*config)

func SkipMetaStore() Option { _ = "STUB: not implemented"; return *new(Option) }

func ReadOnly() Option { _ = "STUB: not implemented"; return *new(Option) }

func Open(volumeDir string, opts ...Option) (*DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *DB) Close() { _ = "STUB: not implemented"; return }

type CompactInfo struct {
	Index uint64 `json:"Index"`
	Term  uint64 `json:"Term"`
}

type RaftDetail struct {
	ConfState raftpb.ConfState `json:"ConfState"`
	HardState raftpb.HardState `json:"HardState"`
	Commit    uint64           `json:"Commit"`
	Apply     uint64           `json:"Apply"`
	Compact   CompactInfo      `json:"Compact"`
}

func (db *DB) GetRaftDetail(node uint64) (d RaftDetail, err error) {
	_ = "STUB: not implemented"
	return *new(RaftDetail), nil
}

//nolint:errorlint // compare to ErrNotFound is ok.

//nolint:errorlint // compare to ErrNotFound is ok.

//nolint:errorlint // compare to ErrNotFound is ok.

//nolint:errorlint // compare to ErrNotFound is ok.

func (db *DB) GetConfState(node uint64) (raftpb.ConfState, error) {
	_ = "STUB: not implemented"
	return *new(raftpb.ConfState), nil
}

func (db *DB) GetHardState(node uint64) (raftpb.HardState, error) {
	_ = "STUB: not implemented"
	return *new(raftpb.HardState), nil
}

// clear commit

func (db *DB) PutHardState(node uint64, hs raftpb.HardState) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *DB) GetCommit(node uint64) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (db *DB) GetApply(node uint64) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (db *DB) PutApply(node uint64, app uint64) error { _ = "STUB: not implemented"; return nil }

func (db *DB) GetCompact(node uint64) (CompactInfo, error) {
	_ = "STUB: not implemented"
	return *new(CompactInfo), nil
}

func (db *DB) PutCompact(node uint64, info CompactInfo) error {
	_ = "STUB: not implemented"
	return nil
}
