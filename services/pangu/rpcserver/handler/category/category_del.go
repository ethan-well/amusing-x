package category

import (
	charonservice "amusingx.fit/amusingx/protos/charons/service"
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu"
	"amusingx.fit/amusingx/rpcclient/charonclient"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerCategoryDelete(ctx context.Context, req *panguservice.CategoryDeleteRequest) (*panguservice.CategoryDeleteResponse, aerror.Error) {
	if len(req.Ids) == 0 {
		return &panguservice.CategoryDeleteResponse{Result: true}, nil
	}

	resp, err := charonclient.Client.Delete(ctx, &charonservice.CategoryDeleteRequest{Ids: req.Ids})
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "delete category failed")
	}

	return &panguservice.CategoryDeleteResponse{Result: resp.Succeed}, nil
}
