package plutoservice

import (
	"amusingx.fit/amusingx/protos/amusingplutoservice/plutoservice"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
)

func (s *PlutoService) InventoryLock(ctx context.Context, request *plutoservice.InventoryLockRequest) (*plutoservice.InventoryLockResponse, error) {
	if request == nil {
		return nil, xerror.NewError(nil, xerror.Code.CParamsError, "请求参数为空")
	}

	err := validParams(request)
	if err != nil {
		return nil, err
	}

	err = inventoryLock(ctx, request)
	if err != nil {
		return nil, err
	}

	return &plutoservice.InventoryLockResponse{
		Succeed: true,
		Message: "库存占用成功",
	}, nil
}

func validParams(request *plutoservice.InventoryLockRequest) *xerror.Error {
	if request.Count <= 0 {
		return xerror.NewError(nil, xerror.Code.CParamsError, "count 参数无效")
	}

	if request.Id <= 0 {
		return xerror.NewError(nil, xerror.Code.CParamsError, "id 参数无效")
	}

	logger.Infof("参数验证通过")

	return nil
}

func inventoryLock(ctx context.Context, req *plutoservice.InventoryLockRequest) *xerror.Error {
	return nil
}
