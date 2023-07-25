package uploader

import (
	"amusingx.fit/amusingx/services/product/uploader/comm"
	"amusingx.fit/amusingx/services/product/uploader/localuploader"
	"amusingx.fit/amusingx/services/product/uploader/mysqluploader"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

type Uploader interface {
	UploadBase64(ctx context.Context, base64str string, fileName string) (*comm.UploadInfo, aerror.Error)
	GetPictureInfo(ctx context.Context, imageId int64) (*comm.UploadInfo, aerror.Error)
	DeletePictureInfo(ctx context.Context, imageId []int64) aerror.Error
}

func NewUploader(uploadType int) (Uploader, aerror.Error) {
	switch uploadType {
	case comm.UploadTypeLocal:
		return localuploader.NewUploader("/tmp/amusing-x"), nil
	case comm.UploadTypeMysql:
		return mysqluploader.NewUploader(), nil
	default:
		return nil, aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "uploadType: %d is invalid", uploadType)
	}
}
