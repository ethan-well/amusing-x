package router

import (
	"amusingx.fit/amusingx/services/amusingapiserv/api/home"
	"amusingx.fit/amusingx/services/amusingapiserv/api/pong"
	"github.com/gorilla/mux"
)

func Register(mux *mux.Router) {
	mux.HandleFunc("/v1/pong", pong.Pong)
	mux.HandleFunc("/", home.HandlerIndex)
}
