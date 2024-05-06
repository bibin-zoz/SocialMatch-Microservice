// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: pkg/pb/room/room.proto

package room

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
	RoomService_CreateRoom_FullMethodName          = "/room.RoomService/CreateRoom"
	RoomService_EditRoom_FullMethodName            = "/room.RoomService/EditRoom"
	RoomService_ChangeStatus_FullMethodName        = "/room.RoomService/ChangeStatus"
	RoomService_AddMembers_FullMethodName          = "/room.RoomService/AddMembers"
	RoomService_SeeRoomJoinRequests_FullMethodName = "/room.RoomService/SeeRoomJoinRequests"
	RoomService_GetAllRooms_FullMethodName         = "/room.RoomService/GetAllRooms"
	RoomService_GetGroupMembers_FullMethodName     = "/room.RoomService/GetGroupMembers"
	RoomService_SendMessage_FullMethodName         = "/room.RoomService/SendMessage"
	RoomService_ReadMessages_FullMethodName        = "/room.RoomService/ReadMessages"
)

// RoomServiceClient is the client API for RoomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoomServiceClient interface {
	// CreateRoom creates a new room
	CreateRoom(ctx context.Context, in *RoomCreateRequest, opts ...grpc.CallOption) (*Room, error)
	// EditRoom edits an existing room
	EditRoom(ctx context.Context, in *RoomEditRequest, opts ...grpc.CallOption) (*Room, error)
	// ChangeStatus changes the status of a room
	ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*Room, error)
	// AddMembers adds members to a room
	AddMembers(ctx context.Context, in *AddMembersRequest, opts ...grpc.CallOption) (*Room, error)
	// SeeRoomJoinRequests retrieves join requests for a room
	SeeRoomJoinRequests(ctx context.Context, in *RoomJoinRequestsRequest, opts ...grpc.CallOption) (*RoomJoinRequestsResponse, error)
	// GetAllRooms retrieves all rooms
	GetAllRooms(ctx context.Context, in *GetAllRoomsRequest, opts ...grpc.CallOption) (*GetAllRoomsResponse, error)
	// New RPC method to get group members using room ID
	GetGroupMembers(ctx context.Context, in *GetGroupMembersRequest, opts ...grpc.CallOption) (*GetGroupMembersResponse, error)
	// New RPC method to send a message in a room
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*Message, error)
	// New RPC method to read messages in a room
	ReadMessages(ctx context.Context, in *ReadMessagesRequest, opts ...grpc.CallOption) (*ReadMessagesResponse, error)
}

type roomServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoomServiceClient(cc grpc.ClientConnInterface) RoomServiceClient {
	return &roomServiceClient{cc}
}

func (c *roomServiceClient) CreateRoom(ctx context.Context, in *RoomCreateRequest, opts ...grpc.CallOption) (*Room, error) {
	out := new(Room)
	err := c.cc.Invoke(ctx, RoomService_CreateRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) EditRoom(ctx context.Context, in *RoomEditRequest, opts ...grpc.CallOption) (*Room, error) {
	out := new(Room)
	err := c.cc.Invoke(ctx, RoomService_EditRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*Room, error) {
	out := new(Room)
	err := c.cc.Invoke(ctx, RoomService_ChangeStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) AddMembers(ctx context.Context, in *AddMembersRequest, opts ...grpc.CallOption) (*Room, error) {
	out := new(Room)
	err := c.cc.Invoke(ctx, RoomService_AddMembers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) SeeRoomJoinRequests(ctx context.Context, in *RoomJoinRequestsRequest, opts ...grpc.CallOption) (*RoomJoinRequestsResponse, error) {
	out := new(RoomJoinRequestsResponse)
	err := c.cc.Invoke(ctx, RoomService_SeeRoomJoinRequests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) GetAllRooms(ctx context.Context, in *GetAllRoomsRequest, opts ...grpc.CallOption) (*GetAllRoomsResponse, error) {
	out := new(GetAllRoomsResponse)
	err := c.cc.Invoke(ctx, RoomService_GetAllRooms_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) GetGroupMembers(ctx context.Context, in *GetGroupMembersRequest, opts ...grpc.CallOption) (*GetGroupMembersResponse, error) {
	out := new(GetGroupMembersResponse)
	err := c.cc.Invoke(ctx, RoomService_GetGroupMembers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, RoomService_SendMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) ReadMessages(ctx context.Context, in *ReadMessagesRequest, opts ...grpc.CallOption) (*ReadMessagesResponse, error) {
	out := new(ReadMessagesResponse)
	err := c.cc.Invoke(ctx, RoomService_ReadMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomServiceServer is the server API for RoomService service.
// All implementations must embed UnimplementedRoomServiceServer
// for forward compatibility
type RoomServiceServer interface {
	// CreateRoom creates a new room
	CreateRoom(context.Context, *RoomCreateRequest) (*Room, error)
	// EditRoom edits an existing room
	EditRoom(context.Context, *RoomEditRequest) (*Room, error)
	// ChangeStatus changes the status of a room
	ChangeStatus(context.Context, *ChangeStatusRequest) (*Room, error)
	// AddMembers adds members to a room
	AddMembers(context.Context, *AddMembersRequest) (*Room, error)
	// SeeRoomJoinRequests retrieves join requests for a room
	SeeRoomJoinRequests(context.Context, *RoomJoinRequestsRequest) (*RoomJoinRequestsResponse, error)
	// GetAllRooms retrieves all rooms
	GetAllRooms(context.Context, *GetAllRoomsRequest) (*GetAllRoomsResponse, error)
	// New RPC method to get group members using room ID
	GetGroupMembers(context.Context, *GetGroupMembersRequest) (*GetGroupMembersResponse, error)
	// New RPC method to send a message in a room
	SendMessage(context.Context, *SendMessageRequest) (*Message, error)
	// New RPC method to read messages in a room
	ReadMessages(context.Context, *ReadMessagesRequest) (*ReadMessagesResponse, error)
	mustEmbedUnimplementedRoomServiceServer()
}

// UnimplementedRoomServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoomServiceServer struct {
}

func (UnimplementedRoomServiceServer) CreateRoom(context.Context, *RoomCreateRequest) (*Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedRoomServiceServer) EditRoom(context.Context, *RoomEditRequest) (*Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditRoom not implemented")
}
func (UnimplementedRoomServiceServer) ChangeStatus(context.Context, *ChangeStatusRequest) (*Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeStatus not implemented")
}
func (UnimplementedRoomServiceServer) AddMembers(context.Context, *AddMembersRequest) (*Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMembers not implemented")
}
func (UnimplementedRoomServiceServer) SeeRoomJoinRequests(context.Context, *RoomJoinRequestsRequest) (*RoomJoinRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SeeRoomJoinRequests not implemented")
}
func (UnimplementedRoomServiceServer) GetAllRooms(context.Context, *GetAllRoomsRequest) (*GetAllRoomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRooms not implemented")
}
func (UnimplementedRoomServiceServer) GetGroupMembers(context.Context, *GetGroupMembersRequest) (*GetGroupMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupMembers not implemented")
}
func (UnimplementedRoomServiceServer) SendMessage(context.Context, *SendMessageRequest) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedRoomServiceServer) ReadMessages(context.Context, *ReadMessagesRequest) (*ReadMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadMessages not implemented")
}
func (UnimplementedRoomServiceServer) mustEmbedUnimplementedRoomServiceServer() {}

// UnsafeRoomServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoomServiceServer will
// result in compilation errors.
type UnsafeRoomServiceServer interface {
	mustEmbedUnimplementedRoomServiceServer()
}

func RegisterRoomServiceServer(s grpc.ServiceRegistrar, srv RoomServiceServer) {
	s.RegisterService(&RoomService_ServiceDesc, srv)
}

func _RoomService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomService_CreateRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).CreateRoom(ctx, req.(*RoomCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_EditRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomEditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).EditRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomService_EditRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).EditRoom(ctx, req.(*RoomEditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_ChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).ChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomService_ChangeStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).ChangeStatus(ctx, req.(*ChangeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_AddMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).AddMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomService_AddMembers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).AddMembers(ctx, req.(*AddMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_SeeRoomJoinRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomJoinRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).SeeRoomJoinRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomService_SeeRoomJoinRequests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).SeeRoomJoinRequests(ctx, req.(*RoomJoinRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_GetAllRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).GetAllRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomService_GetAllRooms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).GetAllRooms(ctx, req.(*GetAllRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_GetGroupMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).GetGroupMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomService_GetGroupMembers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).GetGroupMembers(ctx, req.(*GetGroupMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomService_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_ReadMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).ReadMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoomService_ReadMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).ReadMessages(ctx, req.(*ReadMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoomService_ServiceDesc is the grpc.ServiceDesc for RoomService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoomService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "room.RoomService",
	HandlerType: (*RoomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _RoomService_CreateRoom_Handler,
		},
		{
			MethodName: "EditRoom",
			Handler:    _RoomService_EditRoom_Handler,
		},
		{
			MethodName: "ChangeStatus",
			Handler:    _RoomService_ChangeStatus_Handler,
		},
		{
			MethodName: "AddMembers",
			Handler:    _RoomService_AddMembers_Handler,
		},
		{
			MethodName: "SeeRoomJoinRequests",
			Handler:    _RoomService_SeeRoomJoinRequests_Handler,
		},
		{
			MethodName: "GetAllRooms",
			Handler:    _RoomService_GetAllRooms_Handler,
		},
		{
			MethodName: "GetGroupMembers",
			Handler:    _RoomService_GetGroupMembers_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _RoomService_SendMessage_Handler,
		},
		{
			MethodName: "ReadMessages",
			Handler:    _RoomService_ReadMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/room/room.proto",
}