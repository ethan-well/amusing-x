package attributemapping

import (
	"amusingx.fit/amusingx/protos/comm/response"
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/rpcclient/charonclient"
	"amusingx.fit/amusingx/services/comm"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerAttributeMappingDelete(ctx context.Context, req *panguservice.AttributeMappingDeleteRequest) *response.CommResponse {
	resp, e := charonclient.Client.AttributeMappingDelete(ctx, req)
	var err aerror.Error
	if e != nil {
		err = aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "unexpected error")
	}

	return comm.Response(resp, err)
}

func HandlerAttributeMappingsDelete(ctx context.Context, req *panguservice.AttributeMappingsDeleteRequest) *response.CommResponse {
	if req.Filter == "" {
		err := aerror.NewErrorf(nil, aerror.Code.CParamsError, "params 'filter' is invalid")
		return comm.Response(nil, err)
	}

	resp, e := charonclient.Client.AttributeMappingsDelete(ctx, req)
	var err aerror.Error
	if e != nil {
		err = aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "unexpected error")
	}

	return comm.Response(resp, err)
}
