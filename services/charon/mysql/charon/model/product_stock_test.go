package model

import (
	charon2 "amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/services/charon/mysql/charon"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestProductStockInsert(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	tx, e := charon.CharonDB.Beginx()
	if e != nil {
		t.Fatal(e)
	}
	defer tx.Rollback()

	productStock := &charon2.ProductStock{
		SubProductId:       11,
		RealInventory:      100,
		AvailableInventory: 50,
	}

	productStock, err := ProductStockInsert(context.Background(), productStock, tx)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	tx.Commit()

	t.Logf("product stock: %s", logger.ToJson(productStock))
}

func TestProductStockDeleteBySubProductIds(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	tx, e := charon.CharonDB.Beginx()
	if e != nil {
		t.Fatal(e)
	}
	defer tx.Rollback()

	err := ProductStockDeleteBySubProductIds(context.Background(), []int64{10}, tx)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}
	tx.Commit()
}

func TestQueryProductStockBySubProductIds(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	tx, e := charon.CharonDB.Beginx()
	if e != nil {
		t.Fatal(e)
	}
	defer tx.Rollback()

	result, err := QueryProductStockBySubProductIds(context.Background(), []int64{11}, tx)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}
	tx.Commit()

	t.Logf("result: %s", logger.ToJson(result))
}

func TestProductStockUpdateWithTx(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	tx, e := charon.CharonDB.Beginx()
	if e != nil {
		t.Fatal(e)
	}
	defer tx.Rollback()

	err := ProductStockUpdateWithTx(context.Background(), &charon2.ProductStock{
		SubProductId:       11,
		RealInventory:      10001,
		AvailableInventory: 888,
	}, tx)
	if err != nil {
		t.Fatal(err)
	}

	tx.Commit()
}
