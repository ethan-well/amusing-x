package riskrpcserver

import (
	"amusingx.fit/amusingx/protos/amusingriskcontrolserv/loginrisk/loginrisk"
	"github.com/ItsWewin/superfactory/xerror"
	"google.golang.org/grpc"
	"sync"
)

type RPCServerClients struct {
	Client loginrisk.LoginRiskClient
	CC     *grpc.ClientConn
}

var RPCClient *RPCServerClients
var once sync.Once

func InitRiskServerRPCClient(cc *grpc.ClientConn) *xerror.Error {
	if cc == nil {
		return xerror.NewError(nil, xerror.Code.BUnexpectedData, "cc is nil")
	}

	CloseRiskServerRPCClient()

	once.Do(func() {
		RPCClient = &RPCServerClients{
			Client: loginrisk.NewLoginRiskClient(cc),
			CC:     cc,
		}
	})

	return nil
}

func CloseRiskServerRPCClient() {
	if RPCClient != nil && RPCClient.CC != nil {
		RPCClient.CC.Close()
	}

	RPCClient = nil
}
