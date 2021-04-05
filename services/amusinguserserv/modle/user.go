package modle

import (
	amusinguser2 "amusingx.fit/amusingx/services/amusinguserserv/mysql/amusinguser"
	"context"
	"errors"
)

type User struct {
	ID             int64  `db:"id"`
	Nickname       string `db:"nickname"`
	Phone          string `db:"phone"`
	PasswordDigest string `db:"password_digest"`
	Password       string
	CreateTime     string `db:"create_time"`
	UpdateTime     string `db:"update_time"`
}

func (u *User) ExistedWithNicknameOrPhone(ctx context.Context) (bool, error) {
	if len(u.Nickname) == 0 || len(u.Phone) == 0 {
		return false, errors.New("nickname or phone number is blank")
	}

	udb, err := amusinguser2.QueryUserByNicknameOrPhone(ctx, u.Nickname, u.Phone)
	if err != nil {
		return false, err
	}

	switch {
	case udb == nil:
		return false, nil
	case udb.Nickname == u.Nickname:
		return true, errors.New("nickname is token")
	case udb.Phone == u.Phone:
		return true, errors.New("phone number is ")
	default:
		return true, errors.New("user is existed")
	}
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
