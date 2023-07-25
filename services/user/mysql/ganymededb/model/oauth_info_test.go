package model

import (
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	"amusingx.fit/amusingx/services/user/conf"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
	"time"
)

func TestInsertOAuthInfo(t *testing.T) {
	conf.Mock()
	Mock()

	oauth := &ganymede.OauthInfo{
		ID:          11,
		UseID:       1111,
		Provider:    "github",
		OuterID:     22222,
		Login:       "Login",
		AvatarUrl:   "AvatarUrl-333",
		Email:       "Email-333",
		AccessToken: "AccessToken-333",
		Code:        "Code-333",
		CreateTime:  "",
	}

	tx, err := UserDB.Beginx()
	if err != nil {
		t.Fatal(err)
	}
	u, xErr := InsertOrUpdateOAuthInfo(context.TODO(), tx, oauth)
	if xErr != nil {
		t.Fatalf("some error: %s", xErr)
	}

	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("oauth: %s", logger.ToJson(u))
}

func TestSelectForUpdatte(t *testing.T) {
	conf.Mock()
	Mock()

	ctx := context.Background()

	tx, err := UserDB.BeginTxx(ctx, nil)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	defer tx.Rollback()

	selectSql := `SELECT id, user_id, provider, access_token
	FROM oauth_info
	WHERE provider = ? AND login = ?
	FOR UPDATE`

	//selectSql2 := `SELECT count(*) FROM oauth_info WHERE provider = ? AND login = ? FOR UPDATE`

	var oauth ganymede.OauthInfo
	//err = tx.QueryRowxContext(ctx, selectSql, "github", "Login").Scan(&oauth.ID, &oauth.UseID, &oauth.Provider, &oauth.AccessToken)
	//if err != nil {
	//	t.Fatal(err)
	//}

	err = tx.GetContext(ctx, &oauth, selectSql, "github", "Login2")
	if err != nil {
		t.Fatal(err)
	}

	update := `UPDATE oauth_info set login = ? WHERE provider = ? AND login = ?`

	if &oauth != nil {
		t.Logf("rowCount > 0: %s", logger.ToJson(oauth))

		_, err := tx.ExecContext(ctx, update, "login-xxxxxxxxxx", "github", "Login2")
		if err != nil {
			t.Logf("tx update error: %s", err)
			return
		}

		logger.Info("update succeed")
		return
	} else {
		t.Logf("rowCount < 0")
		return
	}

	err = tx.Commit()
	if err != nil {
		t.Logf("commit errr: %s", err)
	}
}

func TestQueryOAuthInfoByProviderAndLogin(t *testing.T) {
	conf.Mock()
	Mock()

	ctx := context.Background()

	tx, err := UserDB.BeginTxx(ctx, nil)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	defer tx.Rollback()

	oauth, xErr := QueryOAuthInfoByProviderAndLogin(ctx, tx, "github", "Login")
	if xErr != nil {
		t.Logf("QueryOAuthInfoByProviderAndLogin err: %s", xErr)
		return
	}

	if oauth.UseID > 0 {
		oauth.AccessToken = "889977554444"
		_, xErr = InsertOrUpdateOAuthInfo(ctx, tx, oauth)
		if xErr != nil {
			t.Logf("InsertOAuthInfo err: %s", xErr)
			return
		}
	}

	time.Sleep(10 * time.Minute)

	err = tx.Commit()
	if err != nil {
		t.Logf("tx.Commit() err: %s", err)
		return
	}

	t.Logf("succed")
}
