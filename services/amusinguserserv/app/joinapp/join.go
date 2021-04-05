package joinapp

import (
	"amusingx.fit/amusingx/apistruct/amusinguserserv"
	"amusingx.fit/amusingx/services/amusinguserserv/modle"
	"context"
	"errors"
)

func CreateUser(ctx context.Context, u *amusinguserserv.JoinRequest) error {
	user := modle.User{
		Nickname: u.Nickname,
		Phone:    u.Phone,
		Password: u.Password,
	}

	existed, err := user.ExistedWithNicknameOrPhone(ctx)
	if err != nil {
		return err
	}

	if existed {
		return errors.New("nickname or phone existed")
	}

}
