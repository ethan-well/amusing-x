package oauth

import (
	"amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/europa/rpcclient/ganymede"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
	"net/http"
	"time"
)

func Login(ctx context.Context, w http.ResponseWriter, provider, code string) (*ganymedeservice.OAuthLoginResponse, *xerror.Error) {
	rpcReq := &ganymedeservice.OAuthLoginRequest{
		Provider: provider,
		Code:     code,
	}

	// grpc 调用 ganymede 服务登陆接口
	resp, err := ganymede.RPCClient.Client.OAuthLogin(ctx, rpcReq)
	if err != nil || !resp.Result || resp.LoginInfo == nil {
		return nil, xerror.NewErrorf(err, xerror.Code.BUnexpectedData, "login failed")
	}

	cookie := &http.Cookie{
		Name:       "sid",
		Value:      resp.LoginInfo.SessionId,
		Path:       "",
		Domain:     "",
		Expires:    time.Now().Add(7 * 24 * time.Hour),
		RawExpires: "",
		MaxAge:     0,
		Secure:     true,
		HttpOnly:   true,
		SameSite:   http.SameSiteStrictMode,
		Raw:        "",
		Unparsed:   nil,
	}

	http.SetCookie(w, cookie)

	return resp, nil
}
