package authen

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/europa/currentinfo/currentuser"
	"amusingx.fit/amusingx/services/europa/rpcclient/ganymede"
	"amusingx.fit/amusingx/services/europa/session"
	"context"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
	"net/http"
)

type handlerFunc func(http.ResponseWriter, *http.Request)

func LoginAuthentication(f handlerFunc) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		if !authenticated(ctx, r) {
			rest.FailJsonResponse(w, xerror.Code.CForbidden, xerror.Message.CForbidden)
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

	sid, ok := session.GetSessionID(r)
	if !ok || len(sid) == 0 {
		return false
	}

	resp, err := ganymede.RPCClient.Client.IsLogin(ctx, &ganymedeservice.IsLoginRequest{SessionID: cook.Value})
	if err != nil {
		return false
	}

	if !resp.Login {
		return false
	}

	currentuser.SetCurrentUser(&currentuser.CurrentUser{
		Id:    resp.UserInfo.Id,
		Name:  resp.UserInfo.Name,
		Login: resp.UserInfo.Login,
		Email: resp.UserInfo.Email,
	})

	return resp.Login
}
