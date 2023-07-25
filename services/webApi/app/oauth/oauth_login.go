package oauth

import (
	ganymedeservice "amusingx.fit/amusingx/protos/user/service/user/proto"
	"amusingx.fit/amusingx/services/webApi/rpcclient/user"
	"amusingx.fit/amusingx/services/webApi/session"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"net/http"
)

func Login(ctx context.Context, w http.ResponseWriter, provider, code, service string) (*ganymedeservice.OAuthLoginResponse, aerror.Error) {
	rpcReq := &ganymedeservice.OAuthLoginRequest{
		Provider: provider,
		Code:     code,
		Service:  service,
	}

	// grpc 调用 ganymede 服务登陆接口
	resp, err := user.RPCClient.Client.OAuthLogin(ctx, rpcReq)
	if err != nil || !resp.Result || resp.LoginInfo == nil {
		return nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "login failed")
	}

	session.SetSession(w, resp.LoginInfo.SessionInfo)

	return resp, nil
}
