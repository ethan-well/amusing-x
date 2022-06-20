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

func TestSubProduct(t *testing.T) {
	ClientMock(t)

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

	w := sync.WaitGroup{}
	w.Add(4)
	go _ping(ctx, Client, &w)
	go _SubProducts(ctx, Client, &w)
	go _SubProductPictures(ctx, Client, &w)
	go _SubProductUpdate(ctx, Client, &w)
	w.Wait()
}

func _SubProductPictures(ctx context.Context, client charonservice.CharonServClient, w *sync.WaitGroup) {
	defer w.Done()
	pictures, err := client.SubProductPictures(ctx, &proto.SubProductPicturesRequest{SubProductIds: []int64{9}})
	if err != nil {
		logger.Errorf("err: %s", err)
	}

	logger.Infof("pictures: %s", logger.ToJson(pictures))
}

func _SubProducts(ctx context.Context, client charonservice.CharonServClient, w *sync.WaitGroup) {
	defer w.Done()

	in := &proto.SubProductListRequest{
		Page:   0,
		Limit:  10,
		Sort:   "",
		Filter: "{}",
		Offset: 0,
	}

	resp, err := client.SubProducts(ctx, in)
	if err != nil {
		logger.Errorf("err: %s", err)
	}

	logger.Infof("\nresp: %s", logger.ToJson(resp))
}

func _SubProductUpdate(ctx context.Context, client charonservice.CharonServClient, w *sync.WaitGroup) {
	defer w.Done()

	in := &proto.SubProductUpdateRequest{
		Id:                 16,
		Name:               "name new",
		Desc:               "desc new",
		Currency:           "CNY",
		Price:              1111,
		Stock:              2222,
		RealInventory:      3333,
		AvailableInventory: 4444,
	}

	resp, err := client.SubProductUpdate(ctx, in)
	if err != nil {
		logger.Errorf("err: %s", err)
	}

	logger.Infof("\nresp: %s", logger.ToJson(resp))
}
