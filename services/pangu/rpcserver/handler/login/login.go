package login

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service/ganymede/proto"
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu"
	"amusingx.fit/amusingx/services/europa/rpcclient/ganymede"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerLogin(ctx context.Context, req *panguservice.OAuthLoginRequest) (*panguservice.OAuthLoginResponse, aerror.Error) {
	// grpc 调用 ganymede 服务登陆接口
	rpcReq := &ganymedeservice.OAuthLoginRequest{
		Provider: req.Provider,
		Code:     req.Code,
	}

	resp, err := ganymede.RPCClient.Client.OAuthLogin(ctx, rpcReq)
	if err != nil || !resp.Result || resp.LoginInfo == nil {
		return nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "login failed")
	}

	return &panguservice.OAuthLoginResponse{
		Result:    true,
		LoginInfo: &panguservice.LoginInfo{SessionId: resp.LoginInfo.SessionId},
	}, nil
}
