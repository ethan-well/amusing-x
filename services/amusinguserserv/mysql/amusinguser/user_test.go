package amusinguser

import (
	"amusingx.fit/amusingx/services/amusinguserserv/conf"
	"context"
	"testing"
)

func TestQueryUserByIdContext(t *testing.T) {
	conf.Mock()

	t.Logf("conf : %s", conf.Conf.MysqlAmusinguserPassword)

	Mock()

	ctx := context.TODO()
	var id int64 = 1

	user, err := QueryUserByIdContext(ctx, id)
	if err != nil {
		t.Fatalf("some error occurred: %s", err)
	}

	if user == nil {
		t.Logf("no user with id = %d", id)
	} else {
		t.Logf("succced, user id: %d", user.ID)
	}
}
