package subproduct

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerCreate(ctx context.Context, in *proto.SubProductCreateRequest) (*proto.SubProduct, aerror.Error) {
	if in.Name == "" {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "name is nil")
	}

	product, err := model.SubProductInsert(ctx, &charon.SubProduct{
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

	return &proto.SubProduct{
		Id:        product.ID,
		Name:      product.Name,
		Desc:      product.Desc,
		ProductId: product.ProductId,
		Currency:  product.Currency,
		Price:     product.Price,
		Stock:     product.Stock,
	}, nil
}
