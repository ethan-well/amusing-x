package amusingxriskrpcserver

import (
	"amusingx.fit/amusingx/protos/amusingriskcontrolserv/loginrisk/loginrisk"
	"github.com/ItsWewin/superfactory/xerror"
	"google.golang.org/grpc"
	"sync"
)

type RPCServerClients struct {
	RiskServerRPCClient loginrisk.LoginRiskClient
	CC                  *grpc.ClientConn
}

var RiskServerRPCClient *RPCServerClients
var once sync.Once

func InitAmusingXRiskServerRPCClient(cc *grpc.ClientConn) *xerror.Error {
	if cc == nil {
		return xerror.NewError(nil, xerror.Code.BUnexpectedData, "cc is nil")
	}

	if RiskServerRPCClient != nil {
		RiskServerRPCClient.CC.Close()
		RiskServerRPCClient = nil
	}

	once.Do(func() {
		RiskServerRPCClient = &RPCServerClients{
			RiskServerRPCClient: loginrisk.NewLoginRiskClient(cc),
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
