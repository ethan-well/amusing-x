package subproduct

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"encoding/json"
	"github.com/ItsWewin/superfactory/aerror"
	"strconv"
)

func HandlerDelete(ctx context.Context, in *proto.SubProductDeleteRequest) (*proto.SubProductDeleteResponse, aerror.Error) {
	if in.Id <= 0 {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "id is 0")
	}

	tx, e := charon.CharonDB.Beginx()
	if e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "sql beginx failed")
	}

	err := model.SubProductDelete(ctx, []int64{in.Id}, tx)
	if err != nil {
		return nil, err
	}

	err = model.ProductStockDeleteBySubProductIds(ctx, []int64{in.Id}, tx)
	if err != nil {
		return nil, err
	}

	if e := tx.Commit(); e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "sql commit failed")
	}

	return &proto.SubProductDeleteResponse{Result: true}, nil
}

func HandlerDeleteMany(ctx context.Context, in *proto.SubProductsDeleteRequest) (*proto.SubProductsDeleteResponse, aerror.Error) {
	if in.Filter == "" {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "prams 'filter' is invalid")
	}

	var filter proto.Filter
	e := json.Unmarshal([]byte(in.Filter), &filter)
	if e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.CParamsError, "get code failed")
	}

	var ids []int64
	for _, a := range filter.Id {
		id, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "filter id to int64 failed")
		}
		ids = append(ids, id)
	}

	tx, e := charon.CharonDB.Beginx()
	if e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "sql begin failed")
	}
	defer tx.Rollback()

	// delete product
	if err := model.SubProductDeleteWithTx(ctx, tx, ids); err != nil {
		return nil, err
	}

	// delete category_product_mapping
	if err := model.AttributeMappingDeleteBySubProductIdWithTx(ctx, tx, ids); err != nil {
		return nil, err
	}

	if err := model.ProductStockDeleteBySubProductIds(ctx, ids, tx); err != nil {
		return nil, err
	}

	if e := tx.Commit(); e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "tx.commit failed")
	}

	return &proto.SubProductsDeleteResponse{Result: true, Ids: ids}, nil
}
