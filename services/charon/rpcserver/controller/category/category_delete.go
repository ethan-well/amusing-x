package category

import (
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
)

func HandlerCategoryDelete(ctx context.Context, req *charonservice.CategoryDeleteRequest) (*charonservice.CategoryDeleteResponse, aerror.Error) {
	logger.Infof("req: %s", logger.ToJson(req))

	if req.Id == 0 {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "ids is invalid")
	}

	if err := model.Delete(ctx, []int64{req.Id}); err != nil {
		return nil, err
	}

	return &charonservice.CategoryDeleteResponse{Succeed: true}, nil
}
