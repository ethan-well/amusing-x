package router

import (
	"amusingx.fit/amusingx/services/amusingwebapiserv/api/generic"
	"amusingx.fit/amusingx/services/amusingwebapiserv/api/join"
	"amusingx.fit/amusingx/services/amusingwebapiserv/api/login"
	"amusingx.fit/amusingx/services/amusingwebapiserv/api/password"
	"amusingx.fit/amusingx/services/amusingwebapiserv/api/pong"
	"amusingx.fit/amusingx/services/amusingwebapiserv/api/regexp"
	"amusingx.fit/amusingx/services/amusingwebapiserv/api/verificationcode"
	"github.com/gorilla/mux"
	"net/http"
)

func Register(mux *mux.Router) {
	mux.HandleFunc("/v1/amusingxwebapiserv/pong", pong.Pong).Methods(http.MethodGet)
	mux.HandleFunc("/v1/amusingxwebapiserv/login/regexp", regexp.HandlerRegexp).Methods(http.MethodGet)
	mux.HandleFunc("/v1/amusingxwebapiserv/login", login.HandlerLogin).Methods(http.MethodPost)
	mux.HandleFunc("/v1/amusingxwebapiserv/join", join.HandleJoin).Methods(http.MethodPost)
	mux.HandleFunc("/v1/amusingxwebapiserv/country/list", generic.HandlerCountryCodeList).Methods(http.MethodGet)
	mux.HandleFunc("/v1/amusingxwebapiserv/verification_code", verificationcode.HandlerVerificationCode).Methods(http.MethodGet)
	mux.HandleFunc("/v1/amusingxwebapiserv/verification_check", verificationcode.HandlerVerificationCheck).Methods(http.MethodGet)
	mux.HandleFunc("/v1/amusingxwebapiserv/reset_password", password.HandlerResetPassword).Methods(http.MethodPost)
}
