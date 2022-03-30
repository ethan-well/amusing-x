package panguserver

import (
	"amusingx.fit/amusingx/protos/comm/response"
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/pangu/rpcserver/handler/attribute"
	"context"
)

func (s *PanguServer) AttributeCreate(ctx context.Context, in *panguservice.AttributeCreateRequest) (*response.CommResponse, error) {
	resp := attribute.HandlerAttributeCreate(ctx, in)
	return resp, nil
}

func (s *PanguServer) Attribute(ctx context.Context, in *panguservice.AttributeRequest) (*response.CommResponse, error) {
	return attribute.HandlerAttribute(ctx, in), nil
}

func (s *PanguServer) Attributes(ctx context.Context, in *panguservice.AttributeListRequest) (*response.CommResponse, error) {
	resp := attribute.HandlerAttributes(ctx, in)
	return resp, nil
}

func (s *PanguServer) AttributeDelete(ctx context.Context, in *panguservice.AttributeDeleteRequest) (*response.CommResponse, error) {
	return attribute.HandlerAttributeDelete(ctx, in), nil
}

func (s *PanguServer) AttributeUpdate(ctx context.Context, in *panguservice.AttributeUpdateRequest) (*response.CommResponse, error) {
	return attribute.HandlerAttributeUpdate(ctx, in), nil
}
