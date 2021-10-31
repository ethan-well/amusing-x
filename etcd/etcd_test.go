package etcd

import (
	"context"
	"fmt"
	"github.com/ItsWewin/superfactory/etcd"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	"testing"
	"time"
)

func TestSomeCode(t *testing.T) {
	etcd.InitEtcdClientV3(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})

	_, err := etcd.Client.Put(context.Background(), "key1", "v1-111111")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("succeed")
}

func TestLockCode(t *testing.T) {
	etcd.InitEtcdClientV3(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})

	// 生成一个30s超时的上下文
	timeout, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()
	// 获取租约
	response, e := etcd.Client.Grant(timeout, 30)
	if e != nil {
		log.Fatal(e.Error())
	}
	// 通过租约创建session
	session, e := concurrency.NewSession(etcd.Client, concurrency.WithLease(response.ID))
	if e != nil {
		log.Fatal(e.Error())
	}
	defer session.Close()
	// 通过session和锁前缀
	mutex := concurrency.NewMutex(session, "/lock")
	e = mutex.Lock(timeout)
	if e != nil {
		log.Fatal(e.Error())
	}

	defer mutex.Unlock(timeout)
}

func TestEtcdV2(t *testing.T) {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
	)
	// 客户端配置
	config = clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	}
	// 建立连接

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	kv := clientv3.NewKV(client)
	var putResp *clientv3.PutResponse
	if putResp, err = kv.Put(context.TODO(), "/school/class/students", "helios0", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(putResp.Header.Revision)
	if putResp.PrevKv != nil {
		fmt.Printf("prev Value: %s \n CreateRevision : %d \n ModRevision: %d \n Version: %d \n",
			string(putResp.PrevKv.Value), putResp.PrevKv.CreateRevision, putResp.PrevKv.ModRevision, putResp.PrevKv.Version)
	}
}
