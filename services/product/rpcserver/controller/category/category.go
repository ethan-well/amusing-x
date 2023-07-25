package category

import (
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	"amusingx.fit/amusingx/services/product/mysql/charon/model"
	"amusingx.fit/amusingx/services/product/rpcserver/controller/app/categoryapp"
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

func HandlerCategory(ctx context.Context, req *charonservice.CategoryRequest) (*charonservice.CategoryResponse, aerror.Error) {
	if req.Id <= 0 {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsErr, "params 'id' is invalid")
	}

	category, err := model.QueryCategoryByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &charonservice.CategoryResponse{Result: true, Category: &charonservice.Category{
		Id:   category.ID,
		Name: category.Name,
		Desc: category.Desc,
	}}, nil
}
