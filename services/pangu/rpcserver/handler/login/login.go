package login

import (
	"amusingx.fit/amusingx/services/europa/rpcclient/ganymede"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerLogin(ctx context.Context) {
	// grpc 调用 ganymede 服务登陆接口
	resp, err := ganymede.RPCClient.Client.OAuthLogin(ctx, rpcReq)
	if err != nil || !resp.Result || resp.LoginInfo == nil {
		return nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "login failed")
	}
}
