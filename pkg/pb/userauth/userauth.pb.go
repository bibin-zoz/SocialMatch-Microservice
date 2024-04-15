// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.3
// source: pkg/pb/userauth/userauth.proto

package userauth

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

type UserOtpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserOtpRequest) Reset() {
	*x = UserOtpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOtpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOtpRequest) ProtoMessage() {}

func (x *UserOtpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOtpRequest.ProtoReflect.Descriptor instead.
func (*UserOtpRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_userauth_userauth_proto_rawDescGZIP(), []int{0}
}

func (x *UserOtpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UserOtpRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Email  string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Otp    int64  `protobuf:"varint,3,opt,name=otp,proto3" json:"otp,omitempty"`
}

func (x *UserOtpRequestResponse) Reset() {
	*x = UserOtpRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOtpRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOtpRequestResponse) ProtoMessage() {}

func (x *UserOtpRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOtpRequestResponse.ProtoReflect.Descriptor instead.
func (*UserOtpRequestResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_userauth_userauth_proto_rawDescGZIP(), []int{1}
}

func (x *UserOtpRequestResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UserOtpRequestResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserOtpRequestResponse) GetOtp() int64 {
	if x != nil {
		return x.Otp
	}
	return 0
}

type UserOtpVerificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Otp   int64  `protobuf:"varint,2,opt,name=otp,proto3" json:"otp,omitempty"`
}

func (x *UserOtpVerificationRequest) Reset() {
	*x = UserOtpVerificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOtpVerificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOtpVerificationRequest) ProtoMessage() {}

func (x *UserOtpVerificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOtpVerificationRequest.ProtoReflect.Descriptor instead.
func (*UserOtpVerificationRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_userauth_userauth_proto_rawDescGZIP(), []int{2}
}

func (x *UserOtpVerificationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserOtpVerificationRequest) GetOtp() int64 {
	if x != nil {
		return x.Otp
	}
	return 0
}

type UserOtpVerificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      int64        `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	UserDetails *UserDetails `protobuf:"bytes,2,opt,name=userDetails,proto3" json:"userDetails,omitempty"`
}

func (x *UserOtpVerificationResponse) Reset() {
	*x = UserOtpVerificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOtpVerificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOtpVerificationResponse) ProtoMessage() {}

func (x *UserOtpVerificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOtpVerificationResponse.ProtoReflect.Descriptor instead.
func (*UserOtpVerificationResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_userauth_userauth_proto_rawDescGZIP(), []int{3}
}

func (x *UserOtpVerificationResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UserOtpVerificationResponse) GetUserDetails() *UserDetails {
	if x != nil {
		return x.UserDetails
	}
	return nil
}

type UserSignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname string `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password  string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Phone     string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Genderid  int32  `protobuf:"varint,7,opt,name=genderid,proto3" json:"genderid,omitempty"`
	Age       int32  `protobuf:"varint,8,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *UserSignUpRequest) Reset() {
	*x = UserSignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSignUpRequest) ProtoMessage() {}

func (x *UserSignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSignUpRequest.ProtoReflect.Descriptor instead.
func (*UserSignUpRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_userauth_userauth_proto_rawDescGZIP(), []int{4}
}

func (x *UserSignUpRequest) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *UserSignUpRequest) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *UserSignUpRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserSignUpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserSignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserSignUpRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserSignUpRequest) GetGenderid() int32 {
	if x != nil {
		return x.Genderid
	}
	return 0
}

func (x *UserSignUpRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type UserEditDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname string `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password  string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Phone     string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Genderid  int32  `protobuf:"varint,7,opt,name=genderid,proto3" json:"genderid,omitempty"`
	Age       int32  `protobuf:"varint,8,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *UserEditDetailsRequest) Reset() {
	*x = UserEditDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEditDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEditDetailsRequest) ProtoMessage() {}

func (x *UserEditDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEditDetailsRequest.ProtoReflect.Descriptor instead.
func (*UserEditDetailsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_userauth_userauth_proto_rawDescGZIP(), []int{5}
}

func (x *UserEditDetailsRequest) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *UserEditDetailsRequest) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *UserEditDetailsRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserEditDetailsRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserEditDetailsRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserEditDetailsRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserEditDetailsRequest) GetGenderid() int32 {
	if x != nil {
		return x.Genderid
	}
	return 0
}

func (x *UserEditDetailsRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type UserDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Firstname string `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *UserDetails) Reset() {
	*x = UserDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetails) ProtoMessage() {}

func (x *UserDetails) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[6]
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
	return file_pkg_pb_userauth_userauth_proto_rawDescGZIP(), []int{6}
}

func (x *UserDetails) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserDetails) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *UserDetails) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *UserDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserDetails) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type UserEditDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      int64        `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	UserDetails *UserDetails `protobuf:"bytes,2,opt,name=userDetails,proto3" json:"userDetails,omitempty"`
}

func (x *UserEditDetailsResponse) Reset() {
	*x = UserEditDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEditDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEditDetailsResponse) ProtoMessage() {}

func (x *UserEditDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEditDetailsResponse.ProtoReflect.Descriptor instead.
func (*UserEditDetailsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_userauth_userauth_proto_rawDescGZIP(), []int{7}
}

func (x *UserEditDetailsResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UserEditDetailsResponse) GetUserDetails() *UserDetails {
	if x != nil {
		return x.UserDetails
	}
	return nil
}

type UserSignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       int64        `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	UserDetails  *UserDetails `protobuf:"bytes,2,opt,name=userDetails,proto3" json:"userDetails,omitempty"`
	AccessToken  string       `protobuf:"bytes,3,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	RefreshToken string       `protobuf:"bytes,4,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
}

func (x *UserSignUpResponse) Reset() {
	*x = UserSignUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSignUpResponse) ProtoMessage() {}

func (x *UserSignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSignUpResponse.ProtoReflect.Descriptor instead.
func (*UserSignUpResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_userauth_userauth_proto_rawDescGZIP(), []int{8}
}

func (x *UserSignUpResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UserSignUpResponse) GetUserDetails() *UserDetails {
	if x != nil {
		return x.UserDetails
	}
	return nil
}

func (x *UserSignUpResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *UserSignUpResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type UserLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserLoginRequest) Reset() {
	*x = UserLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginRequest) ProtoMessage() {}

func (x *UserLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginRequest.ProtoReflect.Descriptor instead.
func (*UserLoginRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_userauth_userauth_proto_rawDescGZIP(), []int{9}
}

func (x *UserLoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       int64        `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	UserDetails  *UserDetails `protobuf:"bytes,2,opt,name=userDetails,proto3" json:"userDetails,omitempty"`
	AccessToken  string       `protobuf:"bytes,3,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	RefreshToken string       `protobuf:"bytes,4,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
}

func (x *UserLoginResponse) Reset() {
	*x = UserLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginResponse) ProtoMessage() {}

func (x *UserLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_userauth_userauth_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginResponse.ProtoReflect.Descriptor instead.
func (*UserLoginResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_userauth_userauth_proto_rawDescGZIP(), []int{10}
}

func (x *UserLoginResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UserLoginResponse) GetUserDetails() *UserDetails {
	if x != nil {
		return x.UserDetails
	}
	return nil
}

func (x *UserLoginResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *UserLoginResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

var File_pkg_pb_userauth_userauth_proto protoreflect.FileDescriptor

var file_pkg_pb_userauth_userauth_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x26, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x74,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x58,
	0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x22, 0x44, 0x0a, 0x1a, 0x55, 0x73, 0x65, 0x72,
	0x4f, 0x74, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03,
	0x6f, 0x74, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x22, 0x6a,
	0x0a, 0x1b, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x74, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0b, 0x75,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xdf, 0x01, 0x0a, 0x11, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0xe4, 0x01, 0x0a,
	0x16, 0x55, 0x73, 0x65, 0x72, 0x45, 0x64, 0x69, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x66, 0x0a, 0x17, 0x55, 0x73, 0x65,
	0x72, 0x45, 0x64, 0x69, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x0b,
	0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x22, 0xa7, 0x01, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x33, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x44, 0x0a, 0x10, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0xa6, 0x01, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x33, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x84, 0x03, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x55,
	0x70, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67,
	0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x45, 0x64,
	0x69, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x64, 0x69, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x45, 0x64, 0x69, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72,
	0x4f, 0x74, 0x70, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4f,
	0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x74, 0x70, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x74, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x74, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_userauth_userauth_proto_rawDescOnce sync.Once
	file_pkg_pb_userauth_userauth_proto_rawDescData = file_pkg_pb_userauth_userauth_proto_rawDesc
)

func file_pkg_pb_userauth_userauth_proto_rawDescGZIP() []byte {
	file_pkg_pb_userauth_userauth_proto_rawDescOnce.Do(func() {
		file_pkg_pb_userauth_userauth_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_userauth_userauth_proto_rawDescData)
	})
	return file_pkg_pb_userauth_userauth_proto_rawDescData
}

var file_pkg_pb_userauth_userauth_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pkg_pb_userauth_userauth_proto_goTypes = []interface{}{
	(*UserOtpRequest)(nil),              // 0: user.UserOtpRequest
	(*UserOtpRequestResponse)(nil),      // 1: user.UserOtpRequestResponse
	(*UserOtpVerificationRequest)(nil),  // 2: user.UserOtpVerificationRequest
	(*UserOtpVerificationResponse)(nil), // 3: user.UserOtpVerificationResponse
	(*UserSignUpRequest)(nil),           // 4: user.UserSignUpRequest
	(*UserEditDetailsRequest)(nil),      // 5: user.UserEditDetailsRequest
	(*UserDetails)(nil),                 // 6: user.UserDetails
	(*UserEditDetailsResponse)(nil),     // 7: user.UserEditDetailsResponse
	(*UserSignUpResponse)(nil),          // 8: user.UserSignUpResponse
	(*UserLoginRequest)(nil),            // 9: user.UserLoginRequest
	(*UserLoginResponse)(nil),           // 10: user.UserLoginResponse
}
var file_pkg_pb_userauth_userauth_proto_depIdxs = []int32{
	6,  // 0: user.UserOtpVerificationResponse.userDetails:type_name -> user.UserDetails
	6,  // 1: user.UserEditDetailsResponse.userDetails:type_name -> user.UserDetails
	6,  // 2: user.UserSignUpResponse.userDetails:type_name -> user.UserDetails
	6,  // 3: user.UserLoginResponse.userDetails:type_name -> user.UserDetails
	4,  // 4: user.User.UserSignUp:input_type -> user.UserSignUpRequest
	9,  // 5: user.User.UserLogin:input_type -> user.UserLoginRequest
	5,  // 6: user.User.UserEditDetails:input_type -> user.UserEditDetailsRequest
	0,  // 7: user.User.UserOtpGeneration:input_type -> user.UserOtpRequest
	2,  // 8: user.User.UserOtpVerification:input_type -> user.UserOtpVerificationRequest
	8,  // 9: user.User.UserSignUp:output_type -> user.UserSignUpResponse
	10, // 10: user.User.UserLogin:output_type -> user.UserLoginResponse
	7,  // 11: user.User.UserEditDetails:output_type -> user.UserEditDetailsResponse
	1,  // 12: user.User.UserOtpGeneration:output_type -> user.UserOtpRequestResponse
	3,  // 13: user.User.UserOtpVerification:output_type -> user.UserOtpVerificationResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_pb_userauth_userauth_proto_init() }
func file_pkg_pb_userauth_userauth_proto_init() {
	if File_pkg_pb_userauth_userauth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_userauth_userauth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserOtpRequest); i {
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
		file_pkg_pb_userauth_userauth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserOtpRequestResponse); i {
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
		file_pkg_pb_userauth_userauth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserOtpVerificationRequest); i {
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
		file_pkg_pb_userauth_userauth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserOtpVerificationResponse); i {
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
		file_pkg_pb_userauth_userauth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSignUpRequest); i {
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
		file_pkg_pb_userauth_userauth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEditDetailsRequest); i {
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
		file_pkg_pb_userauth_userauth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_pb_userauth_userauth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEditDetailsResponse); i {
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
		file_pkg_pb_userauth_userauth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSignUpResponse); i {
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
		file_pkg_pb_userauth_userauth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginRequest); i {
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
		file_pkg_pb_userauth_userauth_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginResponse); i {
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
			RawDescriptor: file_pkg_pb_userauth_userauth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_userauth_userauth_proto_goTypes,
		DependencyIndexes: file_pkg_pb_userauth_userauth_proto_depIdxs,
		MessageInfos:      file_pkg_pb_userauth_userauth_proto_msgTypes,
	}.Build()
	File_pkg_pb_userauth_userauth_proto = out.File
	file_pkg_pb_userauth_userauth_proto_rawDesc = nil
	file_pkg_pb_userauth_userauth_proto_goTypes = nil
	file_pkg_pb_userauth_userauth_proto_depIdxs = nil
}
