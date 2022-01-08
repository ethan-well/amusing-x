package promotionalbook

import (
	"amusingx.fit/amusingx/services/pluto/conf"
	"amusingx.fit/amusingx/services/pluto/xredis"
	"context"
	"testing"
)

func TestBookInventoryCache_InventoryDecrBy(t *testing.T) {
	if testing.Short() {
		t.Skip("skip current func ...")
	}

	conf.Mock()
	xredis.Mock()

	bookInventoryCache := NewBookInventoryCache()

	_, err := bookInventoryCache.InventoryDecrBy(context.Background(), 7772, 10)
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
			_, err := bookInventoryCache.InventoryDecrBy(context.Background(), 7772, 10)
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

	_, err := bookInventoryCache.InventoryIncrBy(context.Background(), 7772, 10)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("succeed")
}
