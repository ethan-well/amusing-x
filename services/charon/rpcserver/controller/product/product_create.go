package product

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerCreate(ctx context.Context, in *proto.ProductCreateRequest) (*proto.Product, aerror.Error) {
	if in.Name == "" {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "name is nil")
	}

	product, err := model.ProductInsert(ctx, &charon.Product{
		Name: in.Name,
		Desc: in.Desc,
	})
	if err != nil {
		return nil, err
	}

	return &proto.Product{
		ID:   product.ID,
		Name: product.Name,
		Desc: product.Desc,
	}, nil
}
