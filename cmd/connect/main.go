package main

import (
	"log"

	"DailysServer/connect"
	"DailysServer/connect/conf"
	"DailysServer/connect/grpc"

	"github.com/freezeChen/studio-library/zlog"
)

type Tag struct {
	RpcAddReSs string
}

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
		return
	}
	log.Printf("%+v,%+v",conf.Conf,conf.Conf)
	//println(conf.Conf)
	zlog.InitLogger(conf.Conf.Log)

	server := connect.NewServer(conf.Conf)
	connect.InitWebSocket(server, conf.Conf.WebSocket.Addr)

	grpc.New(conf.Conf.RpcServer, server)

}
