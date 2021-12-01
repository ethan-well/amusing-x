package amusingcharonserver

import (
	"amusingx.fit/amusingx/protos/charons/charonservice"
	"amusingx.fit/amusingx/rpcclient"
	"github.com/ItsWewin/superfactory/xerror"
	"google.golang.org/grpc"
)

var (
	Conn   *grpc.ClientConn
	Client charonservice.CharonServClient
)

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

	Client = charonservice.NewCharonServClient(Conn)

	return nil
}

func CloseClient(addr string) {
	if Conn != nil {
		Client = nil
		Conn.Close()
	}
}
