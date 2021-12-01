package plutoservice

import (
	"amusingx.fit/amusingx/protos/plutoice/plutoservice"
	"amusingx.fit/amusingx/services/pluto/rpcserver/promotionalbook"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
)

func (s *PlutoService) InventoryUnlock(ctx context.Context, req *plutoservice.InventoryUnlockRequest) (*plutoservice.InventoryUnlockResponse, error) {
	if req == nil {
		return nil, xerror.NewError(nil, xerror.Code.CParamsError, "请求参数为空")
	}

	err := validParams2(req)
	if err != nil {
		return nil, err
	}

	iv, err := inventoryUnLock(ctx, req)
	if err != nil {
		return &plutoservice.InventoryUnlockResponse{
			Succeed:            false,
			Message:            "释放库存失败",
			RemainingInventory: iv,
		}, err
	}

	return &plutoservice.InventoryUnlockResponse{
		Succeed:            true,
		Message:            "释放库存成功",
		RemainingInventory: iv,
	}, nil
}

func validParams2(request *plutoservice.InventoryUnlockRequest) *xerror.Error {
	if request.Count <= 0 {
		return xerror.NewError(nil, xerror.Code.CParamsError, "count 参数无效")
	}

	if request.Id <= 0 {
		return xerror.NewError(nil, xerror.Code.CParamsError, "id 参数无效")
	}

	if request.GetObj() != plutoservice.CacheObjType_Book &&
		request.GetObj() != plutoservice.CacheObjType_Mac {
		return xerror.NewError(nil, xerror.Code.CParamsError, "object type is not support")
	}

	logger.Infof("参数验证通过: %s", logger.ToJson(request))

	return nil
}

// 返回剩余的库存
func inventoryUnLock(ctx context.Context, req *plutoservice.InventoryUnlockRequest) (int64, *xerror.Error) {
	var promotionalObj PromotionalObj

	switch req.GetObj() {
	case plutoservice.CacheObjType_Book:
		promotionalObj = promotionalbook.NewPromotionalBook(req.Id)
	case plutoservice.CacheObjType_Mac:
		return 0, xerror.NewError(nil, xerror.Code.CParamsError, "unexpected obj type")
	default:
		return 0, xerror.NewErrorf(nil, xerror.Code.CParamsError, "unexpected obj type")
	}

	iv, err := promotionalObj.InventoryUnlock(ctx, int(req.Count))
	if err != nil {
		return 0, err
	}

	return iv, nil
}
