package github

import (
	"amusingx.fit/amusingx/services/ganymede/conf"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	conf.Mock()

	gitOAth := conf.Conf.OAuth.Github
	oAuth := New(gitOAth.ClientID, gitOAth.ClientSecret, gitOAth.RedirectUrl)

	resp, err := oAuth.GetAccessToken(gitOAth.AccessTokenUrl, "1de6e15ed2fcc46f2b62")
	if err != nil {
		t.Fatal(err)
	}

	if resp.IsError() {
		t.Logf("request is errr: %s", resp.ErrorDescription)
	}

	t.Logf("resp: %s", logger.ToJson(resp))
}
