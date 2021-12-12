package oauthlogin

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/ganymede/conf"
	"amusingx.fit/amusingx/services/ganymede/mysql/ganymededb/model"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestOauthLogin(t *testing.T) {
	conf.Mock()

	model.InitMySQL()

	req := &ganymedeservice.OAuthLoginRequest{
		Provider: "github",
		Code:     "8cc35348db4ea7efa6c4",
	}

	login, err := oauthLogin(context.Background(), req)
	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	t.Logf("succeed: %s", logger.ToJson(login))
}
