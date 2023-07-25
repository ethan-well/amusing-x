package rpcserver

import (
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/admin/conf"
	"amusingx.fit/amusingx/services/admin/rpcserver/getway"
	"amusingx.fit/amusingx/services/admin/rpcserver/panguserver"
	"amusingx.fit/amusingx/services/admin/rpcserver/servermiddleware"
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
		servermiddleware.StreamServerInterceptorPanicRecover(), servermiddleware.StreamServerInterceptorAuthentication(),
	)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			servermiddleware.UnaryServerInterceptorPanicRecover(), servermiddleware.UnaryServerInterceptorAuthentication(),
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

	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(getway.CustomMatcher), runtime.WithForwardResponseOption(getway.HttpResponseModifier), runtime.WithMetadata(getway.GatewayMetadataAnnotator))
	// Register Greeter
	err = panguservice.RegisterPanGuServiceHandler(context.Background(), mux, conn)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SUnexpectedErr, "Failed register panGu service handler err")
	}

	err = http.ListenAndServe(conf.Conf.Server.HttpServer.Addr, prettier(mux))
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SUnexpectedErr, "Failed ListenAndServe")
	}
	return nil
}

func prettier(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Range", "categories 0-9/20")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Range")
		h.ServeHTTP(w, r)
	})
}

func CloseRPCServer() {
	if Server != nil && Server.Server != nil {
		Server.Server.Stop()
	}
}
