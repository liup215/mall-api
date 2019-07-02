package main

import (
	"mall/app/api/web/member/conf"
	"mall/app/api/web/member/server/http/server"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}

	server.Init(conf.Conf)
}
