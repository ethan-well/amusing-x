package main

import (
	"amusingx.fit/amusingx/services/callisto/conf"
	"amusingx.fit/amusingx/services/callisto/mysql/riskcontrol"
	"amusingx.fit/amusingx/services/callisto/rpcserver"
	"amusingx.fit/amusingx/services/callisto/xredis"
	"github.com/ItsWewin/superfactory/powertrain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
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

// 服务初始化时候执行
func InitFunc() {
	log.Println("amusing api sever listen:", conf.Conf.Addr)
	riskcontrol.InitMySQL()

	xredis.InitRedis(conf.Conf.RedisAddr, conf.Conf.RedisPassword, conf.Conf.RedisDB)

	InitRPCServer()
}

// 服务执行完毕时候执行
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
