// Copyright 2015 The etcd Authors
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
	"io"
	"log"
	"os"
	"sync"
)

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Warning(v ...interface{})
	Warningf(format string, v ...interface{})

	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})

	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
}

func SetLogger(l Logger) { _ = "STUB: not implemented"; return }

func ResetDefaultLogger() { _ = "STUB: not implemented"; return }

func getLogger() Logger { _ = "STUB: not implemented"; return *new(Logger) }

var (
	defaultLogger = &DefaultLogger{Logger: log.New(os.Stderr, "raft", log.LstdFlags)}
	discardLogger = &DefaultLogger{Logger: log.New(io.Discard, "", 0)}
	raftLoggerMu  sync.Mutex
	raftLogger    = Logger(defaultLogger)
)

const (
	calldepth = 2
)

// DefaultLogger is a default implementation of the Logger interface.
type DefaultLogger struct {
	*log.Logger
	debug bool
}

func (l *DefaultLogger) EnableTimestamps() { _ = "STUB: not implemented"; return }

func (l *DefaultLogger) EnableDebug() { _ = "STUB: not implemented"; return }

func (l *DefaultLogger) Debug(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *DefaultLogger) Debugf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *DefaultLogger) Info(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *DefaultLogger) Infof(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *DefaultLogger) Error(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *DefaultLogger) Errorf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *DefaultLogger) Warning(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *DefaultLogger) Warningf(format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l *DefaultLogger) Fatal(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *DefaultLogger) Fatalf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *DefaultLogger) Panic(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *DefaultLogger) Panicf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func header(lvl, msg string) string { _ = "STUB: not implemented"; return "" }
