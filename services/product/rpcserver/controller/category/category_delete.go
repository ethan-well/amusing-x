package category

import (
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/product/mysql/charon/model"
	"context"
	"encoding/json"
	"github.com/ItsWewin/superfactory/aerror"
	"strconv"
)

func HandlerCategoryDelete(ctx context.Context, req *charonservice.CategoryDeleteRequest) (*charonservice.CategoryDeleteResponse, aerror.Error) {
	if req.Id == 0 {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsErr, "ids is invalid")
	}

	if err := model.Delete(ctx, []int64{req.Id}); err != nil {
		return nil, err
	}

	return &charonservice.CategoryDeleteResponse{Succeed: true}, nil
}

func HandlerCategoriesDelete(ctx context.Context, req *proto.CategoriesDeleteRequest) (*proto.CategoriesDeleteResponse, aerror.Error) {
	if req.Filter == "" {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsErr, "prams 'filter' is invalid")
	}

	var filter proto.Filter
	e := json.Unmarshal([]byte(req.Filter), &filter)
	if e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.CParamsErr, "get code failed")
	}

	var ids []int64
	for _, a := range filter.Id {
		id, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			return nil, aerror.NewErrorf(nil, aerror.Code.CParamsErr, "filter id to int64 failed")
		}
		ids = append(ids, id)
	}

	if err := model.Delete(ctx, ids); err != nil {
		return nil, err
	}

	return &proto.CategoriesDeleteResponse{Result: true, Ids: ids}, nil
}
