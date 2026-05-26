// SPDX-FileCopyrightText: 2023 Linkall Inc.
//
// SPDX-License-Identifier: Apache-2.0

package snowflake

import (
	"sync"

	vanus "github.com/vanus-labs/vanus/api/vsr"
)

var lock = sync.Mutex{}

// NewTestID only used for Uint Test.
func NewTestID() vanus.ID { _ = "STUB: not implemented"; return *new(vanus.ID) }

// avoiding same id

// InitializeFake just only used for Uint Test.
func InitializeFake() { _ = "STUB: not implemented"; return }
