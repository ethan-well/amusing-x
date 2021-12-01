package main

import (
	"amusingx.fit/amusingx/etcd"
	"amusingx.fit/amusingx/services/pluto/conf"
	"amusingx.fit/amusingx/services/pluto/mysql"
	"amusingx.fit/amusingx/services/pluto/rpcserver"
	"amusingx.fit/amusingx/services/pluto/xredis"
	"fmt"
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

// 服务初始化时候执行
func InitFunc() {
	mysql.InitMySQL()

	xredis.InitRedis(conf.Conf.RedisAddr, conf.Conf.RedisPassword, conf.Conf.RedisDB)

	etcd.InitEtcdClientV3(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	InitRPCServer()
	ServerInitLog()
}

// 服务执行完毕时候执行
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

func ServerInitLog() {
	fmt.Println("pluto")
}
