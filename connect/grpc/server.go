package grpc

import (
	"DailysServer/connect"
	"DailysServer/connect/conf"
	"github.com/micro/go-micro"
	"time"
)

type server struct {
	srv *connect.Server
}

func New(c *conf.RpcServer, s *connect.Server) *server {
	service := micro.NewService(
		micro.Name(c.Name),
		micro.RegisterTTL(time.Duration(c.TTL)*time.Second),
		micro.RegisterInterval(time.Duration(c.Interval)*time.Second),
	)
	service.Init()

	go func() {
		if err := service.Run(); err != nil {
			panic(err)
		}
	}()

	return &server{
		srv: s,
	}
}
