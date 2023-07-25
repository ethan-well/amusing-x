package charonserver

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/product/rpcserver/controller/attributemaping"
	"context"
)

func (s *CharonServer) AttributeMappingCreate(ctx context.Context, in *proto.AttributeMappingCreateRequest) (*proto.AttributeMapping, error) {
	return attributemaping.HandlerCreate(ctx, in)
}

func (s *CharonServer) AttributeMappingDelete(ctx context.Context, in *proto.AttributeMappingDeleteRequest) (*proto.AttributeMappingDeleteResponse, error) {
	return attributemaping.HandlerDelete(ctx, in)
}

func (s *CharonServer) AttributeMappingsDelete(ctx context.Context, in *proto.AttributeMappingsDeleteRequest) (*proto.AttributeMappingsDeleteResponse, error) {
	return attributemaping.HandlerDeleteMany(ctx, in)
}

func (s *CharonServer) AttributeMappings(ctx context.Context, in *proto.AttributeMappingListRequest) (*proto.AttributeMappingListResponse, error) {
	return attributemaping.HandlerList(ctx, in)
}

func (s *CharonServer) AttributeMapping(ctx context.Context, in *proto.AttributeMappingRequest) (*proto.AttributeMapping, error) {
	return attributemaping.HandlerQuery(ctx, in)
}

func (s *CharonServer) AttributeMappingUpdate(ctx context.Context, in *proto.AttributeMappingUpdateRequest) (*proto.AttributeMapping, error) {
	return attributemaping.HandlerUpdate(ctx, in)
}
