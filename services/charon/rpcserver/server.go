package rpcserver

import (
	charonservice "amusingx.fit/amusingx/protos/charons/service"
	"amusingx.fit/amusingx/services/charon/conf"
	"amusingx.fit/amusingx/services/charon/rpcserver/charonserver"
	"amusingx.fit/amusingx/services/charon/rpcserver/servermiddleware"
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
			logger.Errorf("[InitRPCServer] server is panic: %s", r)
		}
	}()

	rpcServ := grpc.NewServer(grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		servermiddleware.StreamServerInterceptorPanicRecover(),
	)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			servermiddleware.UnaryServerInterceptorPanicRecover(),
		)))

	charonservice.RegisterCharonServServer(rpcServ, new(charonserver.CharonServer))

	rpcConf := conf.Conf.Server.GrpcServer
	logger.Infof("InitRPCServer. network: %s, addr: %s", rpcConf.Network, rpcConf.Address)

	lis, err := net.Listen(rpcConf.Network, rpcConf.Address)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SUnexpectedErr, "RPC Listen failed:  %s", err)
	}

	err = rpcServ.Serve(lis)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SUnexpectedErr, "RPC Server failed: %s", err)
	}

	Server = &RPCService{Server: rpcServ}

	return nil
}

func CloseRPCServer() {
	if Server != nil && Server.Server != nil {
		Server.Server.Stop()
	}
}
