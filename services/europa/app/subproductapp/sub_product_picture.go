package subproductapp

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/rpcclient/charonclient"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

type SubProductPictureDomain struct {
	Request *proto.SubProductPicturesRequest
}

func NewSubProductPictureDomain(request *proto.SubProductPicturesRequest) *SubProductPictureDomain {
	return &SubProductPictureDomain{Request: request}
}

func (d *SubProductPictureDomain) GetSubProductPictures(ctx context.Context) (*proto.SubProductPicturesResponse, aerror.Error) {
	if d.Request == nil {
		return nil, aerror.NewError(nil, aerror.Code.CParamsError, "request is nil")
	}

	resp, err := charonclient.Client.SubProductPictures(ctx, d.Request)
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.SUnexpectedErr, "request sub products error")
	}

	if resp == nil {
		return nil, aerror.NewError(nil, aerror.Code.BUnexpectedData, "response is nil")
	}

	return resp, nil
}
