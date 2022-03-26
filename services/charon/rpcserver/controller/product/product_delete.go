package product

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerDelete(ctx context.Context, in *proto.ProductDeleteRequest) (*proto.ProductDeleteResponse, aerror.Error) {
	return &proto.ProductDeleteResponse{Result: true}, nil
}
