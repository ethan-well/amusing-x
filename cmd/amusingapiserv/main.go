package main

import (
	"amusingx.fit/amusingx/services/amusingapiserv/conf"
	"amusingx.fit/amusingx/services/amusingapiserv/router"
	"github.com/ItsWewin/superfactory/powertrain"
	"github.com/gorilla/mux"
	"log"
	"time"
)

var timeOut = time.Second * 30

func main() {
	powertrain.Run(conf.Conf, func(o *powertrain.Options) {
		o.InitFunc = InitFunc
		o.RegisterRouter = func(mux *mux.Router) {
			router.Register(mux)
		}
	})

	conf.Conf.Print()
}

func InitFunc() {
	log.Println("amusing api sever listen: ", conf.Conf.Addr)
}
