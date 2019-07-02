package service

import (
	"mall/app/api/web/finance/conf"
	"mall/app/api/web/finance/dao"
	memberClient "mall/app/api/web/member/server/http/client"
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
