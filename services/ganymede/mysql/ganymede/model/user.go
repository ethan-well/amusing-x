package model

import (
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
	"github.com/jmoiron/sqlx"
)

func InsertUser(ctx context.Context, tx *sqlx.Tx, user *ganymede.User) (*ganymede.User, *xerror.Error) {
	query := `INSERT INTO user (name, login) VALUES (:name,:login)`
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

//func FindUserBy
