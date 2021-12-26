package category

import (
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu"
	"amusingx.fit/amusingx/rpcclient/charonclient"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerCategoryUpdate(ctx context.Context, req *panguservice.CategoryUpdateRequest) (*panguservice.CategoryUpdateResponse, aerror.Error) {
	resp, err := charonclient.Client.Update(ctx, &charonservice.CategoryUpdateRequest{
		Category: &charonservice.Category{Id: req.Category.ID, Name: req.Category.Name, Desc: req.Category.Desc}})

	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "create category failed")
	}

	return &panguservice.CategoryUpdateResponse{Result: resp.Result}, nil
}
