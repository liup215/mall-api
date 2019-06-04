package service

import (
	"mall/app/service/main/finance/conf"
	"mall/app/service/main/finance/dao"
	memberClient "mall/app/service/main/member/server/http/client"
)

type Service struct {
	d  *dao.Dao
	mc *memberClient.Client
}

func New(c *conf.Config) *Service {
	return &Service{
		d:  dao.New(c),
		mc: memberClient.New(c.MemberClient),
	}
}
