package router

import (
	"amusingx.fit/amusingx/services/amusinguserserv/api/login"
	"amusingx.fit/amusingx/services/amusinguserserv/api/pong"
	"github.com/gorilla/mux"
	"net/http"
)

func Register(mux *mux.Router) {
	mux.HandleFunc("/v1/pong", pong.Pong).Methods(http.MethodGet)
	mux.HandleFunc("/login", login.HandlerLogin).Methods(http.MethodPost)
}
