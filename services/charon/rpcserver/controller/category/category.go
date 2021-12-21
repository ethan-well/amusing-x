package category

import (
	charonservice "amusingx.fit/amusingx/protos/charons/service"
	"amusingx.fit/amusingx/services/charon/rpcserver/controller/app/categoryapp"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerCreateCategory(ctx context.Context, request *charonservice.CategoryCreateRequest) (*charonservice.CategoryCreateResponse, aerror.Error) {
	category, err := categoryapp.CreateCategory(ctx, request)
	if err != nil {
		return nil, aerror.NewErrorf(err, err.Code(), err.Message())
	}

	resp := &charonservice.CategoryCreateResponse{
		Succeed: true,
		Id:      category.ID,
		Name:    category.Name,
		Desc:    category.Desc,
	}

	return resp, nil
}
