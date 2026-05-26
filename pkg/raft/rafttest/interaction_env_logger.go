// Copyright 2019 The etcd Authors
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

package rafttest

import (
	"strings"

	"github.com/vanus-labs/vanus/pkg/raft"
)

type logLevels [6]string

var lvlNames logLevels = [...]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"}

type RedirectLogger struct {
	*strings.Builder
	Lvl int // 0 = DEBUG, 1 = INFO, 2 = WARNING, 3 = ERROR, 4 = FATAL, 5 = NONE
}

var _ raft.Logger = (*RedirectLogger)(nil)

func (l *RedirectLogger) printf(lvl int, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l *RedirectLogger) print(lvl int, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *RedirectLogger) Debug(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *RedirectLogger) Debugf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *RedirectLogger) Info(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *RedirectLogger) Infof(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *RedirectLogger) Warning(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *RedirectLogger) Warningf(format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l *RedirectLogger) Error(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *RedirectLogger) Errorf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *RedirectLogger) Fatal(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *RedirectLogger) Fatalf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *RedirectLogger) Panic(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *RedirectLogger) Panicf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }
