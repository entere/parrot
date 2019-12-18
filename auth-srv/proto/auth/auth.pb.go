// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

package com_island_code_srv_auth

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

type MakeTokenRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	DeviceID             string   `protobuf:"bytes,2,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MakeTokenRequest) Reset()         { *m = MakeTokenRequest{} }
func (m *MakeTokenRequest) String() string { return proto.CompactTextString(m) }
func (*MakeTokenRequest) ProtoMessage()    {}
func (*MakeTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{0}
}

func (m *MakeTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MakeTokenRequest.Unmarshal(m, b)
}
func (m *MakeTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MakeTokenRequest.Marshal(b, m, deterministic)
}
func (m *MakeTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MakeTokenRequest.Merge(m, src)
}
func (m *MakeTokenRequest) XXX_Size() int {
	return xxx_messageInfo_MakeTokenRequest.Size(m)
}
func (m *MakeTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MakeTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MakeTokenRequest proto.InternalMessageInfo

func (m *MakeTokenRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *MakeTokenRequest) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

type MakeTokenResponse struct {
	Error                *Error         `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data                 *MakeTokenData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MakeTokenResponse) Reset()         { *m = MakeTokenResponse{} }
func (m *MakeTokenResponse) String() string { return proto.CompactTextString(m) }
func (*MakeTokenResponse) ProtoMessage()    {}
func (*MakeTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{1}
}

func (m *MakeTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MakeTokenResponse.Unmarshal(m, b)
}
func (m *MakeTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MakeTokenResponse.Marshal(b, m, deterministic)
}
func (m *MakeTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MakeTokenResponse.Merge(m, src)
}
func (m *MakeTokenResponse) XXX_Size() int {
	return xxx_messageInfo_MakeTokenResponse.Size(m)
}
func (m *MakeTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MakeTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MakeTokenResponse proto.InternalMessageInfo

func (m *MakeTokenResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *MakeTokenResponse) GetData() *MakeTokenData {
	if m != nil {
		return m.Data
	}
	return nil
}

type MakeTokenData struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MakeTokenData) Reset()         { *m = MakeTokenData{} }
func (m *MakeTokenData) String() string { return proto.CompactTextString(m) }
func (*MakeTokenData) ProtoMessage()    {}
func (*MakeTokenData) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{2}
}

func (m *MakeTokenData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MakeTokenData.Unmarshal(m, b)
}
func (m *MakeTokenData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MakeTokenData.Marshal(b, m, deterministic)
}
func (m *MakeTokenData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MakeTokenData.Merge(m, src)
}
func (m *MakeTokenData) XXX_Size() int {
	return xxx_messageInfo_MakeTokenData.Size(m)
}
func (m *MakeTokenData) XXX_DiscardUnknown() {
	xxx_messageInfo_MakeTokenData.DiscardUnknown(m)
}

var xxx_messageInfo_MakeTokenData proto.InternalMessageInfo

func (m *MakeTokenData) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DelTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelTokenRequest) Reset()         { *m = DelTokenRequest{} }
func (m *DelTokenRequest) String() string { return proto.CompactTextString(m) }
func (*DelTokenRequest) ProtoMessage()    {}
func (*DelTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{3}
}

func (m *DelTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelTokenRequest.Unmarshal(m, b)
}
func (m *DelTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelTokenRequest.Marshal(b, m, deterministic)
}
func (m *DelTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelTokenRequest.Merge(m, src)
}
func (m *DelTokenRequest) XXX_Size() int {
	return xxx_messageInfo_DelTokenRequest.Size(m)
}
func (m *DelTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DelTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DelTokenRequest proto.InternalMessageInfo

func (m *DelTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DelTokenResponse struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelTokenResponse) Reset()         { *m = DelTokenResponse{} }
func (m *DelTokenResponse) String() string { return proto.CompactTextString(m) }
func (*DelTokenResponse) ProtoMessage()    {}
func (*DelTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{4}
}

func (m *DelTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelTokenResponse.Unmarshal(m, b)
}
func (m *DelTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelTokenResponse.Marshal(b, m, deterministic)
}
func (m *DelTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelTokenResponse.Merge(m, src)
}
func (m *DelTokenResponse) XXX_Size() int {
	return xxx_messageInfo_DelTokenResponse.Size(m)
}
func (m *DelTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DelTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DelTokenResponse proto.InternalMessageInfo

func (m *DelTokenResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type GetCachedTokenRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	DeviceID             string   `protobuf:"bytes,2,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCachedTokenRequest) Reset()         { *m = GetCachedTokenRequest{} }
func (m *GetCachedTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetCachedTokenRequest) ProtoMessage()    {}
func (*GetCachedTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{5}
}

func (m *GetCachedTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCachedTokenRequest.Unmarshal(m, b)
}
func (m *GetCachedTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCachedTokenRequest.Marshal(b, m, deterministic)
}
func (m *GetCachedTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCachedTokenRequest.Merge(m, src)
}
func (m *GetCachedTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetCachedTokenRequest.Size(m)
}
func (m *GetCachedTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCachedTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCachedTokenRequest proto.InternalMessageInfo

func (m *GetCachedTokenRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *GetCachedTokenRequest) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

type GetCachedTokenResponse struct {
	Error                *Error              `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data                 *GetCachedTokenData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetCachedTokenResponse) Reset()         { *m = GetCachedTokenResponse{} }
func (m *GetCachedTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GetCachedTokenResponse) ProtoMessage()    {}
func (*GetCachedTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{6}
}

func (m *GetCachedTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCachedTokenResponse.Unmarshal(m, b)
}
func (m *GetCachedTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCachedTokenResponse.Marshal(b, m, deterministic)
}
func (m *GetCachedTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCachedTokenResponse.Merge(m, src)
}
func (m *GetCachedTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GetCachedTokenResponse.Size(m)
}
func (m *GetCachedTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCachedTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCachedTokenResponse proto.InternalMessageInfo

func (m *GetCachedTokenResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *GetCachedTokenResponse) GetData() *GetCachedTokenData {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetCachedTokenData struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCachedTokenData) Reset()         { *m = GetCachedTokenData{} }
func (m *GetCachedTokenData) String() string { return proto.CompactTextString(m) }
func (*GetCachedTokenData) ProtoMessage()    {}
func (*GetCachedTokenData) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{7}
}

func (m *GetCachedTokenData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCachedTokenData.Unmarshal(m, b)
}
func (m *GetCachedTokenData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCachedTokenData.Marshal(b, m, deterministic)
}
func (m *GetCachedTokenData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCachedTokenData.Merge(m, src)
}
func (m *GetCachedTokenData) XXX_Size() int {
	return xxx_messageInfo_GetCachedTokenData.Size(m)
}
func (m *GetCachedTokenData) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCachedTokenData.DiscardUnknown(m)
}

var xxx_messageInfo_GetCachedTokenData proto.InternalMessageInfo

func (m *GetCachedTokenData) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type QueryUserData struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	LoginName            string   `protobuf:"bytes,2,opt,name=loginName,proto3" json:"loginName,omitempty"`
	LoginPwd             string   `protobuf:"bytes,3,opt,name=loginPwd,proto3" json:"loginPwd,omitempty"`
	Token                string   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	DeviceID             string   `protobuf:"bytes,5,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryUserData) Reset()         { *m = QueryUserData{} }
func (m *QueryUserData) String() string { return proto.CompactTextString(m) }
func (*QueryUserData) ProtoMessage()    {}
func (*QueryUserData) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{8}
}

func (m *QueryUserData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUserData.Unmarshal(m, b)
}
func (m *QueryUserData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUserData.Marshal(b, m, deterministic)
}
func (m *QueryUserData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserData.Merge(m, src)
}
func (m *QueryUserData) XXX_Size() int {
	return xxx_messageInfo_QueryUserData.Size(m)
}
func (m *QueryUserData) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserData.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserData proto.InternalMessageInfo

func (m *QueryUserData) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *QueryUserData) GetLoginName() string {
	if m != nil {
		return m.LoginName
	}
	return ""
}

func (m *QueryUserData) GetLoginPwd() string {
	if m != nil {
		return m.LoginPwd
	}
	return ""
}

func (m *QueryUserData) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *QueryUserData) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

type QueryUserRequest struct {
	LoginName            string   `protobuf:"bytes,1,opt,name=loginName,proto3" json:"loginName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryUserRequest) Reset()         { *m = QueryUserRequest{} }
func (m *QueryUserRequest) String() string { return proto.CompactTextString(m) }
func (*QueryUserRequest) ProtoMessage()    {}
func (*QueryUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{9}
}

func (m *QueryUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUserRequest.Unmarshal(m, b)
}
func (m *QueryUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUserRequest.Marshal(b, m, deterministic)
}
func (m *QueryUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserRequest.Merge(m, src)
}
func (m *QueryUserRequest) XXX_Size() int {
	return xxx_messageInfo_QueryUserRequest.Size(m)
}
func (m *QueryUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserRequest proto.InternalMessageInfo

func (m *QueryUserRequest) GetLoginName() string {
	if m != nil {
		return m.LoginName
	}
	return ""
}

type QueryUserResponse struct {
	Error                *Error         `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data                 *QueryUserData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *QueryUserResponse) Reset()         { *m = QueryUserResponse{} }
func (m *QueryUserResponse) String() string { return proto.CompactTextString(m) }
func (*QueryUserResponse) ProtoMessage()    {}
func (*QueryUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{10}
}

func (m *QueryUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUserResponse.Unmarshal(m, b)
}
func (m *QueryUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUserResponse.Marshal(b, m, deterministic)
}
func (m *QueryUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserResponse.Merge(m, src)
}
func (m *QueryUserResponse) XXX_Size() int {
	return xxx_messageInfo_QueryUserResponse.Size(m)
}
func (m *QueryUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserResponse proto.InternalMessageInfo

func (m *QueryUserResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *QueryUserResponse) GetData() *QueryUserData {
	if m != nil {
		return m.Data
	}
	return nil
}

type LogoutEvent struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SentTime             int64    `protobuf:"varint,2,opt,name=sentTime,proto3" json:"sentTime,omitempty"`
	UserID               string   `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutEvent) Reset()         { *m = LogoutEvent{} }
func (m *LogoutEvent) String() string { return proto.CompactTextString(m) }
func (*LogoutEvent) ProtoMessage()    {}
func (*LogoutEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{11}
}

func (m *LogoutEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutEvent.Unmarshal(m, b)
}
func (m *LogoutEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutEvent.Marshal(b, m, deterministic)
}
func (m *LogoutEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutEvent.Merge(m, src)
}
func (m *LogoutEvent) XXX_Size() int {
	return xxx_messageInfo_LogoutEvent.Size(m)
}
func (m *LogoutEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutEvent.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutEvent proto.InternalMessageInfo

func (m *LogoutEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LogoutEvent) GetSentTime() int64 {
	if m != nil {
		return m.SentTime
	}
	return 0
}

func (m *LogoutEvent) GetUserID() string {
	if m != nil {
		return m.UserID
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
	return fileDescriptor_82b5829f48cfb8e5, []int{12}
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
	proto.RegisterType((*MakeTokenRequest)(nil), "com.island.code.srv.auth.MakeTokenRequest")
	proto.RegisterType((*MakeTokenResponse)(nil), "com.island.code.srv.auth.MakeTokenResponse")
	proto.RegisterType((*MakeTokenData)(nil), "com.island.code.srv.auth.MakeTokenData")
	proto.RegisterType((*DelTokenRequest)(nil), "com.island.code.srv.auth.DelTokenRequest")
	proto.RegisterType((*DelTokenResponse)(nil), "com.island.code.srv.auth.DelTokenResponse")
	proto.RegisterType((*GetCachedTokenRequest)(nil), "com.island.code.srv.auth.GetCachedTokenRequest")
	proto.RegisterType((*GetCachedTokenResponse)(nil), "com.island.code.srv.auth.GetCachedTokenResponse")
	proto.RegisterType((*GetCachedTokenData)(nil), "com.island.code.srv.auth.GetCachedTokenData")
	proto.RegisterType((*QueryUserData)(nil), "com.island.code.srv.auth.QueryUserData")
	proto.RegisterType((*QueryUserRequest)(nil), "com.island.code.srv.auth.QueryUserRequest")
	proto.RegisterType((*QueryUserResponse)(nil), "com.island.code.srv.auth.QueryUserResponse")
	proto.RegisterType((*LogoutEvent)(nil), "com.island.code.srv.auth.LogoutEvent")
	proto.RegisterType((*Error)(nil), "com.island.code.srv.auth.Error")
}

func init() { proto.RegisterFile("proto/auth/auth.proto", fileDescriptor_82b5829f48cfb8e5) }

var fileDescriptor_82b5829f48cfb8e5 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0x6d, 0x6b, 0x13, 0x41,
	0x10, 0xc7, 0x7b, 0xb9, 0x5c, 0xb1, 0x13, 0x6a, 0xe2, 0xd2, 0x96, 0x10, 0x04, 0xe5, 0x40, 0xaa,
	0x51, 0xaf, 0x25, 0xe2, 0x2b, 0xdf, 0x58, 0xbd, 0x2a, 0xc1, 0x07, 0xec, 0x51, 0x3f, 0xc0, 0x7a,
	0x37, 0x24, 0x47, 0x2f, 0xb7, 0xf5, 0x76, 0x2f, 0x52, 0xfc, 0x00, 0xbe, 0x16, 0xbf, 0x97, 0x9f,
	0x49, 0x76, 0xb3, 0xd9, 0xec, 0x25, 0x1e, 0x1c, 0x04, 0xfa, 0x26, 0xdc, 0xec, 0xce, 0xcc, 0x7f,
	0xe6, 0xb7, 0x33, 0x04, 0x0e, 0xaf, 0x0b, 0x26, 0xd8, 0x09, 0x2d, 0xc5, 0x54, 0xfd, 0x04, 0xca,
	0x26, 0xfd, 0x98, 0xcd, 0x82, 0x94, 0x67, 0x34, 0x4f, 0x82, 0x98, 0x25, 0x18, 0xf0, 0x62, 0x1e,
	0xc8, 0x7b, 0xff, 0x1d, 0xf4, 0x3e, 0xd1, 0x2b, 0xbc, 0x64, 0x57, 0x98, 0x47, 0xf8, 0xbd, 0x44,
	0x2e, 0xc8, 0x11, 0xec, 0x96, 0x1c, 0x8b, 0x71, 0xd8, 0x77, 0x1e, 0x3a, 0x8f, 0xf7, 0x22, 0x6d,
	0x91, 0x01, 0xdc, 0x49, 0x70, 0x9e, 0xc6, 0x38, 0x0e, 0xfb, 0x2d, 0x75, 0x63, 0x6c, 0xff, 0x97,
	0x03, 0xf7, 0xac, 0x44, 0xfc, 0x9a, 0xe5, 0x1c, 0xc9, 0x4b, 0xf0, 0xb0, 0x28, 0x58, 0xa1, 0x12,
	0x75, 0x46, 0x0f, 0x82, 0xba, 0x3a, 0x82, 0x73, 0xe9, 0x16, 0x2d, 0xbc, 0xc9, 0x2b, 0x68, 0x27,
	0x54, 0x50, 0x25, 0xd2, 0x19, 0x1d, 0xd7, 0x47, 0x19, 0xc5, 0x90, 0x0a, 0x1a, 0xa9, 0x20, 0xff,
	0x11, 0xec, 0x57, 0x8e, 0xc9, 0x01, 0x78, 0x42, 0x1a, 0xba, 0x9b, 0x85, 0xe1, 0x1f, 0x43, 0x37,
	0xc4, 0xac, 0xd2, 0xf7, 0xff, 0x1d, 0xc7, 0xd0, 0x5b, 0x39, 0x6e, 0xd5, 0x97, 0xff, 0x01, 0x0e,
	0xdf, 0xa3, 0x78, 0x4b, 0xe3, 0x29, 0x26, 0x5b, 0x13, 0xff, 0xed, 0xc0, 0xd1, 0x7a, 0xb6, 0xed,
	0xb0, 0xbf, 0xae, 0x60, 0x7f, 0x56, 0x1f, 0x55, 0x95, 0xb5, 0xd8, 0x0f, 0x81, 0x6c, 0xde, 0xd5,
	0x70, 0xfd, 0xe3, 0xc0, 0xfe, 0x45, 0x89, 0xc5, 0xcd, 0x57, 0x8e, 0x85, 0xf2, 0xab, 0xa3, 0x70,
	0x1f, 0xf6, 0x32, 0x36, 0x49, 0xf3, 0xcf, 0x74, 0x86, 0x1a, 0xc3, 0xea, 0x40, 0x32, 0x52, 0xc6,
	0x97, 0x1f, 0x49, 0xdf, 0x5d, 0x30, 0x5a, 0xda, 0x2b, 0xe5, 0xb6, 0xa5, 0x5c, 0xa1, 0xea, 0xad,
	0x51, 0x3d, 0x85, 0x9e, 0x29, 0x6a, 0xf9, 0x3a, 0x15, 0x7d, 0x67, 0x4d, 0x5f, 0x4d, 0xbe, 0x15,
	0x72, 0x4b, 0x93, 0x5f, 0x21, 0xa7, 0xe9, 0x5f, 0x40, 0xe7, 0x23, 0x9b, 0xb0, 0x52, 0x9c, 0xcf,
	0x31, 0x17, 0xe4, 0x2e, 0xb4, 0xd2, 0x44, 0xd7, 0xdb, 0x4a, 0x13, 0xd9, 0x36, 0xc7, 0x5c, 0x5c,
	0xa6, 0x9a, 0xa2, 0x1b, 0x19, 0xdb, 0x42, 0xef, 0xda, 0xe8, 0xfd, 0xe7, 0xe0, 0xa9, 0xfa, 0x08,
	0x81, 0xb6, 0x2c, 0x40, 0xa5, 0xf3, 0x22, 0xf5, 0x4d, 0x7a, 0xe0, 0xce, 0xf8, 0x44, 0xbf, 0x88,
	0xfc, 0x1c, 0xfd, 0x75, 0xa1, 0x7d, 0x56, 0x8a, 0x29, 0xc9, 0xa0, 0x2b, 0x97, 0xf0, 0x2c, 0x8e,
	0x91, 0x73, 0x35, 0x09, 0x64, 0xd8, 0x60, 0x8d, 0x35, 0xf1, 0xc1, 0xd3, 0x46, 0xbe, 0x0b, 0xd4,
	0xfe, 0x0e, 0x99, 0x01, 0x09, 0x31, 0x93, 0x34, 0x6c, 0xc1, 0x27, 0xf5, 0x49, 0xd6, 0x36, 0x7f,
	0x30, 0x6c, 0xe2, 0x6a, 0xe4, 0x7e, 0xc2, 0x81, 0x99, 0x72, 0x5b, 0xf0, 0xa4, 0xe9, 0xc6, 0x2c,
	0x65, 0x4f, 0x9b, 0x07, 0x18, 0xf1, 0x0c, 0xba, 0xe6, 0xed, 0xdf, 0xdc, 0xa8, 0x0d, 0x18, 0x36,
	0x18, 0x93, 0x06, 0x64, 0x37, 0x86, 0xd8, 0xdf, 0xf9, 0xb6, 0xab, 0xfe, 0x3f, 0x5e, 0xfc, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0x92, 0x1c, 0x35, 0x65, 0x58, 0x06, 0x00, 0x00,
}
