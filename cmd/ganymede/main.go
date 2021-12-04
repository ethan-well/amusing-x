package main

import (
	"amusingx.fit/amusingx/services/ganymede/conf"
	"amusingx.fit/amusingx/services/ganymede/mysql/amusinguser"
	"amusingx.fit/amusingx/services/ganymede/rpcclient"
	rpcserver2 "amusingx.fit/amusingx/services/ganymede/rpcserver"
	"amusingx.fit/amusingx/services/ganymede/session"
	"amusingx.fit/amusingx/services/ganymede/xredis"
	"github.com/ItsWewin/superfactory/powertrain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	powertrain.Run(conf.ConfSect, func(o *powertrain.Options) {
		o.InitFunc = InitFunc
		o.DeferFunc = DeferFunc
		o.RegisterRouter = func(mux *mux.Router) {
			//router.Register(mux)
		}
	})

	conf.ConfSect.Print()
}

// 服务初始化时候执行
func InitFunc() {
	log.Println("amusing api sever listen:", conf.ConfSect.Addr)
	amusinguser.InitMySQL()

	xredis.InitRedis(conf.ConfSect.RedisAddr, conf.ConfSect.RedisPassword, conf.ConfSect.RedisDB)

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

// 服务执行完毕时候执行
func DeferFunc() {
	amusinguser.MysqlDisConnect()

	xredis.CloseRedis()

	// 关闭 rpc 连接
	rpcclient.RPCClientClose()
}
