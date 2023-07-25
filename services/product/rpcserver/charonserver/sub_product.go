package charonserver

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/product/rpcserver/controller/subproduct"
	"context"
	"github.com/ItsWewin/superfactory/logger"
)

func (s *CharonServer) SubProductCreate(ctx context.Context, in *proto.SubProductCreateRequest) (*proto.SubProduct, error) {
	return subproduct.HandlerCreate(ctx, in)
}

func (s *CharonServer) SubProductDelete(ctx context.Context, in *proto.SubProductDeleteRequest) (*proto.SubProductDeleteResponse, error) {
	return subproduct.HandlerDelete(ctx, in)
}

func (s *CharonServer) SubProductsDelete(ctx context.Context, in *proto.SubProductsDeleteRequest) (*proto.SubProductsDeleteResponse, error) {
	return subproduct.HandlerDeleteMany(ctx, in)
}

func (s *CharonServer) SubProducts(ctx context.Context, in *proto.SubProductListRequest) (*proto.SubProductListResponse, error) {
	return subproduct.HandlerList(ctx, in)
}

func (s *CharonServer) SubProduct(ctx context.Context, in *proto.SubProductRequest) (*proto.SubProduct, error) {
	return subproduct.HandlerQuery(ctx, in)
}

func (s *CharonServer) SubProductUpdate(ctx context.Context, in *proto.SubProductUpdateRequest) (*proto.SubProduct, error) {
	resp, err := subproduct.HandlerUpdate(ctx, in)
	if err != nil {
		logger.Errorf("[SubProductUpdate] err: %s", err)
		return nil, err
	}

	return resp, nil
}

func (s *CharonServer) SubProductPictures(ctx context.Context, in *proto.SubProductPicturesRequest) (*proto.SubProductPicturesResponse, error) {
	return subproduct.ProductPictures(ctx, in)
}
