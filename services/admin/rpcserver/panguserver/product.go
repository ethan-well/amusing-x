package panguserver

import (
	"amusingx.fit/amusingx/protos/comm/response"
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/admin/rpcserver/handler/Product"
	"context"
)

func (s *PanguServer) ProductCreate(ctx context.Context, in *panguservice.ProductCreateRequest) (*response.CommResponse, error) {
	resp := product.HandlerProductCreate(ctx, in)
	return resp, nil
}

func (s *PanguServer) Product(ctx context.Context, in *panguservice.ProductRequest) (*response.CommResponse, error) {
	return product.HandlerProduct(ctx, in), nil
}

func (s *PanguServer) Products(ctx context.Context, in *panguservice.ProductListRequest) (*response.CommResponse, error) {
	resp := product.HandlerProducts(ctx, in)
	return resp, nil
}

func (s *PanguServer) ProductDelete(ctx context.Context, in *panguservice.ProductDeleteRequest) (*response.CommResponse, error) {
	return product.HandlerProductDelete(ctx, in), nil
}

func (s *PanguServer) ProductsDelete(ctx context.Context, in *panguservice.ProductsDeleteRequest) (*response.CommResponse, error) {
	return product.HandlerProductsDelete(ctx, in), nil
}

func (s *PanguServer) ProductUpdate(ctx context.Context, in *panguservice.ProductUpdateRequest) (*response.CommResponse, error) {
	return product.HandlerProductUpdate(ctx, in), nil
}
