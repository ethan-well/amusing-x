package category

import (
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/rpcclient/charonclient"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerCategoryList(ctx context.Context, req *panguservice.CategoryListRequest) (*panguservice.CategoryListResponse, aerror.Error) {
	resp, err := charonclient.Client.Categories(ctx, req)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "get category list failed")
	}

	var categoryList []*panguservice.Category
	for _, c := range resp.Categories {
		categoryList = append(categoryList, &panguservice.Category{
			ID:   c.ID,
			Name: c.Name,
			Desc: c.Desc,
		})
	}

	return &panguservice.CategoryListResponse{
		Page:       resp.Page,
		Limit:      resp.Limit,
		Total:      resp.Total,
		HasNext:    resp.HasNext,
		Categories: categoryList,
		Data:       categoryList,
	}, nil
}
