package main

import (
	"amusingx.fit/amusingx/rpcclient/charonclient"
	"amusingx.fit/amusingx/rpcclient/ganymede"
	pangu2 "amusingx.fit/amusingx/rpcclient/pangu"
	"amusingx.fit/amusingx/services/pangu/conf"
	"amusingx.fit/amusingx/services/pangu/mysql/pangu"
	"amusingx.fit/amusingx/services/pangu/rpcserver"
	"amusingx.fit/amusingx/services/pangu/xredis"
	"github.com/ItsWewin/superfactory/powertrain"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	powertrain.Run(conf.Conf, func(o *powertrain.Options) {
		o.InitFunc = InitFunc
		o.DeferFunc = DeferFunc
		o.InitHttpServer = false
		//o.RegisterRouter = func(mux *mux.Router) {
		//	//router.Register(mux)
		//}
	})

	conf.Conf.Print()
}

// InitFunc 服务初始化时候执行
func InitFunc() {
	pangu.InitMySQL()

	redis0 := conf.Conf.Redis.RedisO
	xredis.InitRedis(redis0.Addr, redis0.Password, redis0.DBNo)

	err := charonclient.InitClient(conf.Conf.GrpcClient.Charon.Addr)
	if err != nil {
		panic(err)
	}

	err = ganymede.InitClient(conf.Conf.GrpcClient.Ganymede.Addr)
	if err != nil {
		panic(err)
	}

	err = rpcserver.InitRPCServer()
	if err != nil {
		panic(err)
	}
}

// DeferFunc 服务执行完毕时候执行
func DeferFunc() {
	pangu.MysqlDisConnect()

	xredis.CloseRedis()
}

func initRpcClient() {
	pangu2.InitClient(conf.Conf.GrpcClient.Charon.Addr)
}
