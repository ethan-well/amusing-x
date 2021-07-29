package amusinguser

import (
	"amusingx.fit/amusingx/mysqlstruct/amusinguser"
	"context"
	"database/sql"
	"github.com/ItsWewin/superfactory/xerror"
)

// 通过 id 查询用户
func QueryUserByIdContext(ctx context.Context, id int64) (*amusinguser.User, *xerror.Error) {
	user := &amusinguser.User{}

	query := `SELECT  id, nickname, phone, password_digest, create_time, update_time FROM user WHERE id= ?`
	err := AmusingUserDB.GetContext(ctx, user, query, id)
	if err != nil {
		return nil, xerror.NewError(err, xerror.Code.SSqlExecuteErr, "Query user failed. ")
	}

	return user, nil
}

func QueryUserByPhone(ctx context.Context, phone string) (*amusinguser.User, *xerror.Error) {
	user := &amusinguser.User{}

	query := `SELECT id, nickname, phone, password_digest, salt FROM user WHERE phone = ?`
	err := AmusingUserDB.GetContext(ctx, user, query, phone)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, xerror.NewError(err, xerror.Code.SSqlExecuteErr, "sql execute error")
	}

	return user, nil
}

func QueryUserByNicknameOrPhone(ctx context.Context, nickName, phone string) (*amusinguser.User, *xerror.Error) {
	user := &amusinguser.User{}

	query := `SELECT id, nickname, phone, password_digest FROM user WHERE nickname = ? OR phone = ?`
	err := AmusingUserDB.GetContext(ctx, user, query, nickName, phone)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, xerror.NewError(err, xerror.Code.SSqlExecuteErr, "sql execute error")
	}

	return user, nil
}

func Insert(ctx context.Context, user *amusinguser.User) (*amusinguser.User, *xerror.Error) {
	defer clearPasswordDigest(user)
	query := `INSERT INTO user (nickname, phone, password_digest, salt) VALUES (:nickname,:phone,:password_digest, :salt)`
	result, err := AmusingUserDB.NamedExecContext(ctx, query, user)
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

func clearPasswordDigest(user *amusinguser.User) {
	user.PasswordDigest = ""
}

func UpdatePassword(ctx context.Context, user *amusinguser.User) (*amusinguser.User, *xerror.Error) {
	defer clearPasswordDigest(user)

	query := `UPDATE user set password_digest = :password_digest, salt = :salt WHERE phone = :phone`
	_, err := AmusingUserDB.NamedExecContext(ctx, query, user)
	if err != nil {
		return nil, xerror.NewError(err, xerror.Code.SSqlExecuteErr, "reset password failed. ")
	}

	return user, nil
}
