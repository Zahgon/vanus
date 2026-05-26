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

	"time"

	// third-party libraries.
	"github.com/spf13/cobra"

	// first-party libraries.
	metapb "github.com/vanus-labs/vanus/api/meta"

	// this project.
	"github.com/vanus-labs/vanus/server/store/vsb"
	"github.com/vanus-labs/vanus/tool/vsrepair/meta"
)

const (
	defaultTimeout = 10 * time.Second
)

func DescribeCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

type blockDetail struct {
	ID     uint64                    `json:"ID"`
	VSB    *vsb.Header               `json:"VSB,omitempty"`
	Raft   *meta.RaftDetail          `json:"Raft,omitempty"`
	Status *metapb.SegmentHealthInfo `json:"Status,omitempty"`
}

func describe(_ *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

//nolint:errorlint // compare to meta.ErrNotFound is ok.
