package authen

import (
	"github.com/ItsWewin/superfactory/set/stringSet"
	"net/http"
	"net/url"
)

var skipAuthenticationRouters = []string{
	"/v1/europa/pong",
	"/v1/europa/login",
	"/v1/europa/login/regexp",
	"/v1/europa/country/list",
	"/v1/europa/verification_code",
	"/v1/europa/verification_check",
	"/v1/europa/reset_password",
	"/v1/europa/login/oauth",
	"/v1/europa/oauth/info",
}

var SkipRouterSet = &stringSet.StringSet{}

func init() {
	SkipRouterSet = stringSet.NewSet(skipAuthenticationRouters...)
}

func SkipRouter(r *http.Request) bool {
	u, err := url.Parse(r.URL.Path)
	if err != nil {
		return false
	}

	return SkipRouterSet.HasMember(u.Path)
}
