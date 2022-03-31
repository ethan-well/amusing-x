package model

import (
	charon2 "amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/services/charon/mysql/charon"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestSubProductInsert(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product := &charon2.SubProduct{
		Name:      "name",
		Desc:      "desc",
		ProductId: 1,
		Currency:  "CNY",
		Price:     10000,
		Stock:     111111,
	}
	product, err := SubProductInsert(context.Background(), product)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("sub product: %s", logger.ToJson(product))
}

func TestSubProductQueryById(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product, err := SubProductQueryById(context.Background(), 3)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("product: %s", logger.ToJson(product))
}

func TestSubProductUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product := &charon2.SubProduct{
		ID:        4,
		Name:      "name 2",
		Desc:      "desc 2",
		ProductId: 222,
		Currency:  "USD",
		Price:     100001,
		Stock:     222222,
	}
	err := SubProductUpdate(context.Background(), product)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("sub product: %s", logger.ToJson(product))
}

func TestSubProductListQuery(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	total, subProducts, err := SubProductSearch(context.Background(), "name", 0, 10)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("total: %d", total)
	t.Logf("products: %s", logger.ToJson(subProducts))
}

func TestSubProductDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	err := SubProductDelete(context.Background(), []int64{3})
	if err != nil {
		t.Fatalf("some err: %s", err)
	}
}
