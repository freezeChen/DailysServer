package main

import (
	"DailysServer/logic"
	"DailysServer/logic/conf"

	"github.com/freezeChen/studio-library/zlog"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
		return
	}
	zlog.InitLogger(conf.GetConf().Log)
	logic.New(conf.GetConf().RpcServer)
}
