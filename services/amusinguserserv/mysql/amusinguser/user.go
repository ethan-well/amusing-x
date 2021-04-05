package amusinguser

import (
	"amusingx.fit/amusingx/mysqlstruct/amusinguser"
	"context"
	"database/sql"
)

// 通过 id 查询用户
func QueryUserByIdContext(ctx context.Context, id int64) (*amusinguser.User, error) {
	user := &amusinguser.User{}

	query := `SELECT  id, nickname, phone, password_digest, create_time, update_time FROM user WHERE id= ?`
	err := AmusingUserDB.GetContext(ctx, user, query, id)
	if err != nil {
		return nil, err
	}

	return user, err
}

func QueryUserByNicknameOrPhone(ctx context.Context, nickName, phone string) (*amusinguser.User, error) {
	user := &amusinguser.User{}

	query := `SELECT id, nickname, phone, password_digest FROM user where nickname = ? or phone = ?`
	err := AmusingUserDB.GetContext(ctx, user, query, nickName, phone)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}

	return user, err
}
