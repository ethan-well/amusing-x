package lock

import (
	"amusingx.fit/amusingx/services/amusinguserserv/xredis"
	"context"
	"testing"
)

var l = NewLocker("key1", "value1", 100)

func TestLock(t *testing.T) {
	xredis.Mock()

	if err := l.Lock(context.TODO()); err != nil {
		t.Fatalf("some error: %s", err)
	}
}

func TestUnlock(t *testing.T) {
	xredis.Mock()

	if err := l.Unlock(context.TODO()); err != nil {
		t.Fatalf("some error: %s", err)
	}

	t.Log("succeed")
}

func TestRefresh(t *testing.T) {
	xredis.Mock()

	if err := l.Refresh(context.TODO()); err != nil {
		t.Fatalf("some error: %s", err)
	}

	t.Logf("succeed")
}
