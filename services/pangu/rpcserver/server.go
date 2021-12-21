package rpcserver

import (
	"amusingx.fit/amusingx/protos/pangu/service"
	"amusingx.fit/amusingx/services/pangu/conf"
	"amusingx.fit/amusingx/services/pangu/rpcserver/panguserver"
	"amusingx.fit/amusingx/services/pangu/rpcserver/servermiddleware"
	"context"
	"fmt"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"net"
	"net/http"
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
		logger.Errorf("net.Listen err: %s", err)
		return aerror.NewErrorf(err, aerror.Code.SUnexpectedErr, "RPC Listen failed:  %s", err)
	}

	go func() {
		err := rpcServ.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("0.0.0.0%s", rpcConf.Address),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SUnexpectedErr, "Failed to dial server: %s", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = panguservice.RegisterPanGuServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SUnexpectedErr, "Failed register panGu service handler err")
	}

	httpServ := conf.Conf.Server.HttpServer
	gwServer := &http.Server{
		Addr:    httpServ.Addr,
		Handler: gwmux,
	}

	logger.Infof("Serving gRPC-Gateway on %s", httpServ.Addr)
	err = gwServer.ListenAndServe()
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SUnexpectedErr, "Failed ListenAndServe")
	}

	return nil
}

func CloseRPCServer() {
	if Server != nil && Server.Server != nil {
		Server.Server.Stop()
	}
}
