/*
   @Time : 2019-07-25 14:32
   @Author : frozenChen
   @File : Server
   @Software: KingMaxWMS_APP_API
*/
package connect

import (
	"context"
	"encoding/json"
	"time"

	"DailysServer/connect/conf"
	models "DailysServer/logic/model"
	"DailysServer/proto"

	"github.com/freezeChen/studio-library/zlog"
	"github.com/micro/go-micro/client"
)

const (
	_HandshakeTimeout = 8 * time.Second
	_HeartBeat        = 8 * time.Minute
)

type Server struct {
	Bucket      *Bucket
	round       *Round
	logicClient proto.LogicService
}

func NewServer(c *conf.Config) *Server {
	return &Server{
		Bucket:      NewBucket(),
		round:       NewRound(),
		logicClient: proto.NewLogicService("vip.frozen.srv.logic", client.DefaultClient),
	}

}

func (server *Server) BatchPush(ids []int64, notice []byte) {
	var msg = new(proto.Proto)
	msg.Opr = proto.OpSendMsg

	msg.Body = notice

	for _, v := range ids {
		ch := server.Bucket.Get(v)
		if ch != nil {
			zlog.Debug("Push")
			ch.Push(msg)
		}
	}
}

func (server *Server) Push(id int64, notice []byte) {
	if ch := server.Bucket.Get(id); ch != nil {
		var msg = new(proto.Proto)
		msg.Opr = proto.OpSendMsg
		msg.Body = notice
		ch.Push(msg)
	}

}

func (server *Server) Operate(ctx context.Context, p *proto.Proto) (err error) {

	switch p.Opr {
	case proto.OpSendMsg:
		var msg models.Message
		err = json.Unmarshal(p.Body, &msg)
		if err != nil {
			return
		}
		_, err = server.logicClient.SendMessage(ctx, &proto.MessageReq{
			SenderId:    msg.SenderId,
			RecipientId: msg.RecipientId,
			Type:        msg.Type,
			Content:     msg.Content,
		})
	}

	return
}

func (server *Server) Connect(ctx context.Context, sid string, uid int64) error {
	_, err := server.logicClient.Connect(ctx, &proto.ConnectReq{Uid: uid, Sid: sid})
	return err
}

func (server *Server) DisConnect(ctx context.Context, sid string, uid int64) error {
	_, err := server.logicClient.DisConnect(ctx, &proto.DisConnectReq{Uid: uid, Sid: sid})
	return err
}
