package joinapp

import (
	"amusingx.fit/amusingx/apistruct/apiService"
	"amusingx.fit/amusingx/services/webApi/model"
	"amusingx.fit/amusingx/services/webApi/xredis/lock"
	"context"
	"fmt"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
)

func CreateUser(ctx context.Context, u *apiService.JoinRequest) (*model.User, aerror.Error) {
	var (
		err  aerror.Error
		user = &model.User{
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

		return nil, aerror.NewError(err, aerror.Code.SGetLockeErr, "Try again. ")
	}
	// 释放锁
	defer locker.Unlock(ctx)

	existed, err := user.ExistedWithNicknameOrPhone(ctx)
	if err != nil {
		return nil, err
	}

	if existed {
		return nil, aerror.NewError(err, aerror.Code.BDataIsNotAllow, "Phone or nickname is taken. ")
	}

	user, err = model.Create(ctx, user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func clearPassword(r *apiService.JoinRequest, u *model.User) {
	r.Password = ""
	u.Password = ""
	u.PasswordDigest = ""
}
