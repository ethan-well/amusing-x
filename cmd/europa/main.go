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
	"log"
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

// 服务初始化时候执行
func InitFunc() {
	log.Println("amusing api sever listen:", conf.Conf.Addr)
	amusingxwebapi.InitMySQL()

	xredis.InitRedis(conf.Conf.RedisAddr, conf.Conf.RedisPassword, conf.Conf.RedisDB)

	err := session.InitSessionManager("redis", "sid", 24*60*60)
	if err != nil {
		panic(err)
	}

	initRPCClient()
}

// 服务执行完毕时候执行
func DeferFunc() {
	amusingxwebapi.MysqlDisConnect()

	xredis.CloseRedis()

	closeRPCClient()
}

func initRPCClient() {
	xErr := rpcclient.InitRiskServerRPCClient()
	if xErr != nil {
		panic(xErr)
	}

	xErr = rpcclient.InitUserServerRPCClient()
	if xErr != nil {
		panic(xErr)
	}
}

// 关闭 rpc 连接
func closeRPCClient() {
	rpcclient.RiskServerRPCClientClose()

	rpcclient.UserServerRPCClientClose()
}
