package main

import (
	"amusingx.fit/amusingx/protos/amusingriskcontrolserv/loginrisk/loginrisk"
	"amusingx.fit/amusingx/services/amusingriskserv/conf"
	"amusingx.fit/amusingx/services/amusingriskserv/rpcserver/loginriskserver"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
	"google.golang.org/grpc"
	"net"
)

func main() {
	conf.Conf.RPCNetwork = "tcp"
	conf.Conf.RPCAddress = "localhost:11002"

	rpcServ := grpc.NewServer()

	loginrisk.RegisterLoginRiskServer(rpcServ, new(loginriskserver.LoginRiskServer))

	logger.Infof("InitRPCServer. network: %s, addr: %s", conf.Conf.RPCNetwork, conf.Conf.RPCAddress)
	lis, err := net.Listen(conf.Conf.RPCNetwork, conf.Conf.RPCAddress)
	if err != nil {
		err = xerror.NewErrorf(err, xerror.Code.SUnexpectedErr, "RPC Listen failed:  %s", err)
		logger.Infof("err: %s", err)
		return
	}

	err = rpcServ.Serve(lis)
	if err != nil {
		err = xerror.NewErrorf(err, xerror.Code.SUnexpectedErr, "RPC Server failed: %s", err)
		logger.Infof("err: %s", err)
		return
	}

	logger.Infof("succeed")

	//rpcserv.InitRPCServer()
}
