package router

import (
	"amusingx.fit/amusingx/services/amusinguserserv/api/generic"
	"amusingx.fit/amusingx/services/amusinguserserv/api/join"
	"amusingx.fit/amusingx/services/amusinguserserv/api/login"
	"amusingx.fit/amusingx/services/amusinguserserv/api/pong"
	"github.com/gorilla/mux"
	"net/http"
)

func Register(mux *mux.Router) {
	mux.HandleFunc("/v1/pong", pong.Pong).Methods(http.MethodGet)
	mux.HandleFunc("/v1/login", login.HandlerLogin).Methods(http.MethodPost)
	mux.HandleFunc("/v1/join", join.HandleJoin).Methods(http.MethodPost)
	mux.HandleFunc("/v1/country/list", generic.HandlerCountryCodeList).Methods(http.MethodGet)
}
