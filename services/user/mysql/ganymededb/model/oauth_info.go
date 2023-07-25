package model

import (
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	"context"
	"database/sql"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/jmoiron/sqlx"
)

func InsertOrUpdateOAuthInfo(ctx context.Context, tx *sqlx.Tx, oauth *ganymede.OauthInfo) (*ganymede.OauthInfo, aerror.Error) {
	query := `INSERT INTO oauth_info (user_id, provider, access_token, code, outer_id, login, avatar_url, email)
	VALUES (:user_id, :provider, :access_token, :code, :outer_id, :login, :avatar_url, :email)
	ON DUPLICATE KEY UPDATE access_token = :access_token,
							code = :code,
							avatar_url = :avatar_url,
							email = :email
	`
	result, err := tx.NamedExecContext(ctx, query, oauth)
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.SSqlExecuteErr, "Unexpected error. ")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.SSqlExecuteErr, "Unexpected error. ")
	}
	oauth.ID = id

	return oauth, nil
}

func UpdateOAuthUserID(ctx context.Context, tx *sqlx.Tx, userID int64, provider, login string) aerror.Error {
	query := `UPDATE oauth_info set user_id = ? WHERE provider = ? AND login = ?`

	_, err := tx.ExecContext(ctx, query, userID, provider, login)
	if err != nil {
		return aerror.NewError(err, aerror.Code.SSqlExecuteErr, "Unexpected error. ")
	}

	return nil
}

func QueryOAuthInfoByProviderAndLogin(ctx context.Context, tx *sqlx.Tx, provider, login string) (*ganymede.OauthInfo, aerror.Error) {
	query := `SELECT id, user_id, provider, access_token, code, outer_id, login, avatar_url, email, create_time, update_time
	FROM oauth_info
	WHERE provider = ? AND login = ?
	FOR UPDATE
	`

	var oauthInfo ganymede.OauthInfo
	err := tx.GetContext(ctx, &oauthInfo, query, provider, login)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, aerror.NewError(err, aerror.Code.SSqlExecuteErr, aerror.Message.SqlExecErr)
	}

	return &oauthInfo, nil
}
