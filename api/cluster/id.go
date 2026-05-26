package cluster

import (
	ctrlpb "github.com/vanus-labs/vanus/api/controller"

	"github.com/vanus-labs/vanus/api/cluster/raw_client"
)

type idService struct {
	client ctrlpb.SnowflakeControllerClient
}

func newIDService(cc *raw_client.Conn) IDService { _ = "STUB: not implemented"; return *new(IDService) }

func (es *idService) RawClient() ctrlpb.SnowflakeControllerClient {
	_ = "STUB: not implemented"
	return *new(ctrlpb.SnowflakeControllerClient)
}
