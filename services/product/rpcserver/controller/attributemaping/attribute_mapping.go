package attributemaping

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/product/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerQuery(ctx context.Context, in *proto.AttributeMappingRequest) (*proto.AttributeMapping, aerror.Error) {
	attr, err := model.AttributeMappingQueryById(ctx, in.Id)
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
