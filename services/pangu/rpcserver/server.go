package rpcserver

import (
	"amusingx.fit/amusingx/protos/pangu/service"
	"amusingx.fit/amusingx/services/pangu/conf"
	"amusingx.fit/amusingx/services/pangu/rpcserver/panguserver"
	"amusingx.fit/amusingx/services/pangu/rpcserver/servermiddleware"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"net"
)

type RPCService struct {
	Server *grpc.Server
}

var Server *RPCService

func InitRPCServer() *xerror.Error {
	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("[InitRPCServer] server is panic")
		}
	}()

	rpcServ := grpc.NewServer(grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		servermiddleware.StreamServerInterceptorPanicRecover(),
	)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			servermiddleware.UnaryServerInterceptorPanicRecover(),
		)))

	panguservice.RegisterPanGuServiceServer(rpcServ, new(panguserver.PanguServer))

	rpcConf := conf.Conf.Server.GrpcServer
	logger.Infof("InitRPCServer. network: %s, addr: %s", rpcConf.Network, rpcConf.Address)

	lis, err := net.Listen(rpcConf.Network, rpcConf.Address)
	if err != nil {
		return xerror.NewErrorf(err, xerror.Code.SUnexpectedErr, "RPC Listen failed:  %s", err)
	}

	err = rpcServ.Serve(lis)
	if err != nil {
		return xerror.NewErrorf(err, xerror.Code.SUnexpectedErr, "RPC Server failed: %s", err)
	}

	Server = &RPCService{Server: rpcServ}

	return nil
}

func CloseRPCServer() {
	if Server != nil && Server.Server != nil {
		Server.Server.Stop()
	}
}
