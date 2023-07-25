package rpcserver

import (
	riskservice2 "amusingx.fit/amusingx/protos/callisto/service"
	"amusingx.fit/amusingx/services/risk/conf"
	"amusingx.fit/amusingx/services/risk/rpcserver/callistoservice"
	"github.com/ItsWewin/superfactory/aerror"
	"google.golang.org/grpc"
	"net"
)

type RPCService struct {
	Server *grpc.Server
}

var Server *RPCService

func InitRPCServer() aerror.Error {
	rpcServ := grpc.NewServer()

	riskservice2.RegisterRiskServiceServer(rpcServ, new(callistoservice.LoginRiskService))

	rpcConf := conf.Conf.Server.GrpcServer
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
