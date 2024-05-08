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
	User_UserSignUp_FullMethodName           = "/user.User/UserSignUp"
	User_UserLogin_FullMethodName            = "/user.User/UserLogin"
	User_UserEditDetails_FullMethodName      = "/user.User/UserEditDetails"
	User_UserOtpGeneration_FullMethodName    = "/user.User/UserOtpGeneration"
	User_UserOtpVerification_FullMethodName  = "/user.User/UserOtpVerification"
	User_GetUsers_FullMethodName             = "/user.User/GetUsers"
	User_UpdateUserStatus_FullMethodName     = "/user.User/UpdateUserStatus"
	User_AddUserInterest_FullMethodName      = "/user.User/AddUserInterest"
	User_EditUserInterest_FullMethodName     = "/user.User/EditUserInterest"
	User_DeleteUserInterest_FullMethodName   = "/user.User/DeleteUserInterest"
	User_GetUserInterests_FullMethodName     = "/user.User/GetUserInterests"
	User_AddUserPreference_FullMethodName    = "/user.User/AddUserPreference"
	User_EditUserPreference_FullMethodName   = "/user.User/EditUserPreference"
	User_DeleteUserPreference_FullMethodName = "/user.User/DeleteUserPreference"
	User_GetUserPreferences_FullMethodName   = "/user.User/GetUserPreferences"
	User_FollowUser_FullMethodName           = "/user.User/FollowUser"
	User_BlockUser_FullMethodName            = "/user.User/BlockUser"
	User_SendMessage_FullMethodName          = "/user.User/SendMessage"
	User_ReadMessages_FullMethodName         = "/user.User/ReadMessages"
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
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	UpdateUserStatus(ctx context.Context, in *UpdateUserStatusRequest, opts ...grpc.CallOption) (*UpdateUserStatusResponse, error)
	AddUserInterest(ctx context.Context, in *AddUserInterestRequest, opts ...grpc.CallOption) (*AddUserInterestResponse, error)
	EditUserInterest(ctx context.Context, in *EditUserInterestRequest, opts ...grpc.CallOption) (*EditUserInterestResponse, error)
	DeleteUserInterest(ctx context.Context, in *DeleteUserInterestRequest, opts ...grpc.CallOption) (*DeleteUserInterestResponse, error)
	GetUserInterests(ctx context.Context, in *GetUserInterestsRequest, opts ...grpc.CallOption) (*GetUserInterestsResponse, error)
	AddUserPreference(ctx context.Context, in *AddUserPreferenceRequest, opts ...grpc.CallOption) (*AddUserPreferenceResponse, error)
	EditUserPreference(ctx context.Context, in *EditUserPreferenceRequest, opts ...grpc.CallOption) (*EditUserPreferenceResponse, error)
	DeleteUserPreference(ctx context.Context, in *DeleteUserPreferenceRequest, opts ...grpc.CallOption) (*DeleteUserPreferenceResponse, error)
	GetUserPreferences(ctx context.Context, in *GetUserPreferencesRequest, opts ...grpc.CallOption) (*GetUserPreferencesResponse, error)
	FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*FollowUserResponce, error)
	BlockUser(ctx context.Context, in *BlockUserRequest, opts ...grpc.CallOption) (*BlockUserResponce, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponce, error)
	// New RPC method to read messages in a room
	ReadMessages(ctx context.Context, in *ReadMessagesRequest, opts ...grpc.CallOption) (*ReadMessagesResponse, error)
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

func (c *userClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, User_GetUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUserStatus(ctx context.Context, in *UpdateUserStatusRequest, opts ...grpc.CallOption) (*UpdateUserStatusResponse, error) {
	out := new(UpdateUserStatusResponse)
	err := c.cc.Invoke(ctx, User_UpdateUserStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AddUserInterest(ctx context.Context, in *AddUserInterestRequest, opts ...grpc.CallOption) (*AddUserInterestResponse, error) {
	out := new(AddUserInterestResponse)
	err := c.cc.Invoke(ctx, User_AddUserInterest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) EditUserInterest(ctx context.Context, in *EditUserInterestRequest, opts ...grpc.CallOption) (*EditUserInterestResponse, error) {
	out := new(EditUserInterestResponse)
	err := c.cc.Invoke(ctx, User_EditUserInterest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteUserInterest(ctx context.Context, in *DeleteUserInterestRequest, opts ...grpc.CallOption) (*DeleteUserInterestResponse, error) {
	out := new(DeleteUserInterestResponse)
	err := c.cc.Invoke(ctx, User_DeleteUserInterest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserInterests(ctx context.Context, in *GetUserInterestsRequest, opts ...grpc.CallOption) (*GetUserInterestsResponse, error) {
	out := new(GetUserInterestsResponse)
	err := c.cc.Invoke(ctx, User_GetUserInterests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AddUserPreference(ctx context.Context, in *AddUserPreferenceRequest, opts ...grpc.CallOption) (*AddUserPreferenceResponse, error) {
	out := new(AddUserPreferenceResponse)
	err := c.cc.Invoke(ctx, User_AddUserPreference_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) EditUserPreference(ctx context.Context, in *EditUserPreferenceRequest, opts ...grpc.CallOption) (*EditUserPreferenceResponse, error) {
	out := new(EditUserPreferenceResponse)
	err := c.cc.Invoke(ctx, User_EditUserPreference_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteUserPreference(ctx context.Context, in *DeleteUserPreferenceRequest, opts ...grpc.CallOption) (*DeleteUserPreferenceResponse, error) {
	out := new(DeleteUserPreferenceResponse)
	err := c.cc.Invoke(ctx, User_DeleteUserPreference_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserPreferences(ctx context.Context, in *GetUserPreferencesRequest, opts ...grpc.CallOption) (*GetUserPreferencesResponse, error) {
	out := new(GetUserPreferencesResponse)
	err := c.cc.Invoke(ctx, User_GetUserPreferences_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*FollowUserResponce, error) {
	out := new(FollowUserResponce)
	err := c.cc.Invoke(ctx, User_FollowUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) BlockUser(ctx context.Context, in *BlockUserRequest, opts ...grpc.CallOption) (*BlockUserResponce, error) {
	out := new(BlockUserResponce)
	err := c.cc.Invoke(ctx, User_BlockUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponce, error) {
	out := new(SendMessageResponce)
	err := c.cc.Invoke(ctx, User_SendMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ReadMessages(ctx context.Context, in *ReadMessagesRequest, opts ...grpc.CallOption) (*ReadMessagesResponse, error) {
	out := new(ReadMessagesResponse)
	err := c.cc.Invoke(ctx, User_ReadMessages_FullMethodName, in, out, opts...)
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
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	UpdateUserStatus(context.Context, *UpdateUserStatusRequest) (*UpdateUserStatusResponse, error)
	AddUserInterest(context.Context, *AddUserInterestRequest) (*AddUserInterestResponse, error)
	EditUserInterest(context.Context, *EditUserInterestRequest) (*EditUserInterestResponse, error)
	DeleteUserInterest(context.Context, *DeleteUserInterestRequest) (*DeleteUserInterestResponse, error)
	GetUserInterests(context.Context, *GetUserInterestsRequest) (*GetUserInterestsResponse, error)
	AddUserPreference(context.Context, *AddUserPreferenceRequest) (*AddUserPreferenceResponse, error)
	EditUserPreference(context.Context, *EditUserPreferenceRequest) (*EditUserPreferenceResponse, error)
	DeleteUserPreference(context.Context, *DeleteUserPreferenceRequest) (*DeleteUserPreferenceResponse, error)
	GetUserPreferences(context.Context, *GetUserPreferencesRequest) (*GetUserPreferencesResponse, error)
	FollowUser(context.Context, *FollowUserRequest) (*FollowUserResponce, error)
	BlockUser(context.Context, *BlockUserRequest) (*BlockUserResponce, error)
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponce, error)
	// New RPC method to read messages in a room
	ReadMessages(context.Context, *ReadMessagesRequest) (*ReadMessagesResponse, error)
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
func (UnimplementedUserServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUserServer) UpdateUserStatus(context.Context, *UpdateUserStatusRequest) (*UpdateUserStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserStatus not implemented")
}
func (UnimplementedUserServer) AddUserInterest(context.Context, *AddUserInterestRequest) (*AddUserInterestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserInterest not implemented")
}
func (UnimplementedUserServer) EditUserInterest(context.Context, *EditUserInterestRequest) (*EditUserInterestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUserInterest not implemented")
}
func (UnimplementedUserServer) DeleteUserInterest(context.Context, *DeleteUserInterestRequest) (*DeleteUserInterestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserInterest not implemented")
}
func (UnimplementedUserServer) GetUserInterests(context.Context, *GetUserInterestsRequest) (*GetUserInterestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInterests not implemented")
}
func (UnimplementedUserServer) AddUserPreference(context.Context, *AddUserPreferenceRequest) (*AddUserPreferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserPreference not implemented")
}
func (UnimplementedUserServer) EditUserPreference(context.Context, *EditUserPreferenceRequest) (*EditUserPreferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUserPreference not implemented")
}
func (UnimplementedUserServer) DeleteUserPreference(context.Context, *DeleteUserPreferenceRequest) (*DeleteUserPreferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserPreference not implemented")
}
func (UnimplementedUserServer) GetUserPreferences(context.Context, *GetUserPreferencesRequest) (*GetUserPreferencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPreferences not implemented")
}
func (UnimplementedUserServer) FollowUser(context.Context, *FollowUserRequest) (*FollowUserResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")
}
func (UnimplementedUserServer) BlockUser(context.Context, *BlockUserRequest) (*BlockUserResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockUser not implemented")
}
func (UnimplementedUserServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedUserServer) ReadMessages(context.Context, *ReadMessagesRequest) (*ReadMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadMessages not implemented")
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

func _User_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUserStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUserStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UpdateUserStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUserStatus(ctx, req.(*UpdateUserStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AddUserInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserInterestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AddUserInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_AddUserInterest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AddUserInterest(ctx, req.(*AddUserInterestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_EditUserInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditUserInterestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).EditUserInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_EditUserInterest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).EditUserInterest(ctx, req.(*EditUserInterestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteUserInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserInterestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteUserInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_DeleteUserInterest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteUserInterest(ctx, req.(*DeleteUserInterestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserInterests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInterestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserInterests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserInterests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserInterests(ctx, req.(*GetUserInterestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AddUserPreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserPreferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AddUserPreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_AddUserPreference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AddUserPreference(ctx, req.(*AddUserPreferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_EditUserPreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditUserPreferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).EditUserPreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_EditUserPreference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).EditUserPreference(ctx, req.(*EditUserPreferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteUserPreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserPreferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteUserPreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_DeleteUserPreference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteUserPreference(ctx, req.(*DeleteUserPreferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserPreferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPreferencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserPreferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserPreferences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserPreferences(ctx, req.(*GetUserPreferencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_FollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).FollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_FollowUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).FollowUser(ctx, req.(*FollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_BlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).BlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_BlockUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).BlockUser(ctx, req.(*BlockUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ReadMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ReadMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ReadMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ReadMessages(ctx, req.(*ReadMessagesRequest))
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
		{
			MethodName: "GetUsers",
			Handler:    _User_GetUsers_Handler,
		},
		{
			MethodName: "UpdateUserStatus",
			Handler:    _User_UpdateUserStatus_Handler,
		},
		{
			MethodName: "AddUserInterest",
			Handler:    _User_AddUserInterest_Handler,
		},
		{
			MethodName: "EditUserInterest",
			Handler:    _User_EditUserInterest_Handler,
		},
		{
			MethodName: "DeleteUserInterest",
			Handler:    _User_DeleteUserInterest_Handler,
		},
		{
			MethodName: "GetUserInterests",
			Handler:    _User_GetUserInterests_Handler,
		},
		{
			MethodName: "AddUserPreference",
			Handler:    _User_AddUserPreference_Handler,
		},
		{
			MethodName: "EditUserPreference",
			Handler:    _User_EditUserPreference_Handler,
		},
		{
			MethodName: "DeleteUserPreference",
			Handler:    _User_DeleteUserPreference_Handler,
		},
		{
			MethodName: "GetUserPreferences",
			Handler:    _User_GetUserPreferences_Handler,
		},
		{
			MethodName: "FollowUser",
			Handler:    _User_FollowUser_Handler,
		},
		{
			MethodName: "BlockUser",
			Handler:    _User_BlockUser_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _User_SendMessage_Handler,
		},
		{
			MethodName: "ReadMessages",
			Handler:    _User_ReadMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/userauth/userauth.proto",
}
