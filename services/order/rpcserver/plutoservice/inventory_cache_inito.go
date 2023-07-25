package plutoservice

import (
	"amusingx.fit/amusingx/protos/pluto/service"
	"amusingx.fit/amusingx/services/order/rpcserver/promotionalbook"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func (s *PlutoService) InventoryCacheInit(ctx context.Context, req *plutoservice.InventoryCacheInitRequest) (*plutoservice.InventoryCacheInitResponse, error) {
	var promotionalObj PromotionalObj

	if req == nil {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsErr, "req is nil")
	}

	switch req.GetObj() {
	case plutoservice.CacheObjType_Book:
		promotionalObj = promotionalbook.NewPromotionalBook(-1)
	case plutoservice.CacheObjType_Mac:
		return &plutoservice.InventoryCacheInitResponse{Succeed: false, Message: "unsupported type"}, nil
	default:
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsErr, "unexpected obj type")
	}

	err := promotionalObj.InventoryCacheInit(ctx)
	if err != nil {
		return nil, err
	}

	return &plutoservice.InventoryCacheInitResponse{Succeed: true}, nil
}
