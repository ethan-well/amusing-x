package amusinguser

import (
	"amusingx.fit/amusingx/mysqlstruct/amusinguser"
	"context"
)

// 通过 id 查询用户
func QueryUserByIdContext(ctx context.Context, id int64) (*amusinguser.User, error) {
	user := &amusinguser.User{}

	err := AmusingUserDB.GetContext(ctx, user, "SELECT * FROM user WHERE id= ?", id)
	if err != nil {
		return nil, err
	}

	return user, err
}
