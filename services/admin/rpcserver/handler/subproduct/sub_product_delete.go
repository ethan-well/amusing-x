package subproduct

import (
	"amusingx.fit/amusingx/protos/comm/response"
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/rpcclient/charonclient"
	"amusingx.fit/amusingx/services/comm"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerSubProductDelete(ctx context.Context, req *panguservice.SubProductDeleteRequest) *response.CommResponse {
	resp, e := charonclient.Client.SubProductDelete(ctx, req)
	var err aerror.Error
	if e != nil {
		err = aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "unexpected error")
	}

	return comm.Response(resp, err)
}

func HandlerSubProductsDelete(ctx context.Context, req *panguservice.SubProductsDeleteRequest) *response.CommResponse {
	if req.Filter == "" {
		err := aerror.NewErrorf(nil, aerror.Code.CParamsErr, "params 'filter' is invalid")
		return comm.Response(nil, err)
	}

	resp, e := charonclient.Client.SubProductsDelete(ctx, req)
	var err aerror.Error
	if e != nil {
		err = aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "unexpected error")
	}

	return comm.Response(resp, err)
}
