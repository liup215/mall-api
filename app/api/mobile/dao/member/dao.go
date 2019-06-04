package member

import (
	"mall/app/api/mobile/conf"
	mmbClient "mall/app/service/main/member/server/grpc/client"
)

type Dao struct {
	mclit *mmbClient.Client
}

func New(c *conf.Config) *Dao {
	return &Dao{
		mclit: mmbClient.New(conf.MmbServerAddress),
	}
}
