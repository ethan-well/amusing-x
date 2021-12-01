package plutoclient

import (
	"amusingx.fit/amusingx/protos/pluto/service"
	"context"
	"fmt"
	"github.com/ItsWewin/superfactory/logger"
	"sync"
	"testing"
	"time"
)

func TestInitClient(t *testing.T) {
	InitClient(":11002")

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

	w := sync.WaitGroup{}
	w.Add(4)
	go _ping(ctx, Client, &w)
	go _bookInventoryCacheInit(ctx, Client, &w)
	go _bookInventoryLock(ctx, Client, &w)
	go _bookInventoryUnLock(ctx, Client, &w)

	w.Wait()
}

func _ping(ctx context.Context, client plutoservice.AmusingxPlutoServiceClient, w *sync.WaitGroup) {
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

func _bookInventoryCacheInit(ctx context.Context, client plutoservice.AmusingxPlutoServiceClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		req := &plutoservice.InventoryCacheInitRequest{Obj: plutoservice.CacheObjType_Book}
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

func _bookInventoryLock(ctx context.Context, client plutoservice.AmusingxPlutoServiceClient, w *sync.WaitGroup) {
	defer w.Done()

	var w2 = sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		w2.Add(1)
		go func() {
			defer w2.Done()
			in := &plutoservice.InventoryLockRequest{
				Id:    7772,
				Count: 10,
				Obj:   plutoservice.CacheObjType_Book,
			}

			resp, err := client.InventoryLock(ctx, in)
			if err != nil {
				logger.Waringf("InventoryLock failed: %s", err)
			} else {
				logger.Info("resp: %s", logger.ToJson(resp))
			}
		}()
	}

	w2.Wait()
}

func _bookInventoryUnLock(ctx context.Context, client plutoservice.AmusingxPlutoServiceClient, w *sync.WaitGroup) {
	defer w.Done()

	var w2 = sync.WaitGroup{}
	for i := 10; i < 20; i++ {
		w2.Add(1)
		go func() {
			defer w2.Done()
			in := &plutoservice.InventoryUnlockRequest{
				Id:    7772,
				Count: 10,
				Obj:   plutoservice.CacheObjType_Book,
			}

			resp, err := client.InventoryUnlock(ctx, in)
			if err != nil {
				logger.Waringf("InventoryLock failed: %s", err)
			} else {
				logger.Infof("%s", logger.ToJson(resp))
			}
		}()
	}

	w2.Wait()
}
