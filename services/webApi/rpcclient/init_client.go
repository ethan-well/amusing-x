package rpcclient

import (
	"amusingx.fit/amusingx/services/webApi/conf"
	"amusingx.fit/amusingx/services/webApi/rpcclient/callisto"
	"amusingx.fit/amusingx/services/webApi/rpcclient/user"
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

	err = callisto.InitCallistoRPCClient(conn)
	if err != nil {
		return aerror.NewError(err, err.(aerror.Error).Code(), "InitRiskServerRPCClient failed")
	}

	return nil
}

func CallistoRPCClientClose() {
	callisto.CloseCallistoRPCClient()
}

func InitUserServerRPCClient() aerror.Error {
	ganymedeClient := conf.Conf.GrpcClient.UserServer

	conn, err := grpc.Dial(ganymedeClient.Addr, grpc.WithInsecure())
	if err != nil {
		return aerror.NewError(err, aerror.Code.SUnexpectedErr, "grpc.Dial failed")
	}

	err = user.InitUserServerRPCClient(conn)
	if err != nil {
		return aerror.NewError(err, err.(aerror.Error).Code(), "InitUserServerRPCClient failed")
	}

	return nil
}

func UserServerRPCClientClose() {
	user.CloseUserServerPRCClient()
}
