package category

import (
	charonservice "amusingx.fit/amusingx/protos/charons/service"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerCategoryDelete(ctx context.Context, req *charonservice.CategoryDeleteRequest) (*charonservice.CategoryDeleteResponse, aerror.Error) {
	if len(req.Ids) == 0 {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "ids is invalid")
	}

	if err := model.Delete(ctx, req.Ids); err != nil {
		return nil, err
	}

	return &charonservice.CategoryDeleteResponse{Succeed: true}, nil
}
