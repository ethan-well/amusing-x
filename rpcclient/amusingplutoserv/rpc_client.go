package amusingplutoserv

import (
	"amusingx.fit/amusingx/protos/amusingplutoservice/plutoservice"
	"amusingx.fit/amusingx/rpcclient"
	"amusingx.fit/amusingx/services/amusingplutoserv/conf"
	"github.com/ItsWewin/superfactory/xerror"
	"google.golang.org/grpc"
)

var (
	Conn   *grpc.ClientConn
	Client plutoservice.AmusingxPlutoServiceClient
)

func init() {
	if Client == nil {
		InitClient(conf.Conf.RPCAddress)
	}
}

func InitClient(addr string) *xerror.Error {
	if Conn != nil {
		return nil
	}

	var err *xerror.Error
	if len(addr) == 0 {
		return xerror.NewErrorf(nil, xerror.Code.SUnexpectedErr, "addr is blank")
	}

	Conn, err = rpcclient.InitRPCClientConn(addr)
	if err != nil {
		return err
	}

	Client = plutoservice.NewAmusingxPlutoServiceClient(Conn)

	return nil
}

func CloseClient(addr string) {
	if Conn != nil {
		Client = nil
		Conn.Close()
	}
}
