package amusingxriskrpcserver

import (
	"amusingx.fit/amusingx/protos/callisto/service"
	"github.com/ItsWewin/superfactory/aerror"
	"google.golang.org/grpc"
	"sync"
)

type RPCServerClients struct {
	RiskServerRPCClient riskservice.RiskServiceClient
	CC                  *grpc.ClientConn
}

var RiskServerRPCClient *RPCServerClients
var once sync.Once

func InitAmusingXRiskServerRPCClient(cc *grpc.ClientConn) aerror.Error {
	if cc == nil {
		return aerror.NewError(nil, aerror.Code.BUnexpectedData, "cc is nil")
	}

	if RiskServerRPCClient != nil {
		RiskServerRPCClient.CC.Close()
		RiskServerRPCClient = nil
	}

	once.Do(func() {
		RiskServerRPCClient = &RPCServerClients{
			RiskServerRPCClient: riskservice.NewRiskServiceClient(cc),
			CC:                  cc,
		}
	})

	return nil
}

func CloseAmusingXRiskRPCClient() {
	if RiskServerRPCClient != nil {
		if RiskServerRPCClient.CC != nil {
			RiskServerRPCClient.CC.Close()
		}

		RiskServerRPCClient = nil
	}
}
