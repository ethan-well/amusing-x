package router

import (
	"amusingx.fit/amusingx/services/webApi/api/generic"
	"amusingx.fit/amusingx/services/webApi/api/join"
	"amusingx.fit/amusingx/services/webApi/api/login"
	"amusingx.fit/amusingx/services/webApi/api/logout"
	"amusingx.fit/amusingx/services/webApi/api/oauth"
	"amusingx.fit/amusingx/services/webApi/api/oauth/github"
	"amusingx.fit/amusingx/services/webApi/api/password"
	"amusingx.fit/amusingx/services/webApi/api/pong"
	"amusingx.fit/amusingx/services/webApi/api/regexp"
	"amusingx.fit/amusingx/services/webApi/api/subproduct"
	"amusingx.fit/amusingx/services/webApi/api/verificationcode"
	"amusingx.fit/amusingx/services/webApi/authen"
	"github.com/gorilla/mux"
	"net/http"
)

func Register(mux *mux.Router) {
	mux.HandleFunc("/v1/europa/pong", authen.LoginAuthentication(pong.Pong)).Methods(http.MethodGet)
	mux.HandleFunc("/v1/europa/login/regexp", regexp.HandlerRegexp).Methods(http.MethodGet)
	mux.HandleFunc("/v1/europa/login", login.HandlerLogin).Methods(http.MethodPost)
	mux.HandleFunc("/v1/europa/join", join.HandleJoin).Methods(http.MethodPost)
	mux.HandleFunc("/v1/europa/country/list", generic.HandlerCountryCodeList).Methods(http.MethodGet)
	mux.HandleFunc("/v1/europa/verification_code", verificationcode.HandlerVerificationCode).Methods(http.MethodGet)
	mux.HandleFunc("/v1/europa/verification_check", verificationcode.HandlerVerificationCheck).Methods(http.MethodGet)
	mux.HandleFunc("/v1/europa/reset_password", password.HandlerResetPassword).Methods(http.MethodPost)
	mux.HandleFunc("/v1/europa/login/oauth", github.HandlerOauthLogin).Methods(http.MethodPost)
	mux.HandleFunc("/v1/europa/oauth/info", oauth.HandlerGetOAuthInfo).Methods(http.MethodGet)

	// 商品服务
	mux.HandleFunc("/v1/europa/product/categories", oauth.HandlerGetOAuthInfo).Methods(http.MethodGet)
	mux.HandleFunc("/v1/europa/product/sub_product/list", subproduct.HandlerSubProductsList).Methods(http.MethodGet)

	mux.HandleFunc("/v1/europa/product/sub_product/pictures", subproduct.HandlerSubProductPictures).Methods(http.MethodGet)

	mux.HandleFunc("/v1/europa/logout", logout.HandlerLogOut).Methods(http.MethodGet)
}
