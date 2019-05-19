package dao

import (
	"github.com/jinzhu/gorm"
	"mall/app/service/main/member/conf"
	"mall/lib/database/orm"
)

type Dao struct {
	orm *gorm.DB
}

func New(c *conf.Config) *Dao {
	db := orm.New(c.Orm)
	return &Dao{
		orm: db,
	}
}
