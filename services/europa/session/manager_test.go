package session

import (
	"amusingx.fit/amusingx/services/europa/conf"
	"context"
	"github.com/satori/go.uuid"
	"testing"
)

func TestNewManager(t *testing.T) {
	conf.Mock()

	err := InitSessionManager("redis", "uid", 60)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	session, err := GlobalSessionManager.Store.SessionInit(context.TODO(), uuid.NewV4().String())
	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	key1 := "key1"
	value1 := "value1"
	err = session.Set(context.TODO(), key1, value1)
	if err != nil {
		t.Fatalf("some error occured when session.Set: %s", err)
	}

	v, err := session.Get(context.TODO(), key1)
	if err != nil {
		t.Fatalf("some error occured when session.Get: %s", err)
	}
	t.Logf("key: %s, value: %s", key1, v)

	key2 := "key2"
	value2 := "value2"
	err = session.Set(context.TODO(), key2, value2)
	if err != nil {
		t.Fatalf("some error occured when session.Set: %s", err)
	}

	v, err = session.Get(context.TODO(), key2)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("key: %s, value: %s", key2, value2)

	key3 := "key3"
	value3 := "value3"
	err = session.Set(context.TODO(), key3, value3)
	if err != nil {
		t.Fatalf("some error occured when session.Set: %s", err)
	}

	session2, err := GlobalSessionManager.Store.SessionInit(context.TODO(), uuid.NewV4().String())
	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	name, err := session2.Get(context.TODO(), "name")
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Fatalf("name: %s", name)
}
