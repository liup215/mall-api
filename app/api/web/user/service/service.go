package service

import (
	"mall/app/api/web/user/conf"
	"mall/app/api/web/user/dao"
)

type Service struct {
	authKey string
	d       *dao.Dao
}

func New(c *conf.Config) *Service {
	return &Service{
		d:       dao.New(c),
		authKey: c.AuthKey,
	}
}
