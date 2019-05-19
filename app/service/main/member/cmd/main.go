package main

import (
	"mall/app/service/main/member/conf"
	"mall/app/service/main/member/server"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}

	server.Init(conf.Conf)
}
