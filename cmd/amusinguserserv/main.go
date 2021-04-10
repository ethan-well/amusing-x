package main

import (
	"amusingx.fit/amusingx/services/amusinguserserv/conf"
	"amusingx.fit/amusingx/services/amusinguserserv/mysql/amusinguser"
	"amusingx.fit/amusingx/services/amusinguserserv/router"
	"amusingx.fit/amusingx/services/amusinguserserv/session"
	"amusingx.fit/amusingx/services/amusinguserserv/xredis"
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
	amusinguser.InitMySQL()

	xredis.InitRedis(conf.Conf.RedisAddr, conf.Conf.RedisPassword, conf.Conf.RedisDB)

	err := session.InitSessionManager("redis", "sid", 24*60*60)
	if err != nil {
		panic(err)
	}
}

// 服务执行完毕时候执行
func DeferFunc() {
	amusinguser.MysqlDisConnect()

	xredis.CloseRedis()
}
