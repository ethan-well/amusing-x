package panguserver

import (
	"amusingx.fit/amusingx/protos/comm/response"
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/admin/rpcserver/handler/subproduct"
	"context"
)

func (s *PanguServer) SubProductCreate(ctx context.Context, in *panguservice.SubProductCreateRequest) (*response.CommResponse, error) {
	resp := subproduct.HandlerSubProductCreate(ctx, in)
	return resp, nil
}

func (s *PanguServer) SubProduct(ctx context.Context, in *panguservice.SubProductRequest) (*response.CommResponse, error) {
	return subproduct.HandlerSubProduct(ctx, in), nil
}

func (s *PanguServer) SubProducts(ctx context.Context, in *panguservice.SubProductListRequest) (*response.CommResponse, error) {
	resp := subproduct.HandlerSubProducts(ctx, in)
	return resp, nil
}

func (s *PanguServer) SubProductDelete(ctx context.Context, in *panguservice.SubProductDeleteRequest) (*response.CommResponse, error) {
	return subproduct.HandlerSubProductDelete(ctx, in), nil
}

func (s *PanguServer) SubProductsDelete(ctx context.Context, in *panguservice.SubProductsDeleteRequest) (*response.CommResponse, error) {
	return subproduct.HandlerSubProductsDelete(ctx, in), nil
}

func (s *PanguServer) SubProductUpdate(ctx context.Context, in *panguservice.SubProductUpdateRequest) (*response.CommResponse, error) {
	return subproduct.HandlerSubProductUpdate(ctx, in), nil
}
