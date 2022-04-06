package model

import (
	charon2 "amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/services/charon/mysql/charon"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestAttributeMappingInsert(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product := &charon2.AttributeMapping{
		AttrId:       1,
		AttrValue:    "111",
		SubProductId: 2,
	}
	product, err := AttributeMappingInsert(context.Background(), product)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("sub product: %s", logger.ToJson(product))
}

func TestAttributeMappingInsertWithTx(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	tx, e := charon.CharonDB.Beginx()
	if e != nil {
		t.Fatal(e)
	}
	defer tx.Rollback()

	var products []*charon2.AttributeMapping
	product := &charon2.AttributeMapping{
		AttrId:       1,
		AttrValue:    "111222200000",
		SubProductId: 2,
	}
	product2 := &charon2.AttributeMapping{
		AttrId:       2,
		AttrValue:    "2222220000",
		SubProductId: 2,
	}
	products = append(products, product, product2)

	err := AttributeMappingInsertWithTx(context.Background(), tx, products)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	if e := tx.Commit(); e != nil {
		t.Fatalf("commit err: %s", err)
	}

	t.Logf("sub product: %s", logger.ToJson(product))
}

func TestAttributeMappingQueryById(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product, err := AttributeMappingQueryById(context.Background(), 2)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("attribute_mapping: %s", logger.ToJson(product))
}

func TestAttributeMappingQueryBySubProductIDWithTx(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	tx, e := charon.CharonDB.Beginx()
	if e != nil {
		t.Fatal(e)
	}
	defer tx.Rollback()

	product, err := AttributeMappingQueryBySubProductIDWithTx(context.Background(), tx, 3)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("attribute_mapping: %s", logger.ToJson(product))
}

func TestAttributeMappingUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product := &charon2.AttributeMapping{
		ID:           3,
		AttrId:       1000001,
		AttrValue:    "10000222",
		SubProductId: 100000111,
	}
	err := AttributeMappingUpdate(context.Background(), product)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("AttributeMapping: %s", logger.ToJson(product))
}

func TestAttributeMappingDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	err := AttributeMappingDelete(context.Background(), []int64{3})
	if err != nil {
		t.Fatalf("some err: %s", err)
	}
}

func TestAttributeMappingDeleteWithTx(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	tx, e := charon.CharonDB.Beginx()
	if e != nil {
		t.Fatal(e)
	}
	defer tx.Rollback()

	err := AttributeMappingDeleteWithTx(context.Background(), tx, []int64{4})
	if err != nil {
		t.Fatalf("some err: %s", err)
	}
	if e := tx.Commit(); e != nil {
		t.Fatal(e)
	}
}

func TestAttributeMappingDeleteBySubProductIdWithTx(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	tx, e := charon.CharonDB.Beginx()
	if e != nil {
		t.Fatal(e)
	}
	defer tx.Rollback()

	err := AttributeMappingDeleteBySubProductIdWithTx(context.Background(), tx, []int64{2})
	if err != nil {
		t.Fatalf("some err: %s", err)
	}
	if e := tx.Commit(); e != nil {
		t.Fatal(e)
	}
}

func TestAttributeMappingDeleteByAttributeIdWithTx(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	tx, e := charon.CharonDB.Beginx()
	if e != nil {
		t.Fatal(e)
	}
	defer tx.Rollback()

	err := AttributeMappingDeleteByAttributeIdWithTx(context.Background(), tx, []int64{6})
	if err != nil {
		t.Fatalf("some err: %s", err)
	}
	if e := tx.Commit(); e != nil {
		t.Fatal(e)
	}
}

func TestAttributeMappingSearch(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	total, attrs, err := AttributeMappingSearch(context.Background(), "", 0, 10)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("total: %d", total)
	t.Logf("attrs: %s", logger.ToJson(attrs))
}
