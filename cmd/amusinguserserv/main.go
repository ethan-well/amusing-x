package main

import (
	"amusingx.fit/amusingx/services/amusinguserserv/conf"
	"amusingx.fit/amusingx/services/amusinguserserv/mysql/amusinguser"
	"amusingx.fit/amusingx/services/amusinguserserv/router"
	"github.com/ItsWewin/superfactory/powertrain"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	powertrain.Run(conf.Conf, func(o *powertrain.Options) {
		o.InitFunc = InitFunc
		o.DeferFunc = DeferFunc
		o.RegisterRouter = func(mux *http.ServeMux) {
			router.Register(mux)
		}
	})

	conf.Conf.Print()
}

// 服务初始化时候执行
func InitFunc() {
	log.Println("amusing api sever listen: ", conf.Conf.Addr)
	amusinguser.InitMySQL()
}

// 服务执行完毕时候执行
func DeferFunc() {
	amusinguser.MysqlDisConnect()
}
