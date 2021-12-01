package userservice

import (
	"amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/regexp"
	"amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/controller/countrycode"
	password "amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/controller/findpassword"
	"amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/controller/join"
	"amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/controller/login"
	"amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/controller/verificationcode"
	"context"
	"errors"
	"github.com/ItsWewin/superfactory/logger"
)

type UserService struct {
	userservice.UnimplementedAmusingxUserServiceServer
}

func (s *UserService) Pong(ctx context.Context, blank *userservice.BlankParams) (*userservice.PongResponse, error) {
	return &userservice.PongResponse{ServerName: "amusing-x-user-service"}, nil
}

func (s *UserService) Regexps(ctx context.Context, blank *userservice.BlankParams) (*userservice.RegexpResponse, error) {
	regexps := make([]*userservice.Regexp, 0)
	for _, regexp := range regexp.AllRegexps {
		regexps = append(regexps, &userservice.Regexp{
			Id:    regexp.ID,
			Name:  regexp.Name,
			Rules: regexp.Rules,
			Desc:  regexp.Desc,
		})
	}

	return &userservice.RegexpResponse{Regexps: regexps}, nil
}

func (s *UserService) Login(ctx context.Context, in *userservice.LoginRequest) (*userservice.LoginResponse, error) {
	sessionID, err := login.HandlerLogin(ctx, in)
	if err != nil {
		logger.Errorf("Login failed: %s", err)
		return nil, errors.New("Login failed. ")
	}

	return &userservice.LoginResponse{SessionId: sessionID}, err
}

func (s *UserService) Join(ctx context.Context, in *userservice.JoinRequest) (*userservice.JoinResponse, error) {
	joinResponse, err := join.HandleJoin(ctx, in)
	if err != nil {
		logger.Errorf("Join failed: %s", err)
		return nil, errors.New("Join failed. ")
	}

	return joinResponse, nil
}

func (s *UserService) CountryCodes(ctx context.Context, req *userservice.BlankParams) (*userservice.CountryCodeList, error) {
	resp, err := countrycode.HandleCountryCodeList(ctx)
	if err != nil {
		logger.Errorf("[CountryCodes] countrycode.HandleCountryCodeList failed: %s", err)
		return nil, errors.New("get country code failed. ")
	}

	return resp, nil
}

func (s *UserService) GetVerificationCode(ctx context.Context, req *userservice.VerificationCodeRequest) (*userservice.VerificationCodeResponse, error) {
	resp, err := verificationcode.HandlerVerificationCode(ctx, req)
	if err != nil {
		logger.Errorf("[GetVerificationCode] verificationcode.HandlerVerificationCode failed: %s", err)
		return nil, errors.New("get verification code failed. ")
	}

	return resp, nil
}

func (s *UserService) VerificationCodeCheck(ctx context.Context, req *userservice.VerificationCodeCheckRequest) (*userservice.VerificationCodeCheckResponse, error) {
	resp, err := verificationcode.HandlerVerificationCheck(ctx, req)
	if err != nil {
		logger.Errorf("[VerificationCodeCheck] verificationcode.HandlerVerificationCheck failed: %s", err)
		return nil, errors.New("verification code check failed. ")
	}

	return resp, nil
}

func (s *UserService) ResetPassword(ctx context.Context, req *userservice.ResetPasswordRequest) (*userservice.ResetPasswordResponse, error) {
	resp, err := password.HandlerResetPassword(ctx, req)
	if err != nil {
		logger.Errorf("[ResetPassword] password.HandlerResetPassword failed: %s", err)
		return nil, errors.New("reset password failed")
	}

	return resp, nil
}
