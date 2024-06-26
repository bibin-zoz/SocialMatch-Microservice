// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: pkg/pb/admin/pb.proto

package admin

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
	Admin_AdminLogin_FullMethodName       = "/admin.Admin/AdminLogin"
	Admin_GetInterests_FullMethodName     = "/admin.Admin/GetInterests"
	Admin_GetPreferences_FullMethodName   = "/admin.Admin/GetPreferences"
	Admin_AddInterest_FullMethodName      = "/admin.Admin/AddInterest"
	Admin_EditInterest_FullMethodName     = "/admin.Admin/EditInterest"
	Admin_DeleteInterest_FullMethodName   = "/admin.Admin/DeleteInterest"
	Admin_AddPreference_FullMethodName    = "/admin.Admin/AddPreference"
	Admin_EditPreference_FullMethodName   = "/admin.Admin/EditPreference"
	Admin_DeletePreference_FullMethodName = "/admin.Admin/DeletePreference"
)

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error)
	GetInterests(ctx context.Context, in *GetInterestsRequest, opts ...grpc.CallOption) (*GetInterestsResponse, error)
	GetPreferences(ctx context.Context, in *GetPreferencesRequest, opts ...grpc.CallOption) (*GetPreferencesResponse, error)
	AddInterest(ctx context.Context, in *AddInterestRequest, opts ...grpc.CallOption) (*AddInterestResponse, error)
	EditInterest(ctx context.Context, in *EditInterestRequest, opts ...grpc.CallOption) (*EditInterestResponse, error)
	DeleteInterest(ctx context.Context, in *DeleteInterestRequest, opts ...grpc.CallOption) (*DeleteInterestResponse, error)
	AddPreference(ctx context.Context, in *AddPreferenceRequest, opts ...grpc.CallOption) (*AddPreferenceResponse, error)
	EditPreference(ctx context.Context, in *EditPreferenceRequest, opts ...grpc.CallOption) (*EditPreferenceResponse, error)
	DeletePreference(ctx context.Context, in *DeletePreferenceRequest, opts ...grpc.CallOption) (*DeletePreferenceResponse, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error) {
	out := new(AdminLoginResponse)
	err := c.cc.Invoke(ctx, Admin_AdminLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetInterests(ctx context.Context, in *GetInterestsRequest, opts ...grpc.CallOption) (*GetInterestsResponse, error) {
	out := new(GetInterestsResponse)
	err := c.cc.Invoke(ctx, Admin_GetInterests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetPreferences(ctx context.Context, in *GetPreferencesRequest, opts ...grpc.CallOption) (*GetPreferencesResponse, error) {
	out := new(GetPreferencesResponse)
	err := c.cc.Invoke(ctx, Admin_GetPreferences_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AddInterest(ctx context.Context, in *AddInterestRequest, opts ...grpc.CallOption) (*AddInterestResponse, error) {
	out := new(AddInterestResponse)
	err := c.cc.Invoke(ctx, Admin_AddInterest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) EditInterest(ctx context.Context, in *EditInterestRequest, opts ...grpc.CallOption) (*EditInterestResponse, error) {
	out := new(EditInterestResponse)
	err := c.cc.Invoke(ctx, Admin_EditInterest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) DeleteInterest(ctx context.Context, in *DeleteInterestRequest, opts ...grpc.CallOption) (*DeleteInterestResponse, error) {
	out := new(DeleteInterestResponse)
	err := c.cc.Invoke(ctx, Admin_DeleteInterest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AddPreference(ctx context.Context, in *AddPreferenceRequest, opts ...grpc.CallOption) (*AddPreferenceResponse, error) {
	out := new(AddPreferenceResponse)
	err := c.cc.Invoke(ctx, Admin_AddPreference_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) EditPreference(ctx context.Context, in *EditPreferenceRequest, opts ...grpc.CallOption) (*EditPreferenceResponse, error) {
	out := new(EditPreferenceResponse)
	err := c.cc.Invoke(ctx, Admin_EditPreference_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) DeletePreference(ctx context.Context, in *DeletePreferenceRequest, opts ...grpc.CallOption) (*DeletePreferenceResponse, error) {
	out := new(DeletePreferenceResponse)
	err := c.cc.Invoke(ctx, Admin_DeletePreference_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	AdminLogin(context.Context, *AdminLoginRequest) (*AdminLoginResponse, error)
	GetInterests(context.Context, *GetInterestsRequest) (*GetInterestsResponse, error)
	GetPreferences(context.Context, *GetPreferencesRequest) (*GetPreferencesResponse, error)
	AddInterest(context.Context, *AddInterestRequest) (*AddInterestResponse, error)
	EditInterest(context.Context, *EditInterestRequest) (*EditInterestResponse, error)
	DeleteInterest(context.Context, *DeleteInterestRequest) (*DeleteInterestResponse, error)
	AddPreference(context.Context, *AddPreferenceRequest) (*AddPreferenceResponse, error)
	EditPreference(context.Context, *EditPreferenceRequest) (*EditPreferenceResponse, error)
	DeletePreference(context.Context, *DeletePreferenceRequest) (*DeletePreferenceResponse, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) AdminLogin(context.Context, *AdminLoginRequest) (*AdminLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogin not implemented")
}
func (UnimplementedAdminServer) GetInterests(context.Context, *GetInterestsRequest) (*GetInterestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInterests not implemented")
}
func (UnimplementedAdminServer) GetPreferences(context.Context, *GetPreferencesRequest) (*GetPreferencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPreferences not implemented")
}
func (UnimplementedAdminServer) AddInterest(context.Context, *AddInterestRequest) (*AddInterestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInterest not implemented")
}
func (UnimplementedAdminServer) EditInterest(context.Context, *EditInterestRequest) (*EditInterestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditInterest not implemented")
}
func (UnimplementedAdminServer) DeleteInterest(context.Context, *DeleteInterestRequest) (*DeleteInterestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInterest not implemented")
}
func (UnimplementedAdminServer) AddPreference(context.Context, *AddPreferenceRequest) (*AddPreferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPreference not implemented")
}
func (UnimplementedAdminServer) EditPreference(context.Context, *EditPreferenceRequest) (*EditPreferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditPreference not implemented")
}
func (UnimplementedAdminServer) DeletePreference(context.Context, *DeletePreferenceRequest) (*DeletePreferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePreference not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_AdminLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminLogin(ctx, req.(*AdminLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetInterests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInterestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetInterests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_GetInterests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetInterests(ctx, req.(*GetInterestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetPreferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPreferencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetPreferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_GetPreferences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetPreferences(ctx, req.(*GetPreferencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AddInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInterestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AddInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AddInterest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AddInterest(ctx, req.(*AddInterestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_EditInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditInterestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).EditInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_EditInterest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).EditInterest(ctx, req.(*EditInterestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_DeleteInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInterestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).DeleteInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_DeleteInterest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).DeleteInterest(ctx, req.(*DeleteInterestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AddPreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPreferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AddPreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AddPreference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AddPreference(ctx, req.(*AddPreferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_EditPreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditPreferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).EditPreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_EditPreference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).EditPreference(ctx, req.(*EditPreferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_DeletePreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePreferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).DeletePreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_DeletePreference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).DeletePreference(ctx, req.(*DeletePreferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminLogin",
			Handler:    _Admin_AdminLogin_Handler,
		},
		{
			MethodName: "GetInterests",
			Handler:    _Admin_GetInterests_Handler,
		},
		{
			MethodName: "GetPreferences",
			Handler:    _Admin_GetPreferences_Handler,
		},
		{
			MethodName: "AddInterest",
			Handler:    _Admin_AddInterest_Handler,
		},
		{
			MethodName: "EditInterest",
			Handler:    _Admin_EditInterest_Handler,
		},
		{
			MethodName: "DeleteInterest",
			Handler:    _Admin_DeleteInterest_Handler,
		},
		{
			MethodName: "AddPreference",
			Handler:    _Admin_AddPreference_Handler,
		},
		{
			MethodName: "EditPreference",
			Handler:    _Admin_EditPreference_Handler,
		},
		{
			MethodName: "DeletePreference",
			Handler:    _Admin_DeletePreference_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/admin/pb.proto",
}
