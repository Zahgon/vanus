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

package etcd

import (
	"context"
	"time"

	v3client "go.etcd.io/etcd/client/v3"

	kvdef "github.com/vanus-labs/vanus/pkg/kv"
)

const (
	dialTimeout          = 5 * time.Second
	dialKeepAliveTime    = 1 * time.Second
	dialKeepAliveTimeout = 3 * time.Second
)

type etcdClient3 struct {
	client    *v3client.Client
	keyPrefix string
}

func NewEtcdClientV3(endpoints []string, keyPrefix string) (kvdef.Client, error) {
	_ = "STUB: not implemented"
	return *new(kvdef.Client), nil
}

func (c *etcdClient3) Get(ctx context.Context, key string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *etcdClient3) Create(ctx context.Context, key string, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *etcdClient3) Set(ctx context.Context, key string, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *etcdClient3) Update(ctx context.Context, key string, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *etcdClient3) Exists(ctx context.Context, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *etcdClient3) SetWithTTL(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *etcdClient3) Delete(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// return ErrKeyNotFound as need .

func (c *etcdClient3) DeleteDir(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// return ErrKeyNotFound as need .

func (c *etcdClient3) List(ctx context.Context, key string) ([]kvdef.Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *etcdClient3) ListKey(ctx context.Context, path string) (map[string]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *etcdClient3) watch(
	ctx context.Context,
	key string,
	stopCh <-chan struct{},
	isTree bool,
) (chan kvdef.Pair, chan error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *etcdClient3) Watch(ctx context.Context, key string, stopCh <-chan struct{}) (chan kvdef.Pair, chan error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *etcdClient3) WatchTree(ctx context.Context, key string, stopCh <-chan struct{}) (chan kvdef.Pair, chan error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *etcdClient3) CompareAndSwap(ctx context.Context, key string, preValue, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *etcdClient3) CompareAndDelete(ctx context.Context, key string, preValue []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *etcdClient3) Close() error { _ = "STUB: not implemented"; return nil }
