package amusingplutoserv

import (
	"amusingx.fit/amusingx/protos/amusingplutoservice/plutoservice"
	"context"
	"fmt"
	"github.com/ItsWewin/superfactory/logger"
	"sync"
	"testing"
	"time"
)

func TestInitClient(t *testing.T) {
	InitClient("localhost:11002")

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

	w := sync.WaitGroup{}
	w.Add(2)
	go ping(ctx, Client, &w)
	go bookInventoryCacheInit(ctx, Client, &w)

	w.Wait()
}

func ping(ctx context.Context, client plutoservice.AmusingxPlutoServiceClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.Pong(context.Background(), &plutoservice.BlankParams{})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(resp)
	}
}

func bookInventoryCacheInit(ctx context.Context, client plutoservice.AmusingxPlutoServiceClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		req := &plutoservice.InventoryCacheInitRequest{Obj: plutoservice.InventoryCacheInitRequest_Book}
		resp, err := client.InventoryCacheInit(ctx, req)
		if err != nil {
			fmt.Println(err)
		}

		if resp.Succeed {
			logger.Infof("bookInventoryCacheInit succeed")
		} else {
			logger.Errorf("bookInventoryCacheInit failed")
		}
	}
}
