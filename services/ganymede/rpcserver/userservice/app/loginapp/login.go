package loginapp

import (
	"amusingx.fit/amusingx/protos/ganymede/service/ganymede/proto"
	"amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/model"
	"amusingx.fit/amusingx/services/ganymede/session"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/verificationcode/randomcode"
	uuid "github.com/satori/go.uuid"
)

type LoginDomain interface {
}

type Domain struct {
	LoginInfo     *ganymedeservice.LoginRequest
	UserModelInfo *model.UserComplex
	SessionID     string
}

func NewDomain() *Domain {
	return &Domain{}
}

// 设置登录信息
func (d *Domain) SetLoginRequestInfo(loginRequest *ganymedeservice.LoginRequest) aerror.Error {
	if loginRequest == nil {
		return aerror.NewError(nil, aerror.Code.BUnexpectedBlankVariable, "login request is blank")
	}
	d.LoginInfo = loginRequest

	return nil
}

// 设置用户的 DB 信息
func (d *Domain) SetUserModelInfo(ctx context.Context) aerror.Error {
	if d.LoginInfo == nil {
		return aerror.NewError(nil, aerror.Code.BUnexpectedBlankVariable, "login info is blank")
	}

	user, err := model.FindUserByPhone(ctx, d.LoginInfo.AreaCode, d.LoginInfo.Phone)
	if err != nil {
		return err
	}

	if user == nil {
		return aerror.NewError(err, aerror.Code.CUnexpectRequestDate, "账号密码错误")
	}

	d.UserModelInfo = user

	return nil
}

func (d *Domain) LoginAuthentication(ctx context.Context) aerror.Error {
	//if d.LoginInfo.LoginByPassword() {
	//	return d.ValidPassword(ctx)
	//}

	return d.ValidVerificationCode(ctx)
}

func (d *Domain) ValidVerificationCode(ctx context.Context) aerror.Error {
	if d.UserModelInfo == nil || d.LoginInfo == nil {
		return aerror.NewError(nil, aerror.Code.BUnexpectedBlankVariable, "user_model_info or log_info is nil")
	}

	codeStore := randomcode.RandomCodeStoreInit()
	if !codeStore.Check(d.LoginInfo.VerificationCode) {
		return aerror.NewError(nil, aerror.Code.CParamsError, "验证码错误")
	}

	return nil
}

func (d *Domain) ValidPassword(ctx context.Context) aerror.Error {
	if d.UserModelInfo == nil || d.LoginInfo == nil {
		return aerror.NewError(nil, aerror.Code.BUnexpectedBlankVariable, "user_model_info or log_info is nil")
	}

	result, err := d.UserModelInfo.ComparePassword(d.LoginInfo.Password)
	if err != nil {
		return err
	}

	if !result {
		return aerror.NewError(nil, aerror.Code.CUnexpectRequestDate, "用户名或密码错误")
	}

	return nil
}

func (d *Domain) SetSession(ctx context.Context) aerror.Error {
	if d.UserModelInfo == nil {
		return aerror.NewError(nil, aerror.Code.BUnexpectedBlankVariable, "no user info")
	}

	uid := uuid.NewV4().String()

	sess, err := session.GlobalSessionManager.Store.SessionInit(ctx, uuid.NewV4().String())
	if err != nil {
		return aerror.NewError(nil, aerror.Code.SUnexpectedErr, "get session failed")
	}

	err = sess.Set(ctx, "id", d.UserModelInfo.ID)
	if err != nil {
		return aerror.NewError(nil, aerror.Code.SUnexpectedErr, "set session id failed")
	}

	d.SessionID = uid

	return nil
}
