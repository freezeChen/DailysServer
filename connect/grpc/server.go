package grpc

import (
	"time"

	"DailysServer/connect"
	"DailysServer/connect/conf"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
)

type server struct {
	connectSrv *connect.Server
}

func New(c *conf.RpcServer, s *connect.Server) {

	service := grpc.NewService(
		micro.Name(c.Name),
		micro.RegisterTTL(time.Duration(c.TTL)*time.Second),
		micro.RegisterInterval(time.Duration(c.Interval)*time.Second),
	)
	service.Init()

	if err := service.Run(); err != nil {
		panic(err)
	}

}
