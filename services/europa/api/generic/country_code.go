package generic

import (
	"amusingx.fit/amusingx/services/europa/mysql/amusingxwebapi"
	"context"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"net/http"
)

func HandlerCountryCodeList(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	codeList, err := amusingxwebapi.QueryCountryCodes(ctx)
	if err != nil {
		w.Write([]byte("some error"))
		return
	}

	_, err = rest.SucceedJsonResponse(w, codeList)
	if err != nil {
		logger.Errorf("HandlerCountryCodeList failed")
	}
}
