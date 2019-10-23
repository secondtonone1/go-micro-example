// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: db.proto

package dbproto

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

// Client API for DBService service

type DBService interface {
	CheckLogin(ctx context.Context, in *CheckLoginReq, opts ...client.CallOption) (*CheckLoginRsp, error)
	RegisterUsr(ctx context.Context, in *RegisterUsrReq, opts ...client.CallOption) (*RegisterUsrRsp, error)
	CheckUsrEmail(ctx context.Context, in *CheckUsrEmailReq, opts ...client.CallOption) (*CheckUsrEmailRsp, error)
	//修改密码
	ChangePwd(ctx context.Context, in *ResetPwdReq, opts ...client.CallOption) (*ResetPwdRsp, error)
}

type dBService struct {
	c    client.Client
	name string
}

func NewDBService(name string, c client.Client) DBService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "dbproto"
	}
	return &dBService{
		c:    c,
		name: name,
	}
}

func (c *dBService) CheckLogin(ctx context.Context, in *CheckLoginReq, opts ...client.CallOption) (*CheckLoginRsp, error) {
	req := c.c.NewRequest(c.name, "DBService.CheckLogin", in)
	out := new(CheckLoginRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBService) RegisterUsr(ctx context.Context, in *RegisterUsrReq, opts ...client.CallOption) (*RegisterUsrRsp, error) {
	req := c.c.NewRequest(c.name, "DBService.RegisterUsr", in)
	out := new(RegisterUsrRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBService) CheckUsrEmail(ctx context.Context, in *CheckUsrEmailReq, opts ...client.CallOption) (*CheckUsrEmailRsp, error) {
	req := c.c.NewRequest(c.name, "DBService.CheckUsrEmail", in)
	out := new(CheckUsrEmailRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBService) ChangePwd(ctx context.Context, in *ResetPwdReq, opts ...client.CallOption) (*ResetPwdRsp, error) {
	req := c.c.NewRequest(c.name, "DBService.ChangePwd", in)
	out := new(ResetPwdRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DBService service

type DBServiceHandler interface {
	CheckLogin(context.Context, *CheckLoginReq, *CheckLoginRsp) error
	RegisterUsr(context.Context, *RegisterUsrReq, *RegisterUsrRsp) error
	CheckUsrEmail(context.Context, *CheckUsrEmailReq, *CheckUsrEmailRsp) error
	//修改密码
	ChangePwd(context.Context, *ResetPwdReq, *ResetPwdRsp) error
}

func RegisterDBServiceHandler(s server.Server, hdlr DBServiceHandler, opts ...server.HandlerOption) error {
	type dBService interface {
		CheckLogin(ctx context.Context, in *CheckLoginReq, out *CheckLoginRsp) error
		RegisterUsr(ctx context.Context, in *RegisterUsrReq, out *RegisterUsrRsp) error
		CheckUsrEmail(ctx context.Context, in *CheckUsrEmailReq, out *CheckUsrEmailRsp) error
		ChangePwd(ctx context.Context, in *ResetPwdReq, out *ResetPwdRsp) error
	}
	type DBService struct {
		dBService
	}
	h := &dBServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DBService{h}, opts...))
}

type dBServiceHandler struct {
	DBServiceHandler
}

func (h *dBServiceHandler) CheckLogin(ctx context.Context, in *CheckLoginReq, out *CheckLoginRsp) error {
	return h.DBServiceHandler.CheckLogin(ctx, in, out)
}

func (h *dBServiceHandler) RegisterUsr(ctx context.Context, in *RegisterUsrReq, out *RegisterUsrRsp) error {
	return h.DBServiceHandler.RegisterUsr(ctx, in, out)
}

func (h *dBServiceHandler) CheckUsrEmail(ctx context.Context, in *CheckUsrEmailReq, out *CheckUsrEmailRsp) error {
	return h.DBServiceHandler.CheckUsrEmail(ctx, in, out)
}

func (h *dBServiceHandler) ChangePwd(ctx context.Context, in *ResetPwdReq, out *ResetPwdRsp) error {
	return h.DBServiceHandler.ChangePwd(ctx, in, out)
}
