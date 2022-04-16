package attributemaping

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerUpdate(ctx context.Context, in *proto.AttributeMappingUpdateRequest) (*proto.AttributeMapping, aerror.Error) {
	if in == nil || in.Id == 0 {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "id in request info is 0")
	}

	// sub product update
	err := model.AttributeMappingUpdate(ctx, &charon.AttributeMapping{
		ID:           in.Id,
		AttrId:       in.AttrId,
		AttrValue:    in.AttrValue,
		SubProductId: in.SubProductId,
	})
	if err != nil {
		return nil, err
	}

	return &proto.AttributeMapping{
		Id:           in.Id,
		AttrId:       in.AttrId,
		SubProductId: in.SubProductId,
		AttrValue:    in.AttrValue,
	}, nil
}
