package charonserver

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/rpcserver/controller/product"
	"context"
)

func ProductCreate(ctx context.Context, in *proto.ProductCreateRequest) (*proto.Product, error) {
	return product.HandlerCreate(ctx, in)
}

func ProductDelete(ctx context.Context, in *proto.ProductDeleteRequest) (*proto.ProductDeleteResponse, error) {
	return product.HandlerDelete(ctx, in)
}

func Products(ctx context.Context, in *proto.ProductListRequest) (*proto.ProductListResponse, error) {
	return product.HandlerList(ctx, in)
}

func Product(ctx context.Context, in *proto.ProductRequest) (*proto.Product, error) {
	return product.HandlerQuery(ctx, in)
}

func ProductUpdate(ctx context.Context, in *proto.ProductUpdateRequest) (*proto.Product, error) {
	return product.HandlerUpdate(ctx, in)
}
