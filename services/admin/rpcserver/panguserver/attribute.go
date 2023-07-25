package panguserver

import (
	"amusingx.fit/amusingx/protos/comm/response"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/admin/rpcserver/handler/attribute"
	"context"
)

func (s *PanguServer) AttributeCreate(ctx context.Context, in *proto.AttributeCreateRequest) (*response.CommResponse, error) {
	resp := attribute.HandlerAttributeCreate(ctx, in)
	return resp, nil
}

func (s *PanguServer) Attribute(ctx context.Context, in *proto.AttributeRequest) (*response.CommResponse, error) {
	return attribute.HandlerAttribute(ctx, in), nil
}

func (s *PanguServer) Attributes(ctx context.Context, in *proto.AttributeListRequest) (*response.CommResponse, error) {
	resp := attribute.HandlerAttributes(ctx, in)
	return resp, nil
}

func (s *PanguServer) AttributeDelete(ctx context.Context, in *proto.AttributeDeleteRequest) (*response.CommResponse, error) {
	return attribute.HandlerAttributeDelete(ctx, in), nil
}

func (s *PanguServer) AttributesDelete(ctx context.Context, in *proto.AttributesDeleteRequest) (*response.CommResponse, error) {
	return attribute.HandlerAttributesDelete(ctx, in), nil
}

func (s *PanguServer) AttributeUpdate(ctx context.Context, in *proto.AttributeUpdateRequest) (*response.CommResponse, error) {
	return attribute.HandlerAttributeUpdate(ctx, in), nil
}
