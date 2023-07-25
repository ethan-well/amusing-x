package model

import (
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	"amusingx.fit/amusingx/services/user/conf"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestInsert(t *testing.T) {
	conf.Mock()
	Mock()

	u := &ganymede.User{
		Name:  "wei.wei",
		Login: "wei.wei.3",
	}

	tx, err := UserDB.Beginx()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()
	u, xErr := InsertUser(context.TODO(), tx, u)
	if xErr != nil {
		t.Fatalf("some error: %s", xErr)
	}
	if err := tx.Commit(); err != nil {
		t.Fatalf("commit err")
	}

	t.Logf("user: %s", logger.ToJson(u))
}
