package grpc

import (
	"context"

	"DailysServer/connect"
	"DailysServer/proto"
)

type ConnectHandler struct {
	server *connect.Server
}

func (c *ConnectHandler) PushContact(ctx context.Context, req *proto.PushContactReq, reply *proto.EmptyReply) error {
	c.server.PushContact(req.Uid, []byte(req.Content))
	return nil
}

func NewConnectHandler(conn *connect.Server) *ConnectHandler {
	return &ConnectHandler{
		server: conn,
	}
}

func (c *ConnectHandler) PushMessage(ctx context.Context, req *proto.PushMessageReq, reply *proto.EmptyReply) error {
	c.server.Push(req.Uid, []byte(req.Content))
	return nil
}

func (c *ConnectHandler) BatchMessage(context.Context, *proto.BatchMessageReq, *proto.EmptyReply) error {
	panic("implement me")
}
