package attribute

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

func HandlerList(ctx context.Context, in *proto.AttributeListRequest) (*proto.AttributeListResponse, aerror.Error) {
	if in == nil {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "request is invalid")
	}

	total, products, err := model.AttributeSearch(ctx, in.Query, (in.Page-1)*in.Limit, in.Limit)
	if err != nil {
		return nil, err
	}

	var attributes []*proto.Attribute
	for _, p := range products {
		attributes = append(attributes, &proto.Attribute{
			ID:   p.ID,
			Name: p.Name,
			Desc: p.Desc,
		})
	}

	resp := &proto.AttributeListResponse{
		Page:    in.Page,
		Limit:   in.Limit,
		Total:   total,
		HasNext: (in.Page-1)*in.Limit+int64(len(attributes)) < total,
		Data:    attributes,
	}

	return resp, nil
}

func HandlerListV2(ctx context.Context, in *proto.AttributeListRequest) (*proto.AttributeListResponse, aerror.Error) {
	if in == nil {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "request is invalid")
	}

	filter, err := getParams(in)
	if err != nil {
		return nil, err
	}

	total, attributes, err := model.AttributeSearchV2(ctx, in, filter)
	if err != nil {
		return nil, err
	}

	if len(attributes) == 0 {
		return &proto.AttributeListResponse{
			Page:    in.Page,
			Limit:   in.Limit,
			Total:   total,
			HasNext: false,
			Data:    make([]*proto.Attribute, 1),
		}, nil
	}

	var attributeList []*proto.Attribute
	for _, p := range attributes {
		attributeList = append(attributeList, &proto.Attribute{
			ID:   p.ID,
			Name: p.Name,
			Desc: p.Desc,
		})
	}

	resp := &proto.AttributeListResponse{
		Page:    in.Page,
		Limit:   in.Limit,
		Total:   total,
		HasNext: (in.Page-1)*in.Limit+int64(len(attributes)) < total,
		Data:    attributeList,
	}

	return resp, nil
}

func getParams(req *proto.AttributeListRequest) (*charonservice.SearchFilter, aerror.Error) {
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
