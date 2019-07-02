package main

import (
	"mall/app/api/web/account/conf"
	"mall/app/api/web/account/server/http/server"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	server.Init(conf.Conf)
}
