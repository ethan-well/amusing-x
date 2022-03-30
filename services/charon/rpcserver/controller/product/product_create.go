package product

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	charon2 "amusingx.fit/amusingx/services/charon/mysql/charon"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerCreate(ctx context.Context, in *proto.ProductCreateRequest) (*proto.Product, aerror.Error) {
	if in.Name == "" {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "name is nil")
	}

	product, err := model.ProductInsert(ctx, &charon.Product{
		Name: in.Name,
		Desc: in.Desc,
	})
	if err != nil {
		return nil, err
	}

	return &proto.Product{
		ID:   product.ID,
		Name: product.Name,
		Desc: product.Desc,
	}, nil
}

func HandlerCreateWithCategory(ctx context.Context, in *proto.ProductCreateRequest) (*proto.Product, aerror.Error) {
	if in.Name == "" {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "name is nil")
	}

	tx, e := charon2.CharonDB.Beginx()
	if e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "sql beginx failed")
	}
	defer tx.Rollback()

	product, err := model.ProductInsertWithTx(ctx, tx, &charon.Product{
		Name: in.Name,
		Desc: in.Desc,
	})
	if err != nil {
		return nil, err
	}

	mapping, err := model.CategoryProductMappingInsertWithTx(ctx, tx, &charon.CategoryProductMapping{
		CategoryId: in.CategoryId,
		ProductId:  product.ID,
	})
	if err != nil {
		return nil, err
	}

	if e := tx.Commit(); e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "sql commit failed")
	}

	return &proto.Product{
		ID:         product.ID,
		Name:       product.Name,
		Desc:       product.Desc,
		CategoryId: mapping.CategoryId,
	}, nil
}
