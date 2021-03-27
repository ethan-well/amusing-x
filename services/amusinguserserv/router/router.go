package router

import (
	"amusingx.fit/amusingx/services/amusinguserserv/api/login"
	"amusingx.fit/amusingx/services/amusinguserserv/api/pong"
	"net/http"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/v1/pong", pong.Pong)
	mux.HandleFunc("/login", login.HandlerLogin)
}
