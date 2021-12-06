package rpcclient

import (
	"amusingx.fit/amusingx/services/ganymede/conf"
	"amusingx.fit/amusingx/services/ganymede/rpcclient/amusingxriskrpcserver"
	"fmt"
	"github.com/ItsWewin/superfactory/xerror"
	"google.golang.org/grpc"
)

func InitRPCClient() *xerror.Error {
	grpcConf := conf.Conf.GrpcClient.Callisto
	conn, err := grpc.Dial(grpcConf.Addr, grpc.WithInsecure())
	if err != nil {
		return xerror.NewError(err, xerror.Code.SUnexpectedErr, fmt.Sprintf("grpc dial error: %s", err))
	}
	//defer conn.Close()

	xErr := amusingxriskrpcserver.InitAmusingXRiskServerRPCClient(conn)
	if err != nil {
		return xerror.NewError(xErr, xErr.Code, "InitRPCServer failed")
	}

	return nil
}

func RPCClientClose() {
	amusingxriskrpcserver.CloseAmusingXRiskRPCClient()
}
