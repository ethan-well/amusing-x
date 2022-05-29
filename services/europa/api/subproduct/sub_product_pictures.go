package subproduct

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/europa/app/subproductapp"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"net/http"
	"strconv"
	"strings"
)

func HandlerSubProductPictures(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	request, err := getAndValidParamsV2(r)
	if err != nil {
		logger.Errorf("params is invalid: %s", err.Error())

		rest.FailJsonResponse(w, aerror.Code.CUnexpectRequestDate, err.Message())
		return
	}

	resp, err := subproductapp.NewDomain(request).GetSubProducts(ctx)
	if err != nil {
		logger.Errorf("GetSubProducts failed: %s", err.Error())

		rest.FailJsonResponse(w, aerror.Code.CUnexpectRequestDate, err.Message())
		return
	}

	rest.SucceedJsonResponse(w, resp)
}

func getAndValidParamsV2(r *http.Request) (*proto.SubProductPicturesRequest, aerror.Error) {
	subProductIdsStr := r.FormValue("sub_product_ids")
	if subProductIdsStr == "" {
		return nil, aerror.NewError(nil, aerror.Code.CParamsError, "params 'sub_product_ids' is nil")
	}

	subProductIdsStrArr := strings.Split(subProductIdsStr, ",")
	var ids []int64
	for _, idStr := range subProductIdsStrArr {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return nil, aerror.NewError(err, aerror.Code.CParamsError, "params 'sub_product_ids' is invalid")
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return nil, aerror.NewError(nil, aerror.Code.CParamsError, "params 'sub_product_ids' is invalid")
	}

	return &proto.SubProductPicturesRequest{SubProductIds: ids}, nil
}
