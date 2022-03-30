package attribute

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerCreate(ctx context.Context, in *proto.AttributeCreateRequest) (*proto.Attribute, aerror.Error) {
	if in.Name == "" {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "name is nil")
	}

	product, err := model.AttributeInsert(ctx, &charon.Attribute{
		Name: in.Name,
		Desc: in.Desc,
	})
	if err != nil {
		return nil, err
	}

	return &proto.Attribute{
		ID:   product.ID,
		Name: product.Name,
		Desc: product.Desc,
	}, nil
}
