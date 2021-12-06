package main

import (
	"amusingx.fit/amusingx/services/europa/conf"
	"amusingx.fit/amusingx/services/europa/mysql/amusingxwebapi"
	"amusingx.fit/amusingx/services/europa/router"
	"amusingx.fit/amusingx/services/europa/rpcclient"
	"amusingx.fit/amusingx/services/europa/session"
	"amusingx.fit/amusingx/services/europa/xredis"
	"github.com/ItsWewin/superfactory/powertrain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	powertrain.Run(conf.Conf, func(o *powertrain.Options) {
		o.InitFunc = InitFunc
		o.DeferFunc = DeferFunc
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

	err := session.InitSessionManager("redis", "sid", 24*60*60)
	if err != nil {
		panic(err)
	}

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

	err = rpcclient.InitGanymedeRPCClient()
	if err != nil {
		panic(err)
	}
}

// 关闭 rpc 连接
func closeRPCClient() {
	rpcclient.CallistoRPCClientClose()

	rpcclient.GanymedeRPCClientClose()
}
