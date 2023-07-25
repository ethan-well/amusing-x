package subproductapp

import (
	"amusingx.fit/amusingx/protos/europa/service/europa/proto"
	proto2 "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/rpcclient/charonclient"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

type Domain struct {
	Request *proto.SubProductListRequest
}

func NewDomain(request *proto.SubProductListRequest) *Domain {
	return &Domain{Request: request}
}

func (d *Domain) GetSubProducts(ctx context.Context) (*proto.SubProductListResponse, aerror.Error) {
	if d.Request == nil {
		return nil, aerror.NewError(nil, aerror.Code.CParamsErr, "request is nil")
	}

	req := &proto2.SubProductListRequest{
		Page:   d.Request.Page,
		Limit:  d.Request.Limit,
		Offset: d.Request.Offset,
		Filter: "{}",
	}

	subResp, err := charonclient.Client.SubProducts(ctx, req)
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.SUnexpectedErr, "request sub products error")
	}

	if subResp == nil {
		return nil, aerror.NewError(nil, aerror.Code.BUnexpectedData, "response is nil")
	}

	resp := &proto.SubProductListResponse{
		SubProducts: nil,
		Page:        subResp.Page,
		Limit:       subResp.Limit,
		Total:       subResp.Total,
		HasNext:     subResp.HasNext,
	}

	var subProductIds []int64
	var subProducts []*proto.SubProductDetail
	for _, p := range subResp.Data {
		subProductIds = append(subProductIds, p.Id)
		subProducts = append(subProducts, &proto.SubProductDetail{
			SubProductInfo: p,
		})
	}

	resp.SubProducts = subProducts

	// 获取 attribute 详情
	attrResp, err := charonclient.Client.AttributesWithSubProduct(ctx, &proto2.AttributeListRequest{SubProductIds: subProductIds})
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.SUnexpectedErr, "query attributes with sub product error")
	}

	if attrResp == nil {
		return resp, nil
	}

	subProductIdAttrsMap := make(map[int64][]*proto2.AttributeWithSubProduct)
	for _, attr := range attrResp.Data {
		if _, ok := subProductIdAttrsMap[attr.SubProductId]; !ok {
			subProductIdAttrsMap[attr.SubProductId] = []*proto2.AttributeWithSubProduct{attr}
		} else {
			subProductIdAttrsMap[attr.SubProductId] = append(subProductIdAttrsMap[attr.SubProductId], attr)
		}
	}

	for _, p := range resp.SubProducts {
		if attrs, ok := subProductIdAttrsMap[p.SubProductInfo.Id]; ok {
			p.Attributes = attrs
		}
	}

	return resp, nil
}
