package rpcserver

import (
	"amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/ganymede/conf"
	userservice2 "amusingx.fit/amusingx/services/ganymede/rpcserver/userservice"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
	"google.golang.org/grpc"
	"net"
)

type RPCService struct {
	Server *grpc.Server
}

var Server *RPCService

func InitRPCService() *xerror.Error {
	logger.Infof("InitRPCService")

	rpcService := grpc.NewServer()

	userservice.RegisterAmusingxUserServiceServer(rpcService, new(userservice2.UserService))

	grpcConf := conf.Conf.Server.GrpcServer
	lis, err := net.Listen(grpcConf.Network, grpcConf.Address)
	if err != nil {
		return xerror.NewError(err, xerror.Code.SUnexpectedErr, "RPC net listen failed")
	}
	logger.Infof("Listen")

	err = rpcService.Serve(lis)
	if err != nil {
		logger.Infof("Serve")
		return xerror.NewError(err, xerror.Code.SUnexpectedErr, "RPC serve failed")
	} else {
		logger.Info("Init RPCService succeed")
	}

	Server = &RPCService{Server: rpcService}

	return nil
}
