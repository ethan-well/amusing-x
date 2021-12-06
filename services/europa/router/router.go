package router

import (
	"amusingx.fit/amusingx/services/europa/api/generic"
	"amusingx.fit/amusingx/services/europa/api/join"
	"amusingx.fit/amusingx/services/europa/api/login"
	"amusingx.fit/amusingx/services/europa/api/oauth/github"
	"amusingx.fit/amusingx/services/europa/api/password"
	"amusingx.fit/amusingx/services/europa/api/pong"
	"amusingx.fit/amusingx/services/europa/api/regexp"
	"amusingx.fit/amusingx/services/europa/api/verificationcode"
	"github.com/gorilla/mux"
	"net/http"
)

func Register(mux *mux.Router) {
	mux.HandleFunc("/v1/europa/pong", pong.Pong).Methods(http.MethodGet)
	mux.HandleFunc("/v1/europa/login/regexp", regexp.HandlerRegexp).Methods(http.MethodGet)
	mux.HandleFunc("/v1/europa/login", login.HandlerLogin).Methods(http.MethodPost)
	mux.HandleFunc("/v1/europa/join", join.HandleJoin).Methods(http.MethodPost)
	mux.HandleFunc("/v1/europa/country/list", generic.HandlerCountryCodeList).Methods(http.MethodGet)
	mux.HandleFunc("/v1/europa/verification_code", verificationcode.HandlerVerificationCode).Methods(http.MethodGet)
	mux.HandleFunc("/v1/europa/verification_check", verificationcode.HandlerVerificationCheck).Methods(http.MethodGet)
	mux.HandleFunc("/v1/europa/reset_password", password.HandlerResetPassword).Methods(http.MethodPost)
	mux.HandleFunc("/v1/europa/login/oauth", github.HandlerOauthLogin).Methods(http.MethodPost)
}
