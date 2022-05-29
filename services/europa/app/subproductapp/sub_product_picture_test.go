package subproductapp

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/rpcclient/charonclient"
	"amusingx.fit/amusingx/services/pangu/conf"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestSubProductPictureDomain_GetSubProducts(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip testing.")
	}

	conf.Mock()

	err := charonclient.InitClient(conf.Conf.GrpcClient.Charon.Addr)
	if err != nil {
		panic(err)
	}

	req := &proto.SubProductPicturesRequest{SubProductIds: []int64{9}}
	domain := NewSubProductPictureDomain(req)
	pictures, err := domain.GetSubProductPictures(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("pictures: %s", logger.ToJson(pictures))
}
