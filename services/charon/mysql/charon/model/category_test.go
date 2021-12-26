package model

import (
	charon2 "amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/services/charon/mysql/charon"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestInsetCategory(t *testing.T) {
	charon.Mock()

	category := &charon2.Category{
		Name: "category name 2",
		Desc: "category desc",
	}
	category, err := InsetCategory(context.Background(), category)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("category: %s", logger.ToJson(category))
}

func TestCategoryQuery(t *testing.T) {
	if testing.Short() {
		t.Skip("skip current func")
	}

	charon.Mock()
	var (
		ctx          = context.Background()
		id     int64 = 0
		name         = ""
		desc         = ""
		offSet int64 = 1
		limit  int64 = 2
	)
	category, total, err := CategoryQuery(ctx, id, name, desc, offSet, limit)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("total: %d", total)
	t.Logf("category: %s", logger.ToJson(category))
}

func TestDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("skip current func")
	}

	charon.Mock()
	ctx := context.Background()
	err := Delete(ctx, []int64{1})
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Log("succeed")
}

func TestQueryCategoryByID(t *testing.T) {
	if testing.Short() {
		t.Skip("skip current func")
	}

	charon.Mock()
	ctx := context.Background()

	category, err := QueryCategoryByID(ctx, 11)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	t.Logf("category: %s", logger.ToJson(category))
}
