package xrpc

import (
	"amusingx.fit/amusingx/services/amusinguserserv/conf"
	"amusingx.fit/amusingx/services/amusinguserserv/xrpc/amusingxriskserver"
	"fmt"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
	"google.golang.org/grpc"
)

func InitRPCClient() *xerror.Error {

	logger.Infof("InitRPCServer. addr: %s", conf.Conf.RPCAddress)

	conn, err := grpc.Dial(conf.Conf.RPCAddress, grpc.WithInsecure())
	if err != nil {
		return xerror.NewError(err, xerror.Code.SUnexpectedErr, fmt.Sprintf("grpc dial error: %s", err))
	}
	defer conn.Close()

	xErr := amusingxriskserver.InitAmusingXRiskServerClient(conn)
	if err != nil {
		return xerror.NewError(xErr, xErr.Code, "InitRPCServer failed")
	}

	return nil
}
