package main

import (
	"amusingx.fit/amusingx/services/ganymede/conf"
	"amusingx.fit/amusingx/services/ganymede/mysql/ganymede/model"
	"amusingx.fit/amusingx/services/ganymede/rpcclient"
	rpcserver2 "amusingx.fit/amusingx/services/ganymede/rpcserver"
	"amusingx.fit/amusingx/services/ganymede/session"
	"amusingx.fit/amusingx/services/ganymede/xredis"
	"github.com/ItsWewin/superfactory/powertrain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	powertrain.Run(conf.Conf, func(o *powertrain.Options) {
		o.InitFunc = InitFunc
		o.DeferFunc = DeferFunc
		o.InitHttpServer = false
		o.RegisterRouter = func(mux *mux.Router) {
			//router.Register(mux)
		}
	})

	conf.Conf.Print()
}

// InitFunc 服务初始化时候执行
func InitFunc() {
	model.InitMySQL()

	redis0 := conf.Conf.Redis.RedisO
	xredis.InitRedis(redis0.Addr, redis0.Password, redis0.DBNo)

	err := session.InitSessionManager("redis", "sid", 24*60*60)
	if err != nil {
		panic(err)
	}

	xErr := rpcclient.InitRPCClient()
	if xErr != nil {
		panic(xErr)
	}

	err = rpcserver2.InitRPCService()
	if err != nil {
		panic(err)
	}
}

// DeferFunc 服务执行完毕时候执行
func DeferFunc() {
	model.MysqlDisConnect()

	xredis.CloseRedis()

	// 关闭 rpc 连接
	rpcclient.RPCClientClose()
}
