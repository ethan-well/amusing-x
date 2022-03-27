package subproduct

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerList(ctx context.Context, in *proto.SubProductListRequest) (*proto.SubProductListResponse, aerror.Error) {
	if in == nil {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "request is invalid")
	}

	total, products, err := model.SubProductSearch(ctx, in.Query, (in.Page-1)*in.Limit, in.Limit)
	if err != nil {
		return nil, err
	}

	var productList []*proto.SubProduct
	for _, p := range products {
		productList = append(productList, &proto.SubProduct{
			ID:        p.ID,
			Name:      p.Name,
			Desc:      p.Desc,
			ProductId: p.ProductId,
			Currency:  p.Currency,
			Price:     p.Price,
			Stock:     p.Stock,
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
