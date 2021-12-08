package model

import (
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	"amusingx.fit/amusingx/services/ganymede/conf"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestInsertOAuthInfo(t *testing.T) {
	conf.Mock()
	Mock()

	oauth := &ganymede.OauthInfo{
		ID:          11,
		UseID:       1111,
		Provider:    "github",
		OuterID:     22222,
		Login:       "Login-3",
		AvatarUrl:   "AvatarUrl-333",
		Email:       "Email-333",
		AccessToken: "AccessToken-333",
		Code:        "Code-333",
		CreateTime:  "",
	}

	tx, err := GanymedeDB.Beginx()
	if err != nil {
		t.Fatal(err)
	}
	u, xErr := InsertOAuthInfo(context.TODO(), tx, oauth)
	if xErr != nil {
		t.Fatalf("some error: %s", xErr)
	}

	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("oauth: %s", logger.ToJson(u))
}

func TestQueryOAuthInfoByProviderAndLogin(t *testing.T) {
	conf.Mock()
	Mock()

	tx, err := GanymedeDB.BeginTxx(context.TODO(), nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	info, err := QueryOAuthInfoByProviderAndLogin(context.TODO(), tx, "github", "Login")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("info: %s", logger.ToJson(info))
}
