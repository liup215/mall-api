package service

import (
	"mall/app/service/main/member/conf"
	"mall/app/service/main/member/dao"
)

type Service struct {
	d *dao.Dao
}

func New(c *conf.Config) *Service {
	return &Service{
		d: dao.New(c),
	}
}
