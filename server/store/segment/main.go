// SPDX-FileCopyrightText: 2023 Linkall Inc.
//
// SPDX-License-Identifier: Apache-2.0

package segment

import (
	// standard libraries.
	"context"
	// first-party libraries.
	// this project.
)

const (
	debugModeENV = "SEGMENT_SERVER_DEBUG_MODE"
)

func loadConfig(filename string) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

func Main(configPath string) { _ = "STUB: not implemented"; return }

func MainExt(ctx context.Context, cfg Config, debug bool) { _ = "STUB: not implemented"; return }
