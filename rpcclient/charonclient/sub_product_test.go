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
	w.Add(2)
	go _ping(ctx, Client, &w)
	go _SubProducts(ctx, Client, &w)
	w.Wait()
}

func _SubProducts(ctx context.Context, client charonservice.CharonServClient, w *sync.WaitGroup) {
	w.Done()

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
