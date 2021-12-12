package authen

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/europa/rpcclient/ganymede"
	"amusingx.fit/amusingx/services/europa/session"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"net/http"
)

type handlerFunc func(http.ResponseWriter, *http.Request)

func LoginAuthentication(f handlerFunc) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		if !authenticated(ctx, r) {
			w.WriteHeader(http.StatusForbidden)
			//rest.FailJsonResponse(w, xerror.Code.CForbidden, xerror.Message.CForbidden)
			return
		}

		f(w, r)
	}
}

func authenticated(ctx context.Context, r *http.Request) bool {
	if SkipRouter(r) {
		logger.Infof("skip ...")
		return true
	}

	cook, err := r.Cookie(session.CookieName)
	if err != nil || cook == nil || len(cook.Value) == 0 {
		return false
	}

	resp, err := ganymede.RPCClient.Client.IsLogin(ctx, &ganymedeservice.IsLoginRequest{SessionID: cook.Value})
	if err != nil {
		return false
	}

	return resp.Login
}
