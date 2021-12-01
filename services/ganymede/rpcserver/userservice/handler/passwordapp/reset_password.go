package passwordapp

import (
	"amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/model"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

type ResetPasswordDomain interface {
}

type Domain struct {
	ResetPasswordInfo *userservice.ResetPasswordRequest
	UserModelInfo     *model.User
}

func NewDomain() *Domain {
	return &Domain{}
}

func (d *Domain) SetResetPasswordInfo(request *userservice.ResetPasswordRequest) *xerror.Error {
	if request == nil {
		return xerror.NewError(nil, xerror.Code.CUnexpectRequestDate, "Request is nil. ")
	}

	d.ResetPasswordInfo = request

	return nil
}

func (d *Domain) SetUserModelInfo(ctx context.Context) *xerror.Error {
	if d.ResetPasswordInfo == nil {
		return xerror.NewError(nil, xerror.Code.BUnexpectedBlankVariable, "reset password info is blank")
	}

	user, err := model.FindUserByPhone(ctx, d.ResetPasswordInfo.AreaCode, d.ResetPasswordInfo.Phone)
	if err != nil {
		return err
	}

	if user == nil {
		return xerror.NewError(err, xerror.Code.CUnexpectRequestDate, "账号密码错误")
	}

	d.UserModelInfo = user

	return nil
}

func (d *Domain) ResetPassword(ctx context.Context, password string) *xerror.Error {
	if d.ResetPasswordInfo == nil {
		return xerror.NewError(nil, xerror.Code.BUnexpectedBlankVariable, "reset password info is blank")
	}

	if d.UserModelInfo == nil {
		return xerror.NewError(nil, xerror.Code.BUnexpectedBlankVariable, "user model info is blank")
	}

	err := d.UserModelInfo.ResetPassword(ctx, password)
	if err != nil {
		return xerror.NewError(err, err.Code, err.Message)
	}

	return nil
}
