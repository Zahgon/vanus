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

//go:generate mockgen -source=member.go -destination=mock_member.go -package=member
package member

import (
	"context"
	"errors"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"go.uber.org/atomic"
)

var (
	ErrStartEtcd            = errors.New("start etcd failed")
	ErrStartEtcdCanceled    = errors.New("etcd start canceled")
	defaultEtcdStartTimeout = time.Minute
)

type Member interface {
	Init(context.Context) error
	Start(context.Context) error
	Stop(context.Context)
	RegisterMembershipChangedProcessor(MembershipEventProcessor)
	ResignIfLeader()
	IsLeader() bool
	GetLeaderID() string
	GetLeaderAddr() string
	IsReady() bool
}

var _ Member = &member{}

func New(cfg Config) Member { _ = "STUB: not implemented"; return *new(Member) }

type LeaderInfo struct {
	LeaderID   string
	LeaderAddr string
}

type EventType string

const (
	EventBecomeLeader   EventType = "leader"
	EventBecomeFollower EventType = "follower"
)

type MembershipChangedEvent struct {
	Type EventType
}

type MembershipEventProcessor func(ctx context.Context, event MembershipChangedEvent) error

type member struct {
	cfg             Config
	client          *clientv3.Client
	resourceLockKey string
	session         *concurrency.Session
	mutex           *concurrency.Mutex
	isLeader        atomic.Bool
	handlers        []MembershipEventProcessor
	handlerMu       sync.RWMutex
	sessionMu       sync.RWMutex
	exit            chan struct{}
	isReady         atomic.Bool
	wg              sync.WaitGroup
}

const (
	dialTimeout          = 5 * time.Second
	dialKeepAliveTime    = 1 * time.Second
	dialKeepAliveTimeout = 3 * time.Second
	acquireLockDuration  = 5 * time.Second
)

func (m *member) Init(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (m *member) waitForEtcdReady(ctx context.Context, endpoints []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *member) ready(ctx context.Context, endpoints []string) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *member) Start(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (m *member) leaderElection() { _ = "STUB: not implemented"; return }

// execute after server start

// refresh session until success

func (m *member) tryLock(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (m *member) setLeader(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (m *member) refresh(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func (m *member) Stop(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *member) release(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (m *member) execHandlers(ctx context.Context, event MembershipChangedEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *member) RegisterMembershipChangedProcessor(handler MembershipEventProcessor) {
	_ = "STUB: not implemented"
	return
}

func (m *member) ResignIfLeader() {
	_ = "STUB: not implemented"
	// TODO(jiangkai)
	return
}

func (m *member) IsLeader() bool { _ = "STUB: not implemented"; return false }

func (m *member) GetLeaderID() string {
	_ = "STUB: not implemented"
	// TODO(jiangkai): maybe lookup etcd per call has low performance.
	return ""
}

func (m *member) GetLeaderAddr() string {
	_ = "STUB: not implemented"
	// TODO(jiangkai): maybe lookup etcd per call has low performance.
	return ""
}

func (m *member) IsReady() bool { _ = "STUB: not implemented"; return false }
