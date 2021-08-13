package userservice

import (
	"amusingx.fit/amusingx/protos/amusingxuserserv/userservice"
	"context"
)

type UserService struct {
	userservice.UnimplementedAmusingxUserServiceServer
}

func (s *UserService) Pong(ctx context.Context, blank *userservice.BlankParams) (*userservice.PongResponse, error) {
	return &userservice.PongResponse{ServerName: "amusing-x-user-service"}, nil
}
