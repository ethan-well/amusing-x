package login

import (
	"amusingx.fit/amusingx/apistruct/amusinguserserv"
	"amusingx.fit/amusingx/services/amusinguserserv/modle"
	"context"
	"encoding/json"
	"github.com/ItsWewin/superfactory/logger"
	"io/ioutil"
	"net/http"
)

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	params, err := getAndValidParams(r)
	if err != nil {
		logger.Errorf("%s get and valid prams failed")
		w.Write([]byte(err.Error()))

		return
	}

	date, _ := json.Marshal(params)

	logger.Infof("params: %s", string(date))

	user, err := modle.FindUserByID(ctx, 1)
	if err != nil {
		w.Write([]byte(err.Error()))

		return
	}

	logger.Infof("user: %v\n", user)

	w.Write(date)
	return
}

func getAndValidParams(r *http.Request) (*amusinguserserv.LoginRequest, error) {
	login := &amusinguserserv.LoginRequest{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, login)

	return login, nil
}
