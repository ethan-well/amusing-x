package product

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerQuery(ctx context.Context, in *proto.ProductRequest) (*proto.Product, aerror.Error) {
	return &proto.Product{
		ID:   0,
		Name: "",
		Desc: "",
	}, nil
}
