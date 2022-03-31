package product

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	charon2 "amusingx.fit/amusingx/services/charon/mysql/charon"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerUpdate(ctx context.Context, in *proto.ProductUpdateRequest) (*proto.Product, aerror.Error) {
	if in == nil || in.Id == 0 {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "id in request info is 0")
	}

	err := model.ProductUpdate(ctx, &charon.Product{
		ID:   in.Id,
		Name: in.Name,
		Desc: in.Desc,
	})
	if err != nil {
		return nil, err
	}

	product, err := model.ProductWideInfoById(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, nil
	}

	return &proto.Product{
		ID:   product.ID,
		Name: product.Name,
		Desc: product.Desc,
	}, nil
}

// HandlerUpdateWithCategory
// sql 可能需要调整
func HandlerUpdateWithCategory(ctx context.Context, in *proto.ProductUpdateRequest) (*proto.Product, aerror.Error) {
	if in == nil || in.Id == 0 {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "id in request info is 0")
	}

	tx, e := charon2.CharonDB.Beginx()
	if e != nil {
		return nil, aerror.NewErrorf(nil, aerror.Code.SSqlExecuteErr, "beginx failed")
	}
	defer tx.Rollback()

	err := model.ProductUpdateWithTx(ctx, tx, &charon.Product{
		ID:   in.Id,
		Name: in.Name,
		Desc: in.Desc,
	})
	if err != nil {
		return nil, err
	}

	err = model.CategoryProductMappingDeleteByProductIdWithTX(ctx, tx, in.Id)
	if err != nil {
		return nil, err
	}

	mapping := &charon.CategoryProductMapping{
		CategoryId: in.CategoryId,
		ProductId:  in.Id,
	}
	mapping, err = model.CategoryProductMappingInsertWithTx(ctx, tx, mapping)
	if err != nil {
		return nil, err
	}
	if e := tx.Commit(); e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "tx commit failed")
	}

	product, err := model.ProductWideInfoById(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, nil
	}

	return &proto.Product{
		ID:         product.ID,
		Name:       product.Name,
		Desc:       product.Desc,
		CategoryId: mapping.ID,
	}, nil
}
