package charonserver

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/rpcserver/controller/subproduct"
	"context"
)

func (s *CharonServer) SubProductCreate(ctx context.Context, in *proto.SubProductCreateRequest) (*proto.SubProduct, error) {
	return subproduct.HandlerCreate(ctx, in)
}

func (s *CharonServer) SubProductDelete(ctx context.Context, in *proto.SubProductDeleteRequest) (*proto.SubProductDeleteResponse, error) {
	return subproduct.HandlerDelete(ctx, in)
}

func (s *CharonServer) SubProducts(ctx context.Context, in *proto.SubProductListRequest) (*proto.SubProductListResponse, error) {
	return subproduct.HandlerList(ctx, in)
}

func (s *CharonServer) SubProduct(ctx context.Context, in *proto.SubProductRequest) (*proto.SubProduct, error) {
	return subproduct.HandlerQuery(ctx, in)
}

func (s *CharonServer) SubProductUpdate(ctx context.Context, in *proto.SubProductUpdateRequest) (*proto.SubProduct, error) {
	return subproduct.HandlerUpdate(ctx, in)
}
