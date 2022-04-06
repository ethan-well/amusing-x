package subproduct

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	charon2 "amusingx.fit/amusingx/services/charon/mysql/charon"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/set/intset"
)

func HandlerUpdate(ctx context.Context, in *proto.SubProductUpdateRequest) (*proto.SubProduct, aerror.Error) {
	logger.Errorf("in: %s", logger.ToJson(in))

	if in == nil || in.Id == 0 {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "id in request info is 0")
	}

	tx, e := charon2.CharonDB.Beginx()
	if e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "sql beginx failed")
	}
	defer tx.Rollback()

	// sub product update
	err := model.SubProductUpdateWithTx(ctx, tx, &charon.SubProduct{
		ID:        in.Id,
		Name:      in.Name,
		Desc:      in.Desc,
		ProductId: in.ProductId,
		Currency:  in.Currency,
		Price:     in.Price,
		Stock:     in.Stock,
	})
	if err != nil {
		return nil, err
	}

	subProduct, err := model.SubProductQueryById(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	attrs, err := model.AttributeMappingQueryBySubProductIDWithTx(ctx, tx, in.Id)
	if err != nil {
		return nil, err
	}
	var oldAttrIds []int64
	var attrIdMap = make(map[int64]*charon.AttributeMapping)
	for _, p := range attrs {
		oldAttrIds = append(oldAttrIds, p.AttrId)
		attrIdMap[p.AttrId] = p
	}

	newAttrIds := in.AttributeId

	set1 := intset.NewInt64Set(oldAttrIds...)
	set2 := intset.NewInt64Set(newAttrIds...)
	deleteSet := set1.DifferenceSet(set2)

	err = model.AttributeDeleteWithTx(ctx, tx, deleteSet)
	if err != nil {
		return nil, err
	}

	err = model.AttributeMappingDeleteBySubProductIdWithTx(ctx, tx, []int64{in.Id})
	if err != nil {
		return nil, err
	}

	var newAttrMappings []*charon.AttributeMapping
	for _, attrId := range newAttrIds {
		attr, ok := attrIdMap[attrId]
		if !ok {
			attr = &charon.AttributeMapping{
				AttrId:       attrId,
				AttrValue:    "",
				SubProductId: subProduct.ID,
			}
		}
		newAttrMappings = append(newAttrMappings, attr)
	}

	logger.Errorf("newAttrMappings: %s", logger.ToJson(newAttrMappings))

	err = model.AttributeMappingInsertWithTx(ctx, tx, newAttrMappings)
	if err != nil {
		return nil, err
	}

	if e := tx.Commit(); e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "commit error")
	}

	return &proto.SubProduct{
		ID:          subProduct.ID,
		Name:        subProduct.Name,
		Desc:        subProduct.Desc,
		ProductId:   subProduct.ProductId,
		Currency:    subProduct.Currency,
		Price:       subProduct.Price,
		Stock:       subProduct.Stock,
		AttributeId: in.AttributeId,
	}, nil
}
