package category

import (
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"encoding/json"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/aregexp"
	"strconv"
)

func HandlerCategoryList(ctx context.Context, req *proto.CategoryListRequest) (*proto.CategoryListResponse, aerror.Error) {
	filter, err := getParams(req)
	if err != nil {
		return nil, err
	}

	categories, total, err := model.CategoryQuery(ctx, req, filter)
	if err != nil {
		return nil, err
	}

	var categoryList []*proto.Category
	for _, c := range categories {
		categoryList = append(categoryList, &proto.Category{
			ID:   c.ID,
			Name: c.Name,
			Desc: c.Desc,
		})
	}

	resp := &proto.CategoryListResponse{
		Page:       req.Page,
		Limit:      req.Limit,
		Total:      total,
		HasNext:    (req.Page-1)*req.Limit+int64(len(categoryList)) < total,
		Categories: categoryList,
	}

	return resp, nil
}

func getParams(req *proto.CategoryListRequest) (*charonservice.SearchFilter, aerror.Error) {
	ids, names, desc, err := getQuery(req.Query)
	if err != nil {
		return nil, err
	}

	searchFilter := &charonservice.SearchFilter{
		Id:   ids,
		Name: names,
		Desc: desc,
	}

	var filter proto.Filter
	e := json.Unmarshal([]byte(req.Filter), &filter)
	if e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.CParamsError, "get code failed")
	}

	for _, a := range filter.Id {
		id, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "filter id to int64 failed")
		}
		searchFilter.Id = append(searchFilter.Id, id)
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	req.Offset = (req.Page - 1) * req.Limit
	return searchFilter, nil
}

func getQuery(query string) (ids []int64, name []string, desc []string, aErr aerror.Error) {
	if query == "" {
		return
	}
	var err error
	var id int64
	if aregexp.IsNumber(query) {
		id, err = strconv.ParseInt(query, 10, 64)
		if err != nil {
			aErr = aerror.NewErrorf(err, aerror.Code.CParamsError, "params 'id' is invalid")
		}
		ids = append(ids, id)
	} else {
		name = []string{query}
		desc = []string{query}
	}

	return
}

type SearchFilter struct {
	id   []int64
	name []string
	desc []string
}
