package attribute

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/product/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerUpdate(ctx context.Context, in *proto.AttributeUpdateRequest) (*proto.Attribute, aerror.Error) {
	if in == nil || in.Id == 0 {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsErr, "id in request info is 0")
	}

	err := model.AttributeUpdate(ctx, &charon.Attribute{
		ID:   in.Id,
		Name: in.Name,
		Desc: in.Desc,
	})
	if err != nil {
		return nil, err
	}

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
