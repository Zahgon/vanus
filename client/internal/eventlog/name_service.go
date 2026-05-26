// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eventlog

import (
	// standard libraries.
	"context"

	// third-party libraries.

	// first-party libraries.

	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	metapb "github.com/vanus-labs/vanus/api/meta"
	"github.com/vanus-labs/vanus/pkg/observability/tracing"

	// this project.
	"github.com/vanus-labs/vanus/client/pkg/record"
)

func NewNameService(endpoints []string) *NameService { _ = "STUB: not implemented"; return nil }

type NameService struct {
	client ctrlpb.EventlogControllerClient
	tracer *tracing.Tracer
}

func (ns *NameService) LookupWritableSegment(ctx context.Context, logID uint64) (*record.Segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ns *NameService) LookupReadableSegments(ctx context.Context, logID uint64) ([]*record.Segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toSegments(pbs []*metapb.Segment) []*record.Segment { _ = "STUB: not implemented"; return nil }

// only return first working segment

func toSegment(segment *metapb.Segment) *record.Segment { _ = "STUB: not implemented"; return nil }

// TODO: writable
