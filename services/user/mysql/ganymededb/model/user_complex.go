package model

import (
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	"context"
	"database/sql"
	"github.com/ItsWewin/superfactory/aerror"
)

// 通过 id 查询用户
func QueryUserByIdContext(ctx context.Context, id int64) (*ganymede.UserComplex, aerror.Error) {
	user := &ganymede.UserComplex{}

	query := `SELECT  id, nickname, phone, password_digest, create_time, update_time FROM user WHERE id= ?`
	err := UserDB.GetContext(ctx, user, query, id)
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.SSqlExecuteErr, "Query user failed. ")
	}

	return user, nil
}

func QueryUserByPhone(ctx context.Context, phone string) (*ganymede.UserComplex, aerror.Error) {
	user := &ganymede.UserComplex{}

	query := `SELECT id, nickname, phone, password_digest, salt FROM user WHERE phone = ?`
	err := UserDB.GetContext(ctx, user, query, phone)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, aerror.NewError(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	return user, nil
}

func QueryUserByNicknameOrPhone(ctx context.Context, nickName, phone string) (*ganymede.UserComplex, aerror.Error) {
	user := &ganymede.UserComplex{}

	query := `SELECT id, nickname, phone, password_digest FROM user WHERE nickname = ? OR phone = ?`
	err := UserDB.GetContext(ctx, user, query, nickName, phone)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, aerror.NewError(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	return user, nil
}

func InsertUserComplex(ctx context.Context, user *ganymede.UserComplex) (*ganymede.UserComplex, aerror.Error) {
	defer clearPasswordDigest(user)
	query := `INSERT INTO user (nickname, phone, password_digest, salt) VALUES (:nickname,:phone,:password_digest, :salt)`
	result, err := UserDB.NamedExecContext(ctx, query, user)
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.SSqlExecuteErr, "Unexpected error. ")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.SSqlExecuteErr, "Unexpected error. ")
	}
	user.ID = id

	return user, nil
}

func clearPasswordDigest(user *ganymede.UserComplex) {
	user.PasswordDigest = ""
}

func UpdatePassword(ctx context.Context, user *ganymede.UserComplex) (*ganymede.UserComplex, aerror.Error) {
	defer clearPasswordDigest(user)

	query := `UPDATE user set password_digest = :password_digest, salt = :salt WHERE phone = :phone`
	_, err := UserDB.NamedExecContext(ctx, query, user)
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.SSqlExecuteErr, "reset password failed. ")
	}

	return user, nil
}
