package rpcclient

import (
	"amusingx.fit/amusingx/services/ganymede/conf"
	"amusingx.fit/amusingx/services/ganymede/rpcclient/amusingxriskrpcserver"
	"fmt"
	"github.com/ItsWewin/superfactory/aerror"
	"google.golang.org/grpc"
)

func InitRPCClient() aerror.Error {
	grpcConf := conf.Conf.GrpcClient.Callisto
	conn, err := grpc.Dial(grpcConf.Addr, grpc.WithInsecure())
	if err != nil {
		return aerror.NewError(err, aerror.Code.SUnexpectedErr, fmt.Sprintf("grpc dial error: %s", err))
	}
	//defer conn.Close()

	err = amusingxriskrpcserver.InitAmusingXRiskServerRPCClient(conn)
	if err != nil {
		return aerror.NewError(err, err.(aerror.Error).Code(), "InitRPCServer failed")
	}

	return nil
}

func RPCClientClose() {
	amusingxriskrpcserver.CloseAmusingXRiskRPCClient()
}
