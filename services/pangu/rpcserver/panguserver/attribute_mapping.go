package panguserver

import (
	"amusingx.fit/amusingx/protos/comm/response"
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/pangu/rpcserver/handler/attributemapping"
	"context"
)

func (s *PanguServer) AttributeMappingCreate(ctx context.Context, in *panguservice.AttributeMappingCreateRequest) (*response.CommResponse, error) {
	resp := attributemapping.HandlerHandlerAttributeMappingCreate(ctx, in)
	return resp, nil
}
func (s *PanguServer) AttributeMapping(ctx context.Context, in *panguservice.AttributeMappingRequest) (*response.CommResponse, error) {
	return attributemapping.HandlerAttributeMapping(ctx, in), nil
}
func (s *PanguServer) AttributeMappings(ctx context.Context, in *panguservice.AttributeMappingListRequest) (*response.CommResponse, error) {
	resp := attributemapping.HandlerAttributeMappings(ctx, in)
	return resp, nil
}
func (s *PanguServer) AttributeMappingDelete(ctx context.Context, in *panguservice.AttributeMappingDeleteRequest) (*response.CommResponse, error) {
	return attributemapping.HandlerAttributeMappingDelete(ctx, in), nil
}
func (s *PanguServer) AttributeMappingsDelete(ctx context.Context, in *panguservice.AttributeMappingsDeleteRequest) (*response.CommResponse, error) {
	return attributemapping.HandlerAttributeMappingsDelete(ctx, in), nil
}
func (s *PanguServer) AttributeMappingUpdate(ctx context.Context, in *panguservice.AttributeMappingUpdateRequest) (*response.CommResponse, error) {
	return attributemapping.HandlerAttributeMappingUpdate(ctx, in), nil
}
