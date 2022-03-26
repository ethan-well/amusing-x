package charonserver

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/rpcserver/controller/product"
	"context"
)

func (s *CharonServer) ProductCreate(ctx context.Context, in *proto.ProductCreateRequest) (*proto.Product, error) {
	return product.HandlerCreate(ctx, in)
}

func (s *CharonServer) ProductDelete(ctx context.Context, in *proto.ProductDeleteRequest) (*proto.ProductDeleteResponse, error) {
	return product.HandlerDelete(ctx, in)
}

func (s *CharonServer) Products(ctx context.Context, in *proto.ProductListRequest) (*proto.ProductList, error) {
	products, err := product.HandlerList(ctx, in)
	if err != nil {
		return nil, err
	}
	return &proto.ProductList{Products: products}, nil
}

func (s *CharonServer) Product(ctx context.Context, in *proto.ProductRequest) (*proto.Product, error) {
	return product.HandlerQuery(ctx, in)
}

func (s *CharonServer) ProductUpdate(ctx context.Context, in *proto.ProductUpdateRequest) (*proto.Product, error) {
	return product.HandlerUpdate(ctx, in)
}
