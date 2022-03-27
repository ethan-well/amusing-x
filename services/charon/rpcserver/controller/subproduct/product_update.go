package subproduct

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerUpdate(ctx context.Context, in *proto.SubProductUpdateRequest) (*proto.SubProduct, aerror.Error) {
	if in == nil || in.Id == 0 {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "id in request info is 0")
	}

	err := model.SubProductUpdate(ctx, &charon.SubProduct{
		ID:        in.Id,
		Name:      in.Name,
		Desc:      in.Desc,
		ProductId: in.ProductId,
		Currency:  in.Currency,
		Price:     in.Price,
		Stock:     in.Stock,
	})
	if err != nil {
		return nil, err
	}

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
