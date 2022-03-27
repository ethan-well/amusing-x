package subproduct

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerDelete(ctx context.Context, in *proto.SubProductDeleteRequest) (*proto.SubProductDeleteResponse, aerror.Error) {
	if in.Id <= 0 {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "id is 0")
	}

	err := model.SubProductDelete(ctx, []int64{in.Id})
	if err != nil {
		return nil, err
	}

	return &proto.SubProductDeleteResponse{Result: true}, nil
}
