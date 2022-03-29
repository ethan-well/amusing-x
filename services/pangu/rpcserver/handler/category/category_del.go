package category

import (
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/rpcclient/charonclient"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerCategoryDelete(ctx context.Context, req *panguservice.CategoryDeleteRequest) (*panguservice.CategoryDeleteResponse, aerror.Error) {
	if req.Id == 0 {
		return &panguservice.CategoryDeleteResponse{Result: true}, nil
	}

	resp, err := charonclient.Client.Delete(ctx, &charonservice.CategoryDeleteRequest{Id: req.Id})
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "delete category failed")
	}

	return &panguservice.CategoryDeleteResponse{Result: resp.Succeed}, nil
}

func HandlerCategoriesDelete(ctx context.Context, req *panguservice.CategoriesDeleteRequest) (*panguservice.CategoriesDeleteResponse, aerror.Error) {
	if req.Filter == "" {
		return &panguservice.CategoriesDeleteResponse{Result: true}, nil
	}

	resp, err := charonclient.Client.CategoriesDelete(ctx, req)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "delete category failed")
	}

	return resp, nil
}
