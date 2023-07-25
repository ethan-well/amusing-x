package model

import (
	charon2 "amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/services/product/mysql/charon"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestProductInsert(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product := &charon2.Product{
		Name: "name",
		Desc: "desc",
	}
	product, err := ProductInsert(context.Background(), product)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("product: %s", logger.ToJson(product))
}

func TestProductInsertWithTx(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	tx, err := charon.CharonDB.Beginx()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	defer tx.Rollback()

	product, err := ProductInsertWithTx(context.Background(), tx, &charon2.Product{
		Name: "name",
		Desc: "desc",
	})
	if err != nil {
		t.Fatalf("some err: %s", err)
	}
	tx.Commit()

	t.Logf("product: %s", logger.ToJson(product))
}

func TestProductQueryById(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product, err := ProductQueryById(context.Background(), 3)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("product: %s", logger.ToJson(product))
}

func TestProductDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	err := ProductDelete(context.Background(), []int64{3})
	if err != nil {
		t.Fatalf("some err: %s", err)
	}
}

func TestProductDeleteWithTx(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()
	tx, err := charon.CharonDB.Beginx()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	defer tx.Rollback()

	err = ProductDeleteWithTx(context.Background(), tx, []int64{4, 5})
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	if err := tx.Commit(); err != nil {
		t.Fatalf("tx commit failed: %s", err)
	}
}

func TestProductUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product := &charon2.Product{
		ID:   3,
		Name: "name999",
		Desc: "desc999",
	}
	err := ProductUpdate(context.Background(), product)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("product: %s", logger.ToJson(product))
}

func TestProductUpdateWithTx(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product := &charon2.Product{
		ID:   4,
		Name: "name999",
		Desc: "desc999",
	}

	tx, err := charon.CharonDB.Beginx()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	defer tx.Rollback()

	err = ProductUpdateWithTx(context.Background(), tx, product)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	tx.Commit()

	t.Logf("product: %s", logger.ToJson(product))
}

func TestProductListQuery(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	total, products, err := ProductSearch(context.Background(), "name", 1, 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("total: %d", total)
	t.Logf("products: %s", logger.ToJson(products))
}

func TestProductWideInfoById(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	products, err := ProductWideInfoById(context.Background(), 4)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("products: %s", logger.ToJson(products))
}
