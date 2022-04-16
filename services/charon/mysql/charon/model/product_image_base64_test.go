package model

import (
	charon2 "amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/services/charon/mysql/charon"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestProductImageBase64Insert(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product := &charon2.ProductImageBase64{
		Id:     11,
		Base64: "base64/png,'''''",
	}
	product, err := ProductImageBase64Insert(context.Background(), product)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("result: %s", logger.ToJson(product))
}

func TestImageBase64InsertWithTx(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product := &charon2.ProductImageBase64{
		Id:     11,
		Base64: "base64/png,'''''",
	}

	tx, err := charon.CharonDB.Beginx()
	if err != nil {
		t.Fatalf("beginx failed: %s", err)
	}
	defer tx.Rollback()

	product, err = ProductImageBase64InsertWithTx(context.Background(), tx, product)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	tx.Commit()

	t.Logf("result: %s", logger.ToJson(product))
}

func TestProductImageBase64QueryById(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	image, err := ProductImageBase64QueryById(context.Background(), 5)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("image: %s", logger.ToJson(image))
}

func TestProductImageBase64Delete(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	err := ProductImageBase64Delete(context.Background(), []int64{4, 5})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("succeed")
}
