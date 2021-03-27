package amusinguser

import (
	"amusingx.fit/amusingx/services/amusinguserserv/conf"
	"context"
	"testing"
)

func TestQueryUserByIdContext(t *testing.T) {
	conf.Mock()
	Mock()

	user, err := QueryUserByIdContext(context.TODO(), 1)
	if err != nil {
		t.Fatalf("some error occurred: %s", err)
	}

	if user == nil {
		t.Logf("no user with id = %d", 1)
	} else {
		t.Logf("succced, user id: %d", user.ID)
	}
}
