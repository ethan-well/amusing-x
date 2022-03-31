package model

import (
	charon2 "amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/services/charon/mysql/charon"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestCategoryProductMappingInsert(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	ctx := context.Background()
	mapping := &charon2.CategoryProductMapping{
		CategoryId: 10,
		ProductId:  11,
	}

	category, err := CategoryProductMappingInsert(ctx, mapping)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	t.Logf("category: %s", logger.ToJson(category))
}

func TestCategoryProductMappingInsertWithTx(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	ctx := context.Background()
	mapping := &charon2.CategoryProductMapping{
		CategoryId: 10,
		ProductId:  11,
	}

	tx, err := charon.CharonDB.Beginx()
	if err != nil {
		t.Fatalf("beginx failed: %s", err)
	}
	defer tx.Rollback()

	category, err := CategoryProductMappingInsertWithTx(ctx, tx, mapping)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	tx.Commit()

	t.Logf("category: %s", logger.ToJson(category))
}

func TestCategoryProductMappingQueryById(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	ctx := context.Background()
	category, err := CategoryProductMappingQueryById(ctx, 8)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	t.Logf("category: %s", logger.ToJson(category))
}

func TestCategoryProductMappingDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	ctx := context.Background()
	err := CategoryProductMappingDelete(ctx, []int64{8})
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestCategoryProductMappingUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	ctx := context.Background()

	mapping := &charon2.CategoryProductMapping{
		ID:         8,
		CategoryId: 27,
		ProductId:  43,
	}
	err := CategoryProductMappingUpdate(ctx, mapping)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestCategoryProductMappingUpdateByProductIdWithTx(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	ctx := context.Background()

	mapping := &charon2.CategoryProductMapping{
		CategoryId: 43,
		ProductId:  27,
	}
	tx, e := charon.CharonDB.Beginx()
	if e != nil {
		t.Fatalf("beginx failed: %s", e)
	}
	defer tx.Rollback()

	err := CategoryProductMappingUpdateByProductIdWithTx(ctx, tx, mapping)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if e := tx.Commit(); e != nil {
		t.Fatalf("tx commit failed")
	}
}

func TestCategoryProductMappingDeleteByProductIdWithTX(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	ctx := context.Background()

	tx, e := charon.CharonDB.Beginx()
	if e != nil {
		t.Fatalf("beginx failed: %s", e)
	}
	defer tx.Rollback()

	err := CategoryProductMappingDeleteByProductIdWithTX(ctx, tx, []int64{4, 9})
	if err != nil {
		t.Fatal(err)
	}

	if e := tx.Commit(); e != nil {
		t.Fatal(err)
	}
}
