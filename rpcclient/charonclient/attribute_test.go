package charonclient

import (
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"sync"
	"testing"
	"time"
)

func TestAttribute(t *testing.T) {
	ClientMock(t)

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

	w := sync.WaitGroup{}
	w.Add(1)
	//go _ping(ctx, Client, &w)
	go _attributeWithSubProduct(ctx, Client, &w)
	w.Wait()
}

func _attributeWithSubProduct(ctx context.Context, client charonservice.CharonServClient, w *sync.WaitGroup) {
	defer w.Done()
	products, err := client.AttributesWithSubProduct(ctx, &proto.AttributeListRequest{
		SubProductIds: []int64{9},
	})
	if err != nil {
		panic(err)
	}

	logger.Infof("\nproducts: %s", logger.ToJson(products))
}
