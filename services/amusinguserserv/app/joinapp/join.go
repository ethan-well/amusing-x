package joinapp

import (
	"amusingx.fit/amusingx/apistruct/amusinguserserv"
	"amusingx.fit/amusingx/services/amusinguserserv/modle"
	"amusingx.fit/amusingx/services/amusinguserserv/xredis/lock"
	"amusingx.fit/amusingx/xerror"
	"context"
	"fmt"
	"github.com/ItsWewin/superfactory/logger"
)

func CreateUser(ctx context.Context, u *amusinguserserv.JoinRequest) (*modle.User, *xerror.Error) {
	var (
		err  *xerror.Error
		user = &modle.User{
			Nickname: u.Nickname,
			Phone:    fmt.Sprintf("%s-%s", u.AreaCode, u.Phone),
			Password: u.Password,
		}
	)

	defer clearPassword(u, user)

	// 获取锁
	key := fmt.Sprintf("%s-%s", user.Nickname, user.Phone)
	locker := lock.NewLocker(key, user.Nickname, 30)
	if err := locker.Lock(ctx); err != nil {
		logger.Errorf("\nlock failed: %s\n", err)

		return nil, xerror.NewError(err, xerror.Code.SGetLockeErr, "Try again. ")
	}
	// 释放锁
	defer locker.Unlock(ctx)

	existed, err := user.ExistedWithNicknameOrPhone(ctx)
	if err != nil {
		return nil, err
	}

	if existed {
		return nil, xerror.NewError(err, xerror.Code.BDataIsNotAllow, "Phone or nickname is taken. ")
	}

	user, err = modle.Create(ctx, user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func clearPassword(r *amusinguserserv.JoinRequest, u *modle.User) {
	r.Password = ""
	u.Password = ""
	u.PasswordDigest = ""
}
