package subproduct

import (
	"amusingx.fit/amusingx/protos/europa/service/europa/proto"
	"amusingx.fit/amusingx/services/europa/app/subproductapp"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"net/http"
	"strconv"
)

func HandlerSubProductsList(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	request, err := getAndValidParams(r)
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

func getAndValidParams(r *http.Request) (*proto.SubProductListRequest, aerror.Error) {
	pageStr := r.FormValue("page")
	if pageStr == "" {
		pageStr = "0"
	}

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.CParamsError, "params 'page' is invalid")
	}

	limitStr := r.FormValue("limit")
	if limitStr == "" {
		limitStr = "0"
	}
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.CParamsError, "params 'limit' is invalid")
	}

	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > 50 {
		limit = 50
	}

	return &proto.SubProductListRequest{
		Page:   page,
		Limit:  limit,
		Offset: (page - 1) * limit,
	}, nil
}
