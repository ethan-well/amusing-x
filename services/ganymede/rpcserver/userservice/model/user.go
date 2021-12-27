package model

import (
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/ganymede/authentication"
	"amusingx.fit/amusingx/services/ganymede/mysql/ganymededb/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

type User struct {
	UserInfo *ganymede.User
}

func NewUserModel() *User {
	return &User{}
}

var CurrentUser *User

func InitCurrentUser(info *ganymedeservice.UserInfo) {
	CurrentUser = &User{UserInfo: &ganymede.User{
		ID:    info.Id,
		Name:  info.Name,
		Login: info.Login,
	}}
}

func (u *User) GetUserInfoByID(ctx context.Context, id int64) (*ganymede.User, aerror.Error) {
	tx, err := model.GanymedeDB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "query user info failed")
	}
	defer tx.Rollback()

	tx.Commit()

	return model.QueryUserByID(ctx, tx, id)
}

func (u *User) HavePermission(ctx context.Context, action string) bool {
	ok, _ := authentication.HavePermission(ctx, u.UserInfo.ID, action)
	return ok
}
