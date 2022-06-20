package subproduct

import (
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"encoding/json"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/aregexp"
	"github.com/ItsWewin/superfactory/currency"
	"github.com/ItsWewin/superfactory/logger"
	"strconv"
)

func HandlerList(ctx context.Context, in *proto.SubProductListRequest) (*proto.SubProductListResponse, aerror.Error) {
	if in == nil {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "request is invalid")
	}

	filter, err := getParams(in)
	if err != nil {
		logger.Errorf("[HandlerList] err: %s", err)
		return nil, err
	}

	total, products, err := model.SubProductSearchWithProductStock(ctx, in, filter)
	if err != nil {
		return nil, err
	}

	var productList []*proto.SubProduct
	for _, p := range products {
		productList = append(productList, &proto.SubProduct{
			Id:                 p.ID,
			Name:               p.Name,
			Desc:               p.Desc,
			ProductId:          p.ProductId,
			Currency:           p.Currency,
			Price:              p.Price,
			Stock:              p.Stock,
			CurrencySymbol:     currency.GetSymbol(p.Currency),
			MinNum:             p.MinNum,
			MaxNum:             p.MaxNum,
			RealInventory:      p.RealInventory,
			AvailableInventory: p.AvailableInventory,
		})
	}

	resp := &proto.SubProductListResponse{
		Page:    in.Page,
		Limit:   in.Limit,
		Total:   total,
		HasNext: (in.Page-1)*in.Limit+int64(len(productList)) < total,
		Data:    productList,
	}

	return resp, nil
}

func getParams(req *proto.SubProductListRequest) (*charonservice.SearchFilter, aerror.Error) {
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
