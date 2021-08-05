package router

import (
	"amusingx.fit/amusingx/services/amusingapiserv/api/pong"
	"amusingx.fit/amusingx/services/amusingxwebapiserv/api/mausinguserapiproxy/generic"
	"amusingx.fit/amusingx/services/amusingxwebapiserv/api/mausinguserapiproxy/join"
	"amusingx.fit/amusingx/services/amusingxwebapiserv/api/mausinguserapiproxy/login"
	"amusingx.fit/amusingx/services/amusingxwebapiserv/api/mausinguserapiproxy/password"
	"amusingx.fit/amusingx/services/amusingxwebapiserv/api/mausinguserapiproxy/regexp"
	"amusingx.fit/amusingx/services/amusingxwebapiserv/api/mausinguserapiproxy/verificationcode"
	"github.com/gorilla/mux"
	"net/http"
)

func Register(mux *mux.Router) {
	mux.HandleFunc("/v1/amusingxwebapiserv/pong", pong.Pong).Methods(http.MethodGet)
	mux.HandleFunc("/v1/amusingxwebapiserv/account/login/regexp", regexp.HandlerRegexp).Methods(http.MethodGet)
	mux.HandleFunc("/v1/amusingxwebapiserv/account/login", login.HandlerLogin).Methods(http.MethodPost)
	mux.HandleFunc("/v1/amusingxwebapiserv/account/join", join.HandleJoin).Methods(http.MethodPost)
	mux.HandleFunc("/v1/amusingxwebapiserv/account/country/list", generic.HandlerCountryCodeList).Methods(http.MethodGet)
	mux.HandleFunc("/v1/amusingxwebapiserv/account/verification_code", verificationcode.HandlerVerificationCode).Methods(http.MethodGet)
	mux.HandleFunc("/v1/amusingxwebapiserv/account/verification_check", verificationcode.HandlerVerificationCheck).Methods(http.MethodGet)
	mux.HandleFunc("/v1/amusingxwebapiserv/account/reset_password", password.HandlerResetPassword).Methods(http.MethodPost)
}
