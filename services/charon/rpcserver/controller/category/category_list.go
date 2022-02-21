package category

import (
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/aregexp"
	"github.com/ItsWewin/superfactory/logger"
	"strconv"
)

func HandlerCategoryList(ctx context.Context, req *charonservice.CategoryListRequest) (*charonservice.CategoryListResponse, aerror.Error) {
	id, name, desc, err := getQuery(req.Query)
	if err != nil {
		return nil, err
	}
	if req.Limit <= 0 || req.Limit >= 1000 {
		req.Limit = 100
	}
	if req.Page <= 0 {
		req.Page = 1
	}

	offset := (req.Page - 1) * req.Limit

	logger.Infof("req: %s", logger.ToJson(req))

	categories, total, err := model.CategoryQuery(ctx, id, name, desc, offset, req.Limit)
	if err != nil {
		return nil, err
	}

	var categoryList []*charonservice.Category
	for _, c := range categories {
		categoryList = append(categoryList, &charonservice.Category{
			Id:   c.ID,
			Name: c.Name,
			Desc: c.Desc,
		})
	}

	resp := &charonservice.CategoryListResponse{
		Page:       req.Page,
		Limit:      req.Limit,
		Total:      total,
		HasNext:    false, // req.Page*req.Limit < total,
		Categories: categoryList,
	}

	return resp, nil
}

func getQuery(query string) (id int64, name string, desc string, aErr aerror.Error) {
	var err error
	if aregexp.IsNumber(query) {
		id, err = strconv.ParseInt(query, 10, 64)
		if err != nil {
			aErr = aerror.NewErrorf(err, aerror.Code.CParamsError, "params 'id' is invalid")
		}
	} else {
		name = query
		desc = query
	}

	return
}
