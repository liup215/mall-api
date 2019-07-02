package main

import (
	"mall/app/api/web/user/conf"
	"mall/app/api/web/user/server/http/server"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}

	server.Init(conf.Conf)
}
