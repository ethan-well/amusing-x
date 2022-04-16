package charon

import (
	"amusingx.fit/amusingx/services/charon/uploader/comm"
	"fmt"
	"github.com/ItsWewin/superfactory/aerror"
	"strconv"
	"strings"
	"time"
)

type ProductImage struct {
	Id           int64     `db:"id"`
	ProductId    int64     `db:"product_id"`
	ProductLevel int       `db:"product_level"`
	Url          string    `db:"url"`
	Title        string    `db:"title"`
	UploaderType int       `db:"uploader_type"`
	CreateTime   time.Time `db:"create_time"`
	UpdateTime   time.Time `db:"update_time"`
}

func MysqlUploadUrl(imageId int64) string {
	return fmt.Sprintf("product_image_base64:id:%d", imageId)
}

func (p *ProductImage) ImageBase64Id() (int64, aerror.Error) {
	arr := strings.Split(p.Url, ":")
	if len(arr) == 0 {
		return 0, aerror.NewErrorf(nil, aerror.Code.SSqlExecuteErr, "get image base64 id failed")
	}

	id, e := strconv.ParseInt(arr[len(arr)-1], 10, 64)
	if e != nil {
		return 0, aerror.NewErrorf(e, aerror.Code.BUnexpectedData, "parse int failed")
	}

	return id, nil
}

func (p *ProductImage) IsMysqlUploadType() bool {
	return p.UploaderType == comm.UploadTypeMysql
}
