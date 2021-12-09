package oauthlogin

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/ganymede/conf"
	"amusingx.fit/amusingx/services/ganymede/mysql/ganymededb/model"
	"context"
	"testing"
)

func TestOauthLogin(t *testing.T) {
	conf.Mock()

	model.InitMySQL()

	req := &ganymedeservice.OAuthLoginRequest{
		Provider: "github",
		Code:     "8cc35348db4ea7efa6c4",
	}

	err := oauthLogin(context.Background(), req)
	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	t.Logf("succeed")
}
