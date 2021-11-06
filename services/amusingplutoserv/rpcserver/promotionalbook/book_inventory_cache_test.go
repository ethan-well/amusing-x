package promotionalbook

import (
	"amusingx.fit/amusingx/services/amusingplutoserv/conf"
	"amusingx.fit/amusingx/services/amusingplutoserv/xredis"
	"context"
	"sync"
	"sync/atomic"
	"testing"
)

func TestBookInventoryCache_InventoryDecrBy(t *testing.T) {
	if testing.Short() {
		t.Skip("skip current func ...")
	}

	conf.Mock()
	xredis.Mock()

	bookInventoryCache := NewBookInventoryCache()

	err := bookInventoryCache.InventoryDecrBy(context.Background(), 7772, 10)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("succeed")
}

func TestBookInventoryCache_InventoryDecrByV2(t *testing.T) {
	if testing.Short() {
		t.Skip("skip current func ...")
	}

	conf.Mock()
	xredis.Mock()

	bookInventoryCache := NewBookInventoryCache()

	//w := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		//w.Add(i)
		go func() {
			//defer w.Done()
			err := bookInventoryCache.InventoryDecrBy(context.Background(), 7772, 10)
			if err != nil {
				t.Fatal(err)
			}
		}()

		t.Log("current i: ", i)
	}
}

func TestBookInventoryCache_InventoryIncrBy(t *testing.T) {
	if testing.Short() {
		t.Skip("skip current func ...")
	}

	conf.Mock()
	xredis.Mock()

	bookInventoryCache := NewBookInventoryCache()

	err := bookInventoryCache.InventoryIncrBy(context.Background(), 7772, 10)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("succeed")
}

func TestBookInventoryCache_InventoryIncrByV2(t *testing.T) {
	if testing.Short() {
		t.Skip("skip current func ...")
	}

	conf.Mock()
	xredis.Mock()

	ctx := context.Background()

	bookInventoryCache := NewBookInventoryCache()
	bookInventoryCache.SetInventory(ctx, 7772, 0)

	maxCount := 10000
	incr := 10

	done := make(chan struct{})
	var count atomic.Value
	count.Store(0)
	var mutex sync.Mutex
	for i := 0; i < maxCount; i++ {
		go func(chan struct{}) {
			defer func() {
				mutex.Lock()
				c, _ := count.Load().(int)
				if c == maxCount-1 {
					done <- struct{}{}
				} else {
					count.Store(c + 1)
				}
				mutex.Unlock()
			}()
			err := bookInventoryCache.InventoryIncrBy(context.Background(), 7772, incr)
			if err != nil {
				panic(err)
			}
		}(done)
	}

	<-done

	iv, err := bookInventoryCache.GetInventory(ctx, 7772)
	if err != nil {
		t.Fatal(err)
	}

	if iv != int64(maxCount*incr) {
		t.Fatal("unexpect inventory")
	}
}
