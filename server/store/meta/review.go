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

package meta

import (
	// standard libraries.

	"context"
	"errors"

	// this project.
	walog "github.com/vanus-labs/vanus/server/store/wal"
)

var errFound = errors.New("found")

type ReviewWatcher = func(value interface{}, version int64)

func ReviewSyncStore(ctx context.Context, dir string, key []byte, watcher ReviewWatcher, opts ...walog.Option) error {
	_ = "STUB: not implemented"
	return nil
}

func reviewLatestSnapshot(
	_ context.Context, dir string, unmarshaler Unmarshaler, key []byte, watcher ReviewWatcher,
) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TODO(james.yin): don't skip remaind data?

//nolint:errorlint // compare to errFound is ok.
