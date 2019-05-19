package main

import (
	"mall/app/service/main/finance/conf"
	"mall/app/service/main/finance/server/http/server"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	server.Init(conf.Conf)
}
