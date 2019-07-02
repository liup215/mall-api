package service

import (
	"mall/app/api/web/wechat/conf"
	"mall/app/api/web/wechat/dao"
)

type Service struct {
	d *dao.Dao
}

func New(c *conf.Config) *Service {
	return &Service{
		d: dao.New(c),
	}
}
