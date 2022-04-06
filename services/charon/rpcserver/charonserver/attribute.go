package charonserver

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/rpcserver/controller/attribute"
	"context"
)

func (s *CharonServer) AttributeCreate(ctx context.Context, in *proto.AttributeCreateRequest) (*proto.Attribute, error) {
	return attribute.HandlerCreate(ctx, in)
}

func (s *CharonServer) AttributeDelete(ctx context.Context, in *proto.AttributeDeleteRequest) (*proto.AttributeDeleteResponse, error) {
	return attribute.HandlerDelete(ctx, in)
}

func (s *CharonServer) AttributesDelete(ctx context.Context, in *proto.AttributesDeleteRequest) (*proto.AttributesDeleteResponse, error) {
	return attribute.HandlerDeletes(ctx, in)
}

func (s *CharonServer) Attributes(ctx context.Context, in *proto.AttributeListRequest) (*proto.AttributeListResponse, error) {
	//return attribute.HandlerList(ctx, in)
	return attribute.HandlerListV2(ctx, in)
}

func (s *CharonServer) Attribute(ctx context.Context, in *proto.AttributeRequest) (*proto.Attribute, error) {
	return attribute.HandlerQuery(ctx, in)
}

func (s *CharonServer) AttributeUpdate(ctx context.Context, in *proto.AttributeUpdateRequest) (*proto.Attribute, error) {
	return attribute.HandlerUpdate(ctx, in)
}
