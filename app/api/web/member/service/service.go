package service

import (
	"mall/app/api/web/member/conf"
	"mall/app/api/web/member/dao"
)

type Service struct {
	d *dao.Dao
}

func New(c *conf.Config) *Service {
	return &Service{
		d: dao.New(c),
	}
}
