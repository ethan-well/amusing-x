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
		OuterID:     2222,
		Login:       "Login",
		AvatarUrl:   "AvatarUrl",
		Email:       "Email",
		AccessToken: "AccessToken",
		Code:        "Code",
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
