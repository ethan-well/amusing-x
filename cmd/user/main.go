package main

import (
	"amusingx.fit/amusingx/services/user/conf"
	"amusingx.fit/amusingx/services/user/mysql/ganymededb/model"
	"amusingx.fit/amusingx/services/user/oauth"
	"amusingx.fit/amusingx/services/user/rpcclient"
	"amusingx.fit/amusingx/services/user/rpcserver"
	"amusingx.fit/amusingx/services/user/session"
	"amusingx.fit/amusingx/services/user/xredis"
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

	store := conf.Conf.OAuth.Store
	xErr := oauth.InitOAuthRedisStore(store.Redis.Addr, store.Redis.Password, store.Redis.DBNo, store.Prefix, store.MaxLifeTime)
	if xErr != nil {
		panic(xErr)
	}

	err := session.InitSessionManager("redis", session.MaxAge)
	if err != nil {
		panic(err)
	}

	xErr = rpcclient.InitRPCClient()
	if xErr != nil {
		panic(xErr)
	}

	xErr = rpcserver.InitRPCService()
	if xErr != nil {
		panic(xErr)
	}
}

// DeferFunc 服务执行完毕时候执行
func DeferFunc() {
	model.MysqlDisConnect()

	xredis.CloseRedis()

	// 关闭 rpc 连接
	rpcclient.RPCClientClose()
}
