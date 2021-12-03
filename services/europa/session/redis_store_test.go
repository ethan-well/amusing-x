package session

import (
	"amusingx.fit/amusingx/services/europa/conf"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestSessionInit(t *testing.T) {
	conf.Mock()

	store, err := InitRedisStore(conf.ConfIns.SessionStoreRedisAddr, conf.ConfIns.SessionStoreRedisPassword,
		conf.ConfIns.SessionStoreRedisDB, "session", 60)

	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	sid := "sid:for:test"
	session, err := store.SessionInit(context.Background(), sid)
	if err != nil {
		t.Fatalf("some err: %s", err.Error())
	}

	t.Logf("session: %s", logger.ToJson(session))
}
