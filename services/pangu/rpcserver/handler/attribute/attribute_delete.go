package attribute

import (
	"amusingx.fit/amusingx/protos/comm/response"
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/rpcclient/charonclient"
	"amusingx.fit/amusingx/services/comm"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerAttributeDelete(ctx context.Context, req *panguservice.AttributeDeleteRequest) *response.CommResponse {
	resp, e := charonclient.Client.AttributeDelete(ctx, req)
	var err aerror.Error
	if e != nil {
		err = aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "unexpected error")
	}

	return comm.Response(resp, err)
}
