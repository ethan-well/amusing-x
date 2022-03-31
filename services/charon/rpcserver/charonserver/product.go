package charonserver

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/rpcserver/controller/product"
	"context"
)

func (s *CharonServer) ProductCreate(ctx context.Context, in *proto.ProductCreateRequest) (*proto.Product, error) {
	return product.HandlerCreateWithCategory(ctx, in)
}

func (s *CharonServer) ProductDelete(ctx context.Context, in *proto.ProductDeleteRequest) (*proto.ProductDeleteResponse, error) {
	return product.HandlerDelete(ctx, in)
}

func (s *CharonServer) ProductsDelete(ctx context.Context, in *proto.ProductsDeleteRequest) (*proto.ProductsDeleteResponse, error) {
	return product.HandlerDeleteMany(ctx, in)
}

func (s *CharonServer) Products(ctx context.Context, in *proto.ProductListRequest) (*proto.ProductListResponse, error) {
	return product.HandlerList(ctx, in)
}

func (s *CharonServer) Product(ctx context.Context, in *proto.ProductRequest) (*proto.Product, error) {
	product, err := product.HandlerQuery(ctx, in)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *CharonServer) ProductUpdate(ctx context.Context, in *proto.ProductUpdateRequest) (*proto.Product, error) {
	return product.HandlerUpdateWithCategory(ctx, in)
}
