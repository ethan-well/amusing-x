package pangu

import (
	panguservice "amusingx.fit/amusingx/protos/pangu/service"
	"amusingx.fit/amusingx/services/pangu/conf"
	"amusingx.fit/amusingx/services/pangu/mysql/pangu"
	"fmt"
	"golang.org/x/net/context"
	"sync"
	"testing"
	"time"
)

func TestInitClient(t *testing.T) {
	pangu.Mock()

	InitClient(conf.Conf.Server.GrpcServer.Address)

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

	w := sync.WaitGroup{}
	w.Add(2)
	go _ping(ctx, Client, &w)
	go _create(ctx, Client, &w)

	w.Wait()
}

func _ping(ctx context.Context, client panguservice.PanGuServiceClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.Pong(context.Background(), &panguservice.BlankParams{})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("pong succeed", resp)
	}
}

func _create(ctx context.Context, client panguservice.PanGuServiceClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.CategoryCreate(context.Background(), &panguservice.CategoryCreateRequest{
			Name: "name 9",
			Desc: "name 9 desc",
		})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("pong succeed", resp)
	}
}
