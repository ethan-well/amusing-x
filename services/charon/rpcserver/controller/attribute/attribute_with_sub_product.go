package attribute

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerAttributeWithSubProducts(ctx context.Context, in *proto.AttributeListRequest) (*proto.AttributeWithSubProductListResponse, aerror.Error) {
	if in == nil {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "request is invalid")
	}

	if len(in.SubProductIds) == 0 {
		return nil, nil
	}

	// 通过 sub product ids 查询
	return queryAttributeListBySubProductIds(ctx, in)
}

func queryAttributeListBySubProductIds(ctx context.Context, in *proto.AttributeListRequest) (*proto.AttributeWithSubProductListResponse, aerror.Error) {
	subWithSubProducts, err := model.AttributeQueryBySubProductIds(ctx, in.SubProductIds)
	if err != nil {
		return nil, err
	}

	var productList []*proto.AttributeWithSubProduct

	for _, p := range subWithSubProducts {
		productList = append(productList, &proto.AttributeWithSubProduct{
			ID:           p.ID,
			Name:         p.Name,
			Desc:         p.Desc,
			SubProductId: p.SubProductId,
			AttrValue:    p.AttrValue,
		})
	}

	return &proto.AttributeWithSubProductListResponse{
		Data: productList,
	}, nil
}
