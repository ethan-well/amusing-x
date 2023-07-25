package main

import (
	"amusingx.fit/amusingx/services/risk/conf"
	"amusingx.fit/amusingx/services/risk/mysql/riskcontrol"
	"amusingx.fit/amusingx/services/risk/rpcserver"
	"amusingx.fit/amusingx/services/risk/xredis"
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
			//router.Register(mux)
		}
	})

	conf.Conf.Print()
}

// InitFunc 服务初始化时候执行
func InitFunc() {
	riskcontrol.InitMySQL()

	redis0 := conf.Conf.Redis.RedisO
	xredis.InitRedis(redis0.Addr, redis0.Password, redis0.DBNo)

	InitRPCServer()
}

// DeferFunc 服务执行完毕时候执行
func DeferFunc() {
	riskcontrol.MysqlDisConnect()

	xredis.CloseRedis()

	rpcserver.CloseRPCServer()
}

func InitRPCServer() {
	err := rpcserver.InitRPCServer()
	if err != nil {
		panic(err)
	}
}
