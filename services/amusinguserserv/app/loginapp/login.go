package loginapp

import (
	"amusingx.fit/amusingx/apistruct/amusinguserserv"
	"amusingx.fit/amusingx/services/amusinguserserv/model"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

type LoginDomain interface {
}

type Domain struct {
	LoginInfo     *amusinguserserv.LoginRequest
	UserModelInfo *model.User
}

func NewDomain() *Domain {
	return &Domain{}
}

// 设置登录信息
func (d *Domain) setLoginRequestInfo(loginRequest *amusinguserserv.LoginRequest) *xerror.Error {
	if loginRequest == nil {
		return xerror.NewError(nil, xerror.Code.BUnexpectedBlankVariable, "login request is blank")
	}
	d.LoginInfo = loginRequest

	return nil
}

// 设置用户的 DB 信息
func (d *Domain) setUserModelInfo(ctx context.Context) *xerror.Error {
	if d.LoginInfo == nil {
		return xerror.NewError(nil, xerror.Code.BUnexpectedBlankVariable, "login info is blank")
	}

	user, err := model.FindUserByPhone(ctx, d.LoginInfo.AreaCode, d.LoginInfo.Phone)
	if err != nil {
		return err
	}

	d.UserModelInfo = user

	return nil
}

func (d *Domain) Login(ctx context.Context) error {
	if d.UserModelInfo == nil || d.LoginInfo == nil {
		return xerror.NewError(nil, xerror.Code.BUnexpectedBlankVariable, "user_model_info or log_info is nil")
	}

	result, err := d.UserModelInfo.ComparePassword(d.LoginInfo.Password)
	if err != nil {
		return err
	}

	if !result {
		return xerror.NewError(nil, xerror.Code.CUnexpectRequestDate, "用户名或密码错误")
	}

	return nil
}
