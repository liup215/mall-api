package service

import (
	"mall/app/api/web/order/conf"
	"mall/app/api/web/order/dao"
)

type Service struct {
	d *dao.Dao
}

func New(c *conf.Config) *Service {
	return &Service{
		d: dao.New(c),
	}
}
