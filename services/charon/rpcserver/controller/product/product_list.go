package product

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerList(ctx context.Context, in *proto.ProductListRequest) (*proto.ProductListResponse, aerror.Error) {
	return &proto.ProductListResponse{
		Page:    0,
		Limit:   0,
		Total:   0,
		HasNext: false,
		Data:    nil,
	}, nil
}
