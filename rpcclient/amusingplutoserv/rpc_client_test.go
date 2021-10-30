package amusingplutoserv

import (
	"amusingx.fit/amusingx/protos/amusingplutoservice/plutoservice"
	"amusingx.fit/amusingx/services/amusingplutoserv/conf"
	"context"
	"fmt"
	"github.com/ItsWewin/superfactory/logger"
	"sync"
	"testing"
	"time"
)

func TestInitClient(t *testing.T) {
	conf.Conf.RPCAddress = "localhost:11002"
	InitClient(conf.Conf.RPCAddress)

	ctx, _ := context.WithTimeout(context.Background(), 4*time.Second)

	w := sync.WaitGroup{}
	w.Add(1)
	go ping(ctx, Client, &w)
	//go bookInventoryCacheInit(ctx, Client, &w)

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
