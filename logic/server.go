package logic

import (
	"time"

	"DailysServer/logic/conf"
	"DailysServer/logic/rpc"
	"DailysServer/proto"

	"github.com/micro/go-micro"
)



func New(c *conf.RpcServer) {
	service := micro.NewService(
		micro.Name(c.Name),
		micro.RegisterTTL(time.Duration(c.TTL)*time.Second),
		micro.RegisterInterval(time.Duration(c.Interval)*time.Second))

	service.Init()

	if err := proto.RegisterLogicHandler(service.Server(), &rpc.Logic{}); err != nil {
		panic(err)
		return
	}

	if err := service.Run(); err != nil {
		panic(err)
		return
	}

}
