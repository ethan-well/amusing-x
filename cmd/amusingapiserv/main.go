package main

import (
	"amusingx.fit/amusingx/services/amusingapiserv/api/pong"
	"amusingx.fit/amusingx/services/amusingapiserv/conf"
	"amusingx.fit/amusingx/superfactory/powertrain"
	"net/http"
)


func main() {
	powertrain.Run(conf.Conf)

	conf.Conf.Print()

	http.HandleFunc("/v1/pong", pong.Pong)

	http.ListenAndServe(conf.Conf.Port, nil)
}