package model

import (
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	"context"
	"database/sql"
	"github.com/ItsWewin/superfactory/xerror"
	"github.com/jmoiron/sqlx"
)

func InsertUser(ctx context.Context, tx *sqlx.Tx, user *ganymede.User) (*ganymede.User, *xerror.Error) {
	query := `INSERT INTO user (name, login) VALUES (:name,:login)
		ON DUPLICATE KEY UPDATE login=:login;
	`
	result, err := tx.NamedExecContext(ctx, query, user)
	if err != nil {
		return nil, xerror.NewError(err, xerror.Code.SSqlExecuteErr, "Unexpected error. ")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, xerror.NewError(err, xerror.Code.SSqlExecuteErr, "Unexpected error. ")
	}

	user.ID = id

	return user, nil
}

func QueryUserByLogin(ctx context.Context, tx *sqlx.Tx, login string) (*ganymede.User, *xerror.Error) {
	query := `SELECT id, name, login FROM user WHERE login = ?`

	var user ganymede.User
	err := tx.GetContext(ctx, &user, query, login)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, xerror.NewErrorf(err, xerror.Code.SSqlExecuteErr, "")
	}

	return &user, nil
}

func QueryUserByID(ctx context.Context, tx *sqlx.Tx, id int64) (*ganymede.User, *xerror.Error) {
	query := `SELECT id, name, login FROM user WHERE id = ?`

	var user ganymede.User
	err := tx.GetContext(ctx, &user, query, id)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, xerror.NewErrorf(err, xerror.Code.SSqlExecuteErr, "")
	}

	return &user, nil
}
