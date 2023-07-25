package subproduct

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/product/mysql/charon"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestProductPictures(t *testing.T) {
	if testing.Short() {
		t.Skip("skip testing")
	}

	charon.Mock()

	req := &proto.SubProductPicturesRequest{SubProductIds: []int64{9}}
	subProductPictures, err := ProductPictures(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("subProductPictures: %s", logger.ToJson(subProductPictures))
}
