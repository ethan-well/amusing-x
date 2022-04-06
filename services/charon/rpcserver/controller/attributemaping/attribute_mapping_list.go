package attributemaping

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerList(ctx context.Context, in *proto.AttributeMappingListRequest) (*proto.AttributeMappingListResponse, aerror.Error) {
	if in == nil {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "request is invalid")
	}

	total, products, err := model.AttributeMappingSearch(ctx, in.Query, (in.Page-1)*in.Limit, in.Limit)
	if err != nil {
		return nil, err
	}

	var productList []*proto.AttributeMapping
	for _, p := range products {
		productList = append(productList, &proto.AttributeMapping{
			Id:           p.ID,
			AttrId:       p.AttrId,
			SubProductId: p.SubProductId,
			AttrValue:    p.AttrValue,
		})
	}

	resp := &proto.AttributeMappingListResponse{
		Page:    in.Page,
		Limit:   in.Limit,
		Total:   total,
		HasNext: (in.Page-1)*in.Limit+int64(len(productList)) < total,
		Data:    productList,
	}

	return resp, nil
}
