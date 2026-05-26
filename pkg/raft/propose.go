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

package raft

import (
	// standard libraries.
	"context"

	// this project.
	pb "github.com/vanus-labs/vanus/pkg/raft/raftpb"
)

type ProposeCallback = func(error)

type ProposeData struct {
	Type         pb.EntryType
	Data         []byte
	Callback     ProposeCallback
	NoWaitCommit bool
}

type ProposeDataOption func(cfg *ProposeData)

func Data(data []byte) ProposeDataOption { _ = "STUB: not implemented"; return *new(ProposeDataOption) }

func Callback(cb ProposeCallback) ProposeDataOption {
	_ = "STUB: not implemented"
	return *new(ProposeDataOption)
}

func NoWaitCommit() ProposeDataOption { _ = "STUB: not implemented"; return *new(ProposeDataOption) }

type ProposeOption func(cfg *ProposeData)

func WithData(opts ...ProposeDataOption) ProposeOption {
	_ = "STUB: not implemented"
	return *new(ProposeOption)
}

func Propose(ctx context.Context, n Node, opts ...ProposeOption) { _ = "STUB: not implemented"; return }

type proposeFuture chan error

func newProposeFuture() proposeFuture { _ = "STUB: not implemented"; return *new(proposeFuture) }

func (pf proposeFuture) onProposed(err error) { _ = "STUB: not implemented"; return }

func (pf proposeFuture) wait() error { _ = "STUB: not implemented"; return nil }

func Propose0(ctx context.Context, n Node, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func Propose1(ctx context.Context, n Node, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func Propose2(ctx context.Context, n Node, data []byte) { _ = "STUB: not implemented"; return }
