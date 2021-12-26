package pangu

import (
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu"
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
	//go _create(ctx, Client, &w)
	//go _categoryList(ctx, Client, &w)
	//go _del(ctx, Client, &w)
	go _category(ctx, Client, &w)

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

func _categoryList(ctx context.Context, client panguservice.PanGuServiceClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.CategoryList(context.Background(), &panguservice.CategoryListRequest{
			Query: "4",
			Page:  1,
			Limit: 100,
		})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("pong succeed", resp)
	}
}

func _del(ctx context.Context, client panguservice.PanGuServiceClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.CategoryDelete(context.Background(), &panguservice.CategoryDeleteRequest{Ids: []int64{1, 3, 4, 5, 6}})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("del succeed", resp)
	}
}

func _category(ctx context.Context, client panguservice.PanGuServiceClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.Category(context.Background(), &panguservice.CategoryRequest{Id: 11})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("del succeed", resp)
	}
}
