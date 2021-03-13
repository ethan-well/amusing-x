package planet

import (
	"amusingx.fit/amusingx/superfactory/planet/httpserver"
	"net/http"
)

type Planet struct {
	HttpServer *httpserver.HttpServer
}

func Init() *Planet {
	return &Planet{}
}

func (p *Planet) HttpServerInit(opt *http.Server, register func(mux *http.ServeMux)) {
	p.HttpServer = &httpserver.HttpServer{}

	p.HttpServer.Init(opt).Register(register).HttpStart()
}
