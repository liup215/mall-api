package service

import (
	"mall/app/api/web/account/conf"
	"mall/app/api/web/account/dao"
)

type Service struct {
	d *dao.Dao
}

func New(c *conf.Config) *Service {
	return &Service{
		d: dao.New(c),
	}
}
