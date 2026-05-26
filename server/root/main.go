// SPDX-FileCopyrightText: 2023 Linkall Inc.
//
// SPDX-License-Identifier: Apache-2.0

package root

import (
	// standard libraries.
	"context"

	// third-party libraries.

	// first-party libraries.

	// this project.

	"github.com/vanus-labs/vanus/server/controller"
)

type Config = controller.Config

func loadConfig(filename string) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

func Main(configPath string) { _ = "STUB: not implemented"; return }

func MainExt(ctx context.Context, cfg Config) { _ = "STUB: not implemented"; return }

// for debug in developing stage
