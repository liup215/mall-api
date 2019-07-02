package main

import (
	"mall/app/api/web/wechat/conf"
	"mall/app/api/web/wechat/server/http/server"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	server.Init(conf.Conf)
}
