package main

import (
	"mall/app/api/admin/user/conf"
	"mall/app/api/admin/user/server"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}

	server.Init(conf.Conf)
}
