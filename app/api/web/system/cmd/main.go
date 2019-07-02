package main

import (
	"mall/app/api/web/system/conf"
	"mall/app/api/web/system/server/http/server"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}

	server.Init(conf.Conf)
}
