package rpcserver

import (
	"amusingx.fit/amusingx/protos/pangu/service"
	"amusingx.fit/amusingx/services/pangu/conf"
	"amusingx.fit/amusingx/services/pangu/rpcserver/panguserver"
	"amusingx.fit/amusingx/services/pangu/rpcserver/servermiddleware"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"net"
)

type RPCService struct {
	Server *grpc.Server
}

var Server *RPCService

func InitRPCServer() aerror.Error {
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
		logger.Infof("net.Listen err: %s", err)
		return aerror.NewErrorf(err, aerror.Code.SUnexpectedErr, "RPC Listen failed:  %s", err)
	}

	logger.Infof("net.Listen init succeed")

	err = rpcServ.Serve(lis)
	if err != nil {
		logger.Errorf("RPC Server init err: %s", err)
		return aerror.NewErrorf(err, aerror.Code.SUnexpectedErr, "RPC Server failed: %s", err)
	}

	logger.Infof("rpcServ.Serve succeed")

	Server = &RPCService{Server: rpcServ}

	return nil
}

func CloseRPCServer() {
	if Server != nil && Server.Server != nil {
		Server.Server.Stop()
	}
}
