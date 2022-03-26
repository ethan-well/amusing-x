package product

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerList(ctx context.Context, in *proto.ProductListRequest) ([]*proto.Product, aerror.Error) {
	if in == nil {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "request is invalid")
	}

	products, err := model.ProductSearch(ctx, in.Query, (in.Page-1)*in.Limit, in.Limit)
	if err != nil {
		return nil, err
	}

	var productList []*proto.Product
	for _, p := range products {
		productList = append(productList, &proto.Product{
			ID:   p.ID,
			Name: p.Name,
			Desc: p.Desc,
		})
	}

	return productList, nil
}
