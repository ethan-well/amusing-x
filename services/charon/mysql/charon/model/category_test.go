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

func TestInsetCategory(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

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
		ctx = context.Background()
		req = &proto.CategoryListRequest{
			Query:  "",
			Page:   1,
			Limit:  10,
			Range:  "",
			Sort:   "",
			Offset: 0,
			Ranges: nil,
		}
	)

	filter := &charonservice.SearchFilter{
		Id:   []int64{21},
		Name: nil,
		Desc: nil,
	}

	category, total, err := CategoryQuery(ctx, req, filter)
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

func TestQueryCategoryByName(t *testing.T) {
	if testing.Short() {
		t.Skip("skip current func")
	}

	charon.Mock()
	ctx := context.Background()

	category, err := QueryCategoryByName(ctx, "name1")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	t.Logf("category: %s", logger.ToJson(category))
}

func TestSearchCategoryByProductIds(t *testing.T) {
	if testing.Short() {
		t.Skip("skip current func")
	}

	charon.Mock()
	ctx := context.Background()

	categories, err := SearchCategoryByProductIds(ctx, []int64{4, 5, 6, 7, 8, 9, 10, 12, 19, 20})
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	t.Logf("categories: %s", logger.ToJson(categories))
}
