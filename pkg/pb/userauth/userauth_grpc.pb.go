// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: pkg/pb/userauth/userauth.proto

package userauth

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

const (
	User_UserSignUp_FullMethodName          = "/user.User/UserSignUp"
	User_UserLogin_FullMethodName           = "/user.User/UserLogin"
	User_UserEditDetails_FullMethodName     = "/user.User/UserEditDetails"
	User_UserOtpGeneration_FullMethodName   = "/user.User/UserOtpGeneration"
	User_UserOtpVerification_FullMethodName = "/user.User/UserOtpVerification"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	UserSignUp(ctx context.Context, in *UserSignUpRequest, opts ...grpc.CallOption) (*UserSignUpResponse, error)
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	UserEditDetails(ctx context.Context, in *UserEditDetailsRequest, opts ...grpc.CallOption) (*UserEditDetailsResponse, error)
	UserOtpGeneration(ctx context.Context, in *UserOtpRequest, opts ...grpc.CallOption) (*UserOtpRequestResponse, error)
	UserOtpVerification(ctx context.Context, in *UserOtpVerificationRequest, opts ...grpc.CallOption) (*UserOtpVerificationResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) UserSignUp(ctx context.Context, in *UserSignUpRequest, opts ...grpc.CallOption) (*UserSignUpResponse, error) {
	out := new(UserSignUpResponse)
	err := c.cc.Invoke(ctx, User_UserSignUp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, User_UserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserEditDetails(ctx context.Context, in *UserEditDetailsRequest, opts ...grpc.CallOption) (*UserEditDetailsResponse, error) {
	out := new(UserEditDetailsResponse)
	err := c.cc.Invoke(ctx, User_UserEditDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserOtpGeneration(ctx context.Context, in *UserOtpRequest, opts ...grpc.CallOption) (*UserOtpRequestResponse, error) {
	out := new(UserOtpRequestResponse)
	err := c.cc.Invoke(ctx, User_UserOtpGeneration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserOtpVerification(ctx context.Context, in *UserOtpVerificationRequest, opts ...grpc.CallOption) (*UserOtpVerificationResponse, error) {
	out := new(UserOtpVerificationResponse)
	err := c.cc.Invoke(ctx, User_UserOtpVerification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	UserSignUp(context.Context, *UserSignUpRequest) (*UserSignUpResponse, error)
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	UserEditDetails(context.Context, *UserEditDetailsRequest) (*UserEditDetailsResponse, error)
	UserOtpGeneration(context.Context, *UserOtpRequest) (*UserOtpRequestResponse, error)
	UserOtpVerification(context.Context, *UserOtpVerificationRequest) (*UserOtpVerificationResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) UserSignUp(context.Context, *UserSignUpRequest) (*UserSignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignUp not implemented")
}
func (UnimplementedUserServer) UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserServer) UserEditDetails(context.Context, *UserEditDetailsRequest) (*UserEditDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserEditDetails not implemented")
}
func (UnimplementedUserServer) UserOtpGeneration(context.Context, *UserOtpRequest) (*UserOtpRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserOtpGeneration not implemented")
}
func (UnimplementedUserServer) UserOtpVerification(context.Context, *UserOtpVerificationRequest) (*UserOtpVerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserOtpVerification not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_UserSignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserSignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserSignUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserSignUp(ctx, req.(*UserSignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserEditDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEditDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserEditDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserEditDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserEditDetails(ctx, req.(*UserEditDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserOtpGeneration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserOtpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserOtpGeneration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserOtpGeneration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserOtpGeneration(ctx, req.(*UserOtpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserOtpVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserOtpVerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserOtpVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserOtpVerification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserOtpVerification(ctx, req.(*UserOtpVerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserSignUp",
			Handler:    _User_UserSignUp_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _User_UserLogin_Handler,
		},
		{
			MethodName: "UserEditDetails",
			Handler:    _User_UserEditDetails_Handler,
		},
		{
			MethodName: "UserOtpGeneration",
			Handler:    _User_UserOtpGeneration_Handler,
		},
		{
			MethodName: "UserOtpVerification",
			Handler:    _User_UserOtpVerification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/userauth/userauth.proto",
}
