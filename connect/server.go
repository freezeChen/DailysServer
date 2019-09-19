/*
   @Time : 2019-07-25 14:32
   @Author : frozenChen
   @File : Server
   @Software: KingMaxWMS_APP_API
*/
package connect

import (
	"context"
	"time"

	"DailysServer/connect/conf"
	"DailysServer/proto"

	"github.com/freezeChen/studio-library/zlog"
)

const (
	_HandshakeTimeout = 8 * time.Second
	_HeartBeat        = 8 * time.Minute
)

type Server struct {

	bucket *Bucket
	round  *Round
}

func NewServer(c *conf.Config) *Server {
	return &Server{
		bucket: NewBucket(),
		round:  NewRound(),
	}

}

//func (s *Server) Start() {
//	micro.NewService()
//
//	service := web.NewService(
//		web.Name("vip.frozen.api.mail"),
//		web.RegisterTTL(75*time.Second),
//		web.Address(":9091"),
//		web.RegisterInterval(60*time.Second),
//	)
//
//	if err := service.Init(); err != nil {
//		panic(err)
//		return
//	}
//
//	service.HandleFunc("/mail/socket", func(writer http.ResponseWriter, request *http.Request) {
//		InitWebSocket(s, writer, request)
//	})
//
//	go func() {
//		if err := service.Run(); err != nil {
//			panic(err)
//			return
//		}
//	}()
//}

func (s *Server) batchPush(ids []string, notice []byte) {
	var msg = new(proto.Proto)
	msg.Opr = proto.OpSendMsg

	msg.Body = notice

	for _, v := range ids {
		ch := s.bucket.Get(v)
		if ch != nil {
			zlog.Debug("push")
			ch.Push(msg)
		}
	}
}

func (s *Server) push(id string, notice []byte) {
	if ch := s.bucket.Get(id); ch != nil {
		var msg = new(proto.Proto)
		msg.Opr = proto.OpSendMsg
		msg.Body = notice
		ch.Push(msg)
	}

}

func (s *Server) Operate(ctx context.Context, p *proto.Proto) error {

}
