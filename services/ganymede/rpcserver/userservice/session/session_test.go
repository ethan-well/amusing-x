package session

import (
	"amusingx.fit/amusingx/services/ganymede/conf"
	"amusingx.fit/amusingx/services/ganymede/session"
	"context"
	"testing"
)

func TestModel_SetSession(t *testing.T) {
	if testing.Short() {
		t.Skip("skip func")
	}
	conf.Mock()
	err := session.InitSessionManager("redis", 120)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	model := New()
	ctx := context.Background()

	sid := "55ea737d-1888-459d-90fd-2c50e5927233"
	idCache, xErr := model.GetUserID(ctx, sid)
	if xErr != nil {
		t.Fatalf("err: %s", xErr)
	}

	t.Logf("idCache: %d", idCache)
}
