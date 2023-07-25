package userservice

import (
	service "amusingx.fit/amusingx/protos/user/service/user/proto"
	"amusingx.fit/amusingx/regexp"
	"amusingx.fit/amusingx/services/user/rpcserver/userservice/controller/authentication"
	"amusingx.fit/amusingx/services/user/rpcserver/userservice/controller/countrycode"
	password "amusingx.fit/amusingx/services/user/rpcserver/userservice/controller/findpassword"
	"amusingx.fit/amusingx/services/user/rpcserver/userservice/controller/join"
	"amusingx.fit/amusingx/services/user/rpcserver/userservice/controller/login"
	"amusingx.fit/amusingx/services/user/rpcserver/userservice/controller/logout"
	"amusingx.fit/amusingx/services/user/rpcserver/userservice/controller/oauthlogin"
	"amusingx.fit/amusingx/services/user/rpcserver/userservice/controller/verificationcode"
	"context"
	"errors"
	"github.com/ItsWewin/superfactory/logger"
)

type UserService struct {
	service.UnimplementedUserServiceServer
}

func (s *UserService) Pong(ctx context.Context, blank *service.BlankParams) (*service.PongResponse, error) {
	return &service.PongResponse{ServerName: "amusing-x-user-service"}, nil
}

func (s *UserService) Regexps(ctx context.Context, blank *service.BlankParams) (*service.RegexpResponse, error) {
	regexps := make([]*service.Regexp, 0)
	for _, regexp := range regexp.AllRegexps {
		regexps = append(regexps, &service.Regexp{
			Id:    regexp.ID,
			Name:  regexp.Name,
			Rules: regexp.Rules,
			Desc:  regexp.Desc,
		})
	}

	return &service.RegexpResponse{Regexps: regexps}, nil
}

func (s *UserService) Login(ctx context.Context, in *service.LoginRequest) (*service.LoginResponse, error) {
	sessionID, err := login.HandlerLogin(ctx, in)
	if err != nil {
		logger.Errorf("Login failed: %s", err)
		return nil, errors.New("Login failed. ")
	}

	return &service.LoginResponse{SessionId: sessionID}, err
}

func (s *UserService) Join(ctx context.Context, in *service.JoinRequest) (*service.JoinResponse, error) {
	joinResponse, err := join.HandleJoin(ctx, in)
	if err != nil {
		logger.Errorf("Join failed: %s", err)
		return nil, errors.New("Join failed. ")
	}

	return joinResponse, nil
}

func (s *UserService) CountryCodes(ctx context.Context, req *service.BlankParams) (*service.CountryCodeList, error) {
	resp, err := countrycode.HandleCountryCodeList(ctx)
	if err != nil {
		logger.Errorf("[CountryCodes] countrycode.HandleCountryCodeList failed: %s", err)
		return nil, errors.New("get country code failed. ")
	}

	return resp, nil
}

func (s *UserService) GetVerificationCode(ctx context.Context, req *service.VerificationCodeRequest) (*service.VerificationCodeResponse, error) {
	resp, err := verificationcode.HandlerVerificationCode(ctx, req)
	if err != nil {
		logger.Errorf("[GetVerificationCode] verificationcode.HandlerVerificationCode failed: %s", err)
		return nil, errors.New("get verification code failed. ")
	}

	return resp, nil
}

func (s *UserService) VerificationCodeCheck(ctx context.Context, req *service.VerificationCodeCheckRequest) (*service.VerificationCodeCheckResponse, error) {
	resp, err := verificationcode.HandlerVerificationCheck(ctx, req)
	if err != nil {
		logger.Errorf("[VerificationCodeCheck] verificationcode.HandlerVerificationCheck failed: %s", err)
		return nil, errors.New("verification code check failed. ")
	}

	return resp, nil
}

func (s *UserService) ResetPassword(ctx context.Context, req *service.ResetPasswordRequest) (*service.ResetPasswordResponse, error) {
	resp, err := password.HandlerResetPassword(ctx, req)
	if err != nil {
		logger.Errorf("[ResetPassword] password.HandlerResetPassword failed: %s", err)
		return nil, errors.New("reset password failed")
	}

	return resp, nil
}

func (s *UserService) OAuthLogin(ctx context.Context, req *service.OAuthLoginRequest) (*service.OAuthLoginResponse, error) {
	resp, err := oauthlogin.HandlerOAuthLogin(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *UserService) OAuthInfo(ctx context.Context, req *service.OAuthInfoRequest) (*service.OAuthInfoResponse, error) {
	resp, err := oauthlogin.HandlerGetOAuthInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *UserService) IsLogin(ctx context.Context, req *service.IsLoginRequest) (*service.IsLoginResponse, error) {
	resp, err := login.HandlerIsLogin(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *UserService) LogOut(ctx context.Context, req *service.LogoutRequest) (*service.LogoutResponse, error) {
	resp, err := logout.HandlerIsLogOut(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (s *UserService) Authentication(ctx context.Context, req *service.AuthenticationRequest) (*service.AuthenticationResponse, error) {
	resp, err := authentication.HandlerAuthentication(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
