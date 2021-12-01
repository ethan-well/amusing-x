package rpcclient

import (
	"amusingx.fit/amusingx/services/europa/conf"
	"amusingx.fit/amusingx/services/europa/rpcclient/riskrpcserver"
	"amusingx.fit/amusingx/services/europa/rpcclient/userrpcserver"
	"fmt"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
	"google.golang.org/grpc"
)

func InitRiskServerRPCClient() *xerror.Error {
	logger.Infof("InitRPCClient. addr: %s", conf.Conf.RiskServiceRPCAddress)

	conn, err := grpc.Dial(conf.Conf.RiskServiceRPCAddress, grpc.WithInsecure())
	if err != nil {
		return xerror.NewError(err, xerror.Code.SUnexpectedErr, fmt.Sprintf("grpc dial error: %s", err))
	}

	xErr := riskrpcserver.InitRiskServerRPCClient(conn)
	if xErr != nil {
		return xerror.NewError(xErr, xErr.Code, "InitRiskServerRPCClient failed")
	}

	return nil
}

func RiskServerRPCClientClose() {
	riskrpcserver.CloseRiskServerRPCClient()
}

func InitUserServerRPCClient() *xerror.Error {
	logger.Infof("InitUserServerRPCClient. addr: %s", conf.Conf.UserServiceRPCAddress)

	conn, err := grpc.Dial(conf.Conf.UserServiceRPCAddress, grpc.WithInsecure())
	if err != nil {
		return xerror.NewError(err, xerror.Code.SUnexpectedErr, "grpc.Dial failed")
	}

	xErr := userrpcserver.InitUserServerRPCClient(conn)
	if xErr != nil {
		return xerror.NewError(xErr, xErr.Code, "InitUserServerRPCClient failed")
	}

	return nil
}

func UserServerRPCClientClose() {
	userrpcserver.CloseUserServerPRCClient()
}
