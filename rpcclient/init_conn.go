package rpcclient

import (
	"fmt"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
	"google.golang.org/grpc"
)

func InitRPCClientConn(rpcServerAddr string) (*grpc.ClientConn, *xerror.Error) {
	logger.Infof("InitRPCClient. addr: %s", rpcServerAddr)

	conn, err := grpc.Dial(rpcServerAddr, grpc.WithInsecure())
	if err != nil {
		return nil, xerror.NewError(err, xerror.Code.SUnexpectedErr, fmt.Sprintf("grpc dial error: %s", err))
	}

	return conn, nil
}
