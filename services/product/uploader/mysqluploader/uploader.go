package mysqluploader

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/services/product/mysql/charon/model"
	"amusingx.fit/amusingx/services/product/uploader/comm"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

type Uploader struct {
	uploadType int
}

func NewUploader() *Uploader {
	return &Uploader{uploadType: comm.UploadTypeMysql}
}

func (u *Uploader) UploadBase64(ctx context.Context, base64str string, fileName string) (*comm.UploadInfo, aerror.Error) {
	image := &charon.ProductImageBase64{
		Id:     0,
		Base64: base64str,
	}

	info, err := model.ProductImageBase64Insert(ctx, image)
	if err != nil || info == nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "product image base64 insert error")
	}

	return &comm.UploadInfo{
		Url:        charon.MysqlUploadUrl(info.Id),
		UploadType: u.uploadType,
	}, nil
}

func (u *Uploader) GetPictureInfo(ctx context.Context, imageId int64) (*comm.UploadInfo, aerror.Error) {
	info, err := model.ProductImageQueryById(ctx, imageId)
	if err != nil {
		return nil, err
	}

	if info == nil {
		return nil, nil
	}

	id, err := info.ImageBase64Id()
	if err != nil {
		return nil, err
	}

	image, err := model.ProductImageBase64QueryById(ctx, id)
	if err != nil || image == nil {
		return nil, err
	}

	return &comm.UploadInfo{
		Url:        image.Base64,
		UploadType: info.UploaderType,
	}, nil
}

func (u *Uploader) DeletePictureInfo(ctx context.Context, imageIds []int64) aerror.Error {
	infos, err := model.ProductImagesQueryByIds(ctx, imageIds)
	if err != nil {
		return err
	}

	var imageBase64Ids []int64
	for _, info := range infos {
		if !info.IsMysqlUploadType() {
			continue
		}

		id, err := info.ImageBase64Id()
		if err != nil {
			return err
		}

		imageBase64Ids = append(imageBase64Ids, id)
	}

	return model.ProductImageBase64Delete(ctx, imageBase64Ids)
}
