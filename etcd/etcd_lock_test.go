package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	InitEtcdClientV3(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})

	lock, err := Lock("/mylock", 60*time.Second, 60*1000)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("succeed")

	time.Sleep(20 * time.Second)

	err = lock.Unlock(context.Background())
	if err != nil {
		t.Fatal("unlock failed", err)
	}

	t.Log("del succeed")
}
