package cluster

import (
	"context"
	"time"

	ctrlpb "github.com/vanus-labs/vanus/api/controller"

	"github.com/vanus-labs/vanus/api/cluster/raw_client"
)

type segmentService struct {
	client ctrlpb.SegmentControllerClient
}

func newSegmentService(cc *raw_client.Conn) SegmentService {
	_ = "STUB: not implemented"
	return *new(SegmentService)
}

func (es *segmentService) RawClient() ctrlpb.SegmentControllerClient {
	_ = "STUB: not implemented"
	return *new(ctrlpb.SegmentControllerClient)
}

func (es *segmentService) RegisterHeartbeat(ctx context.Context, interval time.Duration, reqFunc func() interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
