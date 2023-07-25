package main

import (
	"amusingx.fit/amusingx/services/product/conf"
	"amusingx.fit/amusingx/services/product/mysql/charon"
	"amusingx.fit/amusingx/services/product/rpcserver"
	"amusingx.fit/amusingx/services/product/xredis"
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

// InitFunc 服务初始化时候执行
func InitFunc() {
	redis0 := conf.Conf.Redis.RedisO
	xredis.InitRedis(redis0.Addr, redis0.Password, redis0.DBNo)
	charon.InitMySQL()

	InitRPCServer()
}

// DeferFunc 服务执行完毕时候执行
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
