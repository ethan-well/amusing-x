package pangu

import (
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/admin/conf"
	"amusingx.fit/amusingx/services/admin/mysql/pangu"
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
	//go _update(ctx, Client, &w)
	//go _oauthProviderInfo(ctx, Client, &w)
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
		resp, err := client.CategoryDelete(context.Background(), &panguservice.CategoryDeleteRequest{Id: 1})
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

func _update(ctx context.Context, client panguservice.PanGuServiceClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.CategoryUpdate(context.Background(),
			&panguservice.CategoryUpdateRequest{Category: &panguservice.Category{
				ID:   11,
				Name: "new name update",
				Desc: "new desc update",
			}})
		if err != nil {
			fmt.Println(err)
		}

		if !resp.Result {
			fmt.Println("update failed")
		} else {
			fmt.Println("update succeed", resp)
		}
	}
}

func _oauthProviderInfo(ctx context.Context, client panguservice.PanGuServiceClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.OauthProviderInfo(context.Background(), &panguservice.OauthProviderInfoRequest{
			Provider: "github",
			Service:  "pangu",
		})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("get oauth provider succeed", resp)
	}
}
