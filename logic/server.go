package logic

import (
	"context"
	"time"

	"DailysServer/logic/conf"
	"DailysServer/proto"

	"github.com/micro/go-micro"
)

type Server struct {
}

func New(c *conf.RpcServer) {
	service := micro.NewService(
		micro.Name(c.Name),
		micro.RegisterTTL(time.Duration(c.TTL)*time.Second),
		micro.RegisterInterval(time.Duration(c.Interval)*time.Second))

	service.Init()

	if err := proto.RegisterLogicHandler(service.Server(), &Server{}); err != nil {
		panic(err)
		return
	}

	if err := service.Run(); err != nil {
		panic(err)
		return
	}

}
func (s Server) Single(context.Context, *proto.SingleReq, *proto.SingleReply) error {
	panic("implement me")
}
