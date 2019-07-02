package main

import (
	"mall/app/api/web/order/conf"
	"mall/app/api/web/order/server/http/server"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}

	server.Init(conf.Conf)
}
