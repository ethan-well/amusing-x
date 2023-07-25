package subproduct

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	charon2 "amusingx.fit/amusingx/services/product/mysql/charon"
	"amusingx.fit/amusingx/services/product/mysql/charon/model"
	"amusingx.fit/amusingx/services/product/uploader"
	"amusingx.fit/amusingx/services/product/uploader/comm"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/set/intset"
	"strings"
	"sync"
)

func HandlerUpdate(ctx context.Context, in *proto.SubProductUpdateRequest) (*proto.SubProduct, aerror.Error) {
	if in == nil || in.Id == 0 {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsErr, "id in request info is 0")
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
	})
	if err != nil {
		return nil, aerror.NewErrorf(err, err.Code(), err.Message())
	}

	err = model.ProductStockUpdateWithTx(ctx, &charon.ProductStock{
		SubProductId:       in.Id,
		RealInventory:      in.RealInventory,
		AvailableInventory: in.AvailableInventory,
	}, tx)

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

	err = model.AttributeMappingInsertWithTx(ctx, tx, newAttrMappings)
	if err != nil {
		return nil, err
	}

	if e := tx.Commit(); e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "commit error")
	}

	// Todo: 存在分布式事务问题
	pictures, err := uploadImage(ctx, subProduct, in.Pictures)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "upload image failed")
	}

	return &proto.SubProduct{
		Id:                 subProduct.ID,
		Name:               subProduct.Name,
		Desc:               subProduct.Desc,
		ProductId:          subProduct.ProductId,
		Currency:           subProduct.Currency,
		Price:              subProduct.Price,
		AttributeId:        in.AttributeId,
		Pictures:           pictures,
		RealInventory:      100,
		AvailableInventory: 10,
	}, nil
}

func uploadImage(ctx context.Context, subProduct *charon.SubProduct, pictures []*proto.Picture) ([]*proto.Picture, aerror.Error) {
	//if len(pictures) == 0 {
	//	return nil, deletePicture(ctx, subProduct.ID)
	//}
	//
	err := deletePicture(ctx, subProduct.ID)
	if err != nil {
		return nil, err
	}

	w := sync.WaitGroup{}
	errChan := make(chan aerror.Error, 10)
	pictureListChan := make(chan *proto.Picture, len(pictures))
	for _, picture := range pictures {
		w.Add(1)

		go func(w *sync.WaitGroup) {
			defer w.Done()
			info, err := uploadToStore(ctx, picture)
			if err != nil {
				errChan <- err
				return
			}

			pictureListChan <- &proto.Picture{
				RawFile:    picture.RawFile,
				Src:        picture.Src,
				Title:      picture.Title,
				Url:        info.Url,
				UploadType: int64(info.UploadType),
			}
		}(&w)
	}

	w.Wait()

	close(errChan)
	var msg []string
	for e := range errChan {
		msg = append(msg, e.Message())
	}

	if len(msg) != 0 {
		return nil, aerror.NewError(nil, aerror.Code.SSqlExecuteErr, strings.Join(msg, ","))
	}

	close(pictureListChan)
	var pictureList []*proto.Picture
	for info := range pictureListChan {
		pictureList = append(pictureList, info)
	}

	_, err = savePictureInfo(ctx, subProduct, pictureList)
	if err != nil {
		return nil, err
	}

	return pictureList, nil
}

func uploadToStore(ctx context.Context, picture *proto.Picture) (*comm.UploadInfo, aerror.Error) {
	uploader, err := uploader.NewUploader(comm.UploadTypeMysql)
	if err != nil {
		return nil, err
	}

	info, err := uploader.UploadBase64(ctx, picture.Src, picture.Title)
	if err != nil {
		return nil, aerror.NewErrorf(err, err.Code(), "picture upload failed, picture: %s", logger.ToJson(picture))
	}

	return info, nil
}

func savePictureInfo(ctx context.Context, subProduct *charon.SubProduct, pictureList []*proto.Picture) ([]*charon.ProductImage, aerror.Error) {
	tx, e := charon2.CharonDB.Beginx()
	if e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "begin tx failed")
	}
	defer tx.Rollback()

	err := model.ProductImageDeleteBySubProductIdAndLevelWithTx(ctx, []int64{subProduct.ID}, tx)
	if err != nil {
		return nil, err
	}

	var images []*charon.ProductImage
	for _, picture := range pictureList {
		images = append(images, &charon.ProductImage{
			ProductId:    subProduct.ID,
			ProductLevel: charon.ProductLevelSubProduct,
			Url:          picture.Url,
			Title:        picture.Title,
			UploaderType: comm.UploadTypeMysql,
		})
	}

	err = model.ProductImagesInsertWithTx(ctx, tx, images)
	if err != nil {
		return nil, err
	}

	if e := tx.Commit(); e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "commit failed")
	}

	return model.ProductImageQueryByProductIdAndLevelWithTx(ctx, subProduct.ID, charon.ProductLevelSubProduct)
}

func deletePicture(ctx context.Context, subProductId int64) aerror.Error {
	tx, e := charon2.CharonDB.Beginx()
	if e != nil {
		return aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "begin tx failed")
	}
	defer tx.Rollback()

	pictures, err := model.ProductImageQueryByProductIdAndLevelWithTx(ctx, subProductId, charon.ProductLevelSubProduct, tx)
	if err != nil || len(pictures) == 0 {
		return err
	}

	var ids []int64
	for _, p := range pictures {
		ids = append(ids, p.Id)
	}

	err = model.ProductImageDeleteWithTx(ctx, tx, ids)
	if err != nil {
		return err
	}

	uploader, err := uploader.NewUploader(comm.UploadTypeMysql)
	if err != nil {
		return err
	}

	err = uploader.DeletePictureInfo(ctx, ids)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "DeletePictureInfo failed")
	}

	if e := tx.Commit(); e != nil {
		return aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "tx commit failed")
	}

	return nil
}
