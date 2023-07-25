package rpcserver

import (
	service "amusingx.fit/amusingx/protos/user/service/user/proto"
	"amusingx.fit/amusingx/services/user/conf"
	"amusingx.fit/amusingx/services/user/rpcserver/userservice"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
	"google.golang.org/grpc"
	"net"
)

type RPCService struct {
	Server *grpc.Server
}

var Server *RPCService

func InitRPCService() aerror.Error {
	rpcService := grpc.NewServer()

	service.RegisterUserServiceServer(rpcService, new(userservice.UserService))

	grpcConf := conf.Conf.Server.GrpcServer
	lis, err := net.Listen(grpcConf.Network, grpcConf.Address)
	if err != nil {
		return aerror.NewError(err, aerror.Code.SUnexpectedErr, "RPC net listen failed")
	}

	err = rpcService.Serve(lis)
	if err != nil {
		return aerror.NewError(err, aerror.Code.SUnexpectedErr, "RPC serve failed")
	} else {
		logger.Info("Init RPCService succeed")
	}

	Server = &RPCService{Server: rpcService}

	return nil
}
