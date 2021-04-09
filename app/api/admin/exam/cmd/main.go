package main

import (
	"mall/app/api/admin/exam/conf"
	"mall/app/api/admin/exam/server"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	server.Init(conf.Conf)
}
