package main

import (
	"mall/app/api/web/pay/conf"
	"mall/app/api/web/pay/server"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}

	server.Init(conf.Conf)
}
