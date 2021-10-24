package amusingplutoserv

import (
	"amusingx.fit/amusingx/protos/amusingplutoservice/plutoservice"
	"amusingx.fit/amusingx/services/amusingplutoserv/conf"
	"amusingx.fit/amusingx/tool/grpctester"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestPlutoService(t *testing.T) {
	clientConn, xErr := grpctester.InitRPCClient(conf.Conf.RPCAddress)
	if xErr != nil {
		t.Fatal(xErr)
	}

	client := plutoservice.NewAmusingxPlutoServiceClient(clientConn.Conn)

	response, err := client.Pong(context.Background(), &plutoservice.BlankParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("response: %s", logger.ToJson(response))
}
