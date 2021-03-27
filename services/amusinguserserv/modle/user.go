package modle

import (
	amusinguser2 "amusingx.fit/amusingx/services/amusinguserserv/mysql/amusinguser"
	"context"
)

type User struct {
	ID             int64  `db:"id"`
	Nickname       string `db:"nickname"`
	Phone          string `db:"phone"`
	PasswordDigest string `db:"password_digest"`
	CreateTime     string `db:"create_time"`
	UpdateTime     string `db:"update_time"`
}

func FindUserByID(ctx context.Context, id int64) (*User, error) {
	user, err := amusinguser2.QueryUserByIdContext(ctx, id)
	if err != nil {
		return nil, err
	}

	u := &User{
		ID:             user.ID,
		Nickname:       user.Nickname,
		Phone:          user.Phone,
		PasswordDigest: user.PasswordDigest,
		CreateTime:     user.CreateTime,
		UpdateTime:     user.UpdateTime,
	}

	return u, nil
}
