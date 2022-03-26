package charonclient

import (
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	"context"
	"fmt"
	"github.com/ItsWewin/superfactory/logger"
	"sync"
	"testing"
	"time"
)

func TestCategory(t *testing.T) {
	ClientMock(t)

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

	w := sync.WaitGroup{}
	w.Add(6)
	go _ping(ctx, Client, &w)
	go _create(ctx, Client, &w)
	go _categories(ctx, Client, &w)
	go _del(ctx, Client, &w)
	go _category(ctx, Client, &w)
	go _update(ctx, Client, &w)
	w.Wait()
}

func _ping(ctx context.Context, client charonservice.CharonServClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.Pong(context.Background(), &charonservice.BlankParams{})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(resp)
	}
}

func _create(ctx context.Context, client charonservice.CharonServClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.Create(context.Background(), &charonservice.CategoryCreateRequest{
			Name: "name 8",
			Desc: "desc 8",
		})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("resp: %s", logger.ToJson(resp))
	}
}

func _categories(ctx context.Context, client charonservice.CharonServClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.Categories(context.Background(), &charonservice.CategoryListRequest{
			Query: "14",
			Page:  1,
			Limit: 10,
		})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("resp: %s", logger.ToJson(resp))
	}
}

func _del(ctx context.Context, client charonservice.CharonServClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.Delete(context.Background(), &charonservice.CategoryDeleteRequest{Id: 1})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("resp: %s", logger.ToJson(resp))
	}
}
func _category(ctx context.Context, client charonservice.CharonServClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.Category(context.Background(), &charonservice.CategoryRequest{Id: 11})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("resp: %s", logger.ToJson(resp))
	}
}

func _update(ctx context.Context, client charonservice.CharonServClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.Update(context.Background(),
			&charonservice.CategoryUpdateRequest{Category: &charonservice.Category{Id: 11, Name: "new name", Desc: "new desc"}})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("resp: %s", logger.ToJson(resp))
	}
}
