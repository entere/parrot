// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package com_island_code_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for User service

type UserService interface {
	// 获取用户详情
	GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error)
	// 更新用户
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...client.CallOption) (*UpdateUserResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "com.island.code.srv.user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error) {
	req := c.c.NewRequest(c.name, "User.GetUser", in)
	out := new(GetUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...client.CallOption) (*UpdateUserResponse, error) {
	req := c.c.NewRequest(c.name, "User.UpdateUser", in)
	out := new(UpdateUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	// 获取用户详情
	GetUser(context.Context, *GetUserRequest, *GetUserResponse) error
	// 更新用户
	UpdateUser(context.Context, *UpdateUserRequest, *UpdateUserResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		GetUser(ctx context.Context, in *GetUserRequest, out *GetUserResponse) error
		UpdateUser(ctx context.Context, in *UpdateUserRequest, out *UpdateUserResponse) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) GetUser(ctx context.Context, in *GetUserRequest, out *GetUserResponse) error {
	return h.UserHandler.GetUser(ctx, in, out)
}

func (h *userHandler) UpdateUser(ctx context.Context, in *UpdateUserRequest, out *UpdateUserResponse) error {
	return h.UserHandler.UpdateUser(ctx, in, out)
}
