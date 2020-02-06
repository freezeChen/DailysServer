package grpc

import (
	"time"

	"DailysServer/connect"
	"DailysServer/connect/conf"
	"DailysServer/proto"

	"github.com/google/uuid"
	"github.com/micro/go-micro"
)

func New(s *connect.Server) {
	conf.Conf.RpcServer.Id = uuid.New().String()

	service := micro.NewService(
		micro.Name(conf.Conf.RpcServer.Name),
		//micro.Metadata(map[string]string{"id": conf.Conf.RpcServer.Id}),
		micro.RegisterTTL(time.Duration(conf.Conf.RpcServer.TTL)*time.Second),
		micro.RegisterInterval(time.Duration(conf.Conf.RpcServer.Interval)*time.Second),
	)
	service.Init()

	err := proto.RegisterConnectHandler(service.Server(), NewConnectHandler(s))
	if err != nil {
		panic(err)
	}

	if err := service.Run(); err != nil {
		panic(err)
	}

}
