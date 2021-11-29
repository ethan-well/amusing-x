package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	InitEtcdClientV3(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	lock, err := Lock("/mylock", 20*time.Second, 60*1000)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("succeed")
	
	err = lock.Unlock(context.Background())
	if err != nil {
		t.Fatal("unlock failed", err)
	}

	t.Log("del succeed")
}

func TestReLock(t *testing.T) {
	InitEtcdClientV3(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	
	keys := []map[string]time.Duration {
		{"/mylock": 20 * time.Second},
		{"/mylock": 20 * time.Second},
	}
	
	for id, m := range keys {
		for key, d := range m {
			_, err := Lock(key, 1* time.Second, int64(d.Seconds()))
			if id == 0 && err != nil {
				t.Fatal("lock failed")
			}
			
			if id == 1 && err == nil {
				t.Fatal("lock failed")
			}
		}
	}
}
