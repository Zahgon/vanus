package cluster

import (
	"context"
	"sync"

	"github.com/vanus-labs/vanus/api/cluster/raw_client"
	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	"github.com/vanus-labs/vanus/api/meta"
)

var (
	systemEventbusPrefix          = "__"
	defaultSystemEventbusEventlog = 1
)

type eventbusService struct {
	client ctrlpb.EventbusControllerClient
	nsSvc  NamespaceService
	cache  sync.Map
}

func newEventbusService(cc *raw_client.Conn, svc NamespaceService) EventbusService {
	_ = "STUB: not implemented"
	return *new(EventbusService)
}

func (es *eventbusService) GetSystemEventbusByName(ctx context.Context, name string) (*meta.Eventbus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (es *eventbusService) GetEventbusByName(ctx context.Context, ns, name string) (*meta.Eventbus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// es.cache.Store(key, eb) unmask when dirty cache is resolved

func (es *eventbusService) GetEventbus(ctx context.Context, id uint64) (*meta.Eventbus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// es.cache.Store(id, eb) unmask when dirty cache is resolved

func (es *eventbusService) IsSystemEventbusExistByName(ctx context.Context, name string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (es *eventbusService) CreateSystemEventbusIfNotExist(ctx context.Context, name string, desc string) (*meta.Eventbus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (es *eventbusService) Delete(ctx context.Context, id uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (es *eventbusService) RawClient() ctrlpb.EventbusControllerClient {
	_ = "STUB: not implemented"
	return *new(ctrlpb.EventbusControllerClient)
}
