/*
   @Time : 2019-07-25 14:32
   @Author : frozenChen
   @File : server
   @Software: KingMaxWMS_APP_API
*/
package server

import (
	"net/http"
	"time"

	"24on/mail_srv/handler"

	"github.com/freezeChen/studio-library/zlog"
	"github.com/micro/go-micro/web"
)

const (
	_HandshakeTimeout = 8 * time.Second
	_HeartBeat        = 8 * time.Minute
)

type server struct {
	h      *handler.Handler
	tokens *tokens.TokenManager
	bucket *Bucket
	round  *Round
}

func NewServer(handler *handler.Handler) *server {
	return &server{
		h:      handler,
		tokens: tokens.Init(),
		bucket: NewBucket(),
		round:  NewRound(),
	}
}

func (s *server) Start() {
	service := web.NewService(
		web.Name("vip.frozen.api.mail"),
		web.RegisterTTL(75*time.Second),
		web.Address(":9091"),
		web.RegisterInterval(60*time.Second),
	)

	if err := service.Init(); err != nil {
		panic(err)
		return
	}

	service.HandleFunc("/mail/socket", func(writer http.ResponseWriter, request *http.Request) {
		InitWebSocket(s, writer, request)
	})

	go func() {
		if err := service.Run(); err != nil {
			panic(err)
			return
		}
	}()
}

func (s *server) getCount(uid int64) int64 {

	counts, err := s.h.GetMailCounts(uid)
	if err != nil {
		return 0
	}
	return counts
}

func (s *server) batchPush(ids []string, notice []byte) {
	var msg = new(Proto)
	msg.Opr = OpSendMsg

	msg.Body = notice

	for _, v := range ids {
		ch := s.bucket.Get(v)
		if ch != nil {
			zlog.Debug("push")
			ch.Push(msg)
		}
	}
}

func (s *server) push(id string, notice []byte) {
	if ch := s.bucket.Get(id); ch != nil {
		var msg = new(Proto)
		msg.Opr = OpSendMsg
		msg.Body = notice
		ch.Push(msg)
	}

}
