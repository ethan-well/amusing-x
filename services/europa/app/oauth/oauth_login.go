package oauth

import (
	userservice "amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/europa/rpcclient/ganymede"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

func Login(ctx context.Context, provider, code string) *xerror.Error {
	rpcReq := &ganymedeservice.OAuthLoginRequest{
		Provider: provider,
		Code:     code,
	}

	req, err := ganymede.RPCClient.Client.OAuthLogin(ctx, rpcReq)
	if err != nil {
		return xerror.NewError(err, xerror.Code.BUnexpectedData, "oauth failed")
	}

	if !req.Result {
		return xerror.NewError(err, xerror.Code.BCreateFileFailed, "oauth failed")
	}

	return nil
}
