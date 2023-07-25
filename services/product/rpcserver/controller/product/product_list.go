package product

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/product/mysql/charon/model"
	"context"
	"encoding/json"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/aregexp"
	"strconv"
)

func HandlerList(ctx context.Context, in *proto.ProductListRequest) (*proto.ProductListResponse, aerror.Error) {
	if in == nil {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsErr, "request is invalid")
	}

	filter, err := getParams(in)
	if err != nil {
		return nil, err
	}

	total, products, err := model.ProductSearchQuery(ctx, in, filter)
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return &proto.ProductListResponse{
			Page:    in.Page,
			Limit:   in.Limit,
			Total:   total,
			HasNext: false,
			Data:    make([]*proto.Product, 1),
		}, nil
	}

	productCategoryMap, err := getProductIdCategoryMap(ctx, products)
	if err != nil {
		return nil, err
	}

	var productList []*proto.Product
	for _, p := range products {
		var categoryId int64
		category, ok := productCategoryMap[p.ID]
		if ok {
			categoryId = category.ID
		}

		productList = append(productList, &proto.Product{
			ID:         p.ID,
			Name:       p.Name,
			Desc:       p.Desc,
			CategoryId: categoryId,
		})
	}

	resp := &proto.ProductListResponse{
		Page:    in.Page,
		Limit:   in.Limit,
		Total:   total,
		HasNext: (in.Page-1)*in.Limit+int64(len(productList)) < total,
		Data:    productList,
	}

	return resp, nil
}

func getProductIdCategoryMap(ctx context.Context, products []*charon.Product) (map[int64]*charon.CategoryWide, aerror.Error) {
	var productIds []int64
	for _, p := range products {
		productIds = append(productIds, p.ID)
	}

	categories, err := model.SearchCategoryByProductIds(ctx, productIds)
	if err != nil {
		return nil, err
	}

	productIdCategoryMap := make(map[int64]*charon.CategoryWide)
	for _, c := range categories {
		productIdCategoryMap[c.ProductId] = c
	}

	return productIdCategoryMap, nil
}

func getParams(req *proto.ProductListRequest) (*charonservice.SearchFilter, aerror.Error) {
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
		return nil, aerror.NewErrorf(e, aerror.Code.CParamsErr, "get code failed")
	}

	for _, a := range filter.Id {
		id, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			return nil, aerror.NewErrorf(nil, aerror.Code.CParamsErr, "filter id to int64 failed")
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
			aErr = aerror.NewErrorf(err, aerror.Code.CParamsErr, "params 'id' is invalid")
		}
		ids = append(ids, id)
	} else {
		name = []string{query}
		desc = []string{query}
	}

	return
}
