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

//go:generate mockgen -source=eventlog.go -destination=mock_eventlog.go -package=eventlog
package eventlog

import (
	// standard libraries.
	"context"
	"sync"
	"time"

	// third-party libraries.

	"github.com/huandu/skiplist"

	// first-party libraries.

	vanus "github.com/vanus-labs/vanus/api/vsr"

	// this project.
	"github.com/vanus-labs/vanus/pkg/kv"
	"github.com/vanus-labs/vanus/server/controller/eventbus/block"
	"github.com/vanus-labs/vanus/server/controller/eventbus/metadata"
	"github.com/vanus-labs/vanus/server/controller/eventbus/volume"
)

const (
	defaultAppendableSegmentNumber     = 2
	defaultSegmentReplicaNumber        = 3
	defaultSegmentExpiredTime          = 72 * time.Hour
	defaultScaleInterval               = time.Second
	defaultCleanInterval               = time.Second
	defaultCheckExpiredSegmentInterval = time.Minute
)

type Manager interface {
	Run(ctx context.Context, kvClient kv.Client, startTask bool) error
	Stop()
	AcquireEventlog(ctx context.Context, eventbusID vanus.ID, eventbusName string) (*metadata.Eventlog, error)
	GetEventlog(ctx context.Context, id vanus.ID) *metadata.Eventlog
	DeleteEventlog(ctx context.Context, id vanus.ID)
	GetBlock(id vanus.ID) *metadata.Block
	UpdateSegmentReplicas(ctx context.Context, segID vanus.ID, term uint64) error
	GetEventlogSegmentList(elID vanus.ID) []Segment
	GetAppendableSegment(ctx context.Context, eli *metadata.Eventlog, num int) ([]Segment, error)
	UpdateSegment(ctx context.Context, m map[string][]Segment)
	GetSegmentByBlockID(block *metadata.Block) (Segment, error)
}

var mgr = &eventlogManager{
	segmentReplicaNum:           defaultSegmentReplicaNumber,
	scaleInterval:               defaultScaleInterval,
	cleanInterval:               defaultCleanInterval,
	checkSegmentExpiredInterval: defaultCheckExpiredSegmentInterval,
	segmentExpiredTime:          defaultSegmentExpiredTime,
}

type eventlogManager struct {
	allocator block.Allocator

	// string, *eventlog
	eventlogMap sync.Map

	// blockID, *metadata.Block
	globalBlockMap sync.Map

	// segmentID, *segment
	globalSegmentMap sync.Map

	volMgr   volume.Manager
	kvClient kv.Client
	cancel   func()
	mutex    sync.Mutex
	// vanus.ID *Segment
	segmentNeedBeClean          sync.Map
	segmentReplicaNum           uint
	scaleInterval               time.Duration
	cleanInterval               time.Duration
	checkSegmentExpiredInterval time.Duration
	segmentExpiredTime          time.Duration
	createSegmentMutex          sync.Mutex
}

// Make sure eventlogManager implements Manager.
var _ Manager = (*eventlogManager)(nil)

func NewManager(volMgr volume.Manager, replicaNum uint, defaultBlockSize int64) Manager {
	_ = "STUB: not implemented"
	return *new(Manager)
}

func (mgr *eventlogManager) Run(ctx context.Context, kvClient kv.Client, startTask bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (mgr *eventlogManager) Stop() { _ = "STUB: not implemented"; return }

func (mgr *eventlogManager) AcquireEventlog(
	ctx context.Context, eventbusID vanus.ID, eventbusName string,
) (*metadata.Eventlog, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mgr *eventlogManager) GetEventlog(_ context.Context, id vanus.ID) *metadata.Eventlog {
	_ = "STUB: not implemented"
	return nil
}

func (mgr *eventlogManager) getEventlog(id vanus.ID) *eventlog {
	_ = "STUB: not implemented"
	return nil
}

func (mgr *eventlogManager) DeleteEventlog(ctx context.Context, id vanus.ID) {
	_ = "STUB: not implemented"
	return
}

func (mgr *eventlogManager) GetAppendableSegment(
	ctx context.Context, eli *metadata.Eventlog, num int,
) ([]Segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// preparing to cleaning

func (mgr *eventlogManager) UpdateSegment(ctx context.Context, m map[string][]Segment) {
	_ = "STUB: not implemented"
	return
}

// iterate eventlog

// TODO(wenfeng.wang) Don't update state in isNeedUpdate, rename?

func (mgr *eventlogManager) GetEventlogSegmentList(elID vanus.ID) []Segment {
	_ = "STUB: not implemented"
	return nil
}

func (mgr *eventlogManager) GetBlock(id vanus.ID) *metadata.Block {
	_ = "STUB: not implemented"
	return nil
}

func (mgr *eventlogManager) getSegment(id vanus.ID) *Segment { _ = "STUB: not implemented"; return nil }

func (mgr *eventlogManager) UpdateSegmentReplicas(ctx context.Context, leaderID vanus.ID, term uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (mgr *eventlogManager) GetSegmentByBlockID(block *metadata.Block) (Segment, error) {
	_ = "STUB: not implemented"
	return *new(Segment), nil
}

func (mgr *eventlogManager) stop() { _ = "STUB: not implemented"; return }

func (mgr *eventlogManager) getSegmentTopology(_ context.Context, seg Segment) map[uint64]string {
	_ = "STUB: not implemented"
	return nil
}

func (mgr *eventlogManager) initializeEventlog(ctx context.Context, md *metadata.Eventlog) (*eventlog, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mgr *eventlogManager) dynamicScaleUpEventlog(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (mgr *eventlogManager) cleanAbnormalSegment(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (mgr *eventlogManager) checkSegmentExpired(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// StartOffsetInLog must be set when mark previous segment full.
// unreachable

// LastEventBornTime must be set when mark the segment full.
// unreachable

func (mgr *eventlogManager) recordMetrics(ctx context.Context) { _ = "STUB: not implemented"; return }

func (mgr *eventlogManager) createSegment(ctx context.Context, el *eventlog) (*Segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// preparing to cleaning

func (mgr *eventlogManager) generateSegment(ctx context.Context, el *eventlog) (*Segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// make sure segments of one eventlog located in one SegmentServer

func (mgr *eventlogManager) whichIsLeader(raftGroup map[uint64]*metadata.Block) (*metadata.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type eventlog struct {
	// uint64, *Segment
	segmentList *skiplist.SkipList
	md          *metadata.Eventlog
	writePtr    *Segment
	kvClient    kv.Client
	mutex       sync.RWMutex
	// Why
	segments []vanus.ID
}

// newEventlog create an object in memory. if needLoad is true, there will read metadata
// from kv store. the eventlog.segmentList should be built after call this method.
func newEventlog(ctx context.Context, md *metadata.Eventlog, kvClient kv.Client, needLoad bool) (*eventlog, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (el *eventlog) get(segID vanus.ID) *Segment { _ = "STUB: not implemented"; return nil }

func (el *eventlog) appendableSegmentNumber() int { _ = "STUB: not implemented"; return 0 }

func (el *eventlog) currentAppendableSegment() *Segment { _ = "STUB: not implemented"; return nil }

// add a segment to eventlog, the metadata of this eventlog will be updated, but the segment's metadata should be
// updated after call this method.
func (el *eventlog) add(ctx context.Context, seg *Segment) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO clean when failed

func (el *eventlog) markSegmentFull(ctx context.Context, seg *Segment) error {
	_ = "STUB: not implemented"
	// because sync.RWMutex isn't reentrant, so here have to implement *eventlog.nextOf again
	return nil
}

// TODO(wenfeng.wang) update block info at the same time

func (el *eventlog) head() *Segment { _ = "STUB: not implemented"; return nil }

// headAndNext returns copies of head and next segment in the eventlog.
func (el *eventlog) headAndNext() (*Segment, *Segment) { _ = "STUB: not implemented"; return nil, nil }

func (el *eventlog) tail() *Segment { _ = "STUB: not implemented"; return nil }

func (el *eventlog) indexAt(idx int) *Segment { _ = "STUB: not implemented"; return nil }

func (el *eventlog) size() int { _ = "STUB: not implemented"; return 0 }

func (el *eventlog) getAllSegments() []*Segment { _ = "STUB: not implemented"; return nil }

func (el *eventlog) listOfRight(seg *Segment, includeSelf bool) []Segment {
	_ = "STUB: not implemented"
	return nil
}

func (el *eventlog) listOfPrevious(seg *Segment) []*Segment {
	_ = "STUB: not implemented" //nolint:unused // ok
	return nil
}

// reverse slice

func (el *eventlog) deleteHead(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (el *eventlog) updateSegment(ctx context.Context, seg *Segment) error {
	_ = "STUB: not implemented"
	// TODO(wenfeng.wang) use TXN to make sure that update block info at the same time
	return nil
}

func (el *eventlog) lock() { _ = "STUB: not implemented"; return }

func (el *eventlog) unlock() { _ = "STUB: not implemented"; return }

func (el *eventlog) rLock() { _ = "STUB: not implemented"; return }

func (el *eventlog) rUnlock() { _ = "STUB: not implemented"; return }
