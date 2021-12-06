package rpcserver

import (
	"amusingx.fit/amusingx/protos/pluto/service"
	"amusingx.fit/amusingx/services/pluto/conf"
	plutoservice2 "amusingx.fit/amusingx/services/pluto/rpcserver/plutoservice"
	"amusingx.fit/amusingx/services/pluto/rpcserver/servermiddleware"
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

	//rpcServ := grpc.NewServer(grpc.ChainUnaryInterceptor(serverrecover.PanicRecover))
	rpcServ := grpc.NewServer(grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		servermiddleware.StreamServerInterceptorPanicRecover(),
	)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			servermiddleware.UnaryServerInterceptorPanicRecover(),
		)))

	plutoservice.RegisterAmusingxPlutoServiceServer(rpcServ, new(plutoservice2.PlutoService))

	rpcServConf := conf.Conf.Server.GrpcServer
	logger.Infof("InitRPCServer. network: %s, addr: %s", rpcServConf.Network, rpcServConf.Address)

	lis, err := net.Listen(rpcServConf.Network, rpcServConf.Address)
	if err != nil {
		return xerror.NewErrorf(err, xerror.Code.SUnexpectedErr, "RPC Listen failed:  %s", err)
	}

	err = rpcServ.Serve(lis)
	if err != nil {
		return xerror.NewErrorf(err, xerror.Code.SUnexpectedErr, "RPC Server failed: %s", err)
	}
	logger.Info("rpcServ init succeed")

	Server = &RPCService{Server: rpcServ}

	return nil
}

func CloseRPCServer() {
	if Server != nil && Server.Server != nil {
		Server.Server.Stop()
	}
}
