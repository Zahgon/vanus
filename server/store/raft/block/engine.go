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
	// standard libraries.
	"container/list"
	"context"
	"sync"

	// third-party libraries.
	"google.golang.org/grpc"

	// first-party libraries.

	vanus "github.com/vanus-labs/vanus/api/vsr"

	// this project.

	"github.com/vanus-labs/vanus/server/store/block"
	"github.com/vanus-labs/vanus/server/store/meta"
	"github.com/vanus-labs/vanus/server/store/raft/storage"
	"github.com/vanus-labs/vanus/server/store/raft/transport"
	walog "github.com/vanus-labs/vanus/server/store/wal"
)

type Engine struct {
	mu        sync.RWMutex
	appenders list.List

	closeC chan struct{}

	wal        *storage.WAL
	stateStore *meta.SyncStore
	hintStore  *meta.AsyncStore

	resolver *transport.SimpleResolver
	host     transport.Host

	executor *executorFactory

	leaderLis LeaderChangedListener
	appendLis EntryAppendedListener

	dir     string
	walOpts []walog.Option
}

func (e *Engine) Init(dir string, localAddr string, opts ...Option) *Engine {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) init(dir string, localAddr string, cfg config) *Engine {
	e.appenders.Init()
	e.closeC = make(chan struct{})

	e.dir = dir
	e.walOpts = cfg.walOpts
	e.leaderLis = cfg.leaderLis
	e.appendLis = cfg.appendLis

	e.stateStore = cfg.stateStore
	e.hintStore = cfg.hintStore

	e.resolver = transport.NewSimpleResolver()
	e.host = transport.NewHost(e.resolver, localAddr)

	// TODO(james.yin): lazy starting.
	e.executor = newExecutorFactory(cfg.executorCfg, true)
	go e.runTick()

	return e
}

func (e *Engine) Recover(ctx context.Context, raws map[vanus.ID]block.Raw) (map[vanus.ID]Appender, error) {
	_ = "STUB: not implemented"
	// Recover wal and raft storages.
	return nil, nil
}

// Create Appender for every Raw.

// Raft log has been compacted.

// TODO(james.yin): clean expired metadata

// Clean expired storages.

func (e *Engine) Start() { _ = "STUB: not implemented"; return }

func (e *Engine) Close(ctx context.Context) {
	_ = "STUB: not implemented"
	// Close WAL, stateStore, hintStore.
	return
}

// Make sure WAL is closed before close stateStore.

// Close grpc connections for raft.

// FIXME(james.yin): close all executors.

func (e *Engine) RegisterServer(srv *grpc.Server) { _ = "STUB: not implemented"; return }

func (e *Engine) NewAppender(ctx context.Context, raw block.Raw) (Appender, error) {
	_ = "STUB: not implemented"
	return *new(Appender), nil
}

func (e *Engine) newAppender(raw block.Raw, s *storage.Storage) Appender {
	_ = "STUB: not implemented"
	return *new(Appender)
}

func (e *Engine) runTick() { _ = "STUB: not implemented"; return }

func (e *Engine) RegisterNodeRecord(id uint64, val string) error {
	_ = "STUB: not implemented"
	return nil
}
