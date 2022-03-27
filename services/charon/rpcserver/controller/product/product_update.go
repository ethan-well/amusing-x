package product

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerUpdate(ctx context.Context, in *proto.ProductUpdateRequest) (*proto.Product, aerror.Error) {
	if in == nil || in.Id == 0 {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "id in request info is 0")
	}

	err := model.ProductUpdate(ctx, &charon.Product{
		ID:   in.Id,
		Name: in.Name,
		Desc: in.Desc,
	})
	if err != nil {
		return nil, err
	}

	product, err := model.ProductQueryById(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, nil
	}

	return &proto.Product{
		ID:   product.ID,
		Name: product.Name,
		Desc: product.Desc,
	}, nil
}
