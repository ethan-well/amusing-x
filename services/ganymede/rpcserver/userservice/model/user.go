package model

import (
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service/ganymede/proto"
	"amusingx.fit/amusingx/services/ganymede/authentication"
	"amusingx.fit/amusingx/services/ganymede/mysql/ganymededb/model"
	authentication2 "amusingx.fit/amusingx/services/ganymede/mysql/ganymededb/model/authentication"
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
	tx, e := model.GanymedeDB.Beginx()
	if e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "query user info failed")
	}
	defer tx.Rollback()

	user, err := model.QueryUserByID(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	if e := tx.Commit(); e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SRedisExecuteErr, "commit failed")
	}

	return user, nil
}

func (u *User) InitUserInfoByID(ctx context.Context, id int64) aerror.Error {
	tx, err := model.GanymedeDB.BeginTxx(ctx, nil)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "query user info failed")
	}
	defer tx.Rollback()

	tx.Commit()

	user, err := model.QueryUserByID(ctx, tx, id)
	if err != nil {
		return err.(aerror.Error)
	}

	if user == nil {
		return aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "user is not existed")
	}

	u.UserInfo = user

	return nil
}

func (u *User) GetUserRoles(ctx context.Context, id int64, service string) ([]string, aerror.Error) {
	roles, err := authentication2.QueryRolesByUserIdAndService(ctx, id, service)
	if err != nil {
		return nil, err
	}

	var roleNames []string
	for _, role := range roles {
		roleNames = append(roleNames, role.Name)
	}

	return roleNames, nil
}

func (u *User) HavePermission(ctx context.Context, userId int64, action, service string) bool {
	ok, _ := authentication.Permission(ctx, userId, action, service)
	return ok
}
