package model

import (
	charon2 "amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/services/charon/mysql/charon"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestAttributeInsert(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	attribute := &charon2.Attribute{
		Name: "name",
		Desc: "desc",
	}

	attribute, err := AttributeInsert(context.Background(), attribute)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("attribute : %s", logger.ToJson(attribute))
}

func TestAttributeQueryById(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product, err := AttributeQueryById(context.Background(), 1)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("attribute: %s", logger.ToJson(product))
}

func TestAttributeUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product := &charon2.Attribute{
		ID:   2,
		Name: "name 2",
		Desc: "desc 2",
	}
	err := AttributeUpdate(context.Background(), product)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("attribute: %s", logger.ToJson(product))
}

func TestAttributeListQuery(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	total, attributes, err := AttributeSearch(context.Background(), "name", 0, 10)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("total: %d", total)
	t.Logf("attribute: %s", logger.ToJson(attributes))
}

func TestAttributeDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	err := AttributeDelete(context.Background(), []int64{1})
	if err != nil {
		t.Fatalf("some err: %s", err)
	}
}

func TestAttributeDeleteWithTx(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	tx, e := charon.CharonDB.Beginx()
	if e != nil {
		t.Fatal(e)
	}
	defer tx.Rollback()

	err := AttributeDeleteWithTx(context.Background(), tx, []int64{3})
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	if e := tx.Commit(); e != nil {
		t.Fatal(e)
	}
}

func TestAttributeQueryByIds(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	products, err := AttributeQueryByIds(context.Background(), []int64{1, 2, 3, 4, 9, 11})
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("attribute: %s", logger.ToJson(products))
}

func TestAttributeQueryBySubProductIds(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	products, err := AttributeQueryBySubProductIds(context.Background(), []int64{1, 2, 3, 4, 9, 11})
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("attribute: %s", logger.ToJson(products))
}
