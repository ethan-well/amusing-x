package join

import (
	"amusingx.fit/amusingx/apistruct/amusinguserserv"
	"amusingx.fit/amusingx/xerror"
	"encoding/json"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"net/http"
)

func HandleJoin(w http.ResponseWriter, r *http.Request) {
	params, err := getAndValidParams(r)
	if err != nil {
		rest.FailJsonResponse(w, xerror.Code.ParamsError, err.Error())
		return
	}

	err := createUser(ctx, params)
}

func getAndValidParams(r *http.Request) (*amusinguserserv.JoinRequest, error) {
	joinRequest := amusinguserserv.JoinRequest{}

	err := httputil.DecodeJsonBody(r, &joinRequest)
	if err != nil {
		return nil, err
	}

	err = joinRequest.Valid()
	if err != nil {
		return nil, err
	}

	return &joinRequest, nil
}
