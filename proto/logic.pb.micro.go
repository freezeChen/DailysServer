// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: logic.proto

package proto

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// Api Endpoints for Logic service

func NewLogicEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Logic service

type LogicService interface {
	SendMessage(ctx context.Context, in *MessageReq, opts ...client.CallOption) (*EmptyReply, error)
	Connect(ctx context.Context, in *ConnectReq, opts ...client.CallOption) (*ConnectReply, error)
	DisConnect(ctx context.Context, in *DisConnectReq, opts ...client.CallOption) (*DisConnectReply, error)
	GetConversationList(ctx context.Context, in *ConversationListReq, opts ...client.CallOption) (*ConversationListReply, error)
}

type logicService struct {
	c    client.Client
	name string
}

func NewLogicService(name string, c client.Client) LogicService {
	return &logicService{
		c:    c,
		name: name,
	}
}

func (c *logicService) SendMessage(ctx context.Context, in *MessageReq, opts ...client.CallOption) (*EmptyReply, error) {
	req := c.c.NewRequest(c.name, "Logic.SendMessage", in)
	out := new(EmptyReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicService) Connect(ctx context.Context, in *ConnectReq, opts ...client.CallOption) (*ConnectReply, error) {
	req := c.c.NewRequest(c.name, "Logic.Connect", in)
	out := new(ConnectReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicService) DisConnect(ctx context.Context, in *DisConnectReq, opts ...client.CallOption) (*DisConnectReply, error) {
	req := c.c.NewRequest(c.name, "Logic.DisConnect", in)
	out := new(DisConnectReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicService) GetConversationList(ctx context.Context, in *ConversationListReq, opts ...client.CallOption) (*ConversationListReply, error) {
	req := c.c.NewRequest(c.name, "Logic.GetConversationList", in)
	out := new(ConversationListReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Logic service

type LogicHandler interface {
	SendMessage(context.Context, *MessageReq, *EmptyReply) error
	Connect(context.Context, *ConnectReq, *ConnectReply) error
	DisConnect(context.Context, *DisConnectReq, *DisConnectReply) error
	GetConversationList(context.Context, *ConversationListReq, *ConversationListReply) error
}

func RegisterLogicHandler(s server.Server, hdlr LogicHandler, opts ...server.HandlerOption) error {
	type logic interface {
		SendMessage(ctx context.Context, in *MessageReq, out *EmptyReply) error
		Connect(ctx context.Context, in *ConnectReq, out *ConnectReply) error
		DisConnect(ctx context.Context, in *DisConnectReq, out *DisConnectReply) error
		GetConversationList(ctx context.Context, in *ConversationListReq, out *ConversationListReply) error
	}
	type Logic struct {
		logic
	}
	h := &logicHandler{hdlr}
	return s.Handle(s.NewHandler(&Logic{h}, opts...))
}

type logicHandler struct {
	LogicHandler
}

func (h *logicHandler) SendMessage(ctx context.Context, in *MessageReq, out *EmptyReply) error {
	return h.LogicHandler.SendMessage(ctx, in, out)
}

func (h *logicHandler) Connect(ctx context.Context, in *ConnectReq, out *ConnectReply) error {
	return h.LogicHandler.Connect(ctx, in, out)
}

func (h *logicHandler) DisConnect(ctx context.Context, in *DisConnectReq, out *DisConnectReply) error {
	return h.LogicHandler.DisConnect(ctx, in, out)
}

func (h *logicHandler) GetConversationList(ctx context.Context, in *ConversationListReq, out *ConversationListReply) error {
	return h.LogicHandler.GetConversationList(ctx, in, out)
}
