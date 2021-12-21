package rpcclient

import (
	"fmt"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
	"google.golang.org/grpc"
)

func InitRPCClientConn(rpcServerAddr string) (*grpc.ClientConn, aerror.Error) {
	logger.Infof("InitRPCClient. addr: %s", rpcServerAddr)

	conn, err := grpc.Dial(rpcServerAddr, grpc.WithInsecure())
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.SUnexpectedErr, fmt.Sprintf("grpc dial error: %s", err))
	}

	return conn, nil
}
