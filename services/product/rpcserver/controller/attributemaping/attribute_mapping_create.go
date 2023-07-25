package attributemaping

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/product/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerCreate(ctx context.Context, in *proto.AttributeMappingCreateRequest) (*proto.AttributeMapping, aerror.Error) {
	attr, err := model.AttributeMappingInsert(ctx, &charon.AttributeMapping{
		AttrId:       in.AttrId,
		AttrValue:    in.AttrValue,
		SubProductId: in.SubProductId,
	})
	if err != nil {
		return nil, err
	}

	return &proto.AttributeMapping{
		Id:           attr.ID,
		AttrId:       attr.AttrId,
		SubProductId: attr.SubProductId,
		AttrValue:    attr.AttrValue,
	}, nil
}
