package rpcserver

import (
	"amusingx.fit/amusingx/protos/amusingriskcontrolserv/loginrisk/loginrisk"
	"amusingx.fit/amusingx/services/amusingriskserv/conf"
	"amusingx.fit/amusingx/services/amusingriskserv/rpcserver/loginriskserver"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
	"google.golang.org/grpc"
	"net"
)

type RPCService struct {
	Server *grpc.Server
}

var Server *RPCService

func InitRPCServer() *xerror.Error {
	rpcServ := grpc.NewServer()

	loginrisk.RegisterLoginRiskServer(rpcServ, new(loginriskserver.LoginRiskServer))

	logger.Infof("InitRPCServer. network: %s, addr: %s", conf.Conf.RPCNetwork, conf.Conf.RPCAddress)

	lis, err := net.Listen(conf.Conf.RPCNetwork, conf.Conf.RPCAddress)
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
