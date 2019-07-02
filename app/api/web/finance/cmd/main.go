package main

import (
	"mall/app/api/web/finance/conf"
	"mall/app/api/web/finance/server/http/server"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	server.Init(conf.Conf)
}
