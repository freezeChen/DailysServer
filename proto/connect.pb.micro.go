// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: connect.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Connect service

func NewConnectEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Connect service

type ConnectService interface {
	PushMessage(ctx context.Context, in *PushMessageReq, opts ...client.CallOption) (*EmptyReply, error)
	BatchMessage(ctx context.Context, in *BatchMessageReq, opts ...client.CallOption) (*EmptyReply, error)
	PushContact(ctx context.Context, in *PushContactReq, opts ...client.CallOption) (*EmptyReply, error)
}

type connectService struct {
	c    client.Client
	name string
}

func NewConnectService(name string, c client.Client) ConnectService {
	return &connectService{
		c:    c,
		name: name,
	}
}

func (c *connectService) PushMessage(ctx context.Context, in *PushMessageReq, opts ...client.CallOption) (*EmptyReply, error) {
	req := c.c.NewRequest(c.name, "Connect.pushMessage", in)
	out := new(EmptyReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectService) BatchMessage(ctx context.Context, in *BatchMessageReq, opts ...client.CallOption) (*EmptyReply, error) {
	req := c.c.NewRequest(c.name, "Connect.batchMessage", in)
	out := new(EmptyReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectService) PushContact(ctx context.Context, in *PushContactReq, opts ...client.CallOption) (*EmptyReply, error) {
	req := c.c.NewRequest(c.name, "Connect.pushContact", in)
	out := new(EmptyReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Connect service

type ConnectHandler interface {
	PushMessage(context.Context, *PushMessageReq, *EmptyReply) error
	BatchMessage(context.Context, *BatchMessageReq, *EmptyReply) error
	PushContact(context.Context, *PushContactReq, *EmptyReply) error
}

func RegisterConnectHandler(s server.Server, hdlr ConnectHandler, opts ...server.HandlerOption) error {
	type connect interface {
		PushMessage(ctx context.Context, in *PushMessageReq, out *EmptyReply) error
		BatchMessage(ctx context.Context, in *BatchMessageReq, out *EmptyReply) error
		PushContact(ctx context.Context, in *PushContactReq, out *EmptyReply) error
	}
	type Connect struct {
		connect
	}
	h := &connectHandler{hdlr}
	return s.Handle(s.NewHandler(&Connect{h}, opts...))
}

type connectHandler struct {
	ConnectHandler
}

func (h *connectHandler) PushMessage(ctx context.Context, in *PushMessageReq, out *EmptyReply) error {
	return h.ConnectHandler.PushMessage(ctx, in, out)
}

func (h *connectHandler) BatchMessage(ctx context.Context, in *BatchMessageReq, out *EmptyReply) error {
	return h.ConnectHandler.BatchMessage(ctx, in, out)
}

func (h *connectHandler) PushContact(ctx context.Context, in *PushContactReq, out *EmptyReply) error {
	return h.ConnectHandler.PushContact(ctx, in, out)
}
