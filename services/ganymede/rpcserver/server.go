package rpcserver

import (
	"amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/ganymede/conf"
	userservice2 "amusingx.fit/amusingx/services/ganymede/rpcserver/userservice"
	"github.com/ItsWewin/superfactory/xerror"
	"google.golang.org/grpc"
	"net"
)

type RPCService struct {
	Server *grpc.Server
}

var Server *RPCService

func InitRPCService() *xerror.Error {
	rpcService := grpc.NewServer()

	userservice.RegisterAmusingxUserServiceServer(rpcService, new(userservice2.UserService))

	lis, err := net.Listen(conf.Conf.RPCNetwork, conf.Conf.UserServiceRPCAddress)
	if err != nil {
		return xerror.NewError(err, xerror.Code.SUnexpectedErr, "RPC net listen failed")
	}

	err = rpcService.Serve(lis)
	if err != nil {
		return xerror.NewError(err, xerror.Code.SUnexpectedErr, "RPC serve failed")
	}

	Server = &RPCService{Server: rpcService}

	return nil
}
