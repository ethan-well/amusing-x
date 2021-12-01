package model

import (
	"amusingx.fit/amusingx/services/pluto/mysql"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestQueryBookInventoryByLimit(t *testing.T) {
	if testing.Short() {
		t.Skip("skip")
	}

	mysql.Mock()

	ivs, err := QueryBookInventoryByLimit(context.Background(), 0, 10)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("ivs: %s", logger.ToJson(ivs))
}

func TestQueryBookInventoryByID(t *testing.T) {
	if testing.Short() {
		t.Skip("skip")
	}

	mysql.Mock()

	ivs, err := QueryBookInventoryByID(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("ivs: %s", logger.ToJson(ivs))
}
