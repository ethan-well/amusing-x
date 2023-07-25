package authentication

import (
	service "amusingx.fit/amusingx/protos/user/service/user/proto"
	"amusingx.fit/amusingx/services/user/rpcserver/userservice/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerAuthentication(ctx context.Context, req *service.AuthenticationRequest) (*service.AuthenticationResponse, aerror.Error) {
	model := model.NewUserModel()
	var infos []*service.AuthenticationInfo
	for _, action := range req.Actions {
		ok := model.HavePermission(ctx, req.UserId, action, req.Server)
		infos = append(infos, &service.AuthenticationInfo{
			Action:        action,
			HasPermission: ok,
		})
	}

	return &service.AuthenticationResponse{Result: infos}, nil
}
