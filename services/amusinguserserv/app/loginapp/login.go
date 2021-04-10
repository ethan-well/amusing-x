package loginapp

import (
	"amusingx.fit/amusingx/apistruct/amusinguserserv"
	"amusingx.fit/amusingx/services/amusinguserserv/model"
	"amusingx.fit/amusingx/services/amusinguserserv/session"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
	uuid "github.com/satori/go.uuid"
)

type LoginDomain interface {
}

type Domain struct {
	LoginInfo     *amusinguserserv.LoginRequest
	UserModelInfo *model.User
	SessionID     string
}

func NewDomain() *Domain {
	return &Domain{}
}

// 设置登录信息
func (d *Domain) SetLoginRequestInfo(loginRequest *amusinguserserv.LoginRequest) *xerror.Error {
	if loginRequest == nil {
		return xerror.NewError(nil, xerror.Code.BUnexpectedBlankVariable, "login request is blank")
	}
	d.LoginInfo = loginRequest

	return nil
}

// 设置用户的 DB 信息
func (d *Domain) SetUserModelInfo(ctx context.Context) *xerror.Error {
	if d.LoginInfo == nil {
		return xerror.NewError(nil, xerror.Code.BUnexpectedBlankVariable, "login info is blank")
	}

	user, err := model.FindUserByPhone(ctx, d.LoginInfo.AreaCode, d.LoginInfo.Phone)
	if err != nil {
		return err
	}

	if user == nil {
		return xerror.NewError(err, xerror.Code.CUnexpectRequestDate, "账号密码错误")
	}

	d.UserModelInfo = user

	return nil
}

func (d *Domain) ValidPassword(ctx context.Context) *xerror.Error {
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

func (d *Domain) SetSession(ctx context.Context) *xerror.Error {
	if d.UserModelInfo == nil {
		return xerror.NewError(nil, xerror.Code.BUnexpectedBlankVariable, "no user info")
	}

	uid := uuid.NewV4().String()

	sess, err := session.GlobalSessionManager.Store.SessionInit(ctx, uuid.NewV4().String())
	if err != nil {
		return xerror.NewError(nil, xerror.Code.SUnexpectedErr, "get session failed")
	}

	err = sess.Set(ctx, "id", d.UserModelInfo.ID)
	if err != nil {
		return xerror.NewError(nil, xerror.Code.SUnexpectedErr, "set session id failed")
	}

	d.SessionID = uid

	return nil
}
