package main

import (
	"mall/app/api/web/exam/server"
	"mall/app/service/main/exam/conf"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	server.Init(conf.Conf)
}
