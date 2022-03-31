package product

import (
	"amusingx.fit/amusingx/protos/comm/response"
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/rpcclient/charonclient"
	"amusingx.fit/amusingx/services/comm"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerProductDelete(ctx context.Context, req *panguservice.ProductDeleteRequest) *response.CommResponse {
	resp, e := charonclient.Client.ProductDelete(ctx, req)
	var err aerror.Error
	if e != nil {
		err = aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "unexpected error")
	}

	return comm.Response(resp, err)
}

func HandlerProductsDelete(ctx context.Context, req *panguservice.ProductsDeleteRequest) *response.CommResponse {
	if req.Filter == "" {
		err := aerror.NewErrorf(nil, aerror.Code.CParamsError, "params 'filter' is invalid")
		return comm.Response(nil, err)
	}

	resp, e := charonclient.Client.ProductsDelete(ctx, req)
	var err aerror.Error
	if e != nil {
		err = aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "unexpected error")
	}

	return comm.Response(resp, err)
}
