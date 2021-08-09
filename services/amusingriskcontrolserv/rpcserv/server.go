package rpcserv

import (
	"amusingx.fit/amusingx/protos/amusingriskcontrolserv/loginrisk/loginrisk"
	"amusingx.fit/amusingx/services/amusingriskcontrolserv/conf"
	"amusingx.fit/amusingx/services/amusingriskcontrolserv/rpcserv/loginriskserver"
	"github.com/ItsWewin/superfactory/xerror"
	"google.golang.org/grpc"
	"net"
)

func InitRPCServer() *xerror.Error {
	rpcServ := grpc.NewServer()

	loginrisk.RegisterLoginRiskServer(rpcServ, new(loginriskserver.LoginRiskServer))

	lis, err := net.Listen(conf.Conf.RPCNetwork, conf.Conf.RPCAddress)
	if err != nil {
		return xerror.NewErrorf(err, xerror.Code.SUnexpectedErr, "RPC Listen failed:  %s", err)
	}

	err = rpcServ.Serve(lis)
	if err != nil {
		return xerror.NewErrorf(err, xerror.Code.SUnexpectedErr, "RPC Server failed: %s", err)
	}

	return nil
}
