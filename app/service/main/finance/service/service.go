package service

import (
	"mall/app/service/main/finance/conf"
	"mall/app/service/main/finance/dao"
)

type Service struct {
	d *dao.Dao
}

func New(c *conf.Config) *Service {
	return &Service{
		d: dao.New(c),
	}
}
