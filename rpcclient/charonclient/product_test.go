package charonclient

import (
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"context"
	"fmt"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
	"time"
)

func TestProduct(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	ClientMock(t)

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

	product := &proto.Product{
		Name: "name",
		Desc: "desc",
	}

	_productCreate(ctx, Client, product, t)
	_query(ctx, Client, product, t)
	_productUpdate(ctx, Client, product, t)
	_productDel(ctx, Client, product, t)
}

func _query(ctx context.Context, client charonservice.CharonServClient, product *proto.Product, t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.Product(ctx, &proto.ProductRequest{Id: product.ID})
		if err != nil {
			t.Fatal(err)
		}

		if resp == nil {
			t.Fatal("resp is nil")
		}

		if resp.ID != product.ID || resp.Name != product.Name || resp.Desc != product.Desc {
			t.Logf("resp: %s", logger.ToJson(resp))
			t.Fatal("resp is not expected")
		}
	}
}

func _productCreate(ctx context.Context, client charonservice.CharonServClient, product *proto.Product, t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.ProductCreate(ctx, &proto.ProductCreateRequest{
			Name: product.Name,
			Desc: product.Desc,
		})
		if err != nil {
			t.Fatal(err)
		}

		if resp == nil {
			t.Fatal("resp is nil")
		}

		if resp.ID <= 0 {
			t.Fatal("resp.ID <= 0")
		}

		if resp.Name != product.Name {
			t.Fatalf("resp.Name expected: %s, get: %s", product.Name, resp.Name)
		}

		if resp.Desc != product.Desc {
			t.Fatalf("resp.Desc expected: %s, get: %s", product.Desc, resp.Desc)
		}

		product.ID = resp.ID
	}
}

func _productUpdate(ctx context.Context, client charonservice.CharonServClient, product *proto.Product, t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		productNew := proto.Product{
			ID:   product.ID,
			Name: "new name updated: " + time.Now().String(),
			Desc: "new desc updated: " + time.Now().String(),
		}
		resp, err := client.ProductUpdate(ctx, &proto.ProductUpdateRequest{
			Id:   productNew.ID,
			Name: productNew.Name,
			Desc: productNew.Desc,
		})
		if err != nil {
			t.Fatal(err)
		}

		if resp.ID != productNew.ID || resp.Name != productNew.Name || resp.Desc != productNew.Desc {
			t.Fatal("resp is not expected")
		}
	}
}

func _productDel(ctx context.Context, Client charonservice.CharonServClient, product *proto.Product, t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := Client.ProductDelete(ctx, &proto.ProductDeleteRequest{Id: product.ID})
		if err != nil {
			t.Fatal(err)
		}

		if resp == nil || !resp.Result {
			t.Fatal("delete failed")
		}
	}
}
