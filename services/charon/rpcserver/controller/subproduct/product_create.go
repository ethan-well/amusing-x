package subproduct

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	charon2 "amusingx.fit/amusingx/services/charon/mysql/charon"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerCreate(ctx context.Context, in *proto.SubProductCreateRequest) (*proto.SubProduct, aerror.Error) {
	if in.Name == "" {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "name is nil")
	}
	tx, e := charon2.CharonDB.Beginx()
	if e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "sql Beginx failed")
	}
	defer tx.Rollback()

	product, err := model.SubProductInsert(ctx, &charon.SubProduct{
		Name:      in.Name,
		Desc:      in.Desc,
		ProductId: in.ProductId,
		Currency:  in.Currency,
		Price:     in.Price,
	}, tx)

	if err != nil {
		return nil, err
	}

	_, err = model.ProductStockInsert(ctx, &charon.ProductStock{
		SubProductId:       product.ID,
		RealInventory:      in.Stock,
		AvailableInventory: in.Stock,
	})

	return &proto.SubProduct{
		Id:        product.ID,
		Name:      product.Name,
		Desc:      product.Desc,
		ProductId: product.ProductId,
		Currency:  product.Currency,
		Price:     product.Price,
	}, nil
}

func HandlerCreateV2(ctx context.Context, in *proto.SubProductCreateRequest) (*proto.SubProduct, aerror.Error) {
	if in.Name == "" {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "name is nil")
	}

	tx, e := charon2.CharonDB.Beginx()
	if e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "sql execute error")
	}
	defer tx.Rollback()

	product, err := model.SubProductInsertWithTx(ctx, &charon.SubProduct{
		Name:      in.Name,
		Desc:      in.Desc,
		ProductId: in.ProductId,
		Currency:  in.Currency,
		Price:     in.Price,
	}, tx)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return &proto.SubProduct{
		Id:        product.ID,
		Name:      product.Name,
		Desc:      product.Desc,
		ProductId: product.ProductId,
		Currency:  product.Currency,
		Price:     product.Price,
	}, nil
}
