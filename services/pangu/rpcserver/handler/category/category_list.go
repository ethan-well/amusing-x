package category

import (
	charonservice "amusingx.fit/amusingx/protos/charons/service"
	panguservice "amusingx.fit/amusingx/protos/pangu/service"
	"amusingx.fit/amusingx/rpcclient/charonclient"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
)

func HandlerCategoryList(ctx context.Context, req *panguservice.CategoryListRequest) (*panguservice.CategoryListResponse, aerror.Error) {
	resp, err := charonclient.Client.Categories(ctx, &charonservice.CategoryListRequest{
		Query: req.Query,
		Page:  req.Page,
		Limit: req.Limit,
	})
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "get category list failed")
	}

	logger.Info("categories", logger.ToJson(resp))

	var categoryList []*panguservice.Category
	for _, c := range resp.Categories {
		categoryList = append(categoryList, &panguservice.Category{
			ID:   c.Id,
			Name: c.Name,
			Desc: c.Desc,
		})
	}

	return &panguservice.CategoryListResponse{
		Page:       req.Page,
		Limit:      req.Limit,
		Total:      resp.Total,
		HasNext:    resp.Limit*resp.Page < resp.Total,
		Categories: categoryList,
	}, nil
}
