package joinapp

import (
	"amusingx.fit/amusingx/protos/ganymede/userservice"
	"amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/model"
	"amusingx.fit/amusingx/services/ganymede/xredis/lock"
	"context"
	"fmt"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
)

func CreateUser(ctx context.Context, r *userservice.JoinRequest) (*model.User, *xerror.Error) {
	var (
		err  *xerror.Error
		user = &model.User{
			Nickname: r.Nickname,
			Phone:    fmt.Sprintf("%s-%s", r.AreaCode, r.Phone),
			Password: r.Password,
		}
	)

	defer clearPassword(r, user)

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

	user, err = model.Create(ctx, user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func clearPassword(r *userservice.JoinRequest, u *model.User) {
	r.Password = ""
	u.Password = ""
	u.PasswordDigest = ""
}
