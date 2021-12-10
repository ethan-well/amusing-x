package session

import (
	"amusingx.fit/amusingx/services/ganymede/conf"
	"amusingx.fit/amusingx/services/ganymede/session"
	"context"
	"testing"
)

func TestModel_SetSession(t *testing.T) {
	conf.Mock()
	err := session.InitSessionManager("redis", "uid", 120)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	model := New()
	ctx := context.Background()

	var userID int64 = 1213131
	info := Info{UserID: userID}

	sid, xErr := model.SetSession(ctx, &info)
	if xErr != nil {
		t.Fatalf("err: %s", xErr)
	}

	idCache, xErr := model.GetUserID(ctx, sid)
	if xErr != nil {
		t.Fatalf("err: %s", xErr)
	}

	if idCache != userID {
		t.Fatalf("is is not expected")
	}

	t.Logf("sid: %s", sid)
}
