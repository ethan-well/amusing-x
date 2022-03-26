package product

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerCreate(ctx context.Context, in *proto.ProductCreateRequest) (*proto.Product, aerror.Error) {
	return &proto.Product{
		ID:   0,
		Name: "",
		Desc: "",
	}, nil
}
