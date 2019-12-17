package main

import (
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
	}
	zlog.InitLogger(conf.Conf.Log)
	server := connect.NewServer(conf.Conf)

	if err := connect.InitTCP(server, conf.Conf.TCPConfig); err != nil {
		panic(err)
		return
	}

	connect.InitWebSocket(server, conf.Conf.WebSocket.Addr)

	grpc.New(conf.Conf.RpcServer, server)

}
