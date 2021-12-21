package callisto

import (
	"amusingx.fit/amusingx/protos/callisto/service"
	"github.com/ItsWewin/superfactory/aerror"
	"google.golang.org/grpc"
	"sync"
)

type RPCServerClients struct {
	Client riskservice.RiskServiceClient
	CC     *grpc.ClientConn
}

var RPCClient *RPCServerClients
var once sync.Once

func InitCallistoRPCClient(cc *grpc.ClientConn) aerror.Error {
	if cc == nil {
		return aerror.NewError(nil, aerror.Code.BUnexpectedData, "cc is nil")
	}

	CloseCallistoRPCClient()

	once.Do(func() {
		RPCClient = &RPCServerClients{
			Client: riskservice.NewRiskServiceClient(cc),
			CC:     cc,
		}
	})

	return nil
}

func CloseCallistoRPCClient() {
	if RPCClient != nil && RPCClient.CC != nil {
		RPCClient.CC.Close()
	}

	RPCClient = nil
}
