package logic

import (
	"time"

	"DailysServer/logic/conf"
	"DailysServer/logic/rpc"
	"DailysServer/logic/service"
	"DailysServer/proto"

	"github.com/micro/go-micro"
)

func New(c *conf.RpcServer) {
	s := micro.NewService(
		micro.Name(c.Name),
		micro.RegisterTTL(time.Duration(c.TTL)*time.Second),
		micro.RegisterInterval(time.Duration(c.Interval)*time.Second))

	s.Init()

	connectRpc := proto.NewConnectService("vip.frozen.srv.connect", s.Client())

	logicService := service.NewLogicService(conf.GetConf())



	if err := proto.RegisterLogicHandler(s.Server(), rpc.NewLogicHandler(logicService,connectRpc)); err != nil {
		panic(err)
		return
	}

	if err := s.Run(); err != nil {
		panic(err)
		return
	}

}
