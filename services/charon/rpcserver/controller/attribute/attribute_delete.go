package attribute

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"encoding/json"
	"github.com/ItsWewin/superfactory/aerror"
	"strconv"
)

func HandlerDelete(ctx context.Context, in *proto.AttributeDeleteRequest) (*proto.AttributeDeleteResponse, aerror.Error) {
	if in.Id <= 0 {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "id is 0")
	}

	err := model.AttributeDelete(ctx, []int64{in.Id})
	if err != nil {
		return nil, err
	}

	return &proto.AttributeDeleteResponse{Result: true}, nil
}

func HandlerDeletes(ctx context.Context, req *proto.AttributesDeleteRequest) (*proto.AttributesDeleteResponse, aerror.Error) {
	if req.Filter == "" {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "prams 'filter' is invalid")
	}

	var filter proto.Filter
	e := json.Unmarshal([]byte(req.Filter), &filter)
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

	if err := model.AttributeDeleteWithTx(ctx, tx, ids); err != nil {
		return nil, err
	}

	if err := model.AttributeMappingDeleteByAttributeIdWithTx(ctx, tx, ids); err != nil {
		return nil, err
	}

	if e := tx.Commit(); e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.CParamsError, "tx.Commit failed")
	}

	return &proto.AttributesDeleteResponse{Result: true, Ids: ids}, nil
}
