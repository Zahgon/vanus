// Copyright 2023 Linkall Inc.
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

package block

import (
	// this project.
	"github.com/vanus-labs/vanus/server/store/meta"
	walog "github.com/vanus-labs/vanus/server/store/wal"
)

const (
	defaultRaftExecutorParallel      = 4
	defaultAppendExecutorParallel    = 4
	defaultCommitExecutorParallel    = 4
	defaultPersistExecutorParallel   = 4
	defaultApplyExecutorParallel     = 4
	defaultTransportExecutorParallel = 4
)

type executorConfig struct {
	raftExecutorParallel      int
	appendExecutorParallel    int
	commitExecutorParallel    int
	persistExecutorParallel   int
	applyExecutorParallel     int
	transportExecutorParallel int
}

type config struct {
	stateStore *meta.SyncStore
	hintStore  *meta.AsyncStore

	walOpts     []walog.Option
	executorCfg executorConfig

	leaderLis LeaderChangedListener
	appendLis EntryAppendedListener
}

func defaultConfig() config { _ = "STUB: not implemented"; return *new(config) }

type Option func(*config)

func makeConfig(opts ...Option) config { _ = "STUB: not implemented"; return *new(config) }

//nolint:staticcheck,revive // todo
// TODO(james.yin)

//nolint:staticcheck,revive // todo
// TODO(james.yin)

func WithStateStore(stateStore *meta.SyncStore) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithHintStore(hintStore *meta.AsyncStore) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithWALOptions(opts ...walog.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRaftExecutorParallel(parallel int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithAppendExecutorParallel(parallel int) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithCommitExecutorParallel(parallel int) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithPersistExecutorParallel(parallel int) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithApplyExecutorParallel(parallel int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTransportExecutorParallel(parallel int) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLeaderChangedListener(lis LeaderChangedListener) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithEntryAppendedListener(lis EntryAppendedListener) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
