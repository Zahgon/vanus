// SPDX-FileCopyrightText: 2023 Linkall Inc.
//
// SPDX-License-Identifier: Apache-2.0

package snowflake

import (
	// standard libraries.
	"context"
	"sync"
	"sync/atomic"

	// third-party libraries.
	"github.com/sony/sonyflake"

	// first-party libraries.

	ctrlpb "github.com/vanus-labs/vanus/api/controller"
)

var (
	generator   *snowflake
	once        sync.Once
	fake        bool
	initialized atomic.Bool
)

type snowflake struct {
	snow     *sonyflake.Sonyflake
	client   ctrlpb.SnowflakeControllerClient
	ctrlAddr []string
	n        *node
}

// Initialize refactor in the future.
func Initialize(ctx context.Context, ctrlAddr []string, n *node) error {
	_ = "STUB: not implemented"
	return nil
}

func Destroy() { _ = "STUB: not implemented"; return }
