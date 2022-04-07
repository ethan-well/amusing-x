package model

import (
	charon2 "amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/services/charon/mysql/charon"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestProductImageInsert(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product := &charon2.ProductImage{
		ProductId:    100,
		ProductLevel: 12,
		UploaderType: "local",
		Url:          "/tmp/amusing-x/pictures/local.png",
	}
	product, err := ProductImageInsert(context.Background(), product)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("result: %s", logger.ToJson(product))
}

func TestProductImageQueryById(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product, err := ProductImageQueryById(context.Background(), 3)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("result: %s", logger.ToJson(product))
}

func TestProductImageUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	image := &charon2.ProductImage{
		Id:           3,
		ProductId:    10001,
		ProductLevel: 12,
		UploaderType: "localhost",
		Url:          "/tmp/amusing-x/pictures/local2.png",
	}
	err := ProductImageUpdate(context.Background(), image)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("result: %s", logger.ToJson(image))
}

func TestProductImageUpdateWithTx(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ....")
	}

	charon.Mock()

	tx, err := charon.CharonDB.Beginx()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	image := &charon2.ProductImage{
		Id:           3,
		ProductId:    10000111,
		ProductLevel: 12,
		UploaderType: "dddddddd",
	}

	err = ProductImageUpdateWithTx(context.Background(), tx, image)
	if err != nil {
		t.Fatal(err)
	}

	if e := tx.Commit(); e != nil {
		t.Fatal(e)
	}
}

func TestProductImageListQuery(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	total, images, err := ProductImageSearch(context.Background(), "", 0, 10)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("total: %d", total)
	t.Logf("images: %s", logger.ToJson(images))
}

func TestProductImageDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	err := ProductImageDelete(context.Background(), []int64{3})
	if err != nil {
		t.Fatalf("some err: %s", err)
	}
}

func TestProductImageDeleteWithTx(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	tx, e := charon.CharonDB.Beginx()
	if e != nil {
		t.Fatal(e)
	}
	defer tx.Rollback()

	err := ProductImageDeleteWithTx(context.Background(), tx, []int64{4})
	if err != nil {
		t.Fatalf("some err: %s", err)
	}
	if e := tx.Commit(); e != nil {
		t.Fatal(e)
	}
}
