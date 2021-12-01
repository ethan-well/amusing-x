package userrpcserver

import (
	"amusingx.fit/amusingx/protos/ganymede/userservice"
	"github.com/ItsWewin/superfactory/xerror"
	"google.golang.org/grpc"
	"sync"
)

type UserRPCClient struct {
	Client userservice.AmusingxUserServiceClient
	CC     *grpc.ClientConn
}

var RPCClient *UserRPCClient
var once sync.Once

func InitUserServerRPCClient(cc *grpc.ClientConn) *xerror.Error {
	if cc == nil {
		return xerror.NewError(nil, xerror.Code.BUnexpectedData, "cc is nil")
	}

	CloseUserServerPRCClient()

	once.Do(func() {
		RPCClient = &UserRPCClient{
			Client: userservice.NewAmusingxUserServiceClient(cc),
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
