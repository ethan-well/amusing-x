package subproduct

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	charon2 "amusingx.fit/amusingx/services/charon/mysql/charon"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerQuery(ctx context.Context, in *proto.SubProductRequest) (*proto.SubProduct, aerror.Error) {
	tx, e := charon2.CharonDB.Beginx()
	if e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "sql beginx failed")
	}
	defer tx.Rollback()

	product, err := model.SubProductQueryByIdWithTx(ctx, in.Id, tx)
	if err != nil {
		return nil, err
	}

	attrMappings, err := model.AttributeMappingQueryBySubProductIDWithTx(ctx, tx, product.ID)
	if err != nil {
		return nil, err
	}

	var attrIds []int64
	for _, attr := range attrMappings {
		attrIds = append(attrIds, attr.AttrId)
	}

	return &proto.SubProduct{
		Id:          product.ID,
		Name:        product.Name,
		Desc:        product.Desc,
		ProductId:   product.ProductId,
		Currency:    product.Currency,
		Price:       product.Price,
		Stock:       product.Stock,
		AttributeId: attrIds,
	}, nil
}
