package model

import (
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	"amusingx.fit/amusingx/services/ganymede/mysql/ganymededb/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

type UserModel struct {
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (u *UserModel) GetUserInfoByID(ctx context.Context, id int64) (*ganymede.User, aerror.Error) {
	tx, err := model.GanymedeDB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "query user info failed")
	}
	defer tx.Rollback()

	tx.Commit()

	return model.QueryUserByID(ctx, tx, id)
}
