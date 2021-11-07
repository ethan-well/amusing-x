package main

import (
	"amusingx.fit/amusingx/services/amusingcharonserv/rpcserver"
	"amusingx.fit/amusingx/services/amusingplutoserv/conf"
	"amusingx.fit/amusingx/services/amusingplutoserv/xredis"
	"github.com/ItsWewin/superfactory/powertrain"
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
	xredis.InitRedis(conf.Conf.RedisAddr, conf.Conf.RedisPassword, conf.Conf.RedisDB)
	InitRPCServer()
}

// 服务执行完毕时候执行
func DeferFunc() {
	xredis.CloseRedis()
	rpcserver.CloseRPCServer()
}

func InitRPCServer() {
	err := rpcserver.InitRPCServer()
	if err != nil {
		panic(err)
	}
}
