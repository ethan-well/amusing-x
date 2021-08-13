package regexp

import (
	"amusingx.fit/amusingx/services/amusingwebapiserv/regexp"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"net/http"
)

func HandlerRegexp(w http.ResponseWriter, r *http.Request) {
	regexps := regexp.AllRegexps()
	rest.SucceedJsonResponse(w, regexps)
}
