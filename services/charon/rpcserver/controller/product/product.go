package product

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerQuery(ctx context.Context, in *proto.ProductRequest) (*proto.Product, aerror.Error) {
	product, err := model.ProductWideInfoById(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &proto.Product{
		ID:   product.ID,
		Name: product.Name,
		Desc: product.Desc,
	}, nil
}
