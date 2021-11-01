package etcd

import (
	clientv3 "go.etcd.io/etcd/client/v3"
)

var Client *clientv3.Client

func InitEtcdClientV3(conf clientv3.Config) {
	var err error
	Client, err = clientv3.New(conf)
	if err != nil {
		panic(err)
	}
}

func CloseEtcdClientV3() {
	if Client != nil {
		Client.Close()
		Client = nil
	}
}
