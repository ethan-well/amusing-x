package authentication

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service/ganymede/proto"
	"amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerAuthentication(ctx context.Context, req *ganymedeservice.AuthenticationRequest) (*ganymedeservice.AuthenticationResponse, aerror.Error) {
	model := model.NewUserModel()
	var infos []*ganymedeservice.AuthenticationInfo
	for _, action := range req.Actions {
		ok := model.HavePermission(ctx, req.UserId, action, req.Server)
		infos = append(infos, &ganymedeservice.AuthenticationInfo{
			Action:        action,
			HasPermission: ok,
		})
	}

	return &ganymedeservice.AuthenticationResponse{Result: infos}, nil
}
