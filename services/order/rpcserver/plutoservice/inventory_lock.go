package plutoservice

import (
	"amusingx.fit/amusingx/protos/pluto/service"
	"amusingx.fit/amusingx/services/order/rpcserver/promotionalbook"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
)

// 占用库存
func (s *PlutoService) InventoryLock(ctx context.Context, request *plutoservice.InventoryLockRequest) (*plutoservice.InventoryLockResponse, error) {
	if request == nil {
		return nil, aerror.NewError(nil, aerror.Code.CParamsErr, "请求参数为空")
	}

	err := validParams(request)
	if err != nil {
		return nil, err
	}

	iv, err := inventoryLock(ctx, request)
	if err != nil {
		return &plutoservice.InventoryLockResponse{
			Succeed:            false,
			Message:            "库存占用失败",
			RemainingInventory: iv,
		}, err
	}

	return &plutoservice.InventoryLockResponse{
		Succeed:            true,
		Message:            "库存占用成功",
		RemainingInventory: iv,
	}, nil
}

func validParams(request *plutoservice.InventoryLockRequest) aerror.Error {
	if request.Count <= 0 {
		return aerror.NewError(nil, aerror.Code.CParamsErr, "count 参数无效")
	}

	if request.Id <= 0 {
		return aerror.NewError(nil, aerror.Code.CParamsErr, "id 参数无效")
	}

	if request.GetObj() != plutoservice.CacheObjType_Book &&
		request.GetObj() != plutoservice.CacheObjType_Mac {
		return aerror.NewError(nil, aerror.Code.CParamsErr, "object type is not support")
	}

	logger.Infof("参数验证通过: %s", logger.ToJson(request))

	return nil
}

// 返回剩余的库存
func inventoryLock(ctx context.Context, req *plutoservice.InventoryLockRequest) (int64, aerror.Error) {
	var promotionalObj PromotionalObj

	switch req.GetObj() {
	case plutoservice.CacheObjType_Book:
		promotionalObj = promotionalbook.NewPromotionalBook(req.Id)
	case plutoservice.CacheObjType_Mac:
		return 0, aerror.NewError(nil, aerror.Code.CParamsErr, "unexpected obj type")
	default:
		return 0, aerror.NewErrorf(nil, aerror.Code.CParamsErr, "unexpected obj type")
	}

	iv, err := promotionalObj.InventoryLock(ctx, int(req.Count))
	if err != nil {
		return 0, err
	}

	return iv, nil
}
