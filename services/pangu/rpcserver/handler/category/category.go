package category

import (
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/rpcclient/charonclient"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerCategory(ctx context.Context, req *panguservice.CategoryRequest) (*panguservice.CategoryResponse, aerror.Error) {
	category, err := charonclient.Client.Category(ctx, &charonservice.CategoryRequest{Id: req.Id})
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "unexpected error")
	}

	return &panguservice.CategoryResponse{
		Result: category.Result,
		CategoryInfo: &panguservice.Category{
			ID:   category.Category.Id,
			Name: category.Category.Name,
			Desc: category.Category.Desc,
		},
	}, nil
}
