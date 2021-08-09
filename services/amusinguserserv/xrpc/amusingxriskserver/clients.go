package amusingxriskserver

import (
	"amusingx.fit/amusingx/protos/amusingriskcontrolserv/loginrisk/loginrisk"
	"github.com/ItsWewin/superfactory/xerror"
	"google.golang.org/grpc"
	"sync"
)

var RiskServerRPCClient loginrisk.LoginRiskClient
var once sync.Once

func InitAmusingXRiskServerClient(cc grpc.ClientConnInterface) *xerror.Error {
	if cc == nil {
		return xerror.NewError(nil, xerror.Code.BUnexpectedData, "cc is nil")
	}
	once.Do(func() {
		if RiskServerRPCClient == nil {
			RiskServerRPCClient = loginrisk.NewLoginRiskClient(cc)
		}
	})

	return nil
}
