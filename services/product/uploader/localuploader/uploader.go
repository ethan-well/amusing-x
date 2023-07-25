package localuploader

import (
	"amusingx.fit/amusingx/services/product/uploader/comm"
	"context"
	"encoding/base64"
	"github.com/ItsWewin/superfactory/aerror"
	"os"
	"path"
	"strings"
)

type LocalUploader struct {
	dir        string
	uploadType int
}

func NewUploader(dir string) *LocalUploader {
	return &LocalUploader{dir: dir, uploadType: comm.UploadTypeLocal}
}

func (u *LocalUploader) UploadBase64(ctx context.Context, base64str string, fileName string) (*comm.UploadInfo, aerror.Error) {
	arr := strings.Split(base64str, ";base64,")
	if len(arr) != 2 {
		return nil, aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "base64 string is invalid")
	}

	dec, err := base64.StdEncoding.DecodeString(arr[1])
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "DecodeString failed")
	}

	dirPath := path.Join(u.dir, "pictures")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0777)
		if err != nil {
			return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "create dir failed")
		}
	}

	fileName = path.Join(dirPath, fileName)
	f, err := os.Create(fileName)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "create file failed")
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "file write failed")
	}
	if err := f.Sync(); err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "file load to disk failed")
	}

	return &comm.UploadInfo{
		Url:        fileName,
		UploadType: u.uploadType,
	}, nil
}

func (u *LocalUploader) GetPictureInfo(ctx context.Context, imageId int64) (*comm.UploadInfo, aerror.Error) {
	return nil, nil
}

func (u *LocalUploader) DeletePictureInfo(ctx context.Context, imageId []int64) aerror.Error {
	return nil
}
