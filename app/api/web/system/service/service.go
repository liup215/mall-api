package service

import (
	"mall/app/api/web/system/conf"
	"mall/app/api/web/system/dao"
)

type Service struct {
	d *dao.Dao
}

func New(c *conf.Config) *Service {
	return &Service{
		d: dao.New(c),
	}
}
