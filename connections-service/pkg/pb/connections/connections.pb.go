// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.3
// source: pkg/pb/connections/connections.proto

package connections

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetConnectionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      int64          `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	UserDetails []*UserDetails `protobuf:"bytes,2,rep,name=UserDetails,proto3" json:"UserDetails,omitempty"`
}

func (x *GetConnectionsResponse) Reset() {
	*x = GetConnectionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_connections_connections_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConnectionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConnectionsResponse) ProtoMessage() {}

func (x *GetConnectionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_connections_connections_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConnectionsResponse.ProtoReflect.Descriptor instead.
func (*GetConnectionsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_connections_connections_proto_rawDescGZIP(), []int{0}
}

func (x *GetConnectionsResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetConnectionsResponse) GetUserDetails() []*UserDetails {
	if x != nil {
		return x.UserDetails
	}
	return nil
}

type ConnectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connected bool `protobuf:"varint,1,opt,name=connected,proto3" json:"connected,omitempty"`
}

func (x *ConnectionResponse) Reset() {
	*x = ConnectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_connections_connections_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionResponse) ProtoMessage() {}

func (x *ConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_connections_connections_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionResponse.ProtoReflect.Descriptor instead.
func (*ConnectionResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_connections_connections_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectionResponse) GetConnected() bool {
	if x != nil {
		return x.Connected
	}
	return false
}

type CheckUserConnectionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID     int32 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	ReceiverID int32 `protobuf:"varint,2,opt,name=receiverID,proto3" json:"receiverID,omitempty"`
}

func (x *CheckUserConnectionReq) Reset() {
	*x = CheckUserConnectionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_connections_connections_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserConnectionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserConnectionReq) ProtoMessage() {}

func (x *CheckUserConnectionReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_connections_connections_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserConnectionReq.ProtoReflect.Descriptor instead.
func (*CheckUserConnectionReq) Descriptor() ([]byte, []int) {
	return file_pkg_pb_connections_connections_proto_rawDescGZIP(), []int{2}
}

func (x *CheckUserConnectionReq) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CheckUserConnectionReq) GetReceiverID() int32 {
	if x != nil {
		return x.ReceiverID
	}
	return 0
}

type UserDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserDetails) Reset() {
	*x = UserDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_connections_connections_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetails) ProtoMessage() {}

func (x *UserDetails) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_connections_connections_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetails.ProtoReflect.Descriptor instead.
func (*UserDetails) Descriptor() ([]byte, []int) {
	return file_pkg_pb_connections_connections_proto_rawDescGZIP(), []int{3}
}

func (x *UserDetails) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetConnectionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetConnectionsRequest) Reset() {
	*x = GetConnectionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_connections_connections_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConnectionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConnectionsRequest) ProtoMessage() {}

func (x *GetConnectionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_connections_connections_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConnectionsRequest.ProtoReflect.Descriptor instead.
func (*GetConnectionsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_connections_connections_proto_rawDescGZIP(), []int{4}
}

func (x *GetConnectionsRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type BlockConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId uint64 `protobuf:"varint,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	UserId   uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BlockConnectionRequest) Reset() {
	*x = BlockConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_connections_connections_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockConnectionRequest) ProtoMessage() {}

func (x *BlockConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_connections_connections_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockConnectionRequest.ProtoReflect.Descriptor instead.
func (*BlockConnectionRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_connections_connections_proto_rawDescGZIP(), []int{5}
}

func (x *BlockConnectionRequest) GetSenderId() uint64 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *BlockConnectionRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type BlockConnectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *BlockConnectionResponse) Reset() {
	*x = BlockConnectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_connections_connections_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockConnectionResponse) ProtoMessage() {}

func (x *BlockConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_connections_connections_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockConnectionResponse.ProtoReflect.Descriptor instead.
func (*BlockConnectionResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_connections_connections_proto_rawDescGZIP(), []int{6}
}

func (x *BlockConnectionResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type FollowUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid   int64 `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Senderid int64 `protobuf:"varint,2,opt,name=senderid,proto3" json:"senderid,omitempty"`
}

func (x *FollowUserRequest) Reset() {
	*x = FollowUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_connections_connections_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowUserRequest) ProtoMessage() {}

func (x *FollowUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_connections_connections_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowUserRequest.ProtoReflect.Descriptor instead.
func (*FollowUserRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_connections_connections_proto_rawDescGZIP(), []int{7}
}

func (x *FollowUserRequest) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *FollowUserRequest) GetSenderid() int64 {
	if x != nil {
		return x.Senderid
	}
	return 0
}

type FollowUserResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *FollowUserResponce) Reset() {
	*x = FollowUserResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_connections_connections_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowUserResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowUserResponce) ProtoMessage() {}

func (x *FollowUserResponce) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_connections_connections_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowUserResponce.ProtoReflect.Descriptor instead.
func (*FollowUserResponce) Descriptor() ([]byte, []int) {
	return file_pkg_pb_connections_connections_proto_rawDescGZIP(), []int{8}
}

func (x *FollowUserResponce) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_pkg_pb_connections_connections_proto protoreflect.FileDescriptor

var file_pkg_pb_connections_connections_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x6c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3a, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x22, 0x32, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x50, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x44, 0x22, 0x1d, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x16, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x17, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x47, 0x0a, 0x11, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x12, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x32, 0xf6, 0x02, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x4f, 0x0a, 0x0a, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63,
	0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b,
	0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x2e,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_connections_connections_proto_rawDescOnce sync.Once
	file_pkg_pb_connections_connections_proto_rawDescData = file_pkg_pb_connections_connections_proto_rawDesc
)

func file_pkg_pb_connections_connections_proto_rawDescGZIP() []byte {
	file_pkg_pb_connections_connections_proto_rawDescOnce.Do(func() {
		file_pkg_pb_connections_connections_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_connections_connections_proto_rawDescData)
	})
	return file_pkg_pb_connections_connections_proto_rawDescData
}

var file_pkg_pb_connections_connections_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_pb_connections_connections_proto_goTypes = []interface{}{
	(*GetConnectionsResponse)(nil),  // 0: connections.GetConnectionsResponse
	(*ConnectionResponse)(nil),      // 1: connections.ConnectionResponse
	(*CheckUserConnectionReq)(nil),  // 2: connections.CheckUserConnectionReq
	(*UserDetails)(nil),             // 3: connections.UserDetails
	(*GetConnectionsRequest)(nil),   // 4: connections.GetConnectionsRequest
	(*BlockConnectionRequest)(nil),  // 5: connections.BlockConnectionRequest
	(*BlockConnectionResponse)(nil), // 6: connections.BlockConnectionResponse
	(*FollowUserRequest)(nil),       // 7: connections.FollowUserRequest
	(*FollowUserResponce)(nil),      // 8: connections.FollowUserResponce
}
var file_pkg_pb_connections_connections_proto_depIdxs = []int32{
	3, // 0: connections.GetConnectionsResponse.UserDetails:type_name -> connections.UserDetails
	7, // 1: connections.connections.FollowUser:input_type -> connections.FollowUserRequest
	5, // 2: connections.connections.BlockConnection:input_type -> connections.BlockConnectionRequest
	4, // 3: connections.connections.GetConnections:input_type -> connections.GetConnectionsRequest
	2, // 4: connections.connections.CheckUserConnection:input_type -> connections.CheckUserConnectionReq
	8, // 5: connections.connections.FollowUser:output_type -> connections.FollowUserResponce
	6, // 6: connections.connections.BlockConnection:output_type -> connections.BlockConnectionResponse
	0, // 7: connections.connections.GetConnections:output_type -> connections.GetConnectionsResponse
	1, // 8: connections.connections.CheckUserConnection:output_type -> connections.ConnectionResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_pb_connections_connections_proto_init() }
func file_pkg_pb_connections_connections_proto_init() {
	if File_pkg_pb_connections_connections_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_connections_connections_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConnectionsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_connections_connections_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_connections_connections_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserConnectionReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_connections_connections_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetails); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_connections_connections_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConnectionsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_connections_connections_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockConnectionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_connections_connections_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockConnectionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_connections_connections_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_connections_connections_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowUserResponce); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_pb_connections_connections_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_connections_connections_proto_goTypes,
		DependencyIndexes: file_pkg_pb_connections_connections_proto_depIdxs,
		MessageInfos:      file_pkg_pb_connections_connections_proto_msgTypes,
	}.Build()
	File_pkg_pb_connections_connections_proto = out.File
	file_pkg_pb_connections_connections_proto_rawDesc = nil
	file_pkg_pb_connections_connections_proto_goTypes = nil
	file_pkg_pb_connections_connections_proto_depIdxs = nil
}
