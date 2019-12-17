package rpc

import (
	"context"

	"DailysServer/logic/service"
	"DailysServer/proto"
)

type Logic struct {
	svc *service.LogicService
}

func (l *Logic) SendMessage(ctx context.Context, req *proto.MessageReq, reply *proto.EmptyReply) error {




}

func (l *Logic) Connect(context.Context, *proto.ConnectReq, *proto.ConnectReply) error {
	panic("implement me")
}

func (l *Logic) DisConnect(context.Context, *proto.DisConnectReq, *proto.DisConnectReply) error {
	panic("implement me")
}
