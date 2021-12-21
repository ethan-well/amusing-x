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
