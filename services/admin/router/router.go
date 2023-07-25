package router

import (
	"amusingx.fit/amusingx/services/admin/api"
	"github.com/gorilla/mux"
	"net/http"
)

func Register(mux *mux.Router) {
	mux.HandleFunc("/v1/pangu/pong", api.Pong).Methods(http.MethodGet)
}
