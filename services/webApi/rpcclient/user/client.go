package user

import (
	service "amusingx.fit/amusingx/protos/user/service/user/proto"
	"github.com/ItsWewin/superfactory/aerror"
	"google.golang.org/grpc"
	"sync"
)

type RPCClientEntity struct {
	Client service.UserServiceClient
	CC     *grpc.ClientConn
}

var RPCClient *RPCClientEntity
var once sync.Once

func InitUserServerRPCClient(cc *grpc.ClientConn) aerror.Error {
	if cc == nil {
		return aerror.NewError(nil, aerror.Code.BUnexpectedData, "cc is nil")
	}

	CloseUserServerPRCClient()

	once.Do(func() {
		RPCClient = &RPCClientEntity{
			Client: service.NewUserServiceClient(cc),
			CC:     cc,
		}
	})

	return nil
}

func CloseUserServerPRCClient() {
	if RPCClient != nil && RPCClient.CC != nil {
		RPCClient.CC.Close()
	}

	RPCClient = nil
}
