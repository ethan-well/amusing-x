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
	"github.com/ItsWewin/superfactory/uploader/localuploader"
	"strings"
	"sync"
)

func HandlerUpdate(ctx context.Context, in *proto.SubProductUpdateRequest) (*proto.SubProduct, aerror.Error) {
	if in == nil || in.Id == 0 {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "id in request info is 0")
	}

	logger.Infof("in: %s", logger.ToJson(in))

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
		return nil, aerror.NewErrorf(err, err.Code(), err.Message())
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

	err = model.AttributeMappingInsertWithTx(ctx, tx, newAttrMappings)
	if err != nil {
		return nil, err
	}

	err = model.ProductImageDeleteBySubProductIdAndLevelWithTx(ctx, []int64{subProduct.ProductId}, tx)
	if err != nil {
		return nil, err
	}

	if e := tx.Commit(); e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "commit error")
	}

	pictures, err := uploadImage(ctx, subProduct, in.Pictures)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "upload image failed")
	}

	return &proto.SubProduct{
		Id:          subProduct.ID,
		Name:        subProduct.Name,
		Desc:        subProduct.Desc,
		ProductId:   subProduct.ProductId,
		Currency:    subProduct.Currency,
		Price:       subProduct.Price,
		Stock:       subProduct.Stock,
		AttributeId: in.AttributeId,
		Pictures:    pictures,
	}, nil
}

func uploadImage(ctx context.Context, subProduct *charon.SubProduct, pictures []*proto.Picture) ([]*proto.Picture, aerror.Error) {
	if len(pictures) == 0 {
		return nil, nil
	}

	w := sync.WaitGroup{}
	errChan := make(chan aerror.Error, 10)
	var pictureList []*proto.Picture
	for _, picture := range pictures {
		w.Add(1)

		go func(w *sync.WaitGroup) {
			defer w.Done()
			image, err := uploadImageAndSave(ctx, subProduct, picture)
			if err != nil {
				errChan <- err
				return
			}

			pictureList = append(pictureList, &proto.Picture{
				Url:   image.Url,
				Src:   picture.Src,
				Title: image.Title,
				Id:    image.Id,
			})

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

	return pictureList, nil
}

func uploadImageAndSave(ctx context.Context, subProduct *charon.SubProduct, picture *proto.Picture) (*charon.ProductImage, aerror.Error) {
	uploader := localuploader.NewUploader("/tmp/amusing-x/sub-product")
	filePath, err := uploader.UploadBase64(ctx, picture.Src, picture.Title)
	if err != nil {
		return nil, aerror.NewErrorf(err, err.Code(), "picture upload failed, picture: %s", logger.ToJson(picture))
	}

	return model.ProductImageInsert(ctx, &charon.ProductImage{
		ProductId:    subProduct.ID,
		ProductLevel: 2,
		Url:          filePath,
		Title:        picture.Title,
		UploaderType: "local",
	})
}
