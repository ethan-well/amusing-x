package ganymede

import (
	"amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/rpcclient"
	"github.com/ItsWewin/superfactory/xerror"
	"google.golang.org/grpc"
)

var (
	Conn   *grpc.ClientConn
	Client ganymedeservice.GanymedeServiceClient
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

	Client = ganymedeservice.NewGanymedeServiceClient(Conn)

	return nil
}

func CloseClient(addr string) {
	if Conn != nil {
		Client = nil
		Conn.Close()
	}
}
