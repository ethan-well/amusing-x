package session

import (
	"amusingx.fit/amusingx/services/user/conf"
	"amusingx.fit/amusingx/services/user/mysql/ganymededb/model"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestSessionInit(t *testing.T) {
	conf.Mock()

	model.InitMySQL()

	redis := conf.Conf.SessionStore.Redis
	store, err := InitRedisStore(redis.Addr, redis.Password, redis.DBNo, "session", 60)
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
