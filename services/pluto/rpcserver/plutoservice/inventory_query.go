package plutoservice

import (
	plutoservice "amusingx.fit/amusingx/protos/pluto/service"
	"amusingx.fit/amusingx/services/pluto/rpcserver/promotionalbook"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func (s *PlutoService) InventoryQuery(ctx context.Context, req *plutoservice.InventoryQueryRequest) (*plutoservice.InventoryQueryResponse, error) {
	var promotionalObj PromotionalObj

	if req == nil {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "req is nil")
	}

	if req.Id <= 0 {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "id < 0")
	}

	switch req.GetObj() {
	case plutoservice.CacheObjType_Book:
		promotionalObj = promotionalbook.NewPromotionalBook(-1)
	case plutoservice.CacheObjType_Mac:
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "unexpected obj type")
	default:
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "unexpected obj type")
	}

	iv, err := promotionalObj.InventoryQuery(ctx)
	if err != nil {
		return nil, err
	}

	return &plutoservice.InventoryQueryResponse{
		Id:                 req.Id,
		AvailableInventory: iv,
	}, err
}
