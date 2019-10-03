package main

import (
	"DailysServer/logic"
	"DailysServer/logic/conf"
)

func main() {

	if err := conf.Init(); err != nil {
		panic(err)
		return
	}

	logic.New(conf.Conf.RpcServer)

}
