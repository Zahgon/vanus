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

package meta

import (
	// standard libraries.
	"context"
	"os"

	// third-party libraries.
	"github.com/huandu/skiplist"
)

const (
	snapshotExt         = ".snapshot"
	defaultSnapshotPrem = 0o644
	defaultDirPerm      = 0o755
)

func (s *store) tryCreateSnapshot() { _ = "STUB: not implemented"; return }

func (s *store) needCreateSnapshot() bool {
	_ = "STUB: not implemented"
	// TODO(james.yin): create snapshot condition
	return false
}

func (s *store) createSnapshot() { _ = "STUB: not implemented"; return }

// Write data to file.

// Compact expired wal.

func (s *store) resolveSnapshotPath(version int64) string { _ = "STUB: not implemented"; return "" }

func recoverLatestSnapshot(
	_ context.Context, dir string, unmarshaler Unmarshaler,
) (*skiplist.SkipList, int64, error) {
	_ = "STUB: not implemented"
	// Make sure the snapshot directory exists.
	return nil, 0, nil
}

// Delete expired snapshots.

func filterLatestSnapshot(entries []os.DirEntry) (os.DirEntry, []os.DirEntry) {
	_ = "STUB: not implemented"
	return *new(os.DirEntry), nil
}
