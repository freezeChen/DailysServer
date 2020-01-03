package rpc

import (
	"context"

	"DailysServer/logic/service"
	"DailysServer/proto"
)

type LogicHandler struct {
	svc        *service.LogicService
	connClient proto.ConnectService
}

func NewLogicHandler(svc *service.LogicService, rpc proto.ConnectService) *LogicHandler {
	return &LogicHandler{
		svc:        svc,
		connClient: rpc,
	}
}

func (l *LogicHandler) SendMessage(ctx context.Context, req *proto.MessageReq, reply *proto.EmptyReply) error {
	err := l.svc.SaveMessage(req)
	if err != nil {
		return err
	}

}

func (l *LogicHandler) Connect(ctx context.Context, req *proto.ConnectReq, reply *proto.ConnectReply) error {
	err := l.svc.Online(req.Sid, req.Uid)
	if err != nil {
		return err
	}

	list, err := l.svc.GetContactList(req.Uid)
	if err == nil {
		//l.connClient.PushMessage(ctx,&proto.PushMessageReq{})
	}



	return err
}

func (l *LogicHandler) DisConnect(context.Context, *proto.DisConnectReq, *proto.DisConnectReply) error {
	panic("implement me")
}
