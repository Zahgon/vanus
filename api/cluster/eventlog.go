package cluster

import (
	ctrlpb "github.com/vanus-labs/vanus/api/controller"

	"github.com/vanus-labs/vanus/api/cluster/raw_client"
)

type eventlogService struct {
	client ctrlpb.EventlogControllerClient
}

func newEventlogService(cc *raw_client.Conn) EventlogService {
	_ = "STUB: not implemented"
	return *new(EventlogService)
}

func (es *eventlogService) RawClient() ctrlpb.EventlogControllerClient {
	_ = "STUB: not implemented"
	return *new(ctrlpb.EventlogControllerClient)
}
