package userservice

import (
	"amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/regexp"
	"amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/controller/countrycode"
	password "amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/controller/findpassword"
	"amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/controller/join"
	"amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/controller/login"
	"amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/controller/oauthlogin"
	"amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/controller/verificationcode"
	"context"
	"errors"
	"github.com/ItsWewin/superfactory/logger"
)

type UserService struct {
	ganymedeservice.UnimplementedGanymedeServiceServer
}

func (s *UserService) Pong(ctx context.Context, blank *ganymedeservice.BlankParams) (*ganymedeservice.PongResponse, error) {
	return &ganymedeservice.PongResponse{ServerName: "amusing-x-user-service"}, nil
}

func (s *UserService) Regexps(ctx context.Context, blank *ganymedeservice.BlankParams) (*ganymedeservice.RegexpResponse, error) {
	regexps := make([]*ganymedeservice.Regexp, 0)
	for _, regexp := range regexp.AllRegexps {
		regexps = append(regexps, &ganymedeservice.Regexp{
			Id:    regexp.ID,
			Name:  regexp.Name,
			Rules: regexp.Rules,
			Desc:  regexp.Desc,
		})
	}

	return &ganymedeservice.RegexpResponse{Regexps: regexps}, nil
}

func (s *UserService) Login(ctx context.Context, in *ganymedeservice.LoginRequest) (*ganymedeservice.LoginResponse, error) {
	sessionID, err := login.HandlerLogin(ctx, in)
	if err != nil {
		logger.Errorf("Login failed: %s", err)
		return nil, errors.New("Login failed. ")
	}

	return &ganymedeservice.LoginResponse{SessionId: sessionID}, err
}

func (s *UserService) Join(ctx context.Context, in *ganymedeservice.JoinRequest) (*ganymedeservice.JoinResponse, error) {
	joinResponse, err := join.HandleJoin(ctx, in)
	if err != nil {
		logger.Errorf("Join failed: %s", err)
		return nil, errors.New("Join failed. ")
	}

	return joinResponse, nil
}

func (s *UserService) CountryCodes(ctx context.Context, req *ganymedeservice.BlankParams) (*ganymedeservice.CountryCodeList, error) {
	resp, err := countrycode.HandleCountryCodeList(ctx)
	if err != nil {
		logger.Errorf("[CountryCodes] countrycode.HandleCountryCodeList failed: %s", err)
		return nil, errors.New("get country code failed. ")
	}

	return resp, nil
}

func (s *UserService) GetVerificationCode(ctx context.Context, req *ganymedeservice.VerificationCodeRequest) (*ganymedeservice.VerificationCodeResponse, error) {
	resp, err := verificationcode.HandlerVerificationCode(ctx, req)
	if err != nil {
		logger.Errorf("[GetVerificationCode] verificationcode.HandlerVerificationCode failed: %s", err)
		return nil, errors.New("get verification code failed. ")
	}

	return resp, nil
}

func (s *UserService) VerificationCodeCheck(ctx context.Context, req *ganymedeservice.VerificationCodeCheckRequest) (*ganymedeservice.VerificationCodeCheckResponse, error) {
	resp, err := verificationcode.HandlerVerificationCheck(ctx, req)
	if err != nil {
		logger.Errorf("[VerificationCodeCheck] verificationcode.HandlerVerificationCheck failed: %s", err)
		return nil, errors.New("verification code check failed. ")
	}

	return resp, nil
}

func (s *UserService) ResetPassword(ctx context.Context, req *ganymedeservice.ResetPasswordRequest) (*ganymedeservice.ResetPasswordResponse, error) {
	resp, err := password.HandlerResetPassword(ctx, req)
	if err != nil {
		logger.Errorf("[ResetPassword] password.HandlerResetPassword failed: %s", err)
		return nil, errors.New("reset password failed")
	}

	return resp, nil
}

func (s *UserService) OAuthLogin(ctx context.Context, req *ganymedeservice.OAuthLoginRequest) (*ganymedeservice.OAuthLoginResponse, error) {
	resp, err := oauthlogin.HandlerOAuthLogin(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *UserService) OAuthInfo(ctx context.Context, req *ganymedeservice.OAuthInfoRequest) (*ganymedeservice.OAuthInfoResponse, error) {
	resp, err := oauthlogin.HandlerGetOAuthInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *UserService) IsLogin(ctx context.Context, req *ganymedeservice.IsLoginRequest) (*ganymedeservice.IsLoginResponse, error) {
	resp, err := login.HandlerIsLogin(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
