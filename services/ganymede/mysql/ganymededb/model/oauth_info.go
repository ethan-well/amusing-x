package model

import (
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
	"github.com/jmoiron/sqlx"
)

func InsertOAuthInfo(ctx context.Context, tx *sqlx.Tx, oauth *ganymede.OauthInfo) (*ganymede.OauthInfo, *xerror.Error) {
	query := `INSERT INTO oauth_info (user_id, provider, access_token, code, outer_id, login, avatar_url, email)
	VALUES (:user_id, :provider, :access_token, :code, :outer_id, :login, :avatar_url, :email)`
	result, err := tx.NamedExecContext(ctx, query, oauth)
	if err != nil {
		return nil, xerror.NewError(err, xerror.Code.SSqlExecuteErr, "Unexpected error. ")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, xerror.NewError(err, xerror.Code.SSqlExecuteErr, "Unexpected error. ")
	}
	oauth.ID = id

	return oauth, nil
}
