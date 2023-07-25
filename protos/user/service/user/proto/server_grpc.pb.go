// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Pong(ctx context.Context, in *BlankParams, opts ...grpc.CallOption) (*PongResponse, error)
	Regexps(ctx context.Context, in *BlankParams, opts ...grpc.CallOption) (*RegexpResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	CountryCodes(ctx context.Context, in *BlankParams, opts ...grpc.CallOption) (*CountryCodeList, error)
	GetVerificationCode(ctx context.Context, in *VerificationCodeRequest, opts ...grpc.CallOption) (*VerificationCodeResponse, error)
	VerificationCodeCheck(ctx context.Context, in *VerificationCodeCheckRequest, opts ...grpc.CallOption) (*VerificationCodeCheckResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
	OAuthLogin(ctx context.Context, in *OAuthLoginRequest, opts ...grpc.CallOption) (*OAuthLoginResponse, error)
	OAuthInfo(ctx context.Context, in *OAuthInfoRequest, opts ...grpc.CallOption) (*OAuthInfoResponse, error)
	IsLogin(ctx context.Context, in *IsLoginRequest, opts ...grpc.CallOption) (*IsLoginResponse, error)
	LogOut(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	Authentication(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Pong(ctx context.Context, in *BlankParams, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := c.cc.Invoke(ctx, "/user.userService/Pong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Regexps(ctx context.Context, in *BlankParams, opts ...grpc.CallOption) (*RegexpResponse, error) {
	out := new(RegexpResponse)
	err := c.cc.Invoke(ctx, "/user.userService/Regexps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/user.userService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, "/user.userService/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CountryCodes(ctx context.Context, in *BlankParams, opts ...grpc.CallOption) (*CountryCodeList, error) {
	out := new(CountryCodeList)
	err := c.cc.Invoke(ctx, "/user.userService/CountryCodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetVerificationCode(ctx context.Context, in *VerificationCodeRequest, opts ...grpc.CallOption) (*VerificationCodeResponse, error) {
	out := new(VerificationCodeResponse)
	err := c.cc.Invoke(ctx, "/user.userService/GetVerificationCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) VerificationCodeCheck(ctx context.Context, in *VerificationCodeCheckRequest, opts ...grpc.CallOption) (*VerificationCodeCheckResponse, error) {
	out := new(VerificationCodeCheckResponse)
	err := c.cc.Invoke(ctx, "/user.userService/VerificationCodeCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, "/user.userService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) OAuthLogin(ctx context.Context, in *OAuthLoginRequest, opts ...grpc.CallOption) (*OAuthLoginResponse, error) {
	out := new(OAuthLoginResponse)
	err := c.cc.Invoke(ctx, "/user.userService/OAuthLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) OAuthInfo(ctx context.Context, in *OAuthInfoRequest, opts ...grpc.CallOption) (*OAuthInfoResponse, error) {
	out := new(OAuthInfoResponse)
	err := c.cc.Invoke(ctx, "/user.userService/OAuthInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) IsLogin(ctx context.Context, in *IsLoginRequest, opts ...grpc.CallOption) (*IsLoginResponse, error) {
	out := new(IsLoginResponse)
	err := c.cc.Invoke(ctx, "/user.userService/IsLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LogOut(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/user.userService/LogOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Authentication(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error) {
	out := new(AuthenticationResponse)
	err := c.cc.Invoke(ctx, "/user.userService/Authentication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Pong(context.Context, *BlankParams) (*PongResponse, error)
	Regexps(context.Context, *BlankParams) (*RegexpResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Join(context.Context, *JoinRequest) (*JoinResponse, error)
	CountryCodes(context.Context, *BlankParams) (*CountryCodeList, error)
	GetVerificationCode(context.Context, *VerificationCodeRequest) (*VerificationCodeResponse, error)
	VerificationCodeCheck(context.Context, *VerificationCodeCheckRequest) (*VerificationCodeCheckResponse, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	OAuthLogin(context.Context, *OAuthLoginRequest) (*OAuthLoginResponse, error)
	OAuthInfo(context.Context, *OAuthInfoRequest) (*OAuthInfoResponse, error)
	IsLogin(context.Context, *IsLoginRequest) (*IsLoginResponse, error)
	LogOut(context.Context, *LogoutRequest) (*LogoutResponse, error)
	Authentication(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Pong(context.Context, *BlankParams) (*PongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pong not implemented")
}
func (UnimplementedUserServiceServer) Regexps(context.Context, *BlankParams) (*RegexpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Regexps not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) Join(context.Context, *JoinRequest) (*JoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedUserServiceServer) CountryCodes(context.Context, *BlankParams) (*CountryCodeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountryCodes not implemented")
}
func (UnimplementedUserServiceServer) GetVerificationCode(context.Context, *VerificationCodeRequest) (*VerificationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVerificationCode not implemented")
}
func (UnimplementedUserServiceServer) VerificationCodeCheck(context.Context, *VerificationCodeCheckRequest) (*VerificationCodeCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerificationCodeCheck not implemented")
}
func (UnimplementedUserServiceServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedUserServiceServer) OAuthLogin(context.Context, *OAuthLoginRequest) (*OAuthLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OAuthLogin not implemented")
}
func (UnimplementedUserServiceServer) OAuthInfo(context.Context, *OAuthInfoRequest) (*OAuthInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OAuthInfo not implemented")
}
func (UnimplementedUserServiceServer) IsLogin(context.Context, *IsLoginRequest) (*IsLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsLogin not implemented")
}
func (UnimplementedUserServiceServer) LogOut(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogOut not implemented")
}
func (UnimplementedUserServiceServer) Authentication(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authentication not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Pong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlankParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Pong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userService/Pong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Pong(ctx, req.(*BlankParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Regexps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlankParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Regexps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userService/Regexps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Regexps(ctx, req.(*BlankParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userService/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CountryCodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlankParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CountryCodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userService/CountryCodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CountryCodes(ctx, req.(*BlankParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetVerificationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerificationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetVerificationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userService/GetVerificationCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetVerificationCode(ctx, req.(*VerificationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_VerificationCodeCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerificationCodeCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).VerificationCodeCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userService/VerificationCodeCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).VerificationCodeCheck(ctx, req.(*VerificationCodeCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_OAuthLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OAuthLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).OAuthLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userService/OAuthLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).OAuthLogin(ctx, req.(*OAuthLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_OAuthInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OAuthInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).OAuthInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userService/OAuthInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).OAuthInfo(ctx, req.(*OAuthInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_IsLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).IsLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userService/IsLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).IsLogin(ctx, req.(*IsLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LogOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LogOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userService/LogOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LogOut(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Authentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Authentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userService/Authentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Authentication(ctx, req.(*AuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.userService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pong",
			Handler:    _UserService_Pong_Handler,
		},
		{
			MethodName: "Regexps",
			Handler:    _UserService_Regexps_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _UserService_Join_Handler,
		},
		{
			MethodName: "CountryCodes",
			Handler:    _UserService_CountryCodes_Handler,
		},
		{
			MethodName: "GetVerificationCode",
			Handler:    _UserService_GetVerificationCode_Handler,
		},
		{
			MethodName: "VerificationCodeCheck",
			Handler:    _UserService_VerificationCodeCheck_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _UserService_ResetPassword_Handler,
		},
		{
			MethodName: "OAuthLogin",
			Handler:    _UserService_OAuthLogin_Handler,
		},
		{
			MethodName: "OAuthInfo",
			Handler:    _UserService_OAuthInfo_Handler,
		},
		{
			MethodName: "IsLogin",
			Handler:    _UserService_IsLogin_Handler,
		},
		{
			MethodName: "LogOut",
			Handler:    _UserService_LogOut_Handler,
		},
		{
			MethodName: "Authentication",
			Handler:    _UserService_Authentication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/proto/server.proto",
}
