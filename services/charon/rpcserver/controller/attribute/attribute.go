package attribute

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerQuery(ctx context.Context, in *proto.AttributeRequest) (*proto.Attribute, aerror.Error) {
	product, err := model.AttributeQueryById(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &proto.Attribute{
		ID:   product.ID,
		Name: product.Name,
		Desc: product.Desc,
	}, nil
}
