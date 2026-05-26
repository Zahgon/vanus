package cluster

import (
	"context"
	"time"

	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	metapb "github.com/vanus-labs/vanus/api/meta"

	"github.com/vanus-labs/vanus/api/cluster/raw_client"
)

type triggerService struct {
	client ctrlpb.TriggerControllerClient
}

func newTriggerService(cc *raw_client.Conn) TriggerService {
	_ = "STUB: not implemented"
	return *new(TriggerService)
}

func (es *triggerService) RawClient() ctrlpb.TriggerControllerClient {
	_ = "STUB: not implemented"
	return *new(ctrlpb.TriggerControllerClient)
}

func (es *triggerService) RegisterHeartbeat(ctx context.Context, interval time.Duration, reqFunc func() interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (es *triggerService) GetSubscription(ctx context.Context, id uint64) (*metapb.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
