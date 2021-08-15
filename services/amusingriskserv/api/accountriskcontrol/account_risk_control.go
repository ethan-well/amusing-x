package accountriskcontrol

import (
	"amusingx.fit/amusingx/protos/amusingriskservice/riskservice"
	"amusingx.fit/amusingx/services/amusingriskserv/app/verificationcoderisk"
	"context"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
	"net/http"
)

func LoginControl(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req, err := getAndValidParams(r)
	if err != nil {
		rest.FailJsonResponse(w, err.Code, err.Message)
		return
	}

	err = verificationcoderisk.LoginRiskControl(ctx, req)
	if err != nil {
		logger.Errorf("verificationcoderisk.LoginRiskControl failed: %s", err)
		rest.FailJsonResponse(w, err.Code, err.Message)
		return
	}

	rest.SucceedJsonResponse(w, "succeed")
}

func getAndValidParams(r *http.Request) (*riskservice.LoginRiskRequest, *xerror.Error) {
	req := riskservice.LoginRiskRequest{}

	err := httputil.DecodeJsonBody(r, &req)
	if err != nil {
		return nil, xerror.NewError(err, xerror.Code.CParamsError, "decode request body failed")
	}

	xErr := req.Valid()
	if xErr != nil {
		return nil, xErr
	}

	return &req, nil
}
