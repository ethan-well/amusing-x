package main

import (
	"amusingx.fit/amusingx/services/amusingplutoserv/conf"
	"amusingx.fit/amusingx/services/amusingplutoserv/rpcserver"
	"fmt"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/powertrain"
	//_ "github.com/go-sql-driver/mysql"
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
	//mysql.InitMySQL()

	//xredis.InitRedis(conf.Conf.RedisAddr, conf.Conf.RedisPassword, conf.Conf.RedisDB)

	//etcd.InitEtcdClientV3(clientv3.Config{
	//	Endpoints:   []string{"localhost:2379"},
	//	DialTimeout: 5 * time.Second,
	//})

	InitRPCServer()
	ServerInitLog()
}

// 服务执行完毕时候执行
func DeferFunc() {
	//mysql.DisConnectMySQL()
	//
	//xredis.CloseRedis()
	//
	//etcd.CloseEtcdClientV3()
	//
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
	fmt.Println("amusingplutoserv")
}
