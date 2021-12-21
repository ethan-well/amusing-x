package oauth

import (
	"amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/europa/rpcclient/ganymede"
	"amusingx.fit/amusingx/services/europa/session"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"net/http"
)

func Login(ctx context.Context, w http.ResponseWriter, provider, code string) (*ganymedeservice.OAuthLoginResponse, aerror.Error) {
	rpcReq := &ganymedeservice.OAuthLoginRequest{
		Provider: provider,
		Code:     code,
	}

	// grpc 调用 ganymede 服务登陆接口
	resp, err := ganymede.RPCClient.Client.OAuthLogin(ctx, rpcReq)
	if err != nil || !resp.Result || resp.LoginInfo == nil {
		return nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "login failed")
	}

	session.SetSession(w, resp.LoginInfo.SessionInfo)

	return resp, nil
}
