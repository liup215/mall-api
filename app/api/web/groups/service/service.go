package service

import (
	"mall/app/api/web/groups/conf"
	"mall/app/api/web/groups/dao"
)

type Service struct {
	d *dao.Dao
}

func New(c *conf.Config) *Service {
	return &Service{
		d: dao.New(c),
	}
}
