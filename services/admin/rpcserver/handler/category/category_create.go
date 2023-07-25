package category

import (
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/rpcclient/charonclient"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerCreateCategoryCreate(ctx context.Context, req *panguservice.CategoryCreateRequest) (*panguservice.CategoryCreateResponse, aerror.Error) {
	resp, err := charonclient.Client.Create(ctx, &charonservice.CategoryCreateRequest{
		Name: req.Name,
		Desc: req.Desc,
	})

	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "create category failed")
	}

	return &panguservice.CategoryCreateResponse{
		ID:   resp.Id,
		Name: resp.Name,
		Desc: resp.Desc,
	}, nil
}
