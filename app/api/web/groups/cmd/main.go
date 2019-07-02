package main

import (
	"mall/app/api/web/groups/conf"
	"mall/app/api/web/groups/server/http/server"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	server.Init(conf.Conf)
}
