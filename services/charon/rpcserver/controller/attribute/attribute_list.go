package attribute

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
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
