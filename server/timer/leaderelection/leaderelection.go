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

//go:generate mockgen -source=leaderelection.go -destination=mock_leaderelection.go -package=leaderelection
package leaderelection

import (
	"context"
	"sync"

	v3client "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"go.uber.org/atomic"
)

const (
	dialTimeout          = 5
	dialKeepAliveTime    = 1
	dialKeepAliveTimeout = 3
	acquireLockDuration  = 5
)

var (
	newV3Client = v3client.New
	newSession  = concurrency.NewSession
	newMutex    = concurrency.NewMutex
)

type Manager interface {
	Start(ctx context.Context, callbacks LeaderCallbacks) error
	Stop(ctx context.Context) error
	// IsLeader() bool
}

type Mutex interface {
	TryLock(ctx context.Context) error
	Unlock(ctx context.Context) error
}

type leaderElection struct {
	name          string
	resourceLock  string
	leaseDuration int64
	isLeader      atomic.Bool

	etcdClient *v3client.Client
	callbacks  LeaderCallbacks
	session    *concurrency.Session
	mutex      Mutex
	mu         sync.RWMutex
	wg         sync.WaitGroup
}

type LeaderCallbacks struct {
	// OnStartedLeading is called when starts leading
	OnStartedLeading func(context.Context)
	// OnStoppedLeading is called when stops leading
	OnStoppedLeading func(context.Context)
}

func NewLeaderElection(c *Config) Manager { _ = "STUB: not implemented"; return *new(Manager) }

func (le *leaderElection) Start(ctx context.Context, callbacks LeaderCallbacks) error {
	_ = "STUB: not implemented"
	return nil
}

func (le *leaderElection) Stop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// func (le *leaderElection) IsLeader() bool {
// 	return le.isLeader
// }

func (le *leaderElection) tryAcquireLockLoop(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// refresh session until success

func (le *leaderElection) tryLock(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (le *leaderElection) release(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (le *leaderElection) refresh(_ context.Context) bool { _ = "STUB: not implemented"; return false }
