package uploader

import (
	"amusingx.fit/amusingx/services/charon/mysql/charon"
	"amusingx.fit/amusingx/services/charon/uploader/comm"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestNewUploader(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	uploader, err := NewUploader(comm.UploadTypeMysql)
	if err != nil {
		t.Fatal(err)
	}

	up, err := uploader.UploadBase64(context.Background(), "dddddd", "file1")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("up: %s", logger.ToJson(up))
}
