package ganymede

import (
	"amusingx.fit/amusingx/protos/ganymede/service"
	"github.com/ItsWewin/superfactory/xerror"
	"google.golang.org/grpc"
	"sync"
)

type UserRPCClient struct {
	Client ganymedeservice.GanymedeServiceClient
	CC     *grpc.ClientConn
}

var RPCClient *UserRPCClient
var once sync.Once

func InitGanymedeRPCClient(cc *grpc.ClientConn) *xerror.Error {
	if cc == nil {
		return xerror.NewError(nil, xerror.Code.BUnexpectedData, "cc is nil")
	}

	CloseGanymedePRCClient()

	once.Do(func() {
		RPCClient = &UserRPCClient{
			Client: ganymedeservice.NewGanymedeServiceClient(cc),
			CC:     cc,
		}
	})

	return nil
}

func CloseGanymedePRCClient() {
	if RPCClient != nil && RPCClient.CC != nil {
		RPCClient.CC.Close()
	}

	RPCClient = nil
}
