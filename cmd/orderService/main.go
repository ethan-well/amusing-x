package main

import (
	"amusingx.fit/amusingx/etcd"
	"amusingx.fit/amusingx/services/order/conf"
	"amusingx.fit/amusingx/services/order/mysql"
	"amusingx.fit/amusingx/services/order/rpcserver"
	"amusingx.fit/amusingx/services/order/xredis"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/powertrain"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	powertrain.Run(conf.Conf, func(o *powertrain.Options) {
		o.InitFunc = InitFunc
		o.DeferFunc = DeferFunc
		o.RegisterRouter = func(mux *mux.Router) {
			//router.Register(mux)
		}
	})

	conf.Conf.Print()
}

// InitFunc 服务初始化时候执行
func InitFunc() {
	mysql.InitMySQL()

	redis := conf.Conf.Redis.RedisO
	xredis.InitRedis(redis.Addr, redis.Password, redis.DBNo)

	etcd.InitEtcdClientV3(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	InitRPCServer()
}

// DeferFunc 服务执行完毕时候执行
func DeferFunc() {
	mysql.DisConnectMySQL()

	xredis.CloseRedis()

	etcd.CloseEtcdClientV3()

	rpcserver.CloseRPCServer()
}

func InitRPCServer() {
	err := rpcserver.InitRPCServer()
	if err != nil {
		logger.Errorf("init rpc server failed: %s", err)
		panic(err)
	}
}
