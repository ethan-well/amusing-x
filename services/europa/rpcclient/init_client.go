package rpcclient

import (
	"amusingx.fit/amusingx/services/europa/conf"
	"amusingx.fit/amusingx/services/europa/rpcclient/callisto"
	"amusingx.fit/amusingx/services/europa/rpcclient/ganymede"
	"fmt"
	"github.com/ItsWewin/superfactory/aerror"
	"google.golang.org/grpc"
)

func InitCallistoServerRPCClient() aerror.Error {
	callistoClient := conf.Conf.GrpcClient.Callisto
	conn, err := grpc.Dial(callistoClient.Addr, grpc.WithInsecure())
	if err != nil {
		return aerror.NewError(err, aerror.Code.SUnexpectedErr, fmt.Sprintf("grpc dial error: %s", err))
	}

	xErr := callisto.InitCallistoRPCClient(conn)
	if xErr != nil {
		return aerror.NewError(xErr, xerr.Code(), "InitRiskServerRPCClient failed")
	}

	return nil
}

func CallistoRPCClientClose() {
	callisto.CloseCallistoRPCClient()
}

func InitGanymedeRPCClient() aerror.Error {
	ganymedeClient := conf.Conf.GrpcClient.Ganymede

	conn, err := grpc.Dial(ganymedeClient.Addr, grpc.WithInsecure())
	if err != nil {
		return aerror.NewError(err, aerror.Code.SUnexpectedErr, "grpc.Dial failed")
	}

	xErr := ganymede.InitGanymedeRPCClient(conn)
	if xErr != nil {
		return aerror.NewError(xErr, xerr.Code(), "InitUserServerRPCClient failed")
	}

	return nil
}

func GanymedeRPCClientClose() {
	ganymede.CloseGanymedePRCClient()
}
