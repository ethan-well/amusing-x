package subproduct

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	charon2 "amusingx.fit/amusingx/services/charon/mysql/charon"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"amusingx.fit/amusingx/services/charon/uploader"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerQuery(ctx context.Context, in *proto.SubProductRequest) (*proto.SubProduct, aerror.Error) {
	tx, e := charon2.CharonDB.Beginx()
	if e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "sql beginx failed")
	}
	defer tx.Rollback()

	product, err := model.SubProductWithStockQueryByIdWithTx(ctx, in.Id, tx)
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

	pictures, err := model.ProductImageQueryByProductIdAndLevelWithTx(ctx, product.ID, charon.ProductLevelSubProduct, tx)
	if err != nil {
		return nil, err
	}
	if e := tx.Commit(); e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "commit failed")
	}

	var pictureList []*proto.Picture
	for _, p := range pictures {
		uploader, err := uploader.NewUploader(p.UploaderType)
		if err != nil {
			return nil, err
		}

		image, err := uploader.GetPictureInfo(ctx, p.Id)
		if err != nil {
			return nil, err
		}

		pictureList = append(pictureList, &proto.Picture{
			RawFile:    nil,
			Src:        image.Url,
			Title:      p.Title,
			Id:         p.Id,
			Url:        image.Url,
			UploadType: int64(p.UploaderType),
		})
	}

	return &proto.SubProduct{
		Id:                 product.ID,
		Name:               product.Name,
		Desc:               product.Desc,
		ProductId:          product.ProductId,
		Currency:           product.Currency,
		Price:              product.Price,
		AttributeId:        attrIds,
		Pictures:           pictureList,
		AvailableInventory: product.AvailableInventory,
		RealInventory:      product.RealInventory,
	}, nil
}
