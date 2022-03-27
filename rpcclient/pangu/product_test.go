package pangu

import (
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/pangu/conf"
	"amusingx.fit/amusingx/services/pangu/mysql/pangu"
	"context"
	"fmt"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
	"time"
)

func TestProduct(t *testing.T) {
	pangu.Mock()

	InitClient(conf.Conf.Server.GrpcServer.Address)

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

	product := &panguservice.Product{
		ID:   1,
		Name: "name1",
		Desc: "desc1",
	}

	_productCreate(ctx, Client, product, t)

}

func _productCreate(ctx context.Context, client panguservice.PanGuServiceClient, product *panguservice.Product, t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.ProductCreate(ctx, &panguservice.ProductCreateRequest{
			Name: product.Name,
			Desc: product.Desc,
		})
		if err != nil {
			t.Fatal(err)
		}

		if resp == nil {
			t.Fatal("resp is nil")
		}

		t.Logf("resp: %s", logger.ToJson(resp))
	}
}
