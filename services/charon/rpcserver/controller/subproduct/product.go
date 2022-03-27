package subproduct

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerQuery(ctx context.Context, in *proto.SubProductRequest) (*proto.SubProduct, aerror.Error) {
	product, err := model.SubProductQueryById(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &proto.SubProduct{
		ID:        product.ID,
		Name:      product.Name,
		Desc:      product.Desc,
		ProductId: product.ProductId,
		Currency:  product.Currency,
		Price:     product.Price,
		Stock:     product.Stock,
	}, nil
}
