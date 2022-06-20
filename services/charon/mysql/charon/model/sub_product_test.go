package model

import (
	charon2 "amusingx.fit/amusingx/mysqlstruct/charon"
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
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
		MinNum:    1111,
		MaxNum:    1111111111,
	}
	product, err := SubProductInsert(context.Background(), product)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("sub product: %s", logger.ToJson(product))
}

func TestSubProductInsertWithTx(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	tx, err := charon.CharonDB.Beginx()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	product := &charon2.SubProduct{
		Name:      "name",
		Desc:      "desc",
		ProductId: 1,
		Currency:  "CNY",
		Price:     10000,
		MinNum:    1111,
		MaxNum:    1111111111,
	}
	product, err = SubProductInsertWithTx(context.Background(), product, tx)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	tx.Commit()

	t.Logf("sub product: %s", logger.ToJson(product))
}

func TestSubProductQueryById(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product, err := SubProductQueryById(context.Background(), 9)
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
		ID:        13,
		Name:      "name 2",
		Desc:      "desc 2",
		ProductId: 222,
		Currency:  "USD",
		Price:     11,
		MinNum:    111,
		MaxNum:    222,
	}
	err := SubProductUpdate(context.Background(), product)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("sub product: %s", logger.ToJson(product))
}

func TestSubProductUpdateWithTx(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ....")
	}

	charon.Mock()

	tx, err := charon.CharonDB.Beginx()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	product := &charon2.SubProduct{
		ID:        9,
		Name:      "name",
		Desc:      "desc",
		ProductId: 101,
		Currency:  "CNY",
		Price:     101,
	}

	err = SubProductUpdateWithTx(context.Background(), tx, product)
	if err != nil {
		t.Fatal(err)
	}

	if e := tx.Commit(); e != nil {
		t.Fatal(e)
	}
}

func TestSubProductWithStockQueryByIdWithTx(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ....")
	}

	charon.Mock()

	tx, err := charon.CharonDB.Beginx()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	info, err := SubProductWithStockQueryByIdWithTx(context.Background(), 16, tx)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("info: %s", logger.ToJson(info))

	tx.Commit()
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

func TestSubProductSearchV2(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	in := &proto.SubProductListRequest{
		Query:  "",
		Page:   1,
		Limit:  10,
		Sort:   "",
		Filter: "",
		Offset: 0,
	}
	filter := &charonservice.SearchFilter{
		Id:   nil,
		Name: nil,
		Desc: nil,
	}

	total, subProducts, err := SubProductSearchV2(context.Background(), in, filter)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("total: %d", total)
	t.Logf("products: %s", logger.ToJson(subProducts))
}

func TestSubProductSearchWithProductStock(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	in := &proto.SubProductListRequest{
		Query:  "",
		Page:   1,
		Limit:  10,
		Sort:   "",
		Filter: "",
		Offset: 0,
	}
	filter := &charonservice.SearchFilter{
		Id:   nil,
		Name: nil,
		Desc: nil,
	}

	total, subProducts, err := SubProductSearchWithProductStock(context.Background(), in, filter)
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

func TestSubProductDeleteWithTx(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	tx, e := charon.CharonDB.Beginx()
	if e != nil {
		t.Fatal(e)
	}
	defer tx.Rollback()

	err := SubProductDeleteWithTx(context.Background(), tx, []int64{4})
	if err != nil {
		t.Fatalf("some err: %s", err)
	}
	if e := tx.Commit(); e != nil {
		t.Fatal(e)
	}
}
