package ganymede

import (
	"amusingx.fit/amusingx/protos/ganymede/service/ganymede/proto"
	"github.com/ItsWewin/superfactory/aerror"
	"google.golang.org/grpc"
	"sync"
)

type UserRPCClient struct {
	Client ganymedeservice.GanymedeServiceClient
	CC     *grpc.ClientConn
}

var RPCClient *UserRPCClient
var once sync.Once

func InitGanymedeRPCClient(cc *grpc.ClientConn) aerror.Error {
	if cc == nil {
		return aerror.NewError(nil, aerror.Code.BUnexpectedData, "cc is nil")
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
