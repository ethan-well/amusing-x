package charonclient

import (
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"context"
	"fmt"
	"github.com/ItsWewin/superfactory/logger"
	"sync"
	"testing"
	"time"
)

func TestProduct(t *testing.T) {
	ClientMock(t)

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

	w := sync.WaitGroup{}
	w.Add(2)
	go _ping(ctx, Client, &w)
	go _productCreate(ctx, Client, &w)
	w.Wait()
}

func _productCreate(ctx context.Context, client charonservice.CharonServClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.ProductCreate(ctx, &proto.ProductCreateRequest{
			Name: "name 8",
			Desc: "desc 8",
		})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("resp: %s", logger.ToJson(resp))
	}
}
