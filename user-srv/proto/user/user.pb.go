// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package com_island_code_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UpdateUserRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Mobile               string   `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	AvatarURL            string   `protobuf:"bytes,3,opt,name=avatarURL,proto3" json:"avatarURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UpdateUserRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UpdateUserRequest) GetAvatarURL() string {
	if m != nil {
		return m.AvatarURL
	}
	return ""
}

type UpdateUserResponse struct {
	Error                *Error       `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data                 *GetUserData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateUserResponse) Reset()         { *m = UpdateUserResponse{} }
func (m *UpdateUserResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateUserResponse) ProtoMessage()    {}
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *UpdateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserResponse.Unmarshal(m, b)
}
func (m *UpdateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserResponse.Marshal(b, m, deterministic)
}
func (m *UpdateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserResponse.Merge(m, src)
}
func (m *UpdateUserResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateUserResponse.Size(m)
}
func (m *UpdateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserResponse proto.InternalMessageInfo

func (m *UpdateUserResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *UpdateUserResponse) GetData() *GetUserData {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetUserRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type GetUserResponse struct {
	Error                *Error       `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data                 *GetUserData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *GetUserResponse) GetData() *GetUserData {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetUserData struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Mobile               string   `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	AvatarURL            string   `protobuf:"bytes,3,opt,name=avatarURL,proto3" json:"avatarURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserData) Reset()         { *m = GetUserData{} }
func (m *GetUserData) String() string { return proto.CompactTextString(m) }
func (*GetUserData) ProtoMessage()    {}
func (*GetUserData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *GetUserData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserData.Unmarshal(m, b)
}
func (m *GetUserData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserData.Marshal(b, m, deterministic)
}
func (m *GetUserData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserData.Merge(m, src)
}
func (m *GetUserData) XXX_Size() int {
	return xxx_messageInfo_GetUserData.Size(m)
}
func (m *GetUserData) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserData.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserData proto.InternalMessageInfo

func (m *GetUserData) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *GetUserData) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *GetUserData) GetAvatarURL() string {
	if m != nil {
		return m.AvatarURL
	}
	return ""
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{5}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*UpdateUserRequest)(nil), "com.island.code.srv.user.UpdateUserRequest")
	proto.RegisterType((*UpdateUserResponse)(nil), "com.island.code.srv.user.UpdateUserResponse")
	proto.RegisterType((*GetUserRequest)(nil), "com.island.code.srv.user.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "com.island.code.srv.user.GetUserResponse")
	proto.RegisterType((*GetUserData)(nil), "com.island.code.srv.user.GetUserData")
	proto.RegisterType((*Error)(nil), "com.island.code.srv.user.Error")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x35, 0x36, 0xa9, 0x74, 0x02, 0x7e, 0x0c, 0x28, 0xa1, 0x08, 0x4a, 0x40, 0x88, 0xa8, 0x2b,
	0x44, 0x3c, 0x78, 0xaf, 0x88, 0xe0, 0x69, 0x21, 0x27, 0x2f, 0x6e, 0x9b, 0x45, 0x02, 0x4d, 0x37,
	0xee, 0x6e, 0xfb, 0x07, 0x04, 0xff, 0xa0, 0x7f, 0x48, 0x76, 0xb2, 0xda, 0x82, 0x84, 0xf6, 0x20,
	0x78, 0x09, 0x3b, 0x6f, 0xde, 0xec, 0x7b, 0x33, 0x99, 0x85, 0xc3, 0x46, 0x2b, 0xab, 0xae, 0xe7,
	0x46, 0x6a, 0xfa, 0x30, 0x8a, 0x31, 0x99, 0xa8, 0x9a, 0x55, 0x66, 0x2a, 0x66, 0x25, 0x9b, 0xa8,
	0x52, 0x32, 0xa3, 0x17, 0xcc, 0xe5, 0x53, 0x01, 0x07, 0x45, 0x53, 0x0a, 0x2b, 0x0b, 0x23, 0x35,
	0x97, 0x6f, 0x73, 0x69, 0x2c, 0x1e, 0x41, 0xdf, 0x25, 0x1f, 0x47, 0x49, 0x70, 0x1a, 0x64, 0x03,
	0xee, 0x23, 0x87, 0xd7, 0x6a, 0x5c, 0x4d, 0x65, 0xb2, 0xdd, 0xe2, 0x6d, 0x84, 0xc7, 0x30, 0x10,
	0x0b, 0x61, 0x85, 0x2e, 0xf8, 0x53, 0xd2, 0xa3, 0xd4, 0x12, 0x48, 0x3f, 0x02, 0xc0, 0x55, 0x0d,
	0xd3, 0xa8, 0x99, 0x91, 0x78, 0x0b, 0x91, 0xd4, 0x5a, 0x69, 0xd2, 0x88, 0xf3, 0x13, 0xd6, 0xe5,
	0x91, 0xdd, 0x3b, 0x1a, 0x6f, 0xd9, 0x78, 0x07, 0x61, 0x29, 0xac, 0x20, 0x07, 0x71, 0x7e, 0xd6,
	0x5d, 0xf5, 0x20, 0xad, 0xd3, 0x1b, 0x09, 0x2b, 0x38, 0x95, 0xa4, 0x19, 0xec, 0x7a, 0x70, 0x4d,
	0xa3, 0xe9, 0x7b, 0x00, 0x7b, 0x3f, 0xd4, 0x7f, 0xf3, 0xfb, 0x0c, 0xf1, 0x0a, 0xf8, 0xc7, 0x7f,
	0xe5, 0x0a, 0x22, 0xf2, 0x89, 0x08, 0xa1, 0x33, 0x42, 0x97, 0x46, 0x9c, 0xce, 0xb8, 0x0f, 0xbd,
	0xda, 0xbc, 0xfa, 0xfb, 0xdc, 0x31, 0xff, 0x0c, 0x20, 0x74, 0x4e, 0xf0, 0x05, 0x76, 0xbc, 0x29,
	0xcc, 0xd6, 0x36, 0xe3, 0xe7, 0x3c, 0x3c, 0xdf, 0x80, 0xd9, 0x8e, 0x39, 0xdd, 0xc2, 0x0a, 0x60,
	0xb9, 0x2e, 0x78, 0xd1, 0x5d, 0xfa, 0x6b, 0x71, 0x87, 0x97, 0x9b, 0x91, 0xbf, 0xa5, 0xc6, 0x7d,
	0x7a, 0x1e, 0x37, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x52, 0x63, 0x4f, 0x37, 0x03, 0x00,
	0x00,
}
