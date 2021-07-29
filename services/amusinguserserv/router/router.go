package router

import (
	"amusingx.fit/amusingx/services/amusinguserserv/api/generic"
	"amusingx.fit/amusingx/services/amusinguserserv/api/join"
	"amusingx.fit/amusingx/services/amusinguserserv/api/login"
	"amusingx.fit/amusingx/services/amusinguserserv/api/password"
	"amusingx.fit/amusingx/services/amusinguserserv/api/pong"
	"amusingx.fit/amusingx/services/amusinguserserv/api/regexp"
	"amusingx.fit/amusingx/services/amusinguserserv/api/verificationcode"
	"github.com/gorilla/mux"
	"net/http"
)

func Register(mux *mux.Router) {
	mux.HandleFunc("/v1/amusinguserserv/pong", pong.Pong).Methods(http.MethodGet)
	mux.HandleFunc("/v1/amusinguserserv/login/regexp", regexp.HandlerRegexp).Methods(http.MethodGet)
	mux.HandleFunc("/v1/amusinguserserv/login", login.HandlerLogin).Methods(http.MethodPost)
	mux.HandleFunc("/v1/amusinguserserv/join", join.HandleJoin).Methods(http.MethodPost)
	mux.HandleFunc("/v1/amusinguserserv/country/list", generic.HandlerCountryCodeList).Methods(http.MethodGet)
	mux.HandleFunc("/v1/amusinguserserv/verification_code", verificationcode.HandlerVerificationCode).Methods(http.MethodGet)
	mux.HandleFunc("/v1/amusinguserserv/verification_check", verificationcode.HandlerVerificationCheck).Methods(http.MethodGet)
	mux.HandleFunc("/v1/amusinguserserv/reset_password", password.HandlerResetPassword).Methods(http.MethodPost)
}
