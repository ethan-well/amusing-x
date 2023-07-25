package main

import (
	"amusingx.fit/amusingx/rpcclient/charonclient"
	"amusingx.fit/amusingx/services/webApi/conf"
	"amusingx.fit/amusingx/services/webApi/mysql/amusingxwebapi"
	"amusingx.fit/amusingx/services/webApi/router"
	"amusingx.fit/amusingx/services/webApi/rpcclient"
	"amusingx.fit/amusingx/services/webApi/xredis"
	"github.com/ItsWewin/superfactory/powertrain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	powertrain.Run(conf.Conf, func(o *powertrain.Options) {
		o.InitFunc = InitFunc
		o.DeferFunc = DeferFunc
		o.InitHttpServer = true
		o.RegisterRouter = func(mux *mux.Router) {
			router.Register(mux)
		}
	})

	conf.Conf.Print()
}

// InitFunc 服务初始化时候执行
func InitFunc() {
	amusingxwebapi.InitMySQL()

	redis0 := conf.Conf.Redis.RedisO
	xredis.InitRedis(redis0.Addr, redis0.Password, redis0.DBNo)

	initRPCClient()
}

// DeferFunc 服务执行完毕时候执行
func DeferFunc() {
	amusingxwebapi.MysqlDisConnect()

	xredis.CloseRedis()

	closeRPCClient()
}

func initRPCClient() {
	err := rpcclient.InitCallistoServerRPCClient()
	if err != nil {
		panic(err)
	}

	err = rpcclient.InitUserServerRPCClient()
	if err != nil {
		panic(err)
	}

	err = charonclient.InitClient(conf.Conf.GrpcClient.Charon.Addr)
	if err != nil {
		panic(err)
	}
}

// 关闭 rpc 连接
func closeRPCClient() {
	rpcclient.CallistoRPCClientClose()

	rpcclient.UserServerRPCClientClose()
}
